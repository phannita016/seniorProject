package main

import (
	"context"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/phannita016/seniorProject/apis"
	"github.com/phannita016/seniorProject/config"
	"github.com/phannita016/seniorProject/dtos"
	"github.com/phannita016/seniorProject/services"
	"github.com/phannita016/seniorProject/stores"
	"github.com/phannita016/seniorProject/x/libs"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("ReadInConfig", slog.Any("err", err))
		slog.Info("Load configuration in system environment.")
	}

	opt := config.Option{
		ServiceName: viper.GetString("service"),
		URI:         viper.GetString("mongo.uri"),
		Addr:        viper.GetString("mongo.host"),
		Username:    viper.GetString("mongo.username"),
		Password:    viper.GetString("mongo.password"),
		Database:    viper.GetString("mongo.database"),
	}

	db, err := config.ConnectDB(opt)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	slog.Info("Connection Database-Server...", slog.String("uri", opt.URI), slog.String("addr", opt.Addr))

	storeUser := stores.NewStore[dtos.User](db.Database)
	serviveUser := services.NewUser(storeUser, libs.ContextTimeout())
	username := viper.GetString("admin.username")
	password := viper.GetString("admin.password")
	if len(username) > 0 && len(password) > 0 {
		if err = serviveUser.Mocking(username, password); err != nil {
			slog.Error("serviceUser.Mocking", slog.String("user", username), slog.Any("err", err))
		}
	}

	f := fiber.New(apis.Fc)
	// f.Use(limiter.New(apis.))
	f.Use(recover.New())
	f.Use(cors.New())
	f.Use(compress.New())
	f.Use(requestid.New())

	f.Use(apis.Authorization())

	g := f.Group(apis.APIVersion1)
	apis.NewHandleLogin(g, serviveUser)

	if err := f.Listen(":8081"); err != nil {
		panic(err)
	}
}

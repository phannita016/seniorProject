package apps

import (
	"context"
	"log/slog"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/phannita016/seniorProject/config"
	"github.com/phannita016/seniorProject/dtos"
	"github.com/phannita016/seniorProject/services"
	"github.com/phannita016/seniorProject/stores"
	"github.com/phannita016/seniorProject/x/libs"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

type Apps interface {
	Runner(context.Context) error
	Close(context.Context) error
}

type apps struct {
	fiber  *fiber.App
	config *config.AppConfig
	mgc    *config.MONGOClient
}

func NewApps(conf *config.AppConfig) (Apps, error) {
	var (
		dbs = conf.DatabaseServer
	)

	mgc, err := config.ConnectDB(config.Option(dbs))
	if err != nil {
		return nil, err
	}
	slog.Info("Connection Database-Server...", slog.String("uri", dbs.URI), slog.String("addr", dbs.Addr))

	storeUser := stores.NewStore[dtos.User](mgc.Database)
	serviveUser := services.NewUser(storeUser, libs.ContextTimeout())
	username := viper.GetString("admin.username")
	password := viper.GetString("admin.password")
	if len(username) > 0 && len(password) > 0 {
		if err = serviveUser.Mocking(username, password); err != nil {
			slog.Error("serviceUser.Mocking", slog.String("user", username), slog.Any("err", err))
		}
	}

	var f = NewAPIs(serviveUser)

	return &apps{
		fiber:  f,
		config: conf,
		mgc:    mgc,
	}, nil
}

func (a *apps) Runner(ctx context.Context) error {
	g, c := errgroup.WithContext(ctx)
	g.Go(func() error {
		addr := net.JoinHostPort(a.config.Server.Address, a.config.Server.Port)
		return a.fiber.Listen(addr)
	})
	g.Go(func() error {
		<-c.Done()
		return a.fiber.ShutdownWithContext(c)
	})

	return g.Wait()
}

func (a *apps) Close(ctx context.Context) error {
	g, c := errgroup.WithContext(ctx)
	g.Go(func() error {
		return a.fiber.ShutdownWithContext(c)
	})

	slog.Info("stopping application.")
	return g.Wait()
}

package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	MONGOClient struct {
		*mongo.Client
		*mongo.Database
	}

	Option struct {
		ServiceName string
		URI         string
		Addr        string
		Username    string
		Password    string
		Database    string
	}
)

func ConnectDB(opt Option) (*MONGOClient, error) {
	opts := options.Client()
	if len(opt.URI) > 0 {
		opts.ApplyURI(opt.URI)
	} else {
		opts.ApplyURI("mongodb://" + opt.Addr)
	}

	opts.SetDirect(true)
	opts.SetAppName(opt.ServiceName)
	opts.SetSocketTimeout(20 * time.Second)
	opts.SetServerSelectionTimeout(5 * time.Second)
	opts.SetConnectTimeout(5 * time.Second)
	opts.SetMaxConnIdleTime(5 * time.Second)
	opts.SetMinPoolSize(10)
	opts.SetMaxPoolSize(100)
	if (len(opt.Username) > 0) && (len(opt.Password) > 0) {
		opts.SetAuth(options.Credential{
			Username: opt.Username,
			Password: opt.Password,
		})
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	var ctx = context.Background()
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &MONGOClient{Client: client, Database: client.Database(opt.Database)}, nil
}

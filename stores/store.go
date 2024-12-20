package stores

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var optSort = options.FindOptions{Sort: bson.M{"create_at": -1}}

type Store[T any] interface {
	FindAll(ctx context.Context) ([]*T, error)
	FindOneKeyValue(ctx context.Context, key string, value any) (*T, error)
	Create(ctx context.Context, t T) error
}

type mgo[T any] struct {
	c *mongo.Collection
}

func NewStore[T any](mgd *mongo.Database) *mgo[T] {
	var t T
	return &mgo[T]{c: mgd.Collection(reflect.TypeOf(t).Name())}
}

func (m *mgo[T]) FindAll(ctx context.Context) ([]*T, error) {
	cur, err := m.c.Find(ctx, bson.D{}, &optSort)
	if err != nil {
		return nil, err
	}

	defer func() { err = cur.Close(ctx) }()

	var ts []*T
	if err = cur.All(ctx, &ts); err != nil {
		return nil, err
	}
	if ts == nil {
		ts = []*T{}
	}

	return ts, nil
}

func (m *mgo[T]) FindOneKeyValue(ctx context.Context, key string, value any) (*T, error) {
	var t T
	if err := m.c.FindOne(ctx, primitive.D{bson.E{Key: key, Value: value}}, &options.FindOneOptions{Sort: bson.M{"create_at": -1}}).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *mgo[T]) Create(ctx context.Context, t T) error {
	_, err := m.c.InsertOne(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

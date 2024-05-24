package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"org.idev.bunny/backend/common/errors"
	"org.idev.bunny/backend/utils"
)

type Client struct {
	url    string
	dbName string
	cli    *mongo.Client
	db     *mongo.Database
}

func NewMongoClient(ctx context.Context, url string, dbName string) (*Client, error) {
	opts := options.Client().ApplyURI(url)

	cli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := cli.Ping(ctx, nil); err != nil {
		return nil, errors.New("ping mongo fail")
	}

	db := cli.Database(dbName)

	return &Client{
		url: url,
		cli: cli,
		db:  db,
	}, nil
}

func (m *Client) Close(ctx context.Context) error {
	return m.cli.Disconnect(ctx)
}

// @return inserted id, error
func (m *Client) SaveOne(ctx context.Context, collection string, document any) (string, error) {
	res, err := m.db.Collection(collection).InsertOne(ctx, document)
	if err != nil {
		return "", err
	}
	insertedID := res.InsertedID.(primitive.ObjectID).String()
	return insertedID, nil
}

func (m *Client) SaveMany(ctx context.Context, collection string, documents []any) error {
	res, err := m.db.Collection(collection).InsertMany(ctx, documents)
	fmt.Print(res)
	return err
}

func (m *Client) FindOne(ctx context.Context, collection string, filter map[string]any, target any) error {
	res := m.db.Collection(collection).FindOne(ctx, filter)
	if res.Err() != nil {
		return res.Err()
	}
	if err := res.Decode(target); err != nil {
		return err
	}
	return nil
}

func (m *Client) FindMany(ctx context.Context, collection string, filter map[string]any, page int64, pageSize int64, target any) error {
	opts := &options.FindOptions{
		Skip:  utils.Int64((page - 1) * pageSize),
		Limit: &pageSize,
	}
	res, err := m.db.Collection(collection).Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	if err := res.All(ctx, target); err != nil {
		return err
	}
	return nil
}

func (m *Client) UpdateOne(ctx context.Context, collection string, filter map[string]any, update map[string]any) error {
	res, err := m.db.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (m *Client) UpdateMany(ctx context.Context, collection string, filter map[string]any, update map[string]any) error {
	res, err := m.db.Collection(collection).UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (m *Client) DeleteOne(ctx context.Context, collection string, filter map[string]any) error {
	res, err := m.db.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (m *Client) DeleteMany(ctx context.Context, collection string, filter map[string]any) error {
	res, err := m.db.Collection(collection).DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("delete failed")
	}
	return nil
}

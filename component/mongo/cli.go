package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"org.idev.bunny/backend/common/errors"
)

type Client struct {
	url    string
	dbName string
	cli    *mongo.Client
	db     *mongo.Database
}

func NewMongoClient(ctx context.Context, url, dbName string) (*Client, error) {
	opts := options.Client().ApplyURI(url)

	cli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := cli.Ping(ctx, nil); err != nil {
		return nil, errors.New("failed to ping MongoDB")
	}

	return &Client{
		url:    url,
		dbName: dbName,
		cli:    cli,
		db:     cli.Database(dbName),
	}, nil
}

func (m *Client) Close(ctx context.Context) error {
	return m.cli.Disconnect(ctx)
}

func (m *Client) WithTransaction(ctx context.Context, fn func(cli *Client) error) error {
	session, err := m.cli.StartSession()
	if err != nil {
		return fmt.Errorf("failed to start session: %w", err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sc mongo.SessionContext) (interface{}, error) {
		if err := fn(m); err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

// SaveOne inserts a single document into the specified collection and returns the inserted ID.
func (m *Client) SaveOne(ctx context.Context, collection string, document interface{}) (string, error) {
	res, err := m.db.Collection(collection).InsertOne(ctx, document)
	if err != nil {
		return "", fmt.Errorf("failed to insert document: %w", err)
	}
	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get inserted ID")
	}
	return insertedID.Hex(), nil
}

func (m *Client) SaveMany(ctx context.Context, collection string, documents []interface{}) error {
	_, err := m.db.Collection(collection).InsertMany(ctx, documents)
	return err
}

func (m *Client) FindOne(ctx context.Context, collection string, filter map[string]interface{}, target interface{}) error {
	return m.db.Collection(collection).FindOne(ctx, filter).Decode(target)
}

func (m *Client) FindMany(ctx context.Context, collection string, filter map[string]interface{}, page, pageSize int64, target interface{}) error {
	opts := options.Find().
		SetSkip((page - 1) * pageSize).
		SetLimit(pageSize)

	cursor, err := m.db.Collection(collection).Find(ctx, filter, opts)
	if err != nil {
		return fmt.Errorf("failed to execute find: %w", err)
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, target)
}

func (m *Client) UpdateOne(ctx context.Context, collection string, filter, update map[string]interface{}) error {
	res, err := m.db.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}
	if res.MatchedCount == 0 {
		return errors.New("no document matched the filter")
	}
	return nil
}

func (m *Client) UpdateMany(ctx context.Context, collection string, filter, update map[string]interface{}) error {
	res, err := m.db.Collection(collection).UpdateMany(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update documents: %w", err)
	}
	if res.MatchedCount == 0 {
		return errors.New("no documents matched the filter")
	}
	return nil
}

func (m *Client) DeleteOne(ctx context.Context, collection string, filter map[string]interface{}) error {
	res, err := m.db.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}
	if res.DeletedCount == 0 {
		return errors.New("no document matched the filter")
	}
	return nil
}

func (m *Client) DeleteMany(ctx context.Context, collection string, filter map[string]interface{}) error {
	res, err := m.db.Collection(collection).DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete documents: %w", err)
	}
	if res.DeletedCount == 0 {
		return errors.New("no documents matched the filter")
	}
	return nil
}

func (m *Client) Count(ctx context.Context, collection string, filter map[string]interface{}) (int64, error) {
	return m.db.Collection(collection).CountDocuments(ctx, filter)
}

package redis

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	addr     string
	password string
	db       int

	cli *redis.Client
}

func NewClient(ctx context.Context, connURL string) (*Client, error) {
	opts, err := redis.ParseURL(connURL)
	if err != nil {
		return nil, err
	}
	cli := redis.NewClient(opts)
	if cli == nil {
		return nil, errors.New("failed to create Redis client")
	}

	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Client{
		addr:     opts.Addr,
		password: opts.Password,
		db:       opts.DB,
		cli:      cli,
	}, nil
}

func (r *Client) Close() error {
	return r.cli.Close()
}

func (r *Client) Ping(ctx context.Context) error {
	return r.cli.Ping(ctx).Err()
}

func (r *Client) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	if reflect.ValueOf(value).Kind() == reflect.Struct {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return err
		}
		value = jsonValue
	}
	return r.cli.Set(ctx, key, value, expiration).Err()
}

func (r *Client) SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error) {
	return r.cli.SetNX(ctx, key, value, expiration).Result()
}

func (r *Client) Get(ctx context.Context, key string) (any, error) {
	return r.cli.Get(ctx, key).Result()
}

func (r *Client) GetString(ctx context.Context, key string) (string, error) {
	return r.cli.Get(ctx, key).Result()
}

func (r *Client) GetInt(ctx context.Context, key string) (int, error) {
	return r.cli.Get(ctx, key).Int()
}

func (r *Client) GetInt64(ctx context.Context, key string) (int64, error) {
	return r.cli.Get(ctx, key).Int64()
}

func (r *Client) GetFloat32(ctx context.Context, key string) (float32, error) {
	return r.cli.Get(ctx, key).Float32()
}

func (r *Client) GetFloat64(ctx context.Context, key string) (float64, error) {
	return r.cli.Get(ctx, key).Float64()
}

func (r *Client) GetBool(ctx context.Context, key string) (bool, error) {
	return r.cli.Get(ctx, key).Bool()
}

func (r *Client) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return r.cli.Get(ctx, key).Bytes()
}

func (r *Client) GetTime(ctx context.Context, key string) (time.Time, error) {
	return r.cli.Get(ctx, key).Time()
}

func (r *Client) GetStruct(ctx context.Context, key string, target any) error {
	data, err := r.cli.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func (r *Client) GetUint64(ctx context.Context, key string) (uint64, error) {
	return r.cli.Get(ctx, key).Uint64()
}

func (r *Client) Incr(ctx context.Context, key string) (int64, error) {
	return r.cli.Incr(ctx, key).Result()
}

func (r *Client) Del(ctx context.Context, key string) error {
	return r.cli.Del(ctx, key).Err()
}

func (r *Client) Exists(ctx context.Context, key string) (bool, error) {
	res, err := r.cli.Exists(ctx, key).Result()
	return res == 1, err
}

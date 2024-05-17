package redis

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	addr     string
	password string
	db       int

	cli *redis.Client
}

func NewClient(connUrl string) (*RedisClient, error) {
	opts, err := redis.ParseURL(connUrl)
	if err != nil {
		return nil, err
	}
	cli := redis.NewClient(opts)
	if cli == nil {
		return nil, errors.New("create redis client fail")
	}

	return &RedisClient{
		cli: cli,
	}, nil
}

func (r *RedisClient) Close() error {
	return r.cli.Close()
}

func (r *RedisClient) Ping(ctx context.Context) error {
	_, err := r.cli.Ping(ctx).Result()
	return err
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Struct {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return err
		}
		value = jsonValue
	}
	return r.cli.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.cli.SetNX(ctx, key, value, expiration).Result()
}

func (r *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	return r.cli.Get(ctx, key).Result()
}

func (r *RedisClient) GetString(ctx context.Context, key string) (string, error) {
	return r.cli.Get(ctx, key).Result()
}

func (r *RedisClient) GetInt(ctx context.Context, key string) (int, error) {
	return r.cli.Get(ctx, key).Int()
}

func (r *RedisClient) GetInt64(ctx context.Context, key string) (int64, error) {
	return r.cli.Get(ctx, key).Int64()
}

func (r *RedisClient) GetFloat32(ctx context.Context, key string) (float32, error) {
	return r.cli.Get(ctx, key).Float32()
}

func (r *RedisClient) GetFloat64(ctx context.Context, key string) (float64, error) {
	return r.cli.Get(ctx, key).Float64()
}

func (r *RedisClient) GetBool(ctx context.Context, key string) (bool, error) {
	return r.cli.Get(ctx, key).Bool()
}

func (r *RedisClient) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return r.cli.Get(ctx, key).Bytes()
}

func (r *RedisClient) GetTime(ctx context.Context, key string) (time.Time, error) {
	return r.cli.Get(ctx, key).Time()
}

func (r *RedisClient) GetStruct(ctx context.Context, key string, value any) error {
	data, err := r.cli.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), value)
}

func (r *RedisClient) GetUint64(ctx context.Context, key string) (uint64, error) {
	return r.cli.Get(ctx, key).Uint64()
}

func (r *RedisClient) Incr(ctx context.Context, key string) (int64, error) {
	return r.cli.Incr(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.cli.Del(ctx, key).Err()
}

func (r *RedisClient) Exist(ctx context.Context, key string) (bool, error) {
	res, err := r.cli.Exists(ctx, key).Result()
	if res == 1 {
		return true, nil
	}
	return false, err
}

package redis

import (
	"context"
	"testing"
	"time"
)

func TestRedisSetAString(t *testing.T) {
	ctx := context.TODO()
	redisCli, err := NewClient(ctx, "redis://:bunny@localhost:6379/0")
	if err != nil {
		t.Fatal(err)
	}
	err = redisCli.Set(ctx, "key", "value", time.Second*0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisSetAInterface(t *testing.T) {
	ctx := context.TODO()
	redisCli, err := NewClient(ctx, "redis://:bunny@localhost:6379/0")
	if err != nil {
		t.Fatal(err)
	}
	type testStruct struct {
		Name string
	}
	err = redisCli.Set(ctx, "key", testStruct{Name: "test"}, time.Second*0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisGetStruct(t *testing.T) {
	ctx := context.TODO()
	redisCli, err := NewClient(ctx, "redis://:bunny@localhost:6379/0")
	if err != nil {
		t.Fatal(err)
	}
	type testStruct struct {
		Name string
	}
	var value testStruct
	err = redisCli.GetStruct(ctx, "key-not-found", &value)
	if err != nil {
		t.Fatal(err)
	}
	if value.Name != "test" {
		t.Fatal("get struct fail")
	}
}

func TestRedisSetAInt(t *testing.T) {
	ctx := context.TODO()
	redisCli, err := NewClient(ctx, "redis://:bunny@localhost:6379/0")
	if err != nil {
		t.Fatal(err)
	}

	err = redisCli.Set(ctx, "test", 1, time.Second*0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	ctx := context.TODO()
	redisCli, err := NewClient(ctx, "redis://:bunny@localhost:6379/0")
	if err != nil {
		t.Fatal(err)
	}
	err = redisCli.Del(ctx, "key")
	if err != nil {
		t.Fatal(err)
	}
}

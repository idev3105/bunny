package mongo

import (
	"context"
	"fmt"
	"testing"
)

func TestSaveAMap(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://bunny:bunny@localhost:27017/bunny", "bunny")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	insertedID, err := cli.SaveOne(ctx, "test", map[string]string{"name": "bunny"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Scanf("Inserted ID %v", insertedID))
}

func TestSaveAStruct(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://bunny:bunny@localhost:27017/bunny", "bunny")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	insertedID, err := cli.SaveOne(ctx, "test", Test{"bunny"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Scanf("Inserted ID %v", insertedID))
}

func TestSaveManyStruct(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://bunny:bunny@localhost:27017/bunny", "bunny")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	err = cli.SaveMany(ctx, "test", []interface{}{Test{"bunny2"}, Test{"bunny2"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetStruct(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://bunny:bunny@localhost:27017/bunny", "bunny")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	var target Test
	err = cli.FindOne(ctx, "test", map[string]any{"name": "bunny"}, &target)
	if err != nil {
		t.Fatal(err)
	}
	if target.Name != "bunny" {
		t.Errorf("Expected bunny, got %v", target.Name)
	}
}

func TestGetList(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://bunny:bunny@localhost:27017/bunny", "bunny")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		ID   string `bson:"_id"`
		Name string `bson:"name"`
	}
	var target []Test
	err = cli.FindMany(ctx, "test", map[string]any{"name": "bunny"}, 2, 2, &target)
	if err != nil {
		t.Fatal(err)
	}
	if len(target) != 2 {
		t.Errorf("Expected 2, got %v", len(target))
	}
}

func TestUpdateOne(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://bunny:bunny@localhost:27017/bunny", "bunny")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	err = cli.UpdateOne(ctx, "test", map[string]any{"name": "bunny"}, map[string]any{"$set": map[string]any{"name": "bunny-update"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateMany(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://bunny:bunny@localhost:27017/bunny", "bunny")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	err = cli.UpdateMany(ctx, "test", map[string]any{"name": "bunny"}, map[string]any{"$set": map[string]any{"name": "bunny-update"}})
	if err != nil {
		t.Fatal(err)
	}
}

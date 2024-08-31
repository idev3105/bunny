package elasticsearch

import (
	"context"
	"testing"
	"time"
)

func TestElasticsearchClient(t *testing.T) {
	// Initialize the Elasticsearch client
	ctx := context.Background()
	addresses := []string{"http://localhost:9200"}
	client, err := NewClient(ctx, addresses)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Test document
	testDoc := map[string]interface{}{
		"title":     "Test Document",
		"content":   "This is a test document for Elasticsearch.",
		"timestamp": time.Now().Format(time.RFC3339),
	}

	// Test Index
	t.Run("Index", func(t *testing.T) {
		err := client.Index(ctx, "test-index", "1", testDoc)
		if err != nil {
			t.Errorf("Failed to index document: %v", err)
		}
	})

	// Test Get
	t.Run("Get", func(t *testing.T) {
		doc, err := client.Get(ctx, "test-index", "1")
		if err != nil {
			t.Errorf("Failed to get document: %v", err)
		}
		if doc["_source"].(map[string]interface{})["title"] != testDoc["title"] {
			t.Errorf("Expected title %v, got %v", testDoc["title"], doc["_source"].(map[string]interface{})["title"])
		}
	})

	// Test Search
	t.Run("Search", func(t *testing.T) {
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"title": "Test Document",
				},
			},
		}
		results, err := client.Search(ctx, "test-index", query)
		if err != nil {
			t.Errorf("Failed to search: %v", err)
		}
		if len(results) != 1 {
			t.Errorf("Expected 1 result, got %d", len(results))
		}
		if results[0]["title"] != testDoc["title"] {
			t.Errorf("Expected title %v, got %v", testDoc["title"], results[0]["title"])
		}
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		err := client.Delete(ctx, "test-index", "1")
		if err != nil {
			t.Errorf("Failed to delete document: %v", err)
		}

		// Verify deletion
		_, err = client.Get(ctx, "test-index", "1")
		if err == nil {
			t.Error("Expected error when getting deleted document, got nil")
		}
	})
}

func TestElasticsearchClientErrors(t *testing.T) {
	// Initialize the Elasticsearch client with invalid address
	ctx := context.Background()
	addresses := []string{"http://invalid-host:9200"}
	_, err := NewClient(ctx, addresses)
	if err == nil {
		t.Error("Expected error when creating client with invalid address, got nil")
	}

	// Test Index error
	t.Run("IndexError", func(t *testing.T) {
		client, err := NewClient(ctx, []string{"http://localhost:9200"})
		if err != nil {
			t.Fatalf("Failed to create Elasticsearch client: %v", err)
		}
		defer client.Close()

		err = client.Index(ctx, "test-index", "1", make(chan int)) // Unmarshalable type
		if err == nil {
			t.Error("Expected error when indexing unmarshalable type, got nil")
		}
	})

	// Test Get error
	t.Run("GetError", func(t *testing.T) {
		client, err := NewClient(ctx, []string{"http://localhost:9200"})
		if err != nil {
			t.Fatalf("Failed to create Elasticsearch client: %v", err)
		}
		defer client.Close()

		_, err = client.Get(ctx, "non-existent-index", "1")
		if err == nil {
			t.Error("Expected error when getting from non-existent index, got nil")
		}
	})

	// Test Search error
	t.Run("SearchError", func(t *testing.T) {
		client, err := NewClient(ctx, []string{"http://localhost:9200"})
		if err != nil {
			t.Fatalf("Failed to create Elasticsearch client: %v", err)
		}
		defer client.Close()

		_, err = client.Search(ctx, "non-existent-index", nil)
		if err == nil {
			t.Error("Expected error when searching non-existent index, got nil")
		}
	})

	// Test Delete error
	t.Run("DeleteError", func(t *testing.T) {
		client, err := NewClient(ctx, []string{"http://localhost:9200"})
		if err != nil {
			t.Fatalf("Failed to create Elasticsearch client: %v", err)
		}
		defer client.Close()

		err = client.Delete(ctx, "non-existent-index", "1")
		if err == nil {
			t.Error("Expected error when deleting from non-existent index, got nil")
		}
	})
}

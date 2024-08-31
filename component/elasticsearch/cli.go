package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

// Client wraps the Elasticsearch client
type Client struct {
	cli *elasticsearch.Client
}

// NewClient creates a new Elasticsearch client
func NewClient(ctx context.Context, addresses []string) (*Client, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	cli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}

	info, err := cli.Info(cli.Info.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error getting Elasticsearch info: %w", err)
	}
	defer info.Body.Close()

	if info.IsError() {
		return nil, fmt.Errorf("error connecting to Elasticsearch: %s", info.String())
	}

	return &Client{cli: cli}, nil
}

// Close is a no-op function to maintain consistency with other clients
func (c *Client) Close() {}

// Index adds or updates a document in an index
func (c *Client) Index(ctx context.Context, index, id string, document interface{}) error {
	docBytes, err := json.Marshal(document)
	if err != nil {
		return fmt.Errorf("error marshaling document: %w", err)
	}

	res, err := c.cli.Index(
		index,
		bytes.NewReader(docBytes),
		c.cli.Index.WithContext(ctx),
		c.cli.Index.WithDocumentID(id),
	)
	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.String())
	}

	return nil
}

// Get retrieves a document by its ID
func (c *Client) Get(ctx context.Context, index, id string) (map[string]interface{}, error) {
	res, err := c.cli.Get(index, id, c.cli.Get.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error getting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error getting document: %s", res.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return result, nil
}

// Search performs a search query
func (c *Client) Search(ctx context.Context, index string, query map[string]interface{}) ([]map[string]interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	res, err := c.cli.Search(
		c.cli.Search.WithContext(ctx),
		c.cli.Search.WithIndex(index),
		c.cli.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, fmt.Errorf("error searching documents: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error searching documents: %s", res.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	hits, _ := result["hits"].(map[string]interface{})["hits"].([]interface{})
	documents := make([]map[string]interface{}, 0, len(hits))
	for _, hit := range hits {
		if doc, ok := hit.(map[string]interface{})["_source"].(map[string]interface{}); ok {
			documents = append(documents, doc)
		}
	}

	return documents, nil
}

// Delete removes a document from an index
func (c *Client) Delete(ctx context.Context, index, id string) error {
	res, err := c.cli.Delete(index, id, c.cli.Delete.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error deleting document: %s", res.String())
	}

	return nil
}

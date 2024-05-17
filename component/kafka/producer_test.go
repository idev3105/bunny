package kafka

import "testing"

func TestProducer(t *testing.T) {
	p, err := NewProducer("localhost", 9092)
	if err != nil {
		t.Errorf("Failed to create producer: %v", err)
	}
	defer p.Close()
	err = p.Send([]Message{
		{
			Key:   []byte("test1"),
			Topic: "test1",
			Value: []byte("Hello, World! 1"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = p.Send([]Message{
		{
			Key:   []byte("test2"),
			Topic: "test2",
			Value: []byte("Hello, World! 2"),
		},
	})
}

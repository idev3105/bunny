package kafka

type Message struct {
	Topic     string
	Key       []byte
	Partition int32
	Value     []byte
}

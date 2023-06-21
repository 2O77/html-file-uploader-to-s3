package messagequeue

type Message struct {
	FilePath string
}

type MessageQueue interface {
	PublishMessage(Message) error
}

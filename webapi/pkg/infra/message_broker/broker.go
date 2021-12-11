package messagebroker

type IMessageBroker interface{}

type messageBroker struct{}

func connect() interface{} {
	return nil
}

func (messageBroker) Publisher(exchange, queue, routingKey string, body interface{}, header interface{}) {
	// channel := connect()
}

func NewMessageBroker() IMessageBroker {
	return messageBroker{}
}

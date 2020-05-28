package messaging

type Builder struct {
	messageHandlers map[string]MessageHandler
}

func (builder Builder) WithHandler(handler MessageHandler, messageType string) {
	builder.messageHandlers[messageType] = handler
}

func (builder Builder) TryGetValue(messageType string) MessageHandler{
	return builder.messageHandlers[messageType]
}
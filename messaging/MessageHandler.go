package messaging

type MessageHandler interface {
	HandleMessageAsync([]byte)
}

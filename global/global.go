package global

// topic types
const (
	// Master Topic
	MTopic = 1000
	// Private Chat Topic
	PChatTopic = 2000

	// Group Chat Topic
	GChatTopic = 4000

	// Singe PUsh Topic
	SPushTopic = 3000

	// Broadcast PUsh Topic
	BPushTopic = 5000
)

// payload proto types
const (
	PayloadText     = 1
	PayloadProtoBuf = 2
	PayloadJson     = 3
)

package botapi

type BotApi interface {
	Handle(string, func())
	SendMessage(uint64, string)
}

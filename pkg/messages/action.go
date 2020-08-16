package messages

import "time"

type ActionInterface interface {
	Name() string
	SetMessageId(messageId uint32)
	MessageId() uint32
	SetTime(t time.Time)
	Time() time.Time
}

type Action struct {
	messageId   uint32
	requestTime time.Time
}

func (this *Action) SetMessageId(messageId uint32) {
	this.messageId = messageId
}

func (this *Action) MessageId() uint32 {
	return this.messageId
}

func (this *Action) SetTime(t time.Time) {
	this.requestTime = t
}
func (this *Action) Time() time.Time {
	return this.requestTime
}

func (this *Action) Name() string {
	return ""
}

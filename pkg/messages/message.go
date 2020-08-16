package messages

import (
	"encoding/binary"
	"github.com/vmihailenco/msgpack"
	"math"
	"sync"
)

var messageIdLocker = sync.Mutex{}
var messageId = uint32(1)

func NewMessage() *Message {
	return &Message{}
}

func NewActionMessage(action ActionInterface) *Message {
	m := &Message{
		Action: action.Name(),
		Value:  action,
	}
	return m
}

type Message struct {
	Id     uint32
	Action string
	Value  interface{}
}

func (this *Message) Marshal() ([]byte, error) {
	data, err := msgpack.Marshal(this.Value)
	if err != nil {
		return []byte{}, err
	}

	// id
	b := make([]byte, 8)
	if this.Id == 0 {
		this.Id = this.nextId()
	}
	binary.BigEndian.PutUint32(b, this.Id)

	// action length
	b2 := make([]byte, 8)
	binary.BigEndian.PutUint32(b2, uint32(len(this.Action)))
	b = append(b, b2...)

	// data length
	b3 := make([]byte, 8)
	binary.BigEndian.PutUint32(b3, uint32(len(data)))
	b = append(b, b3...)

	b = append(b, []byte(this.Action)...)
	return append(b, data...), nil
}

func (this *Message) nextId() uint32 {
	messageIdLocker.Lock()
	defer messageIdLocker.Unlock()

	if messageId == math.MaxUint32 {
		messageId = 1
	} else {
		messageId++
	}

	return messageId
}

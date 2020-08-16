package messages

import (
	"encoding/binary"
	"github.com/iwind/TeaGo/assert"
	"github.com/vmihailenco/msgpack"
	"testing"
)

func TestMessage_Marshal(t *testing.T) {
	a := assert.NewAssertion(t).Quiet()

	msg := NewMessage()
	msg.Action = "registerPlugin"
	msg.Value = []string{"a", "b", "c", "d", "e", "f"}
	data, err := msg.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))

	l1 := binary.BigEndian.Uint32(data[:8])
	l2 := binary.BigEndian.Uint32(data[8:16])
	a.Log("header:", l1, l2)

	action := string(data[16 : 16+l1])
	a.IsTrue(action == msg.Action)

	value := []string{}
	msgpack.Unmarshal(data[16+l1:16+l1+l2], &value)
	t.Logf("%#v", value)
}

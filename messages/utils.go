package messages

import (
	"errors"
	"github.com/vmihailenco/msgpack"
	"log"
	"reflect"
	"sync"
)

var actions = map[string]interface{}{}
var actionLocker = sync.Mutex{}

func RegisterAction(actionPtr ...interface{}) {
	actionLocker.Lock()
	defer actionLocker.Unlock()

	for _, action := range actionPtr {
		actionPtr, ok := action.(ActionInterface)
		if !ok {
			log.Println("[error]action '" + reflect.TypeOf(action).String() + "' should implement methods for 'ActionInterface'")
			continue
		}
		actions[actionPtr.Name()] = actionPtr
	}
}

func NewAction(action string) interface{} {
	ptr, found := actions[action]
	if !found {
		return nil
	}
	return reflect.New(reflect.TypeOf(ptr).Elem()).Interface()
}

func Unmarshal(action string, data []byte) (interface{}, error) {
	ptr := NewAction(action)
	if ptr == nil {
		return nil, errors.New("action '" + action + "' not found")
	}

	err := msgpack.Unmarshal(data, ptr)
	if err != nil {
		return nil, err
	}

	return ptr, nil
}

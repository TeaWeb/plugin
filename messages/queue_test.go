package messages

import (
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	reqAction := new(Action)
	reqAction.messageId = 1

	beforeTime := time.Now()

	go func() {
		//time.Sleep(1 * time.Second)
		time.Sleep(100 * time.Millisecond)
		respAction := new(Action)
		respAction.messageId = 1
		ActionQueue.Notify(respAction)
	}()

	respAction := ActionQueue.Wait(reqAction)
	t.Log("cost:", time.Since(beforeTime).Seconds()*1000, "ms")
	t.Logf("resp:%#v", respAction)
	t.Logf("times:%d, done:%d", len(ActionQueue.times), len(ActionQueue.done))
}

package messages

import (
	"sync"
	"time"
)

type Queue struct {
	locker  *sync.Mutex
	done    map[uint32]chan ActionInterface // id => chan
	times   map[uint32]time.Time
	Timeout time.Duration
}

var ActionQueue = NewQueue()

func NewQueue() *Queue {
	q := &Queue{
		locker:  &sync.Mutex{},
		done:    map[uint32]chan ActionInterface{},
		times:   map[uint32]time.Time{},
		Timeout: 10 * time.Second,
	}

	// 超时时间
	go func() {
		for {
			time.Sleep(1 * time.Second)

			q.locker.Lock()
			for messageId, t := range q.times {
				if time.Since(t) >= q.Timeout {
					q.done[messageId] <- new(EmptyAction)
					delete(q.done, messageId)
					delete(q.times, messageId)
				}
			}
			q.locker.Unlock()
		}
	}()

	return q
}

func (this *Queue) Wait(action ActionInterface) (respAction ActionInterface) {
	c := make(chan ActionInterface)

	this.locker.Lock()
	this.done[action.MessageId()] = c
	this.times[action.MessageId()] = time.Now()
	this.locker.Unlock()

	resp := <-c
	return resp
}

func (this *Queue) Notify(resp ActionInterface) {
	messageId := resp.MessageId()
	this.locker.Lock()
	c, found := this.done[messageId]
	this.locker.Unlock()
	if !found {
		return
	}
	c <- resp
	close(c)

	this.locker.Lock()
	delete(this.done, messageId)
	delete(this.times, messageId)
	this.locker.Unlock()
}

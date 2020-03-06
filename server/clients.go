package main

import (
	"log"
	"sync"
	"time"
)

type (
	clientsManager struct {
		sync.Mutex
		chans map[string]chan *comment
	}
	comment struct {
		Type      string
		User      string
		Text      string
		Timestamp string
		Channel   string
	}
)

func newClientsManager() *clientsManager {
	return &clientsManager{
		chans: make(map[string]chan *comment),
	}
}

func (c *clientsManager) add(key string) <-chan *comment {
	log.Printf("add channel: %s", key)
	c.Lock()
	defer c.Unlock()
	ch := make(chan *comment, 1000)
	c.chans[key] = ch

	return ch
}

func (c *clientsManager) delete(key string) {
	go func() {
		c.Lock()
		defer c.Unlock()
		ch, ok := c.chans[key]
		if !ok {
			return
		}
		close(ch)
		delete(c.chans, key)
		log.Printf("deleted channel: %s", key)
	}()
}

func (c *clientsManager) pushMsg(msg *comment) {
	if msg == nil {
		return
	}
	c.Lock()
	defer c.Unlock()
	for key, ch := range c.chans {
		select {
		case ch <- msg:
		case <-time.After(3 * time.Second):
			c.delete(key)
		}
	}
}

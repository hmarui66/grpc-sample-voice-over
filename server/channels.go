package main

import (
	"log"
	"sync"
	"time"
)

type (
	channels struct {
		sync.Mutex
		chans map[string]chan *message
	}
	message struct {
		Type      string
		User      string
		Text      string
		Timestamp string
		Channel   string
	}
)

func newChannels() *channels {
	return &channels{
		chans: make(map[string]chan *message),
	}
}

func (c *channels) add(key string) <-chan *message {
	log.Printf("add channel: %s", key)
	c.Lock()
	defer c.Unlock()
	ch := make(chan *message, 1000)
	c.chans[key] = ch

	return ch
}

func (c *channels) delete(key string) {
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

func (c *channels) pushMsg(msg *message) {
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

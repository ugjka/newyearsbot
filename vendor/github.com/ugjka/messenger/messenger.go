// Package messenger provides a simple broadcasting mechanism
package messenger

import (
	"fmt"
)

// Messenger instance. Must be invoked with New().
type Messenger struct {
	buffer    int
	drop      bool
	get       chan chan interface{}
	del       chan chan interface{}
	broadcast chan interface{}
	pool      map[chan interface{}]struct{}
	reset     chan struct{}
	kill      chan struct{}
	blocked   chan interface{} //for testing
	chanlens  chan []int       //for testing
}

// New creates a new Messenger instance.
//
// Buffer sets the buffer size to client channels, set to 0 if you want unbuffered behaviour.
//
// Drop makes broadcasts to drop if a client's buffer is full. Ignored with 0 buffer size.
func New(buffer uint, drop bool) *Messenger {
	m := setup(buffer, drop)
	go m.monitor()
	return m
}

func setup(buffer uint, drop bool) *Messenger {
	m := &Messenger{}
	m.buffer = int(buffer)
	m.drop = drop
	m.get = make(chan chan interface{})
	m.del = make(chan chan interface{})
	m.broadcast = make(chan interface{})
	m.pool = make(map[chan interface{}]struct{})
	m.reset = make(chan struct{})
	m.kill = make(chan struct{})
	return m
}

// For testing
func newNoMonitor(buffer uint, drop bool) *Messenger {
	m := setup(buffer, drop)
	return m
}

// Main loop where all the action happens.
func (m *Messenger) monitor() {
	client := make(chan interface{}, m.buffer)
	for {
		select {
		case m.get <- client:
			m.pool[client] = struct{}{}
			client = make(chan interface{}, m.buffer)
		case client := <-m.del:
			if _, ok := m.pool[client]; ok {
				close(client)
				delete(m.pool, client)
			}
		case <-m.reset:
			for client := range m.pool {
				close(client)
				delete(m.pool, client)
			}
		case <-m.kill:
			for client := range m.pool {
				close(client)
				delete(m.pool, client)
			}
			close(m.get)
			return
		case msg := <-m.broadcast:
			for client := range m.pool {
				if m.drop && m.buffer != 0 && len(client) == m.buffer {
					continue
				}
				select {
				case client <- msg:
				case m.blocked <- msg: // for testing
				}
			}
		case m.chanlens <- m.getlens(): // for testing
		}
	}
}

// For testing
func (m *Messenger) getlens() (l []int) {
	for k := range m.pool {
		l = append(l, len(k))
	}
	return
}

// Reset closes and removes all clients.
func (m *Messenger) Reset() {
	m.reset <- struct{}{}
}

// Kill closes and removes all clients
// and stops the monitor() goroutine of Messenger,
// making the Messenger instance unusable.
// Kill must only be called when all clients have exited.
func (m *Messenger) Kill() {
	m.kill <- struct{}{}
}

// Sub returns a new client.
// Clients can block Broadcast() unless drop is set.
// Clients should check whether the channel is closed or not.
// Returns an error if its Messenger instance is Killed
func (m *Messenger) Sub() (client chan interface{}, err error) {
	sub, ok := <-m.get
	if !ok {
		return nil, fmt.Errorf("can't subscribe, messenger killed")
	}
	return sub, nil
}

// Unsub unsubscribes a client and closes its channel.
func (m *Messenger) Unsub(client chan interface{}) {
	if m.drop {
		m.del <- client
		return
	}
	for {
		select {
		case <-client:
		case m.del <- client:
			return
		}
	}
}

// Broadcast broadcasts a message to all current clients.
// If a client is not listening and and drop is not set this will block.
func (m *Messenger) Broadcast(msg interface{}) {
	m.broadcast <- msg
}

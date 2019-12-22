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
	killed    chan struct{}
	len       chan int
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
	m := &Messenger{
		buffer:    int(buffer),
		drop:      drop,
		get:       make(chan chan interface{}),
		del:       make(chan chan interface{}),
		broadcast: make(chan interface{}),
		pool:      make(map[chan interface{}]struct{}),
		reset:     make(chan struct{}),
		kill:      make(chan struct{}),
		killed:    make(chan struct{}),
		len:       make(chan int),
	}
	if buffer == 0 {
		m.drop = false
	}
	return m
}

// For testing
func newNoMonitor(buffer uint, drop bool) *Messenger {
	return setup(buffer, drop)
}

// Main loop where all the action happens.
func (m *Messenger) monitor() {
	client := make(chan interface{}, m.buffer)
	for {
		select {
		case <-m.len:
			m.len <- len(m.pool)
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
			close(m.killed)
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
	select {
	case <-m.killed:
	case m.reset <- struct{}{}:
	}
}

// Kill closes and removes all clients
// and stops the monitor() goroutine of Messenger,
// making the Messenger instance unusable.
// Kill should only be called when all clients have exited.
func (m *Messenger) Kill() {
	select {
	case <-m.killed:
	case m.kill <- struct{}{}:
	}
}

// Sub returns a new client.
// Clients can block Broadcast() unless drop is set.
// Clients should check whether the channel is closed or not.
// Returns an error if its Messenger instance is Killed.
// You can ignore the err if you never intend to use Kill()
func (m *Messenger) Sub() (client chan interface{}, err error) {
	select {
	case sub := <-m.get:
		return sub, nil
	case <-m.killed:
		return nil, fmt.Errorf("can't subscribe, messenger killed")
	}
}

// Unsub unsubscribes a client and closes its channel.
func (m *Messenger) Unsub(client chan interface{}) {
	if m.drop {
		select {
		case m.del <- client:
		case <-m.killed:
		}
		return
	}
	for {
		select {
		case _, ok := <-client:
			if !ok {
				return
			}
		case m.del <- client:
			return
		}
	}
}

// Broadcast broadcasts a message to all current clients.
// If a client is not listening and drop is not set this will block.
func (m *Messenger) Broadcast(msg interface{}) {
	select {
	case <-m.killed:
	case m.broadcast <- msg:
	}
}

//Len gets the subscriber count
func (m *Messenger) Len() int {
	select {
	case m.len <- 0:
		return <-m.len
	case <-m.killed:
		return 0
	}
}

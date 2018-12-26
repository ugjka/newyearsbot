// Package messenger provides broadcasting mechanism
package messenger

import "fmt"

// Messenger object
type Messenger struct {
	get       chan chan interface{}
	del       chan chan interface{}
	broadcast chan interface{}
	pool      map[chan interface{}]struct{}
	reset     chan struct{}
	kill      chan struct{}
}

// New creates new Messenger
func New() *Messenger {
	m := &Messenger{}
	m.get = make(chan chan interface{})
	m.del = make(chan chan interface{})
	m.broadcast = make(chan interface{})
	m.pool = make(map[chan interface{}]struct{})
	m.reset = make(chan struct{})
	m.kill = make(chan struct{})
	go m.monitor()
	return m
}

func (m *Messenger) monitor() {
	for {
		tmp := make(chan interface{})
		select {
		case m.get <- tmp:
			m.pool[tmp] = struct{}{}
		case del := <-m.del:
			if _, ok := m.pool[del]; ok {
				close(del)
				delete(m.pool, del)
			}
		case <-m.reset:
			for k := range m.pool {
				close(k)
				delete(m.pool, k)
			}
		case <-m.kill:
			for k := range m.pool {
				close(k)
				delete(m.pool, k)
			}
			close(m.get)
			return
		case msg := <-m.broadcast:
			for k := range m.pool {
				k <- msg
			}
		}
	}
}

// Reset removes all clients
func (m *Messenger) Reset() {
	m.reset <- struct{}{}
}

// Kill removes all clients and stops the reading and writing goroutine
func (m *Messenger) Kill() {
	m.kill <- struct{}{}
}

// Sub subscribes a new client for reading broadcasts
func (m *Messenger) Sub() (client chan interface{}, err error) {
	sub := <-m.get
	if sub == nil {
		return nil, fmt.Errorf("can't sub, messenger stopped")
	}
	return sub, nil
}

// Unsub unsubscribes a client
func (m *Messenger) Unsub(client chan interface{}) {
	del := client
	for {
		select {
		case <-client:
		case m.del <- del:
			return
		}
	}
}

// Broadcast broadcasts a message to all current clients
func (m *Messenger) Broadcast(msg interface{}) {
	m.broadcast <- msg
}

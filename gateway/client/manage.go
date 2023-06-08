package client

import (
	"sync"

	"github.com/markgenuine/ever-client-go/domain"
)

type Manager interface {
	SetChannels(chan<- *domain.ClientResponse, <-chan struct{}) uint32
	DeleteRequestID(uint32)
	GetChannels(requestID uint32, delete bool) (chan<- *domain.ClientResponse, <-chan struct{}, bool)
}

type multiplexer struct {
	sync.Locker
	requestIDCounter uint32
	callbacks        map[uint32]manageChan
}

func NewStore() Manager {
	return &multiplexer{
		Locker:    &sync.Mutex{},
		callbacks: make(map[uint32]manageChan),
	}
}

type manageChan struct {
	responsChan chan<- *domain.ClientResponse
	close       <-chan struct{}
}

func (m *multiplexer) GetChannels(requestID uint32, toDelete bool) (chan<- *domain.ClientResponse, <-chan struct{}, bool) {
	m.Lock()
	defer m.Unlock()
	pair, isFound := m.callbacks[requestID]
	if isFound && toDelete {
		delete(m.callbacks, requestID)
	}

	return pair.responsChan, pair.close, isFound
}

func (m *multiplexer) SetChannels(responses chan<- *domain.ClientResponse, close <-chan struct{}) uint32 {
	m.Lock()
	defer m.Unlock()
	m.requestIDCounter++
	requestID := m.requestIDCounter
	m.callbacks[requestID] = manageChan{
		responsChan: responses,
		close:       close,
	}

	return requestID
}

func (m *multiplexer) DeleteRequestID(requestID uint32) {
	m.Lock()
	defer m.Unlock()
	delete(m.callbacks, requestID)
}

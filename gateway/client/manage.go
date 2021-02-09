package client

import (
	"sync"

	"github.com/move-ton/ton-client-go/domain"
)

// Manager ...
type Manager interface {
	SetChannels(chan<- *domain.ClientResponse, <-chan struct{}) int
	DeleteRequestID(int)
	GetChannels(requestID int, delete bool) (chan<- *domain.ClientResponse, <-chan struct{}, bool)
}

type multiplexer struct {
	sync.Locker
	requestIDCounter int
	callbacks        map[int]manageChan
}

// NewStore ...
func NewStore() Manager {
	return &multiplexer{
		Locker:    &sync.Mutex{},
		callbacks: make(map[int]manageChan),
	}
}

type manageChan struct {
	responsChan chan<- *domain.ClientResponse
	close       <-chan struct{}
}

// GetChannels ...
func (m *multiplexer) GetChannels(requestID int, toDelete bool) (chan<- *domain.ClientResponse, <-chan struct{}, bool) {
	m.Lock()
	defer m.Unlock()
	pair, isFound := m.callbacks[requestID]
	if isFound && toDelete {
		delete(m.callbacks, requestID)
	}

	return pair.responsChan, pair.close, isFound
}

//SetChannels ...
func (m *multiplexer) SetChannels(responses chan<- *domain.ClientResponse, close <-chan struct{}) int {
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

// DeleteRequestID ...
func (m *multiplexer) DeleteRequestID(requestID int) {
	m.Lock()
	defer m.Unlock()
	delete(m.callbacks, requestID)
}

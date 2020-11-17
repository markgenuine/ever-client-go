package client

import "sync"

var globalMultiplexer = NewMultiplexer()

type ResponseCode uint32

type RawResponse struct {
	Data  []byte
	Code  ResponseCode
	Error error
}

type channelPair struct {
	responses chan<- *RawResponse
	close     <-chan struct{}
}

type multiplexer struct {
	sync.Locker
	requestIDCounter uint32
	callbacks        map[uint32]channelPair
}

func (m *multiplexer) DeleteByRequestID(requestID uint32) {
	m.Lock()
	defer m.Unlock()
	delete(m.callbacks, requestID)
}

func (m *multiplexer) SetChannels(responses chan<- *RawResponse, close <-chan struct{}) uint32 {
	m.Lock()
	defer m.Unlock()
	m.requestIDCounter++
	requestID := m.requestIDCounter
	m.callbacks[requestID] = channelPair{
		responses: responses,
		close:     close,
	}

	return requestID
}

func (m *multiplexer) GetChannels(requestID uint32, toDelete bool) (chan<- *RawResponse, <-chan struct{}, bool) {
	m.Lock()
	defer m.Unlock()
	pair, isFound := m.callbacks[requestID]
	if isFound && toDelete {
		delete(m.callbacks, requestID)
	}

	return pair.responses, pair.close, isFound
}

type Multiplexer interface {
	SetChannels(chan<- *RawResponse, <-chan struct{}) uint32
	DeleteByRequestID(uint32)
	GetChannels(requestID uint32, delete bool) (chan<- *RawResponse, <-chan struct{}, bool)
}

func NewMultiplexer() Multiplexer {
	return &multiplexer{
		Locker:    &sync.Mutex{},
		callbacks: make(map[uint32]channelPair),
	}
}

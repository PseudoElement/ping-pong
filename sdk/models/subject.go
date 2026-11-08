package models

type Subscriber[T any] struct {
	Unsubscribe func()
	callback    func(T)
	idx         int16
}

type EventEmitter[T any] struct {
	subs      []Subscriber[T]
	lastValue T
	counter   int16
}

func NewEventEmitter[T any](initialValue T) *EventEmitter[T] {
	return &EventEmitter[T]{
		subs:      make([]Subscriber[T], 0),
		lastValue: initialValue,
	}
}

func (s *EventEmitter[T]) Value() T {
	return s.lastValue
}

func (s *EventEmitter[T]) Emit(value T) {
	s.lastValue = value
	for _, sub := range s.subs {
		sub.callback(value)
	}
}

func (s *EventEmitter[T]) Subscribe(callback func(T)) Subscriber[T] {
	counter := s.counter
	sub := Subscriber[T]{
		idx:      counter,
		callback: callback,
		Unsubscribe: func() {
			filteredSubs := make([]Subscriber[T], 0, len(s.subs))
			for _, sub := range s.subs {
				if sub.idx != counter {
					filteredSubs = append(filteredSubs, sub)
				}
			}
			s.subs = filteredSubs
		},
	}
	s.subs = append(s.subs, sub)
	s.counter++
	return sub
}

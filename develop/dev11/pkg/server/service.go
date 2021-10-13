package server

import (
	"sync"
	"time"
)

//Service consist business logic program
type Service interface {
	CreateEvent(userID int, ev Event) (err error)
	UpdateEvent(userID int, newEv, oldEv Event) (err error)
	DeleteEvent(userID int, ev Event) (err error)
	EventsForDay(userID int, ev Event) (notes []string, err error)
	EventsForWeek(userID int, ev Event) (notes []string, err error)
	EventsForMonth(userID int, ev Event) (notes []string, err error)
}

//Event data unit
type Event struct {
	Note string
	Date time.Time
}

//NewEvent create Event
func NewEvent(str string) (res Event, err error) {
	shortForm := "2006-01-02"

	res.Date, err = time.Parse(shortForm, str)
	if err != nil {
		return
	}
	return
}

type service struct {
	data map[int][]Event //map[user_id][]events
	mu   sync.RWMutex
}

func (s *service) CreateEvent(userID int, ev Event) (err error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[userID] = append(s.data[userID], ev)

	return
}

func (s *service) UpdateEvent(userID int, newEv, oldEv Event) (err error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	for ind, ev := range s.data[userID] {
		if ev == oldEv {
			s.data[userID][ind] = newEv
			return
		}
	}

	return
}

func (s *service) DeleteEvent(userID int, ev Event) (err error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	for ind, cur := range s.data[userID] {
		if cur == ev {
			s.data[userID] = append(s.data[userID][:ind], s.data[userID][ind+1:]...)
			return
		}
	}

	return
}

func (s *service) EventsForDay(userID int, ev Event) (notes []string, err error) {
	notes = []string{}

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, event := range s.data[userID] {
		if event.Date.Equal(ev.Date) {
			notes = append(notes, event.Note)
		}
	}

	return
}

func (s *service) EventsForWeek(userID int, ev Event) (notes []string, err error) {
	notes = []string{}

	s.mu.RLock()
	defer s.mu.RUnlock()

	y1, w1 := ev.Date.ISOWeek()
	for _, event := range s.data[userID] {
		y2, w2 := event.Date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			notes = append(notes, event.Note)
		}
	}

	return
}

func (s *service) EventsForMonth(userID int, ev Event) (notes []string, err error) {
	notes = []string{}

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, event := range s.data[userID] {
		if event.Date.Year() == ev.Date.Year() && event.Date.Month() == ev.Date.Month() {
			notes = append(notes, event.Note)
		}
	}

	return
}

//NewService create Service
func NewService() Service {
	return &service{
		data: map[int][]Event{},
		mu:   sync.RWMutex{},
	}
}

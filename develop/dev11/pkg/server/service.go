package server

type Service interface {
	CreateEvent(ev Event) (err error)
	UpdateEvent(newEv, oldEv Event) (err error)
	DeleteEvent(ev Event) (err error)
	EventsForDay(day string) (evs []Event, err error)
	EventsForWeek(week string) (evs []Event, err error)
	EventsForMonth(month string) (evs []Event, err error)
}

type Event struct {
	Day   string
	Week  string
	Month string
	Task  string
}

type service struct {
	//sync.Mutex
	Events []Event
}

func (s *service) CreateEvent(ev Event) (err error) {
	s.Events = append(s.Events, ev)

	return
}

func (s *service) UpdateEvent(newEv, oldEv Event) (err error) {

	//logic

	return
}

func (s *service) DeleteEvent(ev Event) (err error) {
	/*if idEvent < 0 || idEvent >= len(s.Events) {
		return errors.New("not found Event in the id")
	}*/

	//s.Events = append(s.Events[:idEvent], s.Events[idEvent+1:]...)

	return
}

func (s *service) EventsForDay(day string) (evs []Event, err error) {
	evs = make([]Event, 0, 5)

	for _, d := range s.Events {
		if d.Day == day {
			evs = append(evs, d)
		}
	}

	return
}

func (s *service) EventsForMonth(month string) (evs []Event, err error) {
	evs = make([]Event, 0, 5)

	for _, m := range s.Events {
		if m.Month == month {
			evs = append(evs, m)
		}
	}

	return
}
func (s *service) EventsForWeek(week string) (evs []Event, err error) {
	evs = make([]Event, 0, 5)

	for _, w := range s.Events {
		if w.Week == week {
			evs = append(evs, w)
		}
	}

	return
}

func NewService() Service {
	return &service{
		Events: []Event{},
	}
}

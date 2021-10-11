package server

import (
	"github.com/sirupsen/logrus"
)

type logMiddleware struct {
	serv   Service
	logger logrus.FieldLogger
}

func (lg *logMiddleware) CreateEvent(userID int, ev Event) (err error) {

	defer func() {
		l := lg.logger.WithError(err).WithFields(
			map[string]interface{}{
				"userID": userID,
				"Event":  ev,
			},
		)
		if err == nil {
			l.Info("CreateEvent")
		} else {
			l.Error("CreateEvent")
		}

	}()

	return lg.serv.CreateEvent(userID, ev)
}

func (lg *logMiddleware) UpdateEvent(userID int, newEv, oldEv Event) (err error) {

	defer func() {
		l := lg.logger.WithError(err).WithFields(
			map[string]interface{}{
				"userID":   userID,
				"NewEvent": newEv,
				"OldEvent": oldEv,
			},
		)
		if err == nil {
			l.Info("UpdateEvent")
		} else {
			l.Error("UpdateEvent")
		}

	}()

	return lg.serv.UpdateEvent(userID, newEv, oldEv)
}

func (lg *logMiddleware) DeleteEvent(userID int, ev Event) (err error) {

	defer func() {
		l := lg.logger.WithError(err).WithFields(
			map[string]interface{}{
				"userID": userID,
				"Event":  ev,
			},
		)
		if err == nil {
			l.Info("DeleteEvent")
		} else {
			l.Error("DeleteEvent")
		}

	}()

	return lg.serv.DeleteEvent(userID, ev)
}

func (lg *logMiddleware) EventsForDay(userID int, ev Event) (notes []string, err error) {

	defer func() {
		l := lg.logger.WithError(err).WithFields(
			map[string]interface{}{
				"userID": userID,
				"Event":  ev,
			},
		)
		if err == nil {
			l.Info("EventsForDay")
		} else {
			l.Error("EventsForDay")
		}

	}()

	return lg.serv.EventsForDay(userID, ev)
}

func (lg *logMiddleware) EventsForWeek(userID int, ev Event) (notes []string, err error) {

	defer func() {
		l := lg.logger.WithError(err).WithFields(
			map[string]interface{}{
				"userID": userID,
				"Event":  ev,
			},
		)
		if err == nil {
			l.Info("EventsForWeek")
		} else {
			l.Error("EventsForWeek")
		}

	}()

	return lg.serv.EventsForWeek(userID, ev)
}

func (lg *logMiddleware) EventsForMonth(userID int, ev Event) (notes []string, err error) {

	defer func() {
		l := lg.logger.WithError(err).WithFields(
			map[string]interface{}{
				"userID": userID,
				"Event":  ev,
			},
		)
		if err == nil {
			l.Info("EventsForMonth")
		} else {
			l.Error("EventsForMonth")
		}

	}()

	return lg.serv.EventsForMonth(userID, ev)
}

//NewLogMiddleware create LogMiddleware
func NewLogMiddleware(s Service, l logrus.FieldLogger) Service {
	return &logMiddleware{
		serv:   s,
		logger: l,
	}
}

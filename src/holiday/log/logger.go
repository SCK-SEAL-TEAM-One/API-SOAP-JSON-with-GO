package log

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Logs struct {
	ID          int
	CreatedTime time.Time
	LevelLog    string
	Description string
	Status      int
}

type Logger interface {
	Error(string) bool
	Info(string) bool
}
type LoggerMongo struct {
	Session *mgo.Session
}

func (l LoggerMongo) Error(s string) bool {
	return l.InsertLogs("Error", s, 403)
}

func (l LoggerMongo) Info(s string) bool {
	return l.InsertLogs("Info", s, 200)
}
func (l LoggerMongo) InsertLogs(levelLog, description string, status int) bool {
	fmt.Printf("insert success")
	logs := Logs{
		CreatedTime: time.Now(),
		LevelLog:    levelLog,
		Description: description,
		Status:      status,
	}
	err := l.Session.DB("holidayservice").C("logs").Insert(&logs)
	if err != nil {
		return true
	}
	return false
}

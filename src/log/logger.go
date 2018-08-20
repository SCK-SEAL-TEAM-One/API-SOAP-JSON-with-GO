package log

import "time"

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
	// DB *mgo.Database
}

func (l LoggerMongo) Logging(level, s string) bool {
	// return l.database.Collection("Log").Insert(level, s)
	return true
}

func (l LoggerMongo) Error(s string) bool {
	return l.Logging("Error", s)
}

func (l LoggerMongo) Info(s string) bool {
	return l.Logging("Info", s)
}

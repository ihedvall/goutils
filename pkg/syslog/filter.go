package syslog

import (
	"goutils/pkg/logchain"
	"time"
)

type Filter struct {
	Level    int
	Facility int
	Message  string
	Date     time.Time
	Data     []logchain.StructuredData
	Id       int64
	Offset   int
	Count    int
}

func (f *Filter) Clear() {
	f.Level = -1
	f.Facility = -1
	f.Message = ""
	f.Date = time.Time{}
}

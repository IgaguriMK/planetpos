package model

import (
	"log"
	"time"
)

type Distances struct {
	Distances []Distance `json:"distances"`
}

type Distance struct {
	Time int64   `json:"time"`
	From string  `json:"from"`
	To   string  `json:"to"`
	Dist float64 `json:"dist"`
}

var BaseDate = time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC)

func ParseRealTime(str string) int64 {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err.Error())
	}

	real, err := time.ParseInLocation("20060102150405", str, loc)
	if err != nil {
		log.Fatal("Time parse error: ", err)
	}

	return real.Unix() - BaseDate.Unix()
}

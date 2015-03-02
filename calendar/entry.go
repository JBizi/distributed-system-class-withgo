package main

import(
    "time"
  
)

var entries = make(map[int]Entry)
    index int = iota

type Entry struct {
     ID int
     title string
     starts time.Time
     finishes time.Time
}
func (e Entry) Duration() time.Duration {
     return e.finishes.Sub(e.starts)
}

func Add(e Entry) {
     entries[index] = e
     index++
}

func Count()
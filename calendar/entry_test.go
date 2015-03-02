package main

import (
      "testing"
      "github.com/stretchr/testify/assert"
 )

var (
    hour time.Duration
)

func init() {
  
}

func makeEntry() Entry {
  
}

func TestCreate(t testing.T) {
  hour, _:=time.ParseDuration("1h")
  starts := time.Now()
  finishes := time.Now()
  e := new(Entry)
  e.title = "Test Entry"
  e.starts = time.Now()
  
  e := Entry{
    title: "Test Entry",
    starts: stars,
    finishes: finishes,
  }
  assert.Equal(t, 'Test Entry', e.title)
  assert.Equal(t, stasts, e.starts)
  assert.Equal(t, finishes, e.finishes)
  
}
func TestDuration(t testing.T) {
  hour, _:=time.ParseDuration("1h")
  starts := time.Now()
  finishes := time.Now()
  e := new(Entry)
  e.title = "Test Entry"
  e.starts = time.Now()
  
  e := Entry{
    title: "Test Entry",
    starts: stars,
    finishes: finishes,
}
  
func TestAdd(t testing.T)
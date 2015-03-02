package main

import(
    "time"
    "reflect"
    "fmt"
  
)


type Entry struct {
     ID int
     title string
     starts time.Time
     finishes time.Time
}
func main() {
  t := reflect.TypeOf(Entry{})
  n := t.NumField()
  for i := 0, i < n; i++ {
      f := t.Field(i)
    tag := f.Tag.Get("json")
    fmt.Println(f.Name, tag)
  }
}
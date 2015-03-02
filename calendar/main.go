package main

import (
    "github.com/codegangsta/nigroni"
    "github.com/gorilla/pat"
    "net/http"
    "encoding/json"
)

/*
 * GET /entries
 * GET /entries/:id -- get that particular entry
 * POST /entries -- means creating entry
 * POST /entries/:id
 * DELETE /entries/:id
 */
 
 func main() {
     r := pat.New()
     n := negroni.Classic()
     
     r.Get("/entries", func(res http.ResponseHandler, req http.request) {
     res.Header().Set("content-Type", "application/json; charset")
     })
     
     n.UseHandler(r)
     n.Run( :8080)
 }
package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type User struct {
	Name string 
}

type Response struct {
    Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func message(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user := User{}
    response := Response{}

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
    }

    err = json.Unmarshal(body, &user)
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
    }

	response.Greeting = "Hello , " + user.Name + "!"
	
	uj, _ := json.Marshal(response)	
	rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201)
    fmt.Fprintf(rw, "%s", uj)
}

func main() {
    mux := httprouter.New()
    fmt.Println("Server listening on 8080!")
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", message)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
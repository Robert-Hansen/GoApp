package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"github.com/satori/go.uuid"
)

func NewRouter() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", Index)
	r.GET("/hello/:name", Hello)

	return r
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome !\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	fmt.Fprintf(w, "hej"+u1.String()+", %s!\n", ps.ByName("name"))
}
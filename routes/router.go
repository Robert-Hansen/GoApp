package routes

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/robert-hansen/goapp/controller"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("public"))

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	return router
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	controller.Home(w, r)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	fmt.Fprintf(w, "hej"+u1.String()+", %s!\n", ps.ByName("name"))
}
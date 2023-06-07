package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rabee-inc/push/appengine/src/app"
)

func main() {
	e := &app.Environment{}
	e.Get()

	d := &app.Dependency{}
	d.Inject(e)

	r := chi.NewRouter()
	app.Routing(r, d)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", e.Port), r); err != nil {
		log.Fatal(err)
	}
}

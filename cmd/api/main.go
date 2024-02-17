package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ishanjain1502/leaning-go/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true);
	var r *chi.Mux = chi.newRouter()
	handlers.Handle(r)

	fmt.Println("welcome to go api")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil{
		log.Error(err)
	}
}
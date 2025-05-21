package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PrinceFatmer/booking/pkg/config"
	"github.com/PrinceFatmer/booking/pkg/handlers"
	"github.com/PrinceFatmer/booking/pkg/render"
)

const portNumber = ":8080"

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello,  create world!")
	// 	if err != nil {
	// 		fmt.Println(err)

	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of ytes written: %d", n))
	// })
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	// _ = http.ListenAndServe(":8080", nil)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/film/control", app.filmControl)
	mux.HandleFunc("/film/create", app.createFilm2)
	mux.HandleFunc("/film/createcsv", app.createFilmCSV)
	mux.HandleFunc("/film/findname", app.findByName)
	mux.HandleFunc("/film/findid", app.findByID)
	mux.HandleFunc("/film/delname", app.deleteByName)
	mux.HandleFunc("/film/delid", app.deleteByID)
	mux.HandleFunc("/film/findcont", app.findByContent)

	mux.HandleFunc("/ping", ping)

	return mux
}

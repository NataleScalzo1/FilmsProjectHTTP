package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/natale/codermine/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func (app *application) filmControl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	csvFile, err := os.Open("./web/csv/data.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := csv.NewReader(csvFile)
	reader.Comma = '|'
	var body *models.Film
	var toDelete string
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		body = &models.Film{
			Title:   line[0],
			Content: line[1],
			Expires: line[2],
		}
		toDelete = line[3]
	}
	switch toDelete {

	case "1":
		title := body.Title
		content := body.Content
		expires := body.Expires

		result, err := app.films.Insert(title, content, expires)
		if err != nil {
			app.serverError(w, err)
			w.Write([]byte("Film già esistente"))
			return
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(&result)

		filmJson, _ := json.Marshal(body)
		w.Write(filmJson)
		break

	case "0":
		titlename := body.Title

		_, err := app.films.DeletebyName(titlename)

		if err != nil {
			http.NotFound(w, r)
		}
		break

	default:
		w.Write([]byte("Field non trovata"))
		break

	}

}

func (app *application) createFilmCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	csvFile, err := os.Open("./web/csv/data.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := csv.NewReader(csvFile)
	reader.Comma = '|'
	var body *models.Film
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		body = &models.Film{
			Title:   line[0],
			Content: line[1],
			Expires: line[2],
		}
	}
	title := body.Title
	content := body.Content
	expires := body.Expires

	result, err := app.films.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		w.Write([]byte("Film già esistente"))
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(&result)

	filmJson, _ := json.Marshal(body)
	w.Write(filmJson)

}

func (app *application) createFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := r.URL.Query().Get("titlename")
	content := r.URL.Query().Get("content")
	expires := r.URL.Query().Get("expires")

	result, err := app.films.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		w.Write([]byte("Film già esistente"))
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(&result)

}
func (app *application) findByContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	content := r.URL.Query().Get("content")
	s, err := app.films.GetContent(content)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&s)

}
func (app *application) findByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	titlename := r.URL.Query().Get("titlename")
	s, err := app.films.GetName(titlename)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&s)

}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.films.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&s)

}
func (app *application) deleteByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodDelete {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	titlename := r.URL.Query().Get("titlename")

	_, err := app.films.DeletebyName(titlename)

	if err != nil {
		http.NotFound(w, r)
	}

}

func (app *application) deleteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodDelete {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	_, err = app.films.Delete(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

}
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) createFilm2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	out := make([]byte, 1024)

	//
	bodyLen, err := r.Body.Read(out)

	if err != io.EOF {
		fmt.Println(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	var k *models.Film

	err = json.Unmarshal(out[:bodyLen], &k)

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	type idstruct struct {
		ID string `json:"id"`
	}

	id, err := app.films.Insert(k.Title, k.Content, k.Expires)

	if err != nil {
		w.Write([]byte(err.Error() + ".\n" + "Film già esistente."))
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&idstruct{ID: strconv.Itoa(id)})

}

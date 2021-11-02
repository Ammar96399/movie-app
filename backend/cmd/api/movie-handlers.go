package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.erroJson(w, err)
		return
	}

	app.logger.Println("id is", id)

	movie, err := app.models.DB.Get(id)

	/*	movie := models.Movie{
			ID:          id,
			Title:       "Some movie",
			Description: "Some description",
			Year:        2020,
			ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
			Runtime:     100,
			Rating:      5,
			MPAARating:  "PG-14",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	*/
	err = app.writeJson(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}

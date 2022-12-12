package App

import (
	"encoding/json"
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/models"
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/store"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/users"

type App struct {
	store *store.Store
}

func New(s *store.Store) *App {
	return &App{store: s}
}

func (app *App) LoadPosts(c echo.Context) error {
	var users []models.User

	r, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(data, &users); err != nil {
		log.Fatalln(err)
	}

	if err := app.store.SaveUser(users); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Ok")
}

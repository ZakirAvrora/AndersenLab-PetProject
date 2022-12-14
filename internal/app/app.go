package App

import (
	"fmt"
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/store"
	"github.com/ZakirAvrora/AndersenLab-PetProject/service/postloader"
	"github.com/ZakirAvrora/AndersenLab-PetProject/utility/e"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

const defaultPostNum = 50

type App struct {
	store *store.Store
}

func New(s *store.Store) *App {
	return &App{store: s}
}

func (app *App) LoadPosts(c echo.Context) error {

	postsNum := c.QueryParam("posts")
	n, err := validatePostQuery(postsNum)
	if err != nil {
		return err
	}

	if err := postloader.LoadUsers(app.store); err != nil {
		return err
	}

	if err := postloader.LoadPosts(app.store, n); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Ok")
}

func validatePostQuery(posts string) (n int, err error) {
	defer func() { err = e.WrapIfErr("invalid query parameter", err) }()

	if strings.TrimSpace(posts) == "" {
		return defaultPostNum, nil
	}

	n, err = strconv.Atoi(posts)
	if err != nil {
		return 0, err
	}

	if n < 1 || n > 100 {
		return 0, fmt.Errorf("number of posts cannbe only between 1 and 100")
	}

	return n, nil
}

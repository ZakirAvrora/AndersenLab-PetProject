package main

import (
	"github.com/ZakirAvrora/AndersenLab-PetProject/config"
	"github.com/ZakirAvrora/AndersenLab-PetProject/db"
	App "github.com/ZakirAvrora/AndersenLab-PetProject/internal/app"
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/store"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	conf := config.NewConfig("../.env")
	Db := db.Init(conf.Database)

	s := store.New(Db)
	app := App.New(s)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/load", app.LoadPosts)

	e.Logger.Fatal(e.Start(":8080"))
}

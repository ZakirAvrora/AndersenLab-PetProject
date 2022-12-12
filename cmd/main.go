package main

import (
	"encoding/json"
	"github.com/ZakirAvrora/AndersenLab-PetProject/config"
	"github.com/ZakirAvrora/AndersenLab-PetProject/db"
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/models"
	Store "github.com/ZakirAvrora/AndersenLab-PetProject/internal/store"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/users"

func main() {
	conf := config.NewConfig("../.env")
	Db := db.Init(conf.Database)

	_ = Store.New(Db)

	var users []models.User
	//var posts []models.Post
	//var comments []models.Comment

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

}

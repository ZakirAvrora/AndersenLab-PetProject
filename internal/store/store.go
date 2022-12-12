package store

import (
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/models"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveUser(user models.User) error {
	query := `INSERT INTO users (name, username, email, phone, street, suite, city, zipcode,
	            lat, lng, website, company_name, catch_phase, BS)
				VALUES (:name, :username, :email, :phone, :website, :address.street, :address.suite, :address.city,
				        :address.zipcode, :address.geo.lat, :address.geo.lng,
				        :company.company_name,:company.catch_phase, :company.BS)`

	if _, err := s.db.NamedExec(query, user); err != nil {
		return err
	}

	return nil
}

func (s *Store) SavePost(post models.Post) error {
	query := `INSERT INTO comments (post_id, name, email, body)
				VALUES (:post_id, :name, :email, :body)`

	if _, err := s.db.NamedExec(query, post); err != nil {
		return err
	}

	return nil
}

func (s *Store) SaveComment(comment models.Comment) error {
	query := `INSERT INTO posts(user_id, title, body)
				VALUES (:user_id, :title, :body)`

	if _, err := s.db.NamedExec(query, comment); err != nil {
		return err
	}

	return nil
}

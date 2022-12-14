package store

import (
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/models"
	"github.com/ZakirAvrora/AndersenLab-PetProject/utility/e"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveUser(users []models.User) (err error) {
	defer func() { err = e.WrapIfErr("cannot save users to store:", err) }()

	query := `INSERT INTO users (name, username, email, phone, street, suite, city, zipcode,
	            lat, lng, website, company_name, catch_phase, BS)
				VALUES (:name, :username, :email, :phone, :website, :address.street, :address.suite, :address.city,
				        :address.zipcode, :address.geo.lat, :address.geo.lng,
				        :company.company_name,:company.catch_phase, :company.BS)`

	if _, err = s.db.NamedExec(query, users); err != nil {
		return err
	}

	return nil
}

func (s *Store) SavePost(post models.Post) (err error) {
	defer func() { err = e.WrapIfErr("cannot save post to store:", err) }()

	query := `INSERT INTO posts(id, user_id, title, body)
				VALUES (:id, :user_id, :title, :body)`

	if _, err = s.db.NamedExec(query, post); err != nil {
		return err
	}

	return nil
}

func (s *Store) SaveComment(comment []models.Comment) (err error) {
	defer func() { err = e.WrapIfErr("cannot save comments to store:", err) }()

	query := `INSERT INTO comments (post_id, name, email, body)
			VALUES (:post_id, :name, :email, :body)`

	if _, err = s.db.NamedExec(query, comment); err != nil {
		return err
	}

	return nil
}

//func (s *Store) SavePostAndComments(post models.Post, comments []models.Comment) error {

//tx, err := s.db.Beginx()
//if err != nil {
//	return err
//}
//defer tx.Rollback()

//queryPost := `INSERT INTO posts(user_id, title, body)
//			VALUES (:user_id, :title, :body)`
//
//if _, err := s.db.NamedExec(queryPost, post); err != nil {
//	return err
//}
//
//queryComments := `INSERT INTO comments (post_id, name, email, body)
//			VALUES (:post_id, :name, :email, :body)`
//
//if _, err := s.db.NamedExec(queryComments, comments); err != nil {
//	return err
//}

//if err := tx.Commit(); err != nil {
//	return err
//}

//	return nil
//}

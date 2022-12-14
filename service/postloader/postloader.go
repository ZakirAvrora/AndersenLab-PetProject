package postloader

import (
	"fmt"
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/models"
	"github.com/ZakirAvrora/AndersenLab-PetProject/internal/store"
	"github.com/ZakirAvrora/AndersenLab-PetProject/service/jsonload"
	"github.com/ZakirAvrora/AndersenLab-PetProject/utility/e"
	"strconv"
	"sync"
)

const urlUsers = "https://jsonplaceholder.typicode.com/users"
const urlPosts = "https://jsonplaceholder.typicode.com/posts"

func LoadUsers(s *store.Store) (err error) {
	defer func() { err = e.WrapIfErr("error in loading user:", err) }()
	var users []models.User

	if err = jsonload.GetJson(urlUsers, &users); err != nil {
		return err
	}

	if err = s.SaveUser(users); err != nil {
		return err
	}

	return nil
}

func LoadPosts(s *store.Store, n int) (err error) {
	defer func() { err = e.WrapIfErr("error in loading posts:", err) }()
	var posts []models.Post

	if err = jsonload.GetJson(urlPosts, &posts); err != nil {
		return err
	}

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			post := posts[j]
			id := strconv.Itoa(post.PostId)

			commentUrl := fmt.Sprintf("%s/%s/comments", urlPosts, id)

			var comments []models.Comment
			if err := jsonload.GetJson(commentUrl, &comments); err != nil {
				return
			}

			if err := s.SavePost(post); err != nil {
				return
			}

			if err := s.SaveComment(comments); err != nil {
				return
			}

		}(i)
	}

	wg.Wait()
	return nil
}

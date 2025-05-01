package handler

import (
	"github.com/oussaka/go-chi-micro/model"
)

type BlogStorage interface {
	List() []*model.Post
	Get(string) *model.Post
	Update(string, model.Post) *model.Post
	Create(post model.Post)
	Delete(string) *model.Post
}

type BlogStore struct {
}

func NewService() BlogStore {
	return BlogStore{}
}

func (b BlogStore) List() []*model.Post {
	return model.Posts
}

func (b BlogStore) Get(id string) *model.Post {
	for _, post := range model.Posts {
		if post.ID == id {
			return post
		}
	}

	return nil
}

func (b BlogStore) Update(id string, postUpdate model.Post) *model.Post {
	for i, post := range model.Posts {
		if post.ID == id {
			model.Posts[i] = &postUpdate
			return post
		}
	}

	return nil
}

func (b BlogStore) Create(post model.Post) {
	model.Posts = append(model.Posts, &post)
}

func (b BlogStore) Delete(id string) *model.Post {
	for i, post := range model.Posts {
		if post.ID == id {
			model.Posts = append(model.Posts[:i], (model.Posts)[i+1:]...)
			return &model.Post{}
		}
	}

	return nil
}

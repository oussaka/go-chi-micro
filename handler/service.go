package handler

import (
	"github.com/oussaka/go-chi-micro/db"
	"github.com/oussaka/go-chi-micro/model"
	log "github.com/sirupsen/logrus"
)

type BlogStorage interface {
	List() []*model.Blogs
	Get(string) (model.Blogs, error)
	Create(post *model.Blogs) (model.BlogData, error)
	Update(id string, bl *model.Blogs) (model.BlogData, error)
	Delete(string) (string, error)
}

type BlogStore struct {
	sqlDB db.SqlClient
}

func NewService(sqlDB db.SqlClient) *BlogStore {
	return &BlogStore{
		sqlDB: sqlDB,
	}
}

func (b *BlogStore) List() []*model.Blogs {
	return nil
}

func (b *BlogStore) Get(id string) (model.Blogs, error) {
	posts, err := b.sqlDB.GetAllBlogs(id)

	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return model.Blogs{}, err
	}

	return posts, nil
}

func (b *BlogStore) Create(post *model.Blogs) (model.BlogData, error) {
	postData, err := b.sqlDB.CreateBlogPost(post)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.BlogData{}, err
	}

	return postData, nil
}

func (b *BlogStore) Update(id string, postUpdate *model.Blogs) (model.BlogData, error) {
	post, err := b.sqlDB.UpdateBlogs(id, postUpdate)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.BlogData{}, err
	}

	return post, nil
}

func (b *BlogStore) Delete(id string) (string, error) {
	res, err := b.sqlDB.DeleteBlog(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return "record is not available in the system", err
	}

	return res, nil
}

package db

import (
	"github.com/oussaka/go-chi-micro/model"
)

type SqlClient interface {
	CreateBlogPost(bl *model.Blogs) (model.BlogData, error)
	GetAllBlogs(string) (model.Blogs, error)
	UpdateBlogs(id string, bl *model.Blogs) (model.BlogData, error)
	DeleteBlog(id string) (string, error)
}

func NewClient(config *Config) SqlClient {
	return &sqlClient{
		config: config,
	}
}

type Config struct {
	DBConnection string
	// DbPool *sql.DB
}

type sqlClient struct {
	config *Config
}

func (c *sqlClient) CreateBlogPost(bl *model.Blogs) (model.BlogData, error) {
	return createBlog(bl)
}

func (c *sqlClient) GetAllBlogs(id string) (model.Blogs, error) {
	return getAllBlogs(id)
}

func (c *sqlClient) UpdateBlogs(id string, bl *model.Blogs) (model.BlogData, error) {
	return updateBlog(id, bl)
}

func (c *sqlClient) DeleteBlog(id string) (string, error) {
	return deleteBlog(id)
}

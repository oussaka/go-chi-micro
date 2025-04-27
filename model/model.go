package model

type ResponseMeta struct {
	AppStatusCode int    `json:"code"`
	Message       string `json:"statusType,omitempty"`
	ErrorDetail   string `json:"errorDetail,omitempty"`
	ErrorMessage  string `json:"errorMessage,omitempty"`
	DevMessage    string `json:"devErrorMessage,omitempty"`
}

type ErrResponse struct {
	HTTPStatusCode int          `json:"-"` // http response status code
	Status         ResponseMeta `json:"status"`
	AppCode        int64        `json:"code,omitempty"` // application-specific error code
}

type Post struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

var posts = []*Post{
	{
		ID:               "1",
		Title:            "7 habits of highly effective people",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "english",
	},
}

type BlogStorage interface {
	List() []*Post
	Get(string) *Post
	Update(string, Post) *Post
	Create(Post)
	Delete(string) *Post
}

type BlogStore struct {
}

func (b BlogStore) List() []*Post {
	return posts
}

func (b BlogStore) Get(id string) *Post {
	for _, post := range posts {
		if post.ID == id {
			return post
		}
	}
	return nil
}

func (b BlogStore) Update(id string, postUpdate Post) *Post {
	for i, post := range posts {
		if post.ID == id {
			posts[i] = &postUpdate
			return post
		}
	}
	return nil
}

func (b BlogStore) Create(post Post) {
	posts = append(posts, &post)
}

func (b BlogStore) Delete(id string) *Post {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], (posts)[i+1:]...)
			return &Post{}
		}
	}

	return nil
}

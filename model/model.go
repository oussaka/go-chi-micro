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

//type Post struct {
//	ID               string `json:"id"`
//	Title            string `json:"title"`
//	Author           string `json:"author"`
//	PublishedDate    string `json:"published_date"`
//	OriginalLanguage string `json:"original_language"`
//}
//
//var Posts = []*Post{
//	{
//		ID:               "1",
//		Title:            "7 habits of highly effective people",
//		Author:           "Stephen Covey",
//		PublishedDate:    "15/08/1989",
//		OriginalLanguage: "english",
//	},
//}

type Blogs struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	BlogName        string `json:"blog_name"`
	BlogDetails     string `json:"blog_details,omitempty"`
	BlogDescription string `json:"blog_description,omitempty"`
}

type BlogData struct {
	Blog    Blogs  `json:"blog"`
	Message string `json:"message"`
}

package model

type Blogs struct {
	ID              int    `json:"id"`
	BlogName        string `json:"blog_name"`
	BlogDetails     string `json:"blog_details,omitempty"`
	BlogDescription string `json:"blog_description,omitempty"`
}

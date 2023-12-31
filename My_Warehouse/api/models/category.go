package models

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Title string `json:"title"`
}

type Category struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCategory struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type CategoryGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type CategoryGetListResponse struct {
	Count      int         `json:"count"`
	Categories []*Category `json:"categories"`
}

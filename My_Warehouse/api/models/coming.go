package models

type ComingPrimaryKey struct {
	Id string `json:"id"`
}

type ComingIdKey struct {
	ComingID string `json:"coming_id"`
}

type CreateComing struct {
	ComingID string `json:"coming_id"`
	BranchID string `json:"branch_id"`
	Status   string `json:"status"`
}

type Coming struct {
	Id        string `json:"id"`
	ComingID  string `json:"coming_id"`
	BranchID  string `json:"branch_id"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateComing struct {
	Id       string `json:"id"`
	ComingID string `json:"coming_id"`
	BranchID string `json:"branch_id"`
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
}

type ComingGetListRequest struct {
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	Search_by_comingID string `json:"search_by_comingID"`
	Search_by_branch   string `json:"search_by_branch"`
}

type ComingGetListResponse struct {
	Count   int       `json:"count"`
	Comings []*Coming `json:"comings"`
}

package request

type SiteRequest struct {
	Title string `json:"title" binding:"required,title" example:"lorem ipsum"`
}

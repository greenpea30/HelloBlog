package dto

type SearchRequest struct {
	Query    string `form:"q" binding:"required,min=1"`
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=50"`
}

type SearchResultItem struct {
	PostID    int64   `json:"post_id"`
	Title     string  `json:"title"`
	Summary   string  `json:"summary"`
	CreatedAt string  `json:"created_at"`
	Score     float64 `json:"score"`
}

type SearchResponse struct {
	Query    string             `json:"query"`
	Items    []SearchResultItem `json:"items"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
}

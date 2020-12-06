package response

import "github.com/gin-gonic/gin"

type PageResponse struct {
	Items      interface{} `json:"items"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}

func PageSuccess(c *gin.Context, data interface{}, total int64, page int) {
	Success(c, &PageResponse{
		Items: data,
		Pagination: &Pagination{
			Total:    total,
			PageSize: 10,
			Current:  page,
		},
	})
}

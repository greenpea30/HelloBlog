package post

import (
	"strconv"

	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func parseIDParam(c *gin.Context) (int64, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, response.NewError(response.CodeInvalid, "invalid id")
	}
	return id, nil
}

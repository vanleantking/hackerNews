package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (idx *IndexController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "ping")
}

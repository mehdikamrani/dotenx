package crud

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *CRUDController) GetPipelines() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId := c.MustGet("accountId").(string)
		pipelines, err := mc.Service.GetPipelines(accountId)
		if err != nil {
			log.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, pipelines)
	}
}

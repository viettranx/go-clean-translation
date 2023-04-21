package httpapi

import (
	"github.com/gin-gonic/gin"
	"go-clean-translation/service"
	"net/http"
)

type apiController struct {
	service service.TranslateUseCase
}

func NewAPIController(s service.TranslateUseCase) apiController {
	return apiController{service: s}
}

func (api apiController) translate() func(c *gin.Context) {
	return func(c *gin.Context) {
		var param struct {
			OriginalText string `json:"original_text"`
			Source       string `json:"source"`
			Destination  string `json:"destination"`
		}

		if err := c.ShouldBindJSON(&param); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := api.service.Translate(c.Request.Context(), param.OriginalText, param.Source, param.Destination)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func (api apiController) history() func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := api.service.FetchHistories(c.Request.Context())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func (api apiController) SetUpRoute(group *gin.RouterGroup) {
	group.POST("/translate", api.translate())
	group.GET("/histories", api.history())
}

package controller

import (
	"pet/pkg/models"
	Utility "pet/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetArticles(c *gin.Context)
	GetArticle(c *gin.Context)
}

func GetArticles(c *gin.Context) {
	limit := 10
	if c.Query("limit") != "" {
		newLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 10
		} else {
			limit = newLimit
		}
	}
	if limit > 50 {
		limit = 50
	}
	articles := make([]models.Article, limit)

	for i := 0; i < limit; i++ {
		articles[i] = Utility.GetRandomArticle()
	}

	c.JSON(200, articles)
}

func GetArticle(c *gin.Context) {
	articleID := c.Param("id")
	randomArticle := Utility.GetRandomArticle()
	randomArticle.ID = articleID
	c.JSON(200, randomArticle)
}

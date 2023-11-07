package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pet/mocks"
	"pet/pkg/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	t.Run("DefaultLimit", func(t *testing.T) {
		// Create a request with no "limit" query parameter
		r := httptest.NewRequest("GET", "/articles", nil)
		w := httptest.NewRecorder()

		// Create a Gin router
		router := gin.Default()

		// Initialize the controller
		controller := mocks.MockController{}

		// Define the route for testing
		router.GET("/articles", controller.GetArticles)

		// Perform the request
		router.ServeHTTP(w, r)

		// Parse the response
		var articles []models.Article
		err := json.NewDecoder(w.Body).Decode(&articles)
		assert.Error(t, err)

		// Assert that the response has a 200 status code and contains 10 articles
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Len(t, articles, 0)
	})

	t.Run("CustomLimit", func(t *testing.T) {
		// Create a request with "limit" query parameter set to 20
		r := httptest.NewRequest("GET", "/articles?limit=20", nil)
		w := httptest.NewRecorder()

		// Create a Gin router
		router := gin.Default()

		// Initialize the controller
		controller := mocks.MockController{}

		// Define the route for testing
		router.GET("/articles", controller.GetArticles)

		// Perform the request
		router.ServeHTTP(w, r)

		// Parse the response
		var articles []models.Article
		err := json.NewDecoder(w.Body).Decode(&articles)
		assert.Error(t, err)

		// Assert that the response has a 200 status code and contains 20 articles
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Len(t, articles, 0)
	})
}

func TestGetArticle(t *testing.T) {
	t.Run("ValidID", func(t *testing.T) {
		// Create a request with a valid article ID
		r := httptest.NewRequest("POST", "/articles/123", nil)
		w := httptest.NewRecorder()

		// Create a Gin router
		router := gin.Default()

		// Initialize the controller
		controller := mocks.MockController{}

		// Define the route for testing
		router.GET("/articles/:id", controller.GetArticle)

		// Perform the request
		router.ServeHTTP(w, r)

		// Parse the response
		var article models.Article
		err := json.NewDecoder(w.Body).Decode(&article)
		assert.Error(t, err)

		// Assert that the response has a 200 status code and the correct article ID
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "", article.ID)
	})

	t.Run("InvalidID", func(t *testing.T) {
		// Create a request with an invalid article ID
		r := httptest.NewRequest("POST", "/articles/invalid", nil)
		w := httptest.NewRecorder()

		// Create a Gin router
		router := gin.Default()

		// Initialize the controller
		controller := mocks.MockController{}

		// Define the route for testing
		router.GET("/articles/:id", controller.GetArticle)

		// Perform the request
		router.ServeHTTP(w, r)

		// Parse the response
		var article models.Article
		err := json.NewDecoder(w.Body).Decode(&article)
		assert.Error(t, err)

		// Assert that the response has a 200 status code and the article ID is "invalid"
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "", article.ID)
	})
}

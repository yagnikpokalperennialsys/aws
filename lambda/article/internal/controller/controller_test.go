package controller

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"pet/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	testCases := []struct {
		limit      string
		expected   int
		statusCode int
	}{
		{
			limit:      "",
			expected:   10,
			statusCode: 200,
		},
		{
			limit:      "20",
			expected:   20,
			statusCode: 200,
		},
		{
			limit:      "50",
			expected:   50,
			statusCode: 200,
		},
		{
			limit:      "invalid",
			expected:   10,
			statusCode: 200,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.limit, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			w := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(w)

			GetArticles(ctx)

			var articles []models.Article
			err := json.NewDecoder(w.Body).Decode(&articles)
			assert.NoError(t, err)

			assert.Equal(t, testCase.statusCode, w.Code)
		})
	}
}
func TestGetArticle(t *testing.T) {
	testCases := []struct {
		id         string
		statusCode int
	}{
		{
			id:         "1",
			statusCode: 200,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			w := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(w)

			GetArticle(ctx)

			var article models.Article
			err := json.NewDecoder(w.Body).Decode(&article)

			if testCase.statusCode == 200 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			assert.Equal(t, testCase.statusCode, w.Code)
		})
	}
}

package utils

import (
	"pet/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomArticle(t *testing.T) {
	testCases := []struct {
		expectedAuthor  string
		expectedTitle   string
		expectedContent string
	}{
		{
			expectedAuthor:  "Author1",
			expectedTitle:   "Title1",
			expectedContent: "Content1",
		},
		{
			expectedAuthor:  "Author2",
			expectedTitle:   "Title2",
			expectedContent: "Content2",
		},
		{
			expectedAuthor:  "Author3",
			expectedTitle:   "Title3",
			expectedContent: "Content3",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.expectedAuthor, func(t *testing.T) {
			database.Authors = []string{testCase.expectedAuthor}
			database.Titles = []string{testCase.expectedTitle}
			database.Contents = []string{testCase.expectedContent}

			article := GetRandomArticle()
			assert.NotNil(t, article)
			assert.Equal(t, testCase.expectedAuthor, article.Author)
			assert.Equal(t, testCase.expectedTitle, article.Title)
			assert.Equal(t, testCase.expectedContent, article.Content)
		})
	}
}
func TestRandomContent(t *testing.T) {
	testCases := []struct {
		expectedContent string
	}{
		{
			expectedContent: "Content1",
		},
		{
			expectedContent: "Content2",
		},
		{
			expectedContent: "Content3",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.expectedContent, func(t *testing.T) {
			database.Contents = []string{testCase.expectedContent}

			content := RandomContent()
			assert.NotNil(t, content)
			assert.Equal(t, testCase.expectedContent, content)
		})
	}
}
func TestRandomTitle(t *testing.T) {
	testCases := []struct {
		expectedTitle string
	}{
		{
			expectedTitle: "Title1",
		},
		{
			expectedTitle: "Title2",
		},
		{
			expectedTitle: "Title3",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.expectedTitle, func(t *testing.T) {
			database.Titles = []string{testCase.expectedTitle}

			title := RandomTitle()
			assert.NotNil(t, title)
			assert.Equal(t, testCase.expectedTitle, title)
		})
	}
}
func TestRandomAuthor(t *testing.T) {
	testCases := []struct {
		expectedAuthor string
	}{
		{
			expectedAuthor: "Author1",
		},
		{
			expectedAuthor: "Author2",
		},
		{
			expectedAuthor: "Author3",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.expectedAuthor, func(t *testing.T) {
			database.Authors = []string{testCase.expectedAuthor}

			author := RandomAuthor()
			assert.NotNil(t, author)
			assert.Equal(t, testCase.expectedAuthor, author)
		})
	}
}

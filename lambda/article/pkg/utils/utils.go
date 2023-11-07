package utils

import (
	"log"
	"math/rand"
	"pet/pkg/database"
	"pet/pkg/models"

	"github.com/google/uuid"
)

type Utility interface {
	GetRandomArticle() models.Article
	RandomAuthor() string
	RandomTitle() string
	RandomContent() string
	Random(min int, max int) int
	GetUUID() string
}

func GetRandomArticle() models.Article {
	article := models.Article{}

	article.ID = GetUUID()
	article.Title = RandomTitle()
	article.Content = RandomContent()
	article.Author = RandomAuthor()

	return article
}

func RandomAuthor() string {
	return database.Authors[Random(0, len(database.Authors))]
}

func RandomTitle() string {
	return database.Titles[Random(0, len(database.Titles))]
}

func RandomContent() string {
	return database.Contents[Random(0, len(database.Contents))]
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GetUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return uuid.String()
}

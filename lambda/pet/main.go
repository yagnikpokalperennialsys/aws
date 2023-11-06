package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

var ginLambda *ginadapter.GinLambda

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()
		r.GET("/pets", getPets)
		r.POST("/pets/:id", getPet)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

func getPets(c *gin.Context) {
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
	pets := make([]Pet, limit)

	for i := 0; i < limit; i++ {
		pets[i] = getRandomPet()
	}

	c.JSON(200, pets)
}

func getPet(c *gin.Context) {
	petID := c.Param("id")
	randomPet := getRandomPet()
	randomPet.ID = petID
	c.JSON(200, randomPet)
}

var breeds = []string{"Afghan Hound", "Beagle", "Bernese Mountain Dog", "Bloodhound", "Dalmatian", "Jack Russell Terrier", "Norwegian Elkhound"}
var names = []string{"Bailey", "Bella", "Max", "Lucy", "Charlie", "Molly", "Buddy", "Daisy", "Rocky", "Maggie", "Jake", "Sophie", "Jack", "Sadie", "Toby", "Chloe", "Cody", "Bailey", "Buster", "Lola", "Duke", "Zoe", "Cooper", "Abby", "Riley", "Ginger", "Harley", "Roxy", "Bear", "Gracie", "Tucker", "Coco", "Murphy", "Sasha", "Lucky", "Lily", "Oliver", "Angel", "Sam", "Princess", "Oscar", "Emma", "Teddy", "Annie", "Winston", "Rosie"}

type Pet struct {
	ID          string    `json:"id"`
	Breed       string    `json:"breed"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

func getRandomPet() Pet {
	pet := Pet{}

	pet.ID = getUUID()
	pet.Breed = randomBreed()
	pet.Name = randomName()

	pet.DateOfBirth = randomDate()

	return pet
}

func randomDate() time.Time {
	now := time.Now()
	start := now.AddDate(-15, 0, 0)
	delta := now.Unix() - start.Unix()

	sec := rand.Int63n(delta) + start.Unix()
	return time.Unix(sec, 0)
}

func randomBreed() string {
	return breeds[random(0, len(breeds))]
}

func randomName() string {
	return names[random(0, len(names))]
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return uuid.String()
}

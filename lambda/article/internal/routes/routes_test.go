package routes

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// Define a table of test cases
	testCases := []struct {
		Name           string
		RequestPath    string
		ExpectedStatus int
	}{
		{
			Name:           "ValidPath",
			RequestPath:    "/articles",
			ExpectedStatus: 200,
		},
		// Add more test cases here
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			// Create a context and request based on the test case
			ctx := context.Background()
			req := events.APIGatewayProxyRequest{
				Path: testCase.RequestPath,
			}

			// Call the Routes function with the test case inputs
			resp, err := Routes(ctx, req)

			// Assert the expected status code and error
			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectedStatus, resp.StatusCode)
		})
	}
}

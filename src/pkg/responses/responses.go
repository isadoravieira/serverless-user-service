package responses

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// returns a JSON response to api gateway events request
func DomainJSON(response events.APIGatewayProxyResponse, statusCode int, datas interface{}) events.APIGatewayProxyResponse {
	response.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	response.StatusCode = statusCode

	body, err := json.Marshal(datas)
	if err != nil {
		log.Fatal("Erro para fazer o marshal", err)
	}

	response.Body = string(body)

	return response
}

// returns an error in JSON format
func DomainError(response events.APIGatewayProxyResponse, statusCode int, err error) {
	DomainJSON(response, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

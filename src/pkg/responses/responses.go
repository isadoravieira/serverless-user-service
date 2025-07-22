package responses

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// returns a JSON response to api gateway events request
func DomainJSON(statusCode int, datas interface{}) events.APIGatewayProxyResponse {
	body, err := json.Marshal(datas)
	if err != nil {
		log.Fatal("Erro para fazer o marshal", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}

// returns an error in JSON format
func DomainError(statusCode int, err error) events.APIGatewayProxyResponse {
	return DomainJSON(statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

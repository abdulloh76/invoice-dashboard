package utils

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func Response(code int, object interface{}) events.APIGatewayProxyResponse {
	marshalled, err := json.Marshal(object)
	if err != nil {
		return ErrResponse(http.StatusInternalServerError, err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(marshalled),
		IsBase64Encoded: false,
	}
}

func ErrResponse(status int, body string) events.APIGatewayProxyResponse {
	message := map[string]string{
		"message": body,
	}

	messageBytes, _ := json.Marshal(&message)

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(messageBytes),
	}
}

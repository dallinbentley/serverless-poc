package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/stripe/stripe-go/v73/account"
	"github.com/stripe/stripe-go/v73"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var stripeKey string = "sk_test_51LcahxAkN9g1VjCMeVq46jFCUrx1jfR10XNor8VIUOGK5GKMgy4yCKzeW5sZrDr2E52IrRsboe4Z3aKIL5vQo4AC00m7qRYAKo"

func ExpressConnect() (*stripe.Account, error) {
	stripe.Key = stripeKey

	params := &stripe.AccountParams{Type: stripe.String(string(stripe.AccountTypeExpress))}

	return account.New(params)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	result, err := ExpressConnect()

	body, err := json.Marshal(map[string]interface{}{
		"message": result,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

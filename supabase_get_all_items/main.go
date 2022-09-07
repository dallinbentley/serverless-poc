package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/supabase/postgrest-go"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call

func SupabaseClient() *postgrest.Client {
	client := postgrest.NewClient("http://localhost:54321/rest/v1", "", nil)
	client.TokenAuth("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6InNlcnZpY2Vfcm9sZSJ9.vI9obAHOGyVVKa3pD--kJlyxp-Z2zV9UUMAhKpNLAcU")

	return client
}

func GetItemsForBusinessLocation(business_location_id string) (string, int64, error) {
	if SupabaseClient().ClientError != nil {
		panic(SupabaseClient().ClientError)
	}

	return SupabaseClient().From("items").Select("id,quantity", "exact", false).Filter("business_location_id", "eq", business_location_id).ExecuteString()
}

func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	result, num, err := GetItemsForBusinessLocation("5f3f8f8f-f8f8-f8f8-f8f8-a0a0a0a0a0a0")
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"result": result,
		"num": num,
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

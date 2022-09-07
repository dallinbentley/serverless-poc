package utils

import (

	"github.com/supabase/postgrest-go"

)

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
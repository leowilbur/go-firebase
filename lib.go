package firebase

import (
	"net/http"

	firebase_api "github.com/leowilbur/go-firebase/api"
	firebase_client "github.com/leowilbur/go-firebase/client"
)

func NewFirebaseClient(hClient *http.Client, device *firebase_api.FirebaseDevice) (*firebase_client.FireBaseClient, error) {
	return firebase_client.NewFirebaseClient(hClient, device)
}

func RandomAppFID() string {
	result, _ := firebase_api.RandomAppFID()
	return result
}

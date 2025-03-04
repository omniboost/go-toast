package toast_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	asperion "github.com/omniboost/go-toast"
)

var (
	client *asperion.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	toastRestaurantExternalID := os.Getenv("TOAST_RESTAURANT_EXTERNAL_ID")
	debug := os.Getenv("DEBUG")

	client = asperion.NewClient(nil)
	client.SetClientID(clientID)
	client.SetClientSecret(clientSecret)
	client.SetToastRestaurantExternalID(toastRestaurantExternalID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		if baseURL, err := url.Parse(baseURLString); err != nil {
			log.Fatal(err)
		} else {
			client.SetBaseURL(*baseURL)
		}
	}

	client.SetDisallowUnknownFields(true)
	os.Exit(m.Run())
}

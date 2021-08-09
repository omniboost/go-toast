package asperion_test

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"testing"

	asperion "github.com/omniboost/go-asperion"
	"golang.org/x/oauth2"
)

var (
	client     *asperion.Client
	businessID int
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	debug := os.Getenv("DEBUG")
	businessID, err = strconv.Atoi(os.Getenv("BUSINESS_ID"))
	if err != nil {
		log.Fatal(err)
	}
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	oauthConfig := asperion.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	// b, _ := json.MarshalIndent(oauthConfig, "", "  ")
	// log.Fatal(string(b))

	token := &oauth2.Token{
		AccessToken: accessToken,
		// If zero, TokenSource implementations will reuse the same
		// token forever and RefreshToken or equivalent
		// mechanisms for that TokenSource will not be used.
		// Expiry: time.Zero(),
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client = asperion.NewClient(httpClient)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}

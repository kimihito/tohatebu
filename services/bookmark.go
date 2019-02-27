package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/kimihito/tohatebu/models"
)

var accessToken = oauth.Credentials{
	Token:  os.Getenv("HATENA_ACCESS_TOKEN"),
	Secret: os.Getenv("HATENA_ACCESS_TOKEN_SECRET"),
}

var oauthClient = NewAuthenticator(os.Getenv("HATENA_CONSUMER_KEY"), os.Getenv("HATENA_CONSUMER_SECRET"), &accessToken)

type Authenticator struct {
	client      oauth.Client
	httpClient  *http.Client
	credentials *oauth.Credentials
}

func NewAuthenticator(consumerKey string, consumerSecret string, tokens *oauth.Credentials) *Authenticator {
	oauthClient := oauth.Client{
		Credentials: oauth.Credentials{
			Token:  consumerKey,
			Secret: consumerSecret,
		},
		TemporaryCredentialRequestURI: "https://www.hatena.ne.jp/oauth/initiate",
		ResourceOwnerAuthorizationURI: "https://www.hatena.ne.jp/oauth/authorize",
		TokenRequestURI:               "https://www.hatena.ne.jp/oauth/token",
	}
	return &Authenticator{
		client:      oauthClient,
		credentials: tokens,
	}
}

func AddBookmark(e *models.Entry) (*models.Bookmark, error) {
	values := make(url.Values)
	values.Set("url", e.URL)
	values.Set("post_twitter", "true")

	response, err := oauthClient.client.Post(nil, oauthClient.credentials, "http://api.b.hatena.ne.jp/1/my/bookmark", values)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	b := new(models.Bookmark)
	decodeResponse(response, b)
	return b, err
}

func decodeResponse(resp *http.Response, result interface{}) error {
	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, resultByte)
	}
	return json.NewDecoder(strings.NewReader(string(resultByte))).Decode(result)
}

package keybase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Client = &client{}

type client struct{}

func NewClient() Client {
	return &client{}
}

// GetAvatarURL implements Client
func (c *client) GetAvatarURL(username string) (string, error) {
	if strings.TrimSpace(username) == "" {
		return "", nil
	}

	var response UserNameQueryResponse
	endpoint := fmt.Sprintf("/user/lookup.json?username=%[1]s&fields=basics&fields=pictures", username)
	err := queryKeyBase(endpoint, &response)
	if err != nil {
		return "", fmt.Errorf("error while querying keybase: %s", err)
	}

	// The server responded with an error
	if response.Status.Code != 0 {
		return "", fmt.Errorf("response code not valid: %s", response.Status.ErrDesc)
	}

	// Either the pictures do not exist, or the primary one does not exist, or the URL is empty
	data := response.Object
	if data.Pictures == nil || data.Pictures.Primary == nil || len(data.Pictures.Primary.URL) == 0 {
		return "", nil
	}

	// The picture URL is found
	return data.Pictures.Primary.URL, nil
}

// queryKeyBase queries the Keybase APIs for the given endpoint, and de-serializes
// the response as a JSON object inside the given ptr
func queryKeyBase(endpoint string, ptr interface{}) error {
	resp, err := http.Get("https://keybase.io/_/api/1.0" + endpoint)
	if err != nil {
		return fmt.Errorf("error while querying keybase APIs: %s", err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading response body: %s", err)
	}

	err = json.Unmarshal(bz, &ptr)
	if err != nil {
		return fmt.Errorf("error while unmarshaling response body: %s", err)
	}

	return nil
}

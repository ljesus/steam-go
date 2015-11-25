package steam_go

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type SteamApi struct {
	apiKey		string
}

func NewSteamApi(apiKey string) *SteamApi {

	client := new(SteamApi)
	client.apiKey = apiKey

	return client
}

func (self *SteamApi) performGet(url string, values url.Values) ([]byte, error) {
	values = self.appendClientParams(values)

	url = url + values.Encode()

	r, err := http.Get(url)
	defer r.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Error during get: %s", err.Error)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading stream: %s", err.Error)
	}

	return body, nil
}

func (self *SteamApi) appendClientParams(params url.Values) url.Values {
	params.Set("key", self.apiKey)
	return params
}

/**
Steam Api Methods
 */

func (self *SteamApi) GetPlayerSummaries(steam_ids []string) (response *GetPlayerSummariesResponse, err error) {

	body, err := self.performGet(SteamApiGetPlayerSummaries, url.Values{"steamids": steam_ids})

	if err != nil {
		return nil, err
	}

	response = new(GetPlayerSummariesResponse)
	json.Unmarshal(body, response)

	return response, nil
}
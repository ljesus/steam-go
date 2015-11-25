package steam_go

import (
	"io/ioutil"
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
)

type ClientConfig struct {
	SteamApiKey	 string
}

var CONFIG ClientConfig

func TestGetPlayerSummaries(t *testing.T) {
	api := getSteamApi()
	resp, err := api.GetPlayerSummaries([]string{"76561197960435530"})
	assert.NoError(t, err)
	assert.NotNil(t, resp.Response.Players)
}

func getSteamApi() *SteamApi {
	config := getClientConfig()
	return NewSteamApi(config.SteamApiKey)
}

func getClientConfig() ClientConfig {
	if CONFIG.SteamApiKey == "" {
		r, _ := ioutil.ReadFile("api_test.cfg")
		json.Unmarshal(r, &CONFIG)
	}
	return CONFIG
}
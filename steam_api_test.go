package steam_go

import (
	"io/ioutil"
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
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

func TestGetFriendList(t *testing.T) {
	api := getSteamApi()
	resp, err := api.GetFriendList("76561197960435530")
	fmt.Println(resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp.FriendsList.Friends)
}

func getSteamApi() *SteamApi {
	config := getClientConfig()
	return NewSteamApi(config.SteamApiKey)
}

func getClientConfig() ClientConfig {

	if CONFIG.SteamApiKey == "" {
		CONFIG.SteamApiKey = os.Getenv("STEAM_API_KEY")
		if CONFIG.SteamApiKey == "" {
			r, _ := ioutil.ReadFile("api_test.cfg")
			json.Unmarshal(r, &CONFIG)
		}
	}
	return CONFIG
}

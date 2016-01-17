[![GoDoc](https://godoc.org/github.com/ljesus/steam-go?status.svg)](https://godoc.org/github.com/ljesus/steam-go) [![Build Status](https://travis-ci.org/ljesus/steam-go.svg?branch=master)](https://travis-ci.org/ljesus/steam-go)
# steam-go
Steam API implementation in GoLang

    go get github.com/ljesus/steam-go

Example
-------

```go
import (
	steam "github.com/ljesus/steam-go"
)

steam_api := steam.NewSteamApi(config.Config.SteamApiKey)

if data, err := steam_api.GetPlayerSummaries([]string{steam_id}); err == nil {
	//Do things with data
}
```

Web-Services Implemented
---------

	GetPlayerSummaries (v0002)
	GetFriendList (v0001)

Disclaimer
--------

The API is being actively developed and is nowhere near a release candidate. Use with caution.

# steam-go
Steam API implementation in GoLang

    go get github.com/ljesus/steam-go

Example
-------

```go
steam_api := steam.NewSteamApi(config.Config.SteamApiKey)

if data, err := steam_api.GetPlayerSummaries([]string{steam_id}); err == nil {
	//Do things with data
}
```

Web-Services Implemented
---------

	GetPlayerSummaries (v0002)
	

Disclaimer
--------

The API is being actively developed and is nowhere near a release candidate. Use with caution.

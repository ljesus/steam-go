package steam_go

const (
	SteamApiV1 = "v0001"
	SteamApiV2 = "v0002"
	SteamApiBase = "http://api.steampowered.com/"
	SteamApiUser = SteamApiBase + "ISteamUser/"
)

const (
	SteamApiGetPlayerSummaries = SteamApiUser + "GetPlayerSummaries/" + SteamApiV2 + "/?"
)
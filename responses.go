package steam_go

/**
GetPlayerSummaries
 */

type Player struct {
	SteamId							string		`json:"steamid"`
	CommunityVisibilityState		int			`json:"communityvisibilitystate"`
	ProfileState					int			`json:"profilestate"`
	PersonaName						string		`json:"personaname"`
	LastLogoff						int			`json:"lastlogoff"`
	ProfileUrl						string		`json:"profileurl"`
	Avatar							string		`json:"avatar"`
	AvatarMedium					string		`json:"avatarmedium"`
	AvatarFull						string		`json:"avatarfull"`
	PersonaState					int			`json:"personastate"`
	RealName						string		`json:"realname"`
	PrimaryClanId					int			`json:"primaryclanid"`
	TimeCreated						int			`json:"timecreated"`
	PersonaStateFlags				int			`json:"personastateflags"`
	LocCountryCode					string		`json:"loccountrycode"`
	LocStateCode					string		`json:"locstatecode"`
	LocCityId						int			`json:"loccityid"`
}

type GetPlayerSummariesResponse struct {
	Response	GetPlayerSummariesResponseBody
}

type GetPlayerSummariesResponseBody struct {
	Players		[]Player
}

/**
GetFriendList
 */

type Friend struct {
	SteamId							string		`json:"steamid"`
	Relationship					string		`json:"relationship"`
	FriendSince						int64		`json:"friend_since"`
}

type GetFriendListResponse struct {
	FriendsList		GetFriendListResponseBody	`json:"friendslist"`
}

type GetFriendListResponseBody struct {
	Friends		[]Friend
}
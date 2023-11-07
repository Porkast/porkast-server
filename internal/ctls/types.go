package ctls

var (
	Ctl = controller{}
)

type controller struct {
}

type LoginReqData struct {
	UserId   string `json:"userId" v:"required"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RegisterReqData struct {
	Id             string `json:"id"`
	Nickname       string `json:"nickname"`
	Password       string `json:"password"`
	PasswordVerify string `json:"passwordVerify"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Avatar         string `json:"avatar"`
}

type SyncUserInfoReqData struct {
	UserId         string `json:"userId" v:"required"`
	Nickname       string `json:"nickname"`
	Password       string `json:"password"`
	PasswordVerify string `json:"passwordVerify"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Avatar         string `json:"avatar"`
}

type GetUserInfoReqData struct {
	UserId string `json:"userId" v:"required"`
}

type AddListenLaterReqData struct {
	UserId    string `json:"userId" v:"required"`
	ChannelId string `json:"channelId" v:"required"`
	ItemId    string `json:"itemId" v:"required"`
	Source    string `json:"source"`
}

type GetListenLaterReqData struct {
	UserId    string `json:"userId" v:"required"`
	ChannelId string `json:"channelId" v:"required"`
	ItemId    string `json:"itemId" v:"required"`
}

type GetListenLaterListReqData struct {
	UserId string `json:"userId" v:"required"`
	Offset int    `json:"offset" v:"required"`
	Limit  int    `json:"limit" v:"required"`
}

type SubSearchKeywordReqData struct {
	UserId        string `json:"userId" v:"required"`
	Keyword       string `json:"keyword" v:"required"`
	Country       string `json:"country" v:"required"`
	Source        string `json:"source" v:"required"`
	ExcludeFeedId string `json:"excludeFeedId"`
	SortByDate    int    `json:"sortByDate"`
}

type GetSubKeywordListReqData struct {
	UserId  string `json:"userId" v:"required"`
	Keyword string `json:"keyword" v:"required"`
}

type CreatePlaylistReqData struct {
	UserId      string `json:"userId" v:"required"`
	Name        string `json:"name" v:"required"`
	Description string `json:"description"`
}

type AddFeedItemToPlaylistReqData struct {
	PlaylistId string `json:"playlistId" v:"required"`
	ChannelId  string `json:"channelId" v:"required"`
	Guid       string `json:"guid" v:"required"`
	Source     string `json:"source"`
}

type SubscribePlaylistReqData struct {
	UserId        string `json:"userId" v:"required"`
	PlaylistId    string `json:"playlistId" v:"required"`
	CreatorUserId string `json:"creatorUserId"`
}

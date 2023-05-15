package ctls

var (
	Ctl = controller{}
)

type controller struct {
}

type LoginReqData struct {
	Password string `json:"password" v:"required"`
	Email    string `json:"email" v:"required-without:Phone"`
	Phone    string `json:"phone" v:"required-without:Email"`
}

type RegisterReqData struct {
	Nickname       string `json:"nickname" v:"required"`
	Password       string `json:"password" v:"required|length:6,20"`
	PasswordVerify string `json:"passwordVerify" v:"required|length:6,20|same:password"`
	Email          string `json:"email" v:"required-without:Phone"`
	Phone          string `json:"phone" v:"required-without:Email"`
}

type AddListenLaterReqData struct {
	UserId    string `json:"userId" v:"required"`
	ChannelId string `json:"channelId" v:"required"`
	ItemId    string `json:"itemId" v:"required"`
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

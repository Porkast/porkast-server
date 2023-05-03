package ctls

var (
	Ctl = controller{}
)

type controller struct {
}

type LoginReqData struct {
	Username string
	Password string `json:"password" v:"required"`
	Email    string `json:"email" v:"required-without:Phone"`
	Phone    string `json:"phone" v:"required-without:Email"`
}

type RegisterReqData struct {
	Nickname string `json:"nickname" v:"required"`
	Password string `json:"password" v:"required"`
	Email    string `json:"email" v:"required-without:Phone"`
	Phone    string `json:"phone" v:"required-without:Email"`
}

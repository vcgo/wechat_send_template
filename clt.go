package wechat_send_template

import "gopkg.in/chanxuehong/wechat.v2/mp/core"

var {
	Ats core.DefaultAccessTokenServer
	Clt core.Client
}

// SecretInitClt 有wxAppSecret时
func SecretInitClt(wxAppID string, wxAppSecret string) {
	Ats = core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	Clt = core.NewClient(Ats, nil)
}

// TokenInitClt 没有wxAppSecret时只能通过access_token
func TokenInitClt(token string) {
	Ats = core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	Ats.RefreshToken(token)
	Clt = core.NewClient(Ats, nil)
}

package model

type CloudApi struct {
	Addr string `json:"addr"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type GithubKey struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

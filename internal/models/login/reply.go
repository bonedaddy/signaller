package login

import (
	common "github.com/nxshock/signaller/internal/models/common"
)

type LoginReply struct {
	AccessToken string `json:"access_token"`
	HomeServer  string `json:"home_server,omitempty"` // TODO: check api
	UserID      string `json:"user_id"`
}

type Flow struct {
	Type common.AuthenticationType `json:"type"`
}
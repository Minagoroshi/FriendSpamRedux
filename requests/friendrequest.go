package requests

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
	"time"
)

type FriendRequestResponse struct {
	ID             string    `json:"id"`
	SenderUserID   string    `json:"senderUserId"`
	SenderUsername string    `json:"senderUsername"`
	Type           string    `json:"type"`
	Message        string    `json:"message"`
	Details        string    `json:"details"`
	Seen           bool      `json:"seen"`
	CreatedAt      time.Time `json:"created_at"`
}

//FriendRequest is a http Request to send a friend request to another user, and takes a userID as a parameter
//also takes an auth token as a cookie
//returns a FriendRequestResponse
func FriendRequest(userid string, cookies []*http.Cookie, proxy string, useProxy bool) *resty.Response {

	client := resty.New()

	if useProxy {
		if !strings.Contains(proxy, "http") || !strings.Contains(proxy, "https") {
			proxy = "http://" + proxy
		}
		client.SetProxy(proxy)
	}

	resp, _ := client.R().
		SetCookies(cookies).
		SetHeader("Content-Type", "application/json").
		SetResult(&FriendRequestResponse{}).
		Post("https://api.vrchat.cloud/api/1/user/" + userid + "/friendRequest/")

	return resp
}

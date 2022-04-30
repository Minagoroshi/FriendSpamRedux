package requests

import (
	"encoding/base64"
	"errors"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

var LoginFailed = errors.New("account login failed")

//LoginResponse is the response from the login request
type LoginResponse struct {
	ID                 string   `json:"id"`
	Username           string   `json:"username"`
	DisplayName        string   `json:"displayName"`
	UserIcon           string   `json:"userIcon"`
	Bio                string   `json:"bio"`
	BioLinks           []string `json:"bioLinks"`
	ProfilePicOverride string   `json:"profilePicOverride"`
	StatusDescription  string   `json:"statusDescription"`
	PastDisplayNames   []struct {
		DisplayName string    `json:"displayName"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"pastDisplayNames"`
	HasEmail                       bool     `json:"hasEmail"`
	HasPendingEmail                bool     `json:"hasPendingEmail"`
	ObfuscatedEmail                string   `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string   `json:"obfuscatedPendingEmail"`
	EmailVerified                  bool     `json:"emailVerified"`
	HasBirthday                    bool     `json:"hasBirthday"`
	Unsubscribe                    bool     `json:"unsubscribe"`
	StatusHistory                  []string `json:"statusHistory"`
	StatusFirstTime                bool     `json:"statusFirstTime"`
	Friends                        []string `json:"friends"`
	FriendGroupNames               []string `json:"friendGroupNames"`
	CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatar                  string   `json:"currentAvatar"`
	CurrentAvatarAssetURL          string   `json:"currentAvatarAssetUrl"`
	AcceptedTOSVersion             float64  `json:"acceptedTOSVersion"`
	SteamID                        string   `json:"steamId"`
	SteamDetails                   struct {
	} `json:"steamDetails"`
	OculusID              string    `json:"oculusId"`
	HasLoggedInFromClient bool      `json:"hasLoggedInFromClient"`
	HomeLocation          string    `json:"homeLocation"`
	TwoFactorAuthEnabled  bool      `json:"twoFactorAuthEnabled"`
	State                 string    `json:"state"`
	Tags                  []string  `json:"tags"`
	DeveloperType         string    `json:"developerType"`
	LastLogin             time.Time `json:"last_login"`
	LastPlatform          string    `json:"last_platform"`
	AllowAvatarCopying    bool      `json:"allowAvatarCopying"`
	Status                string    `json:"status"`
	DateJoined            string    `json:"date_joined"`
	IsFriend              bool      `json:"isFriend"`
	FriendKey             string    `json:"friendKey"`
	FallbackAvatar        string    `json:"fallbackAvatar"`
	AccountDeletionDate   string    `json:"accountDeletionDate"`
	OnlineFriends         []string  `json:"onlineFriends"`
	ActiveFriends         []string  `json:"activeFriends"`
	OfflineFriends        []string  `json:"offlineFriends"`
}

func Login(account string, proxy string, useProxy bool) (*resty.Response, error) {

	client := resty.New()

	if useProxy {
		if !strings.Contains(proxy, "http") || !strings.Contains(proxy, "https") {
			proxy = "http://" + proxy
		}
		client.SetProxy(proxy)
	}

	//encode the account into base64
	account = base64.StdEncoding.EncodeToString([]byte(account))

	resp, err := client.R().
		SetHeader("Authorization", "Basic "+account).
		SetResult(&LoginResponse{}).
		Get("https://api.vrchat.cloud/api/1/auth/user")

	if !strings.Contains(string(resp.Body()), "currentAvatar") {
		err = LoginFailed
	}

	return resp, err
}

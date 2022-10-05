package daily

import (
	"net/http"
)

type (
	TokenOptions struct {
		Limit         int    `url:"limit,omitempty"`
		EndingAfter   string `url:"ending_after,omitempty"`
		StartingAfter string `url:"starting_after,omitempty"`
	}

	TokenAuthResponse struct {
		Token string `json:"token"` // "d61cd7b2-a273-42b4-89bd-be763fd562c1",
	}

	TokenResponse struct {
		NotBefore         int    `json:"nbf,omitempty"`
		Expires           int    `json:"exp,omitempty"`
		RoomName          string `json:"room_name,omitempty"`
		IsOwner           bool   `json:"is_owner,omitempty"`
		EnableScreenShare bool   `json:"enable_screeshare,omitempty"`
		EnableRecording   bool   `json:"enable_recording,omitempty"`
	}

	TokenPost struct {
		Properties *TokenProperties `json:"properties,omitempty"`
	}

	TokenProperties struct {
		RoomName        string `json:"room_name,omitempty"`
		IsOwner         bool   `json:"is_owner,omitempty"`
		NotBefore       int    `json:"nbf,omitempty"`
		Expires         int    `json:"exp,omitempty"`
		MaxParticipants int    `json:"max_participants,omitempty"`
		// TODO: Model the rest of values from https://docs.daily.co/reference/rest-api/meeting-tokens/create-meeting-token
	}
)

func (c *Client) MeetingTokenGetByID(token string) (TokenAuthResponse, error) {
	out := TokenAuthResponse{}
	err := c.makeRequest(http.MethodGet, "/meeting-tokens/"+token, nil, &out)
	return out, err
}

func (c *Client) MeetingtokenPost(bod TokenPost) (TokenResponse, error) {
	out := TokenResponse{}
	err := c.makeRequest(http.MethodPost, "/meeting-tokens", bod, &out)
	return out, err
}

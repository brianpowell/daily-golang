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
	
	TokenPermission struct {
		HasPresence bool `json:"hasPresence,omitempty"`
		CanSend bool `json:"canSend,omitempty"`
	}
	TokenProperties struct {
		RoomName        string `json:"room_name,omitempty"`
		IsOwner         bool   `json:"is_owner,omitempty"`
		NotBefore       int    `json:"nbf,omitempty"`
		Expires         int    `json:"exp,omitempty"`
		MaxParticipants int    `json:"max_participants,omitempty"`
		EnableRecording bool   `json:"enable_recording,omitempty"`
		EnableScreenShare bool `json:"enable_screenshare,omitempty"`
		EjectAtTokenExp bool `json:"eject_at_token_exp,omitempty"`
		EjectAfterElapsed int `json:"eject_after_elapsed,omitempty"`
		UserID string `json:"user_id,omitempty"`
		UserName string `json:"user_name,omitempty"`
		StartVideoOff bool `json:"start_video_off,omitempty"`
		StartAudioOff bool `json:"start_audio_off,omitempty"`
		EnableTerseLogging bool `json:"enable_terse_logging,omitempty"`
		CloseTabOnExit bool `json:"close_tab_on_exit,omitempty"`
		RedirectOnMeetingExit string `json:"redirect_on_meeting_exit,omitempty"`
		Lang string `json:"lang,omitempty"`
		Permissions *TokenPermission `json:"permissions,omitempty"`

		// TODO: Model the rest of values from https://docs.daily.co/reference/rest-api/meeting-tokens/create-meeting-token
	}
)

func (c *Client) MeetingTokenGetByID(token string) (TokenResponse, error) {
	out := TokenResponse{}
	err := c.makeRequest(http.MethodGet, "/meeting-tokens/"+token, nil, &out)
	return out, err
}

func (c *Client) MeetingTokenPost(bod *TokenProperties) (TokenAuthResponse, error) {
	out := TokenAuthResponse{}
	temp := TokenPost{
		Properties: bod,
	}
	err := c.makeRequest(http.MethodPost, "/meeting-tokens", temp, &out)
	return out, err
}

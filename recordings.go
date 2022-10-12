package daily

type (
	Recording struct {
		ID string `json:"id"`
		RoomName string `json:"room_name"`
		StartTs int `json:"start_ts"`
		Status string `json:"status"`
		MaxParticipants int `json:"max_participants"`
		Duration int `json:"duration"`
		Tracks []string `json:"tracks"`
		ShareToken string `json:"share_token"`
	}

	RecordingList struct {
		TotalCount int `json:"total_count"`
		Data []Recording `json:"data"`
	}

	// Query Options
	RecordingOptions struct {
		Room string `json:"room_name,omitempty"`
		EndingBefore string `json:"ending_before,omitempty"`
		StartingAfter string `json:"starting_after,omitempty"`
		Limit int32 `json:"limit,omitempty"`
	}
// Model these https://docs.daily.co/reference/rest-api/recordings
)

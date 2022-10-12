package daily

type (
	// Meeting object
	Participant struct {
		UserID string `json:"user_id"`
		ParticipantID string `json:"participant_id"`
		UserName string `json:"user_name"`
		JoinTime int `json:"join_time"`
		Duration int `json:"duration"`
		Room string `json:"room,omitempty"`
	}
	Meeting struct {
		ID string `json:"id"`
		RoomName string `json:"room"`
		Duration int `json:"duration"`
		StartTime int `json:"start_time,omitempty"`
		Ongoing bool `json:"ongoing"`
		MaxParticipants int `json:"max_participants"`
		Participants []Participant `json:"participants"`
	}
	// Get meeting information
	MeetingList struct {
		TotalCount int `json:"total_count"`
		Data []Meeting `json:"data"`
	}

	// Query Options
	MeetingOptions struct {
		Room string `json:"room,omitempty"`
		TimeFrameStart int `json:"timeframe_start,omitempty"`
		TimeFrameEnd int `json:"timeframe_end,omitempty"`
		Limit int `json:"limit,omitempty"`
		StartingAfter string `json:"starting_after,omitempty"`
		EndingBefore string `json:"ending_before,omitempty"`
	}
// Model these https://docs.daily.co/reference/rest-api/meetings/get-meeting-information
)

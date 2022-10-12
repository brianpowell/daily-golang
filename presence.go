package daily

type (
	// Presence object is a key value pair, with the value being an array of participants
	/*
		{
			"cool-room": [
				Participants
			]
			"another-room": [
				Participants
			]
		}
	*/
	Presence struct {
		Rooms map[string][]Participant `json:""`
	}

// Model these https://docs.daily.co/reference/rest-api/presence
)

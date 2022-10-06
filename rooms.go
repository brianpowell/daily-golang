package daily

import (
	"fmt"
	"net/http"
)

type (
	RoomOptions struct {
		Limit         int    `url:"limit,omitempty"`
		EndingAfter   string `url:"ending_after,omitempty"`
		StartingAfter string `url:"starting_after,omitempty"`
	}

	Room struct {
		ID          string         `json:"id"`          // "d61cd7b2-a273-42b4-89bd-be763fd562c1",
		Name        string         `json:"name"`        // "hello",
		ApirCreated bool           `json:"api_created"` // false,
		Privacy     string         `json:"privacy"`     // "public",
		URL         string         `json:"url"`         // "https://your-domain.daily.co/hello",
		CreatedAt   string         `json:"created_at"`  // "2019-01-25T23:49:42.000Z",
		Config      RoomProperties `json:"config"`      // {}
	}

	RoomPost struct {
		Name string `json:"name,omitempty"`
		RoomPut
	}

	RoomPut struct {
		Privacy    string          `json:"privacy,omitempty"`
		Properties *RoomProperties `json:"properties,omitempty"`
	}

	RoomProperties struct {
		NotBefore                int    `json:"nbf,omitempty"`
		Expires                  int    `json:"exp,omitempty"`
		MaxParticipants          int    `json:"max_participants,omitempty"`
		EnablePeopleUI           bool   `json:"enable_people_ui,omitempty"`
		EnablePipUI              bool   `json:"enable_pip_ui,omitempty"`
		EnableEmojiReactions     bool   `json:"enable_emoji_reactions,omitempty"`
		EnableHandRaising        bool   `json:"enable_hand_raising,omitempty"`
		EnablePreJoinUI          bool   `json:"enable_prejoin_ui,omitempty"`
		EnableNetworkUI          bool   `json:"enable_network_ui,omitempty"`
		EnableKnocking           bool   `json:"enable_knocking,omitempty"`
		EnableScreenShare        bool   `json:"enable_screenshare,omitempty"`
		EnableVideoProcessingUI  bool   `json:"enable_video_processing_ui,omitempty"`
		EnableChat               bool   `json:"enable_chat,omitempty"`
		StartVideoOff            bool   `json:"start_video_off,omitempty"`
		StartAudioOff            bool   `json:"start_audio_off,omitempty"`
		OwnerOnlyBroadcast       bool   `json:"owner_only_broadcast,omitempty"`
		EnableRecording          string `json:"enable_recording,omitempty"`
		EjectAtRoomExpire        bool   `json:"eject_at_room_exp,omitempty"`
		EjectAfterElapsed        int    `json:"eject_after_elapsed,omitempty"`
		EnableAdvancedChat       bool   `json:"enable_advanced_chat,omitempty"`
		EnableHiddenParticipants bool   `json:"enable_hidden_participants,omitempty"`
		EnableMeshSFU            bool   `json:"enable_mesh_sfu,omitempty"`
		// TODO: Model the rest of values from https://docs.daily.co/reference/rest-api/rooms/create-room
	}
)

func (c *Client) RoomsGet(opt RoomOptions) ([]Room, error) {
	out := []Room{}
	err := c.makeRequest(http.MethodGet, "/rooms", opt, &out)
	return out, err
}

func (c *Client) RoomsGetByID(name string) (Room, error) {
	out := Room{}
	err := c.makeRequest(http.MethodGet, "/rooms/"+name, nil, &out)
	return out, err
}

func (c *Client) RoomsPost(bod RoomPost) (Room, error) {
	out := Room{}
	fmt.Println("Body", bod)
	err := c.makeBodyRequest(http.MethodPost, "/rooms", &bod, &out)
	fmt.Println("Body Out", out)
	return out, err
}

func (c *Client) RoomsDelete(name string) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	err := c.makeRequest(http.MethodDelete, "/rooms/"+name, nil, &out)
	return out, err
}

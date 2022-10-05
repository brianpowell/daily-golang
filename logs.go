package daily

import (
	"net/http"
	"time"
)

type (
	LogOptions struct {
		IncludeLogs    bool   `url:"includeLogs,omitempty"`
		IncludeMetrics bool   `url:"includeMetrics,omitempty"`
		UserSessionId  string `url:"userSessionId,omitempty"`
		MtgSessionId   string `url:"mtgSessionId,omitempty"`
		LogLevel       string `url:"logLevel,omitempty"`
		Order          string `url:"order,omitempty"`
		StartTime      int    `url:"startTime,omitempty"`
		EndTime        int    `url:"endTime,omitempty"`
		Limit          int    `url:"limit,omitempty"`
		Offset         string `url:"offset,omitempty"`
	}

	LogsResponse struct {
		Logs         []Log    `json:"logs"`
		LogsCount    int      `json:"logs_count"`
		Metrics      []Metric `json:"metrics"`
		MetricsCount int      `json:"metrics_count"`
	}

	Log struct {
		Time          time.Time `json:"time"`          // "2020-09-29T18:01:41.791Z",
		ClientTime    time.Time `json:"clientTime"`    // "2020-09-29T18:01:26.529Z",
		Message       string    `json:"message"`       // "participant joined",
		MtgSessionId  string    `json:"mtgSessionId"`  // "7a99abff-0047-4b27-c6c1-49b4ec46f1de",
		UserSessionId string    `json:"userSessionId"` // "4fde3659-71f6-4c28-9a90-e4c3f08ca611",
		PeerId        string    `json:"peerId"`        // "6cfbb8ec-fdbe-49f4-8a7d-39167f5c7745",
		DomainName    string    `json:"domainName"`    // "your-domain",
		Level         int       `json:"level"`         // 1,
		Code          int       `json:"code"`          // 8010
	}

	Metric struct {
		Time                  time.Time              `json:"time"`                  // "2020-09-29T18:01:41.791Z",
		ClientTime            time.Time              `json:"clientTime"`            // "2020-09-29T18:01:26.529Z",
		Message               string                 `json:"message"`               // "participant joined",
		MtgSessionId          string                 `json:"mtgSessionId"`          // "7a99abff-0047-4b27-c6c1-49b4ec46f1de",
		UserSessionId         string                 `json:"userSessionId"`         // "4fde3659-71f6-4c28-9a90-e4c3f08ca611",
		RoomName              string                 `json:"roomName"`              //"logger",
		DomainName            string                 `json:"domainName"`            //"your-domain",
		VideoRecvPacketLoss   int                    `json:"videoRecvPacketLoss"`   //0,
		UserRecvBitsPerSecAvg int                    `json:"userRecvBitsPerSecAvg"` //1038950,
		UserRecvBitsPerSecMax int                    `json:"userRecvBitsPerSecMax"` //1248527,
		UserRecvBitsPerSecMin int                    `json:"userRecvBitsPerSecMin"` //702322,
		UserSentBitsPerSecAvg int                    `json:"userSentBitsPerSecAvg"` //1059720,
		UserSentBitsPerSecMax int                    `json:"userSentBitsPerSecMax"` //1087262,
		UserSentBitsPerSecMin int                    `json:"userSentBitsPerSecMin"` //1039595,
		Metrics               map[string]interface{} `json:"metrics"`
	}
)

func (c *Client) LogsGet(opt LogOptions) (LogsResponse, error) {
	out := LogsResponse{}
	err := c.makeRequest(http.MethodGet, "/logs", opt, &out)
	return out, err
}

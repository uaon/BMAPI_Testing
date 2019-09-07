type Bmsession struct {
	BM_Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Start     time.Time `json:"start"`
			Stop      time.Time `json:"stop"`
			FirstTime bool      `json:"firstTime"`
			Name      string    `json:"name"`
		} `json:"attributes"`
		Relationships struct {
			Server struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"server"`
			Player struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"player"`
		} `json:"relationships"`
	} `json:"data"`
	Included []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			ID        string    `json:"id"`
			Name      string    `json:"name"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
		} `json:"attributes"`
	} `json:"included"`
}
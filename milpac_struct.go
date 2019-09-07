type Milpacs struct {
	Milpac_Data struct {
		Users []struct {
			UserID          int    `json:"user_id"`
			MilpacID        int    `json:"milpac_id"`
			RealName        string `json:"real_name"`
			Username        string `json:"username"`
			Rank            string `json:"rank"`
			RankShorthand   string `json:"rank_shorthand"`
			Status          string `json:"status"`
			PrimaryPosition string `json:"primary_position"`
			Bio             string `json:"bio"`
			JoinDate        string `json:"join_date"`
			PromotionDate   string `json:"promotion_date"`
		} `json:"users"`
	} `json:"data"`
	Type string `json:"type"`
}
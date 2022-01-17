package respdata


type List struct {
	SearchId string `json:"search_id"`
	QueryId string `json:"query_id"`
	SearchResultId string `json:"search_result_id"`
}

type LuckData struct {
	ErrNo int `json:"err_no"`
	IsWinner bool `json:"is_winner"`
	LotteryTime int `json:"lottery_time"`
	ParticipateType int `json:"participate_type"`
	Reward string `json:"reward"`
	Status int `json:"status"`
	WinnerType int `json:"winner_type"`
	Data struct{
		User struct{
			Info struct{
				UserId int `json:"user_id"`
				Name string `json:"name"`
			} `json:"info"`
		} `json:"user"`
	} `json:"data"`
}

type Api struct {
	Message string `json:"message"`
	data struct{} `json:"data"`
}
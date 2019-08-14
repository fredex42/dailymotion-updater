package dm

type DMChannel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DMChannelList struct {
	Page     int64       `json:"page"`
	Limit    int64       `json:"limit"`
	Explicit bool        `json:"explicit"`
	Total    int64       `json:"total"`
	HasMore  bool        `json:"has_more"`
	List     []DMChannel `json:"list"`
}

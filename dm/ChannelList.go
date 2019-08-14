package dm

import "github.com/fredex42/dailymotion-updater/vidispine"

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

/**
converts a list of DMChannel objects into a list of GenericValue key/value objects
*/
func ChannelListToKV(channels *[]DMChannel) *[]vidispine.GenericValue {
	rtn := make([]vidispine.GenericValue, len(*channels))

	for ctr, entry := range *channels {
		rtn[ctr] = vidispine.GenericValue{entry.Id, entry.Name}
	}
	return &rtn
}

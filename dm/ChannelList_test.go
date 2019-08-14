package dm

import (
	"github.com/fredex42/dailymotion_updater/vidispine"
	"testing"
)

func TestChannelListToKV(t *testing.T) {
	sampleList := []DMChannel{
		{
			Id:          "channel1",
			Name:        "Channel 1",
			Description: "First channel",
		},
		{
			Id:          "channel2",
			Name:        "Channel 2",
			Description: "Second channel",
		},
		{
			Id:          "channel3",
			Name:        "Channel 3",
			Description: "Third channel",
		},
	}

	expectedResult := []vidispine.GenericValue{
		{
			Key:   "channel1",
			Value: "Channel 1",
		},
		{
			Key:   "channel2",
			Value: "Channel 2",
		},
		{
			Key:   "channel3",
			Value: "Channel 3",
		},
	}

	result := ChannelListToKV(&sampleList)

	if len(*result) != len(expectedResult) {
		t.Error("Expected ", expectedResult, " got ", *result)
		t.FailNow()
	}

	for ctr := range *result {
		if (*result)[ctr] != expectedResult[ctr] {
			t.Errorf("Element %d mismatched: ", ctr)
		}
	}
}

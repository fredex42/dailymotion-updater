package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fredex42/dailymotion_updater/dm"
	_ "github.com/fredex42/dailymotion_updater/vidispine"
	"log"
)

func main() {
	fmt.Print("dailymotion_updater by Andy Gallagher - https://github.com/fredex42/dailymotion_updater")

	channelList, chanErr := dm.GetChannels()
	if chanErr != nil {
		log.Fatal("Could not get channel data from Daily Motion API: ", chanErr)
	}

	spew.Dump(channelList)
	log.Printf("Got %d channels returned", len(*channelList))
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fredex42/dailymotion-updater/dm"
	"github.com/fredex42/dailymotion-updater/vidispine"
	"log"
	"os"
	"strconv"
)

/**
gets the string value of the given environment variable and returns the given default value
if none was found
*/
func GetWithDefault(key string, dfl string) string {
	rtn := os.Getenv(key)
	if rtn == "" {
		return dfl
	} else {
		return rtn
	}
}

/**
same as GetWithDefault but converts to an int16. Fatal error logged if the value does not convert.
*/
func GetInt16WithDefault(key string, dfl int) int16 {
	result, err := strconv.Atoi(GetWithDefault(key, string(dfl)))
	if err != nil {
		log.Fatalf("Could not convert integer argument %s: %s", key, os.Getenv(key))
	}
	return int16(result)
}

func main() {
	fmt.Print("dailymotion_updater by Andy Gallagher - https://github.com/fredex42/dailymotion-updater\n")

	channelList, chanErr := dm.GetChannels()
	if chanErr != nil {
		log.Fatal("Could not get channel data from Daily Motion API: ", chanErr)
	}

	log.Printf("Got %d channels returned", len(*channelList))

	groupToFind := os.Getenv("MDGROUP_NAME")
	if groupToFind == "" {
		log.Fatal("You need to specify the group to look up by setting the environment variable MDGROUP_NAME")
	}

	fieldToUpdate := os.Getenv("MDFIELD_NAME")
	if fieldToUpdate == "" {
		log.Fatal("You need to specify the field to update by setting the environment variable MDFIELD_NAME")
	}

	vscomm := &vidispine.VidispineCommunicator{
		Protocol: GetWithDefault("VIDISPINE_PROTOCOL", "http"),
		Hostname: GetWithDefault("VIDISPINE_HOST", "localhost"),
		Port:     GetInt16WithDefault("VIDISPINE_PORT", 8080),
		User:     GetWithDefault("VIDISPINE_USER", "admin"),
		Password: os.Getenv("VIDISPINE_PASSWORD"),
	}

	if os.Getenv("VIDISPINE_PASSWORD") == "" {
		log.Printf("WARNING: Attempting to connect to Vidispine as user %s with no password. Expect this to fail.", vscomm.User)
	}

	log.Printf("Connecting to %s://%s:%d as %s...", vscomm.Protocol, vscomm.Hostname, vscomm.Port, vscomm.User)
	//fieldGroup, fgErr := vidispine.GetFieldGroup(vscomm, groupToFind)
	//
	//if fgErr != nil {
	//	log.Fatal("Could not look up fieldgroup: ", fgErr)
	//}

	mdField, getErr := vidispine.GetMDField(vscomm, fieldToUpdate)

	spew.Dump(mdField)
	if getErr != nil {
		log.Fatal("Could not look up metadata field: ", getErr)
	}
	if mdField == nil {
		log.Fatalf("Could not find field %s in group %s", fieldToUpdate, groupToFind)
	}

	fieldData, fdErr := mdField.GetPortalData()
	if fdErr != nil {
		log.Fatal("Could not locate field data: ", fdErr)
	}

	if fieldData == nil {
		log.Fatal("No extradata field found on ", fieldToUpdate)
	}
	spew.Dump(fieldData)

	newValuesPtr := dm.ChannelListToKV(channelList)

	spew.Dump(newValuesPtr)
	isEqual := vidispine.CompareValuesList(newValuesPtr, &fieldData.Values)

	log.Printf("Existing values and new values equal? %t", isEqual)

	if isEqual == true {
		log.Printf("No update needed\n")
	} else {
		fieldData.Values = *newValuesPtr
		marshalled, marshalErr := json.Marshal(&fieldData)

		if marshalErr != nil {
			log.Fatal("Could not convert back to json ", marshalErr)
		}

		setErr := mdField.SetDataKey("extradata", marshalled)
		mdField.StringRestriction.MaxLength = 2048
		if setErr != nil {
			log.Fatal("Could not set data key ", setErr)
		}

		//writeErr := vidispine.SetFieldGroup(vscomm,groupToFind,fieldGroup)
		//
		writeErr := vidispine.SetMDField(vscomm, fieldToUpdate, mdField)
		if writeErr != nil {
			log.Fatal("Could not update server ", writeErr)
		}

	}

	log.Print("Completed.\n")
}

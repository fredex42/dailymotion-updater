package vidispine

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type GenericValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GroupInformation struct {
	Required       bool `json:"required"`
	HideIfNotSet   bool `json:"hideifnotset"`
	Representative bool `json:"representative"`
}

type PortalExtraFieldData struct {
	Name        string `json:"name"`
	Default     string `json:"default"`
	Groups      map[string]GroupInformation
	Description string         `json:"description"`
	Readonly    bool           `json:"readonly"`
	Values      []GenericValue `json:"values"`
	ExternalId  string         `json:"externalid"`
	Type        string         `json:"type"`
	Reusable    bool           `json:"reusable"`
}

func getValuesList(rawData []interface{}) []GenericValue {
	rtn := make([]GenericValue, len(rawData))

	for i, entry := range rawData {
		rec := entry.(map[string]interface{})
		v := GenericValue{rec["key"].(string), rec["value"].(string)}
		rtn[i] = v
	}
	return rtn
}

/**
custom unmarshal method handles the key-named group info by unmarshalling to a map first and then extracting the values
from that
*/
func (p *PortalExtraFieldData) UnmarshalJSON(bytes []byte) error {
	var dictData = make(map[string]interface{})

	err := json.Unmarshal(bytes, &dictData)
	if err != nil {
		return err
	}

	p.Name = dictData["name"].(string)
	p.Default = dictData["default"].(string)
	p.Description = dictData["description"].(string)
	p.Readonly = dictData["readonly"].(bool)
	p.Values = getValuesList(dictData["values"].([]interface{}))
	p.ExternalId = dictData["externalid"].(string)
	p.Type = dictData["type"].(string)
	p.Reusable = dictData["reusable"].(bool)

	var groups = make(map[string]GroupInformation)

	for key, value := range dictData {
		if key != "name" && key != "default" && key != "description" && key != "readonly" && key != "values" && key != "externalid" && key != "type" && key != "reusable" {
			groupInfoDict := value.(map[string]interface{})
			groups[key] = GroupInformation{groupInfoDict["required"].(bool), groupInfoDict["hideifnotset"].(bool), groupInfoDict["representative"].(bool)}
		}
	}
	p.Groups = groups

	return nil
}

/**
custom marshal method builds a k-v map from the structured data and then marshals that in order to deal with key-named
group information
NOTE: you MUST marshal a POINTER to the struct for this to get called.
i.e. json.Marshal(data) won't work but json.Marshal(&data) will.
*/
func (p *PortalExtraFieldData) MarshalJSON() ([]byte, error) {
	var dictData = make(map[string]interface{})

	dictData["name"] = p.Name
	dictData["default"] = p.Default
	dictData["description"] = p.Description
	dictData["readonly"] = p.Readonly
	dictData["values"] = p.Values
	dictData["externalid"] = p.ExternalId
	dictData["type"] = p.Type
	dictData["reusable"] = p.Reusable

	for key, info := range p.Groups {
		dictData[key] = info
	}

	return json.Marshal(dictData)
}

/**
compares the two sets of values. returns a boolean indicating if they are equal (true)
*/
func CompareValuesList(a *[]GenericValue, b *[]GenericValue) bool {
	if len(*a) != len(*b) {
		log.Printf("Lists were different lengths, %d vs %d", len(*a), len(*b))
		return false
	}

	sortedA := make([]GenericValue, len(*a))
	copy(sortedA, *a)
	sort.Slice(sortedA, func(i, j int) bool {
		return sortedA[i].Key < sortedA[j].Key
	})

	sortedB := make([]GenericValue, len(*b))
	copy(sortedB, *b)

	sort.Slice(sortedB, func(i, j int) bool {
		return sortedB[i].Key < sortedB[j].Key
	})

	log.Print(sortedA)
	log.Print(sortedB)

	//since we are now sorted, if the two are non-equal they must also be in the same order.
	//therefore, for any given ctr if a[ctr] != b[ctr] there is a difference in the list
	for ctr := range sortedA {
		fmt.Printf("Compare %s to %s\n", sortedA[ctr].toString(), sortedB[ctr].toString())
		if sortedA[ctr].Key != sortedB[ctr].Key || sortedA[ctr].Value != sortedB[ctr].Value {
			log.Printf("Got difference at %d: %s -> %s\n", ctr, sortedA[ctr].toString(), sortedB[ctr].toString())
			return false
		}
	}
	return true
}

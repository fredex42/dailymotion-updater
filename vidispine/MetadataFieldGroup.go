package vidispine

import "encoding/json"

type GenericData struct {
	Key   string `xml:"key",json:"key"`
	Value string `xml:"value",json:"value"`
}

type StringRestriction struct {
	MinLength int64 `xml:"minLength"`
	MaxLength int64 `xml:"maxLength"`
}

type Schema struct {
	Min  int64  `xml:"min,attr"`
	Max  int64  `xml:"max,attr"`
	Name string `xml:"name,attr"`
}

type MetadataField struct {
	Name              string            `xml:"name"`
	Type              string            `xml:"type"`
	Schema            Schema            `xml:"schema"`
	Data              []GenericData     `xml:"data"`
	StringRestriction StringRestriction `xml:"stringRestriction"`
	Origin            string            `xml:"origin"`
}

type MetadataFieldGroup struct {
	Name   string          `xml:"name"`
	Schema Schema          `xml:"schema"`
	Data   []GenericData   `xml:"data"`
	Fields []MetadataField `xml:"field"`
	Origin string          `xml:"origin"`
}

/**
returns the data for the given data key, or nil if the key does not exist.
*/
func (group *MetadataFieldGroup) GetDataKey(key string) (bool, string) {
	for _, entry := range group.Data {
		if entry.Key == key {
			return true, entry.Value
		}
	}
	return false, ""
}

/**
returns a pointer to the named field, or nil if none was found
*/
func (group *MetadataFieldGroup) GetFieldByName(name string) *MetadataField {
	for _, entry := range group.Fields {
		if entry.Name == name {
			return &entry
		}
	}
	return nil
}

/**
returns the data for the given data key, or nil if the key does not exist.
*/
func (field *MetadataField) GetDataKey(key string) (bool, string) {
	for _, entry := range field.Data {
		if entry.Key == key {
			return true, entry.Value
		}
	}
	return false, ""
}

/**
convenience function to locate, unmarshal and return any extra field data
*/
func (field *MetadataField) GetPortalData() (*PortalExtraFieldData, error) {
	found, d := field.GetDataKey("extradata")
	if found == false {
		return nil, nil
	}

	var rtn PortalExtraFieldData
	err := json.Unmarshal([]byte(d), &rtn)
	if err != nil {
		return nil, err
	}

	return &rtn, nil
}

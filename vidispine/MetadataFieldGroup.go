package vidispine

import (
	"encoding/json"
	"errors"
	"fmt"
)

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
	Name              string             `xml:"name"`
	Type              string             `xml:"type"`
	Schema            Schema             `xml:"schema"`
	Data              []GenericData      `xml:"data"`
	StringRestriction *StringRestriction `xml:"stringRestriction"`
	Origin            string             `xml:"origin"`
}

type MetadataFieldDocument struct {
	XmlNS             string             `xml:"xmlns,attr"`
	Name              string             `xml:"name"`
	Type              string             `xml:"type"`
	Schema            Schema             `xml:"schema"`
	Data              []GenericData      `xml:"data"`
	StringRestriction *StringRestriction `xml:"stringRestriction"`
	Origin            string             `xml:"origin"`
}

type MetadataFieldGroup struct {
	XmlNS  string          `xml:"xmlns,attr"`
	Name   string          `xml:"name"`
	Schema Schema          `xml:"schema"`
	Data   []GenericData   `xml:"data"`
	Fields []MetadataField `xml:"field"`
	Origin string          `xml:"origin"`
}

func (doc *MetadataFieldDocument) getMetadataField() MetadataField {
	return MetadataField{doc.Name, doc.Type, doc.Schema, doc.Data, doc.StringRestriction, doc.Origin}
}

func (d *GenericValue) toString() string {
	return fmt.Sprintf("%s: %s", d.Key, d.Value)
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
func (field *MetadataFieldDocument) GetDataKey(key string) (bool, string) {
	for _, entry := range field.Data {
		if entry.Key == key {
			return true, entry.Value
		}
	}
	return false, ""
}

/**
set the data for the given data key. Errors if the key does not exist.
*/
func (field *MetadataFieldDocument) SetDataKey(key string, newValue []byte) error {
	for ctr, entry := range field.Data {
		if entry.Key == key {
			field.Data[ctr].Value = string(newValue)
			return nil
		}
	}
	errMsg := fmt.Sprintf("Data key %s not found", key)
	return errors.New(errMsg)
}

/**
convenience function to locate, unmarshal and return any extra field data
*/
func (field *MetadataFieldDocument) GetPortalData() (*PortalExtraFieldData, error) {
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

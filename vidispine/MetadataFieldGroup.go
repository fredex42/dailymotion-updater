package vidispine

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
func (group *MetadataFieldGroup) getDataKey(key string) (bool, string) {
	for _, entry := range group.Data {
		if entry.Key == key {
			return true, entry.Value
		}
	}
	return false, ""
}

package vidispine

import (
	"encoding/json"
	"testing"
)

var sample_extradata = `{"name": "Daily Motion Category", "default": "", "MasterDailyMotion": {"required": false, "hideifnotset": false, "representative": false}, "description": "", "readonly": false, "values": [{"value": "Animals", "key": "animals"}, {"value": "Auto-Moto", "key": "auto"}, {"value": "Celeb", "key": "people"}, {"value": "Comedy &amp; Entertainment", "key": "fun"}, {"value": "Community &amp; Blogs", "key": "webcam"}, {"value": "Creative", "key": "creation"}, {"value": "Education", "key": "school"}, {"value": "Gaming", "key": "videogames"}, {"value": "Kids", "key": "kids"}, {"value": "Lifestyle &amp; How-to", "key": "lifestyle"}, {"value": "Movies", "key": "shortfilms"}, {"value": "Music", "key": "music"}, {"value": "News", "key": "news"}, {"value": "Sports", "key": "sport"}, {"value": "Tech", "key": "tech"}, {"value": "Travel", "key": "travel"}, {"value": "TV", "key": "tv"}], "externalid": "", "type": "dropdown", "reusable": false}`

func TestLoadExtraData(t *testing.T) {
	var loaded_data PortalExtraFieldData

	sample_bytes := []byte(sample_extradata)
	err := json.Unmarshal(sample_bytes, &loaded_data)

	if err != nil {
		t.Error("Could not unmarshal json: ", err)
	}

	if loaded_data.Name != "Daily Motion Category" {
		t.Error("Got unexpected name ", loaded_data.Name)
	}

	if len(loaded_data.Values) != 17 {
		t.Error("Got ", len(loaded_data.Values), "17")
	}
}

func TestMarshalExtraData(t *testing.T) {
	var loaded_data PortalExtraFieldData

	sample_bytes := []byte(sample_extradata)
	err := json.Unmarshal(sample_bytes, &loaded_data)

	if err != nil {
		t.Error("Could not unmarshal json: ", err)
	}

	out, marshalErr := json.Marshal(&loaded_data)
	if marshalErr != nil {
		t.Error("Could not re-marshal json: ", err)
	}

	expected_output := `{"MasterDailyMotion":{"required":false,"hideifnotset":false,"representative":false},"default":"","description":"","externalid":"","name":"Daily Motion Category","readonly":false,"reusable":false,"type":"dropdown","values":[{"key":"animals","value":"Animals"},{"key":"auto","value":"Auto-Moto"},{"key":"people","value":"Celeb"},{"key":"fun","value":"Comedy \u0026amp; Entertainment"},{"key":"webcam","value":"Community \u0026amp; Blogs"},{"key":"creation","value":"Creative"},{"key":"school","value":"Education"},{"key":"videogames","value":"Gaming"},{"key":"kids","value":"Kids"},{"key":"lifestyle","value":"Lifestyle \u0026amp; How-to"},{"key":"shortfilms","value":"Movies"},{"key":"music","value":"Music"},{"key":"news","value":"News"},{"key":"sport","value":"Sports"},{"key":"tech","value":"Tech"},{"key":"travel","value":"Travel"},{"key":"tv","value":"TV"}]}`
	outString := string(out)
	if outString != expected_output {
		t.Log("Expected :  ", expected_output)
		t.Log("Actual   : ", outString)
		t.Error("Output did not match expected result")
	}
}

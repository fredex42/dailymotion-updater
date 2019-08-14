package vidispine

import (
	"encoding/xml"
	"testing"
)

var sample_data = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<MetadataFieldGroupDocument xmlns="http://xml.vidispine.com/schema/vidispine">
  <name>MasterDailyMotion</name>
  <schema min="0" max="-1" name="MasterDailyMotion"/>
  <data>
    <key>extradata</key>
    <value>{"field_order": ["gnm_master_dailymotion_title", "gnm_master_dailymotion_description", "gnm_master_dailymotion_keywords", "gnm_master_dailymotion_category", "gnm_master_dailymotion_containsadultcontent", "gnm_master_dailymotion_publish", "gnm_master_dailymotion_author", "gnm_master_dailymotion_nomobileaccess", "gnm_master_dailymotion_uploadstatus", "gnm_master_dailymotion_uploadlog", "gnm_master_dailymotion_remove", "gnm_master_dailymotion_holdingimage_16x9", "gnm_master_dailymotion_publication_url", "gnm_master_dailymotion_item_published", "gnm_master_dailymotion_publication_time", "gnm_master_dailymotion_dailymotionurl", "gnm_master_dailymotion_dailymotioncategory", "gnm_master_dailymotion_dailymotionchannel", "gnm_master_dailymotion_owner", "gnm_master_dailymotion_holdingimage", "gnm_master_dailymotion_uploadtype", "gnm_master_dailymotion_status", "gnm_master_dailymotion_escalation_notes"]}</value>
  </data>
  <field>
    <name>gnm_master_dailymotion_escalation_notes</name>
    <type>string</type>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>65535</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotionSyndication": {}, "readonly": false, "type": "textarea", "name": "Escalation Notes"}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_description</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_description"/>
    <type>string</type>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>307200</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "textarea", "name": "Description", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_title</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_title"/>
    <type>string</type>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>512</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"name": "Title", "default": "", "pattern": "", "MasterDailyMotion": {"required": false, "hideifnotset": false, "representative": false}, "description": "", "readonly": false, "externalid": "", "type": "string", "reusable": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_publication_url</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_publication_url"/>
    <type>string</type>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>512</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "string", "name": "Publication URL", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_item_published</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_item_published"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "checkbox", "name": "Live on site?", "readonly": true}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_publish</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_publish"/>
    <type>date</type>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "date", "name": "Publish date", "readonly": false}</value>
    </data>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_keywords</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_keywords"/>
    <type>string-exact</type>
    <index>extend</index>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>6400</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "tags", "name": "Tags", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_owner</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_owner"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"externalid": "customlookup:/metadataform/userlookup/", "MasterDailyMotion": {}, "type": "string", "name": "Owner", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
      <name>gnm_master_dailymotion_dailymotioncategory</name>
      <type>string-exact</type>
      <stringRestriction/>
      <data>
        <key>extradata</key>
        <value>{"name": "Daily Motion Category", "default": "", "MasterDailyMotion": {"required": false, "hideifnotset": false, "representative": false}, "description": "", "readonly": false, "values": [{"value": "Animals", "key": "animals"}, {"value": "Auto-Moto", "key": "auto"}, {"value": "Celeb", "key": "people"}, {"value": "Comedy &amp; Entertainment", "key": "fun"}, {"value": "Community &amp; Blogs", "key": "webcam"}, {"value": "Creative", "key": "creation"}, {"value": "Education", "key": "school"}, {"value": "Gaming", "key": "videogames"}, {"value": "Kids", "key": "kids"}, {"value": "Lifestyle &amp; How-to", "key": "lifestyle"}, {"value": "Movies", "key": "shortfilms"}, {"value": "Music", "key": "music"}, {"value": "News", "key": "news"}, {"value": "Sports", "key": "sport"}, {"value": "Tech", "key": "tech"}, {"value": "Travel", "key": "travel"}, {"value": "TV", "key": "tv"}], "externalid": "", "type": "dropdown", "reusable": false}</value>
      </data>
      <defaultValue/>
      <origin>VX</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_holdingimage</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_holdingimage"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"externalid": "imagepicker:", "MasterDailyMotion": {}, "type": "string", "name": "Holding image", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_uploadtype</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_uploadtype"/>
    <type>string-exact</type>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>64</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"values": [{"value": "New", "key": "New"}, {"value": "Update", "key": "Update"}], "MasterDailyMotion": {}, "type": "dropdown", "name": "Type of upload", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_uploadstatus</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_uploadstatus"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"values": [{"value": "Not Ready", "key": "Not Ready"}, {"value": "Transcode in Progress", "key": "Transcode in Progress"}, {"value": "Upload in Progress", "key": "Upload in Progress"}, {"value": "Upload Succeeded", "key": "Upload Succeeded"}, {"value": "Upload Failed", "key": "Upload Failed"}, {"value": "Ready to Upload", "key": "Ready to Upload"},  {"value": "Do Not Send", "key": "Do Not Send"},{"value": "Escalate", "key": "Escalate"}], "MasterDailyMotion": {}, "type": "dropdown", "name": "Upload status", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_nomobileaccess</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_nomobileaccess"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"values": [{"value": "No mobile access", "key": "no_mobile_access"}], "MasterDailyMotion": {}, "type": "checkbox", "name": "No Mobile Access", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_containsadultcontent</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_containsadultcontent"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"values": [{"value": "Contains adult content", "key": "contains_adult_content"}], "MasterDailyMotion": {}, "type": "checkbox", "name": "Contains adult content", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_status</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_status"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"values": [{"value": "Draft", "key": "Draft"}, {"value": "Unpublished", "key": "Unpublished"}, {"value": "Ready to publish", "key": "Ready to Publish"}, {"value": "Published", "key": "Published"},  {"value": "Do Not Send", "key": "Do Not Send"},{"value": "Escalate", "key": "Escalate"}], "MasterDailyMotion": {}, "type": "dropdown", "name": "Status", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_holdingimage_16x9</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_holdingimage_16x9"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"externalid": "imagepicker:", "MasterDailyMotion": {}, "type": "string", "name": "Holding image 16x9", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_uploadlog</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_uploadlog"/>
    <type>string</type>
    <index>noindex</index>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>40960</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "textarea", "name": "Upload log", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_dailymotionchannel</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_dailymotionchannel"/>
    <type>string</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "dropdown", "choices": {}, "name": "Daily Motion Channel", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_publication_time</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_publication_time"/>
    <type>date</type>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {"hideifnotset": true}, "description": "If this item was published, when was it launched?", "type": "timestamp", "name": "Publication Date and Time", "readonly": true}</value>
    </data>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_author</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_author"/>
    <type>string-exact</type>
    <stringRestriction/>
    <data>
      <key>extradata</key>
      <value>{"externalid": "customlookup:/metadataform/userlookup/", "MasterDailyMotion": {}, "type": "string", "name": "Author", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_remove</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_remove"/>
    <type>date</type>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "date", "name": "Expiry date", "readonly": false}</value>
    </data>
    <origin>KP</origin>
  </field>
  <field>
    <name>gnm_master_dailymotion_dailymotionurl</name>
    <schema min="0" max="-1" name="gnm_master_dailymotion_dailymotionurl"/>
    <type>string</type>
    <stringRestriction>
      <minLength>0</minLength>
      <maxLength>512</maxLength>
    </stringRestriction>
    <data>
      <key>extradata</key>
      <value>{"MasterDailyMotion": {}, "type": "string", "name": "Daily Motion URL", "readonly": false}</value>
    </data>
    <defaultValue/>
    <origin>KP</origin>
  </field>
  <origin>KP</origin>
</MetadataFieldGroupDocument>
`

func TestLoadFieldGroup(t *testing.T) {
	var test MetadataFieldGroup

	sample_data_bytes := []byte(sample_data)

	err := xml.Unmarshal(sample_data_bytes, &test)

	if err != nil {
		t.Error("Could not unmarshal xml: ", err)
	}
	if test.Name != "MasterDailyMotion" {
		t.Error("Got unexpected group name ", test.Name)
	}

	if len(test.Fields) != 22 {
		t.Error("Got unexpected field count ", len(test.Fields), "expected 22")
	}
}

func TestGetDataKeyFound(t *testing.T) {
	var test MetadataFieldGroup

	sample_data_bytes := []byte(sample_data)
	err := xml.Unmarshal(sample_data_bytes, &test)

	if err != nil {
		t.Error("Could not unmarshal xml: ", err)
	}

	found, result := test.GetDataKey("extradata")
	if found != true {
		t.Error("getDataKey didn't find existing data")
	}
	if len(result) < 10 {
		t.Error("getDataKey returned something too short")
	}
}

func TestGetDataKeyAbsent(t *testing.T) {
	var test MetadataFieldGroup

	sample_data_bytes := []byte(sample_data)
	err := xml.Unmarshal(sample_data_bytes, &test)

	if err != nil {
		t.Error("Could not unmarshal xml: ", err)
	}

	found, result := test.GetDataKey("dadsjkhads")
	if found != false {
		t.Error("getDataKey claimed to find non-existing data")
	}
	if len(result) > 0 {
		t.Error("getDataKey returned data when none was expected")
	}
}

func TestGetFieldByName(t *testing.T) {
	var test MetadataFieldGroup

	sample_data_bytes := []byte(sample_data)
	err := xml.Unmarshal(sample_data_bytes, &test)

	if err != nil {
		t.Error("Could not unmarshal xml: ", err)
	}

	result := test.GetFieldByName("gnm_master_dailymotion_title")
	if result == nil {
		t.Error("GetFieldByName found nothing when it should have found gnm_master_dailymotion_title")
	}

	if result.Name != "gnm_master_dailymotion_title" {
		t.Errorf("GetFieldByName returned wrong field, got name %s", result.Name)
	}
}

func TestGetFieldByNameNone(t *testing.T) {
	var test MetadataFieldGroup

	sample_data_bytes := []byte(sample_data)
	err := xml.Unmarshal(sample_data_bytes, &test)

	if err != nil {
		t.Error("Could not unmarshal xml: ", err)
	}

	result := test.GetFieldByName("fsdsdfgasga")
	if result != nil {
		t.Error("GetFieldByName a field when it should not have found anything")
	}
}

/**
structure should marshal back out to XML and convert to string
*/
func TestMarshalData(t *testing.T) {
	var test MetadataFieldGroup

	sample_data_bytes := []byte(sample_data)
	err := xml.Unmarshal(sample_data_bytes, &test)

	if err != nil {
		t.Error("Could not unmarshal xml: ", err)
	}

	_, marshalErr := xml.Marshal(test)

	if marshalErr != nil {
		t.Error("Failed to marshal data structure ", err)
	}

	//outstring := string(out)
	//
	//fmt.Printf("%s", outstring)
}

func TestMetadataFieldDocument_SetDataKey(t *testing.T) {
	test := MetadataFieldDocument{"", "test_field", "string", Schema{}, []GenericData{
		{
			Key:   "testkey",
			Value: "before value",
		},
	}, nil, "VX"}

	err := test.SetDataKey("testkey", []byte("after value"))
	if err != nil {
		t.Error("SetDataKey errored: ", err)
	}

	if test.Data[0].Value != "after value" {
		t.Error("SetDataKey did not set the key. Expected 'after value', got '" + test.Data[0].Value + "'")
	}
}

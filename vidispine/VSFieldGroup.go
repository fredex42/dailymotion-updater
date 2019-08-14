package vidispine

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
)

/**
contact the server and download information about the named metadata field group
*/
func GetFieldGroup(comm *VidispineCommunicator, groupName string) (*MetadataFieldGroup, error) {
	var rtn MetadataFieldGroup
	subpath := fmt.Sprintf("/API/metadata-field/field-group/%s", url.PathEscape(groupName))

	headers := map[string]string{"Accept": "application/xml"}
	empty := map[string]string{}

	byteData, requestErr := comm.MakeRequest("GET", subpath, empty, empty, headers, nil)

	if requestErr != nil {
		return nil, requestErr
	}

	marshalErr := xml.Unmarshal(byteData, &rtn)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return &rtn, nil
}

func GetMDField(comm *VidispineCommunicator, fieldName string) (*MetadataFieldDocument, error) {
	var rtn MetadataFieldDocument
	subpath := fmt.Sprintf("/API/metadata-field/%s", url.PathEscape(fieldName))

	headers := map[string]string{"Accept": "application/xml"}
	queryParams := map[string]string{"data": "all"}
	empty := map[string]string{}

	byteData, requestErr := comm.MakeRequest("GET", subpath, empty, queryParams, headers, nil)

	if requestErr != nil {
		return nil, requestErr
	}

	marshalErr := xml.Unmarshal(byteData, &rtn)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return &rtn, nil
}

/**
contact the server and upload the given data to the named field group
*/
func SetFieldGroup(comm *VidispineCommunicator, groupName string, group *MetadataFieldGroup) error {
	subpath := fmt.Sprintf("/API/metadata-field/field-group/%s", url.PathEscape(groupName))

	headers := map[string]string{"Accept": "application/xml", "Content-Type": "application/xml"}
	empty := map[string]string{}

	marshalledData, marshalErr := xml.Marshal(group)
	if marshalErr != nil {
		return marshalErr
	}

	log.Printf("Going to write %s", string(marshalledData))

	bodyBuffer := bytes.NewReader(marshalledData)

	_, requestErr := comm.MakeRequest("PUT", subpath, empty, empty, headers, bodyBuffer)

	if requestErr != nil {
		return requestErr
	}

	return nil
}

func SetMDField(comm *VidispineCommunicator, fieldName string, field *MetadataFieldDocument) error {
	subpath := fmt.Sprintf("/API/metadata-field/%s", url.PathEscape(fieldName))

	headers := map[string]string{"Accept": "application/xml", "Content-Type": "application/xml"}
	empty := map[string]string{}

	marshalledData, marshalErr := xml.Marshal(field)
	if marshalErr != nil {
		return marshalErr
	}

	log.Printf("Going to write %s", string(marshalledData))

	bodyBuffer := bytes.NewReader(marshalledData)

	returnedContent, requestErr := comm.MakeRequest("PUT", subpath, empty, empty, headers, bodyBuffer)

	if requestErr != nil {
		return requestErr
	}

	returnedString := string(returnedContent)
	log.Printf("Server returned %s", returnedString)
	return nil
}

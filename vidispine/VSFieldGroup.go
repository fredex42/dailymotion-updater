package vidispine

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

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

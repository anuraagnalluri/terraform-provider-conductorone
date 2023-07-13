// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// RequestCatalogManagementServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type RequestCatalogManagementServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _RequestCatalogManagementServiceGetResponseExpanded RequestCatalogManagementServiceGetResponseExpanded

func (c *RequestCatalogManagementServiceGetResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _RequestCatalogManagementServiceGetResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = RequestCatalogManagementServiceGetResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c RequestCatalogManagementServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_RequestCatalogManagementServiceGetResponseExpanded(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

// RequestCatalogManagementServiceGetResponse - The RequestCatalogManagementServiceGetResponse message.
type RequestCatalogManagementServiceGetResponse struct {
	// The RequestCatalogView message.
	RequestCatalogView *RequestCatalogView `json:"requestCatalogView,omitempty"`
	// The expanded field.
	Expanded []RequestCatalogManagementServiceGetResponseExpanded `json:"expanded,omitempty"`
}

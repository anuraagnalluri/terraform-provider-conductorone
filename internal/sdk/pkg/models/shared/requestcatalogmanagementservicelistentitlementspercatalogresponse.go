// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded

func (c *RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded(c))
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

// RequestCatalogManagementServiceListEntitlementsPerCatalogResponse - The RequestCatalogManagementServiceListEntitlementsPerCatalogResponse message.
type RequestCatalogManagementServiceListEntitlementsPerCatalogResponse struct {
	// The expanded field.
	Expanded []RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded `json:"expanded,omitempty"`
	// The list field.
	List []AppEntitlementView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

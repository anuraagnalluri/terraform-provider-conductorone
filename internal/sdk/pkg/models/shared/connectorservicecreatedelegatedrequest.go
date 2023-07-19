// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorServiceCreateDelegatedRequest - The ConnectorServiceCreateDelegatedRequest message.
type ConnectorServiceCreateDelegatedRequest struct {
	// The ConnectorExpandMask message.
	ConnectorExpandMask *ConnectorExpandMask `json:"expandMask,omitempty"`
	// The catalogId field.
	CatalogID *string `json:"catalogId,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The userIds field.
	UserIds []string `json:"userIds,omitempty"`
}

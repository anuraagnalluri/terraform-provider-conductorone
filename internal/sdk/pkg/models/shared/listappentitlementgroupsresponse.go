// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListAppEntitlementGroupsResponse - The ListAppEntitlementGroupsResponse message.
type ListAppEntitlementGroupsResponse struct {
	// The list field.
	List []AppGroup `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The notificationToken field.
	NotificationToken *string `json:"notificationToken,omitempty"`
}

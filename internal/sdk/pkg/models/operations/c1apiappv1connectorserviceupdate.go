// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIAppV1ConnectorServiceUpdateRequest struct {
	ConnectorServiceUpdateRequest *shared.ConnectorServiceUpdateRequest `request:"mediaType=application/json"`
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
	ID                            string                                `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1ConnectorServiceUpdateResponse struct {
	// Successful response
	ConnectorServiceUpdateResponse *shared.ConnectorServiceUpdateResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}

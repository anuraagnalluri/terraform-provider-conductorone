// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const coupaCatalogID = "2DEn2tGMGD2KR6uKAAAWS3VVJV6"

func (r *IntegrationCoupaResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(coupaCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Coupa"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationCoupaResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	coupaDomain := new(string)
	if !r.CoupaDomain.IsUnknown() && !r.CoupaDomain.IsNull() {
		*coupaDomain = r.CoupaDomain.ValueString()
	} else {
		coupaDomain = nil
	}

	oauth2ClientCredGrantClientId := new(string)
	if !r.Oauth2ClientCredGrantClientId.IsUnknown() && !r.Oauth2ClientCredGrantClientId.IsNull() {
		*oauth2ClientCredGrantClientId = r.Oauth2ClientCredGrantClientId.ValueString()
	} else {
		oauth2ClientCredGrantClientId = nil
	}

	oauth2ClientCredGrantClientSecret := new(string)
	if !r.Oauth2ClientCredGrantClientSecret.IsUnknown() && !r.Oauth2ClientCredGrantClientSecret.IsNull() {
		*oauth2ClientCredGrantClientSecret = r.Oauth2ClientCredGrantClientSecret.ValueString()
	} else {
		oauth2ClientCredGrantClientSecret = nil
	}

	configValues := map[string]*string{
		"coupa-domain":                           coupaDomain,
		"oauth2_client_cred_grant_client_id":     oauth2ClientCredGrantClientId,
		"oauth2_client_cred_grant_client_secret": oauth2ClientCredGrantClientSecret,
	}

	configOut := make(map[string]string)
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = *configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}

	out := shared.Connector{
		DisplayName: sdk.String("Coupa"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(coupaCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationCoupaResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationCoupaResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationCoupaResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
	if resp == nil {
		return
	}
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
	}

	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	r.UserIds = nil
	for _, v := range resp.UserIds {
		r.UserIds = append(r.UserIds, types.StringValue(v))
	}

	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["coupa-domain"]; ok {
					r.CoupaDomain = types.StringValue(v.(string))
				}

				if v, ok := values["oauth2_client_cred_grant_client_id"]; ok {
					r.Oauth2ClientCredGrantClientId = types.StringValue(v.(string))
				}

				if v, ok := values["oauth2_client_cred_grant_client_secret"]; ok {
					r.Oauth2ClientCredGrantClientSecret = types.StringValue(v.(string))
				}

			}
		}
	}
}

func (r *IntegrationCoupaResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationCoupaResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
	}
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	r.UserIds = nil
	for _, v := range resp.UserIds {
		r.UserIds = append(r.UserIds, types.StringValue(v))
	}

	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["coupa-domain"]; ok {
					r.CoupaDomain = types.StringValue(v.(string))
				}

				if v, ok := values["oauth2_client_cred_grant_client_id"]; ok {
					r.Oauth2ClientCredGrantClientId = types.StringValue(v.(string))
				}

				if v, ok := values["oauth2_client_cred_grant_client_secret"]; ok {
					r.Oauth2ClientCredGrantClientSecret = types.StringValue(v.(string))
				}

			}
		}
	}
}

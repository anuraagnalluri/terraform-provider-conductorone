// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"conductorone/internal/sdk"
	"context"
	"fmt"

	"conductorone/internal/sdk/pkg/models/operations"
	"conductorone/internal/sdk/pkg/models/shared"
	"conductorone/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppResource{}
var _ resource.ResourceWithImportState = &AppResource{}

func NewAppResource() resource.Resource {
	return &AppResource{}
}

// AppResource defines the resource implementation.
type AppResource struct {
	client *sdk.ConductoroneAPI
}

// AppResourceModel describes the resource data model.
type AppResourceModel struct {
	AppAccountID    types.String   `tfsdk:"app_account_id"`
	AppAccountName  types.String   `tfsdk:"app_account_name"`
	CertifyPolicyID types.String   `tfsdk:"certify_policy_id"`
	CreatedAt       types.String   `tfsdk:"created_at"`
	DeletedAt       types.String   `tfsdk:"deleted_at"`
	Description     types.String   `tfsdk:"description"`
	DisplayName     types.String   `tfsdk:"display_name"`
	FieldMask       types.String   `tfsdk:"field_mask"`
	GrantPolicyID   types.String   `tfsdk:"grant_policy_id"`
	IconURL         types.String   `tfsdk:"icon_url"`
	ID              types.String   `tfsdk:"id"`
	LogoURI         types.String   `tfsdk:"logo_uri"`
	MonthlyCostUsd  types.Number   `tfsdk:"monthly_cost_usd"`
	Owners          []types.String `tfsdk:"owners"`
	ParentAppID     types.String   `tfsdk:"parent_app_id"`
	RevokePolicyID  types.String   `tfsdk:"revoke_policy_id"`
	UpdatedAt       types.String   `tfsdk:"updated_at"`
	UserCount       types.String   `tfsdk:"user_count"`
}

func (r *AppResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

func (r *AppResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "App Resource",

		Attributes: map[string]schema.Attribute{
			"app_account_id": schema.StringAttribute{
				Computed:    true,
				Description: `The appAccountId field.`,
			},
			"app_account_name": schema.StringAttribute{
				Computed:    true,
				Description: `The appAccountName field.`,
			},
			"certify_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The certifyPolicyId is the ID of the policy that will be used for access review certify tasks.`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The description field.`,
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: `The displayName field.`,
			},
			"field_mask": schema.StringAttribute{
				Computed: true,
			},
			"grant_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The grantPolicyId field is the policy that will be used for access request grant tasks.`,
			},
			"icon_url": schema.StringAttribute{
				Computed:    true,
				Description: `The iconUrl field.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The id field.`,
			},
			"logo_uri": schema.StringAttribute{
				Computed:    true,
				Description: `The logoUri field.`,
			},
			"monthly_cost_usd": schema.NumberAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The monthlyCostUsd field is the monthly cost per seat for the given app.`,
			},
			"owners": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `The owners field is a list of user IDs indicating the app owners.`,
			},
			"parent_app_id": schema.StringAttribute{
				Computed:    true,
				Description: `The parentAppId field is the ID of the parent app if one exists.`,
			},
			"revoke_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The revokePolicyId is the ID of the policy that will be used for revoke access tasks.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"user_count": schema.StringAttribute{
				Computed:    true,
				Description: `The userCount field is the number of app users that are associated with the app.`,
			},
		},
	}
}

func (r *AppResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AppResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToCreateSDKType()
	res, err := r.client.Apps.Create(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.CreateAppResponse.App == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.CreateAppResponse.App)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	request := operations.C1APIAppV1AppsGetRequest{
		ID: id,
	}
	res, err := r.client.Apps.Get(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.GetAppResponse.App == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if res.GetAppResponse.App.DeletedAt != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	data.RefreshFromGetResponse(res.GetAppResponse.App)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var updateAppRequest *shared.UpdateAppRequest
	app := data.ToUpdateSDKType()
	updateMask := "displayName,monthlyCostUsd"
	updateAppRequest = &shared.UpdateAppRequest{
		App:        app,
		UpdateMask: &updateMask,
	}
	id := data.ID.ValueString()
	request := operations.C1APIAppV1AppsUpdateRequest{
		UpdateAppRequest: updateAppRequest,
		ID:               id,
	}
	res, err := r.client.Apps.Update(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.UpdateAppResponse.App == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.UpdateAppResponse.App)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	deleteAppRequest := data.ToDeleteSDKType()
	id := data.ID.ValueString()
	request := operations.C1APIAppV1AppsDeleteRequest{
		DeleteAppRequest: deleteAppRequest,
		ID:               id,
	}
	res, err := r.client.Apps.Delete(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *AppResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

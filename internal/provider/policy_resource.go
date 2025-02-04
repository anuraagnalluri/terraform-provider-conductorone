package provider

import (
	"conductorone/internal/sdk"
	"context"
	"fmt"

	"conductorone/internal/sdk/pkg/models/operations"
	"conductorone/internal/sdk/pkg/models/shared"
	"conductorone/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicyResource{}
var _ resource.ResourceWithImportState = &PolicyResource{}

func NewPolicyResource() resource.Resource {
	return &PolicyResource{}
}

// PolicyResource defines the resource implementation.
type PolicyResource struct {
	client *sdk.ConductoroneAPI
}

// PolicyResourceModel describes the resource data model.
type PolicyResourceModel struct {
	CreatedAt                types.String           `tfsdk:"created_at"`
	DeletedAt                types.String           `tfsdk:"deleted_at"`
	Description              types.String           `tfsdk:"description"`
	DisplayName              types.String           `tfsdk:"display_name"`
	ID                       types.String           `tfsdk:"id"`
	PolicySteps              map[string]PolicySteps `tfsdk:"policy_steps"`
	PolicyType               types.String           `tfsdk:"policy_type"`
	PostActions              []PolicyPostActions    `tfsdk:"post_actions"`
	ReassignTasksToDelegates types.Bool             `tfsdk:"reassign_tasks_to_delegates"`
	SystemBuiltin            types.Bool             `tfsdk:"system_builtin"`
	UpdatedAt                types.String           `tfsdk:"updated_at"`
}

func (r *PolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

func (r *PolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Policy Resource",

		Attributes: map[string]schema.Attribute{
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
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The id field.`,
			},
			"policy_steps": schema.MapNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"steps": schema.ListNestedAttribute{
							Computed: true,
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"approval": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"allow_reassignment": schema.BoolAttribute{
												Computed:    true,
												Optional:    true,
												Description: `The allowReassignment field.`,
											},
											"app_group_approval": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The allowSelfApproval field.`,
													},
													"app_group_id": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The appGroupId field.`,
													},
													"app_id": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The appId field.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														Optional:    true,
														ElementType: types.StringType,
														Description: `The fallbackUserIds field.`,
													},
												},
												Description: `The AppGroupApproval message.`,
											},
											"app_owner_approval": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed: true,
														Optional: true,
														MarkdownDescription: ` App owner is based on the app id and doesn't need to have self-contained data` + "\n" +
															``,
													},
												},
												Description: `The AppOwnerApproval message.`,
											},
											"assigned": schema.BoolAttribute{
												Computed:    true,
												Optional:    true,
												Description: `The assigned field.`,
											},
											"entitlement_owner_approval": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed: true,
														Optional: true,
														MarkdownDescription: ` Entitlement owner is based on the current entitlement's id and doesn't need to have self-contained data` + "\n" +
															``,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														Optional:    true,
														ElementType: types.StringType,
														Description: `The fallbackUserIds field.`,
													},
												},
												Description: `The EntitlementOwnerApproval message.`,
											},
											"manager_approval": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The allowSelfApproval field.`,
													},
													"assigned_user_ids": schema.ListAttribute{
														Computed:    true,
														Optional:    true,
														ElementType: types.StringType,
														Description: `The assignedUserIds field.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														Optional:    true,
														ElementType: types.StringType,
														Description: `The fallbackUserIds field.`,
													},
												},
												Description: `The ManagerApproval message.`,
											},
											"require_approval_reason": schema.BoolAttribute{
												Computed:    true,
												Optional:    true,
												Description: `The requireApprovalReason field.`,
											},
											"require_reassignment_reason": schema.BoolAttribute{
												Computed:    true,
												Optional:    true,
												Description: `The requireReassignmentReason field.`,
											},
											"self_approval": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"assigned_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The assignedUserIds field.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														Optional:    true,
														ElementType: types.StringType,
														MarkdownDescription: ` Self approval is the target of the ticket` + "\n" +
															``,
													},
												},
												Description: `The SelfApproval message.`,
											},
											"user_approval": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Description: `The allowSelfApproval field.`,
													},
													"user_ids": schema.ListAttribute{
														Computed:    true,
														Optional:    true,
														ElementType: types.StringType,
														Description: `The userIds field.`,
													},
												},
												Description: `The UserApproval message.`,
											},
										},
										MarkdownDescription: `The Approval field is used to define who should perform the review.` + "\n" +
											`` +
											`This message contains a oneof. Only a single field of the following list may be set at a time:` + "\n" +
											`  - users` + "\n" +
											`  - manager` + "\n" +
											`  - appOwners` + "\n" +
											`  - group` + "\n" +
											`  - self` + "\n" +
											`  - entitlementOwners` + "\n" +
											"\n" +
											``,
									},
									"provision": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"assigned": schema.BoolAttribute{
												Computed:    true,
												Optional:    true,
												Description: `The assigned field.`,
											},
											"provision_policy": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"connector_provision": schema.SingleNestedAttribute{
														Optional:    true,
														Attributes:  map[string]schema.Attribute{},
														Description: `The ConnectorProvision message.`,
													},
													"delegated_provision": schema.SingleNestedAttribute{
														Optional: true,
														Attributes: map[string]schema.Attribute{
															"app_id": schema.StringAttribute{
																Computed:    true,
																Optional:    true,
																Description: `The appId field.`,
															},
															"entitlement_id": schema.StringAttribute{
																Computed:    true,
																Optional:    true,
																Description: `The entitlementId field.`,
															},
														},
														Description: `The DelegatedProvision message.`,
													},
													"manual_provision": schema.SingleNestedAttribute{
														Optional: true,
														Attributes: map[string]schema.Attribute{
															"instructions": schema.StringAttribute{
																Computed:    true,
																Optional:    true,
																Description: `The instructions field.`,
															},
															"user_ids": schema.ListAttribute{
																Computed:    true,
																Optional:    true,
																ElementType: types.StringType,
																Description: `The userIds field.`,
															},
														},
														Description: `The ManualProvision message.`,
													},
												},
												MarkdownDescription: `The ProvisionPolicy message.` + "\n" +
													`` +
													`This message contains a oneof. Only a single field of the following list may be set at a time:` + "\n" +
													`  - connector` + "\n" +
													`  - manual` + "\n" +
													`  - delegated` + "\n" +
													"\n" +
													``,
											},
										},
										Description: `The Provision message.`,
									},
								},
							},
							Description: `The steps field.`,
						},
					},
				},
				Description: `The policySteps field.`,
			},
			"policy_type": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"POLICY_TYPE_UNSPECIFIED",
						"POLICY_TYPE_GRANT",
						"POLICY_TYPE_REVOKE",
						"POLICY_TYPE_CERTIFY",
						"POLICY_TYPE_ACCESS_REQUEST",
						"POLICY_TYPE_PROVISION",
					),
				},
				MarkdownDescription: `must be one of [POLICY_TYPE_UNSPECIFIED, POLICY_TYPE_GRANT, POLICY_TYPE_REVOKE, POLICY_TYPE_CERTIFY, POLICY_TYPE_ACCESS_REQUEST, POLICY_TYPE_PROVISION]` + "\n" +
					`The policyType field.`,
			},
			"post_actions": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"certify_remediate_immediately": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							MarkdownDescription: ` ONLY valid when used in a CERTIFY Ticket Type:` + "\n" +
								` Causes any deprovision or change in a grant to be applied when Certify Ticket is closed.` + "\n" +
								`` +
								`See the documentation for ` + "`" + `c1.api.policy.v1.PolicyPostActions` + "`" + ` for more details.`,
						},
					},
				},
				Description: `The postActions field.`,
			},
			"reassign_tasks_to_delegates": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The reassignTasksToDelegates field.`,
			},
			"system_builtin": schema.BoolAttribute{
				Computed:    true,
				Description: `The systemBuiltin field.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *PolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PolicyResourceModel
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
	res, err := r.client.Policies.Create(ctx, request)
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
	if res.CreatePolicyResponse.Policy == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.CreatePolicyResponse.Policy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PolicyResourceModel
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
	request := operations.C1APIPolicyV1PoliciesGetRequest{
		ID: id,
	}
	res, err := r.client.Policies.Get(ctx, request)
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
	if res.GetPolicyResponse.Policy == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if res.GetPolicyResponse.Policy.DeletedAt != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	data.RefreshFromGetResponse(res.GetPolicyResponse.Policy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PolicyResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var updatePolicyRequest *shared.UpdatePolicyRequest
	policy := data.ToUpdateSDKType()
	updatePolicyRequest = &shared.UpdatePolicyRequest{
		Policy: policy,
	}
	id := data.ID.ValueString()
	request := operations.C1APIPolicyV1PoliciesUpdateRequest{
		UpdatePolicyRequest: updatePolicyRequest,
		ID:                  id,
	}
	res, err := r.client.Policies.Update(ctx, request)
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
	if res.UpdatePolicyResponse.Policy == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.UpdatePolicyResponse.Policy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PolicyResourceModel
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

	deletePolicyRequest := data.ToDeleteSDKType()
	id := data.ID.ValueString()
	request := operations.C1APIPolicyV1PoliciesDeleteRequest{
		DeletePolicyRequest: deletePolicyRequest,
		ID:                  id,
	}
	res, err := r.client.Policies.Delete(ctx, request)
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

func (r *PolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

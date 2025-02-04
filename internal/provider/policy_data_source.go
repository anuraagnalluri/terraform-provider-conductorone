package provider

import (
	"conductorone/internal/sdk"
	"conductorone/internal/validators"
	"context"
	"fmt"

	"conductorone/internal/sdk/pkg/models/operations"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NewPolicyDataSource() datasource.DataSource {
	return &PolicyDataSource{}
}

// PolicyDataSource defines the data source implementation.
type PolicyDataSource struct {
	client *sdk.ConductoroneAPI
}

// PolicyDataSourceModel describes the data source data model.
type PolicyDataSourceModel struct {
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

func (r *PolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

func (r *PolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Policy DataSource",

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
				Description: `The description field.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The displayName field.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The id field.`,
			},
			"policy_steps": schema.MapNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"steps": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"approval": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"allow_reassignment": schema.BoolAttribute{
												Computed:    true,
												Description: `The allowReassignment field.`,
											},
											"app_group_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `The allowSelfApproval field.`,
													},
													"app_group_id": schema.StringAttribute{
														Computed:    true,
														Description: `The appGroupId field.`,
													},
													"app_id": schema.StringAttribute{
														Computed:    true,
														Description: `The appId field.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The fallbackUserIds field.`,
													},
												},
												Description: `The AppGroupApproval message.`,
											},
											"app_owner_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed: true,
														MarkdownDescription: ` App owner is based on the app id and doesn't need to have self-contained data` + "\n" +
															``,
													},
												},
												Description: `The AppOwnerApproval message.`,
											},
											"assigned": schema.BoolAttribute{
												Computed:    true,
												Description: `The assigned field.`,
											},
											"entitlement_owner_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed: true,
														MarkdownDescription: ` Entitlement owner is based on the current entitlement's id and doesn't need to have self-contained data` + "\n" +
															``,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The fallbackUserIds field.`,
													},
												},
												Description: `The EntitlementOwnerApproval message.`,
											},
											"manager_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `The allowSelfApproval field.`,
													},
													"assigned_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The assignedUserIds field.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The fallbackUserIds field.`,
													},
												},
												Description: `The ManagerApproval message.`,
											},
											"require_approval_reason": schema.BoolAttribute{
												Computed:    true,
												Description: `The requireApprovalReason field.`,
											},
											"require_reassignment_reason": schema.BoolAttribute{
												Computed:    true,
												Description: `The requireReassignmentReason field.`,
											},
											"self_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"assigned_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The assignedUserIds field.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `The fallback field.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														MarkdownDescription: ` Self approval is the target of the ticket` + "\n" +
															``,
													},
												},
												Description: `The SelfApproval message.`,
											},
											"user_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `The allowSelfApproval field.`,
													},
													"user_ids": schema.ListAttribute{
														Computed:    true,
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
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"assigned": schema.BoolAttribute{
												Computed:    true,
												Description: `The assigned field.`,
											},
											"provision_policy": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"connector_provision": schema.SingleNestedAttribute{
														Computed:    true,
														Attributes:  map[string]schema.Attribute{},
														Description: `The ConnectorProvision message.`,
													},
													"delegated_provision": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"app_id": schema.StringAttribute{
																Computed:    true,
																Description: `The appId field.`,
															},
															"entitlement_id": schema.StringAttribute{
																Computed:    true,
																Description: `The entitlementId field.`,
															},
														},
														Description: `The DelegatedProvision message.`,
													},
													"manual_provision": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"instructions": schema.StringAttribute{
																Computed:    true,
																Description: `The instructions field.`,
															},
															"user_ids": schema.ListAttribute{
																Computed:    true,
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
				MarkdownDescription: `must be one of [POLICY_TYPE_UNSPECIFIED, POLICY_TYPE_GRANT, POLICY_TYPE_REVOKE, POLICY_TYPE_CERTIFY, POLICY_TYPE_ACCESS_REQUEST, POLICY_TYPE_PROVISION]` + "\n" +
					`The policyType field.`,
			},
			"post_actions": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"certify_remediate_immediately": schema.BoolAttribute{
							Computed: true,
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

func (r *PolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PolicyDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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
	displayName := data.DisplayName.ValueStringPointer()
	if id == "" && (displayName == nil || *displayName == "") {
		resp.Diagnostics.AddError("either id or display_name must be set", "")
		return
	}

	// If we have an ID, we can use the Get API to fetch the latest data
	if id != "" {
		req := operations.C1APIPolicyV1PoliciesGetRequest{
			ID: id,
		}
		res, err := r.client.Policies.Get(ctx, req)
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
		data.RefreshFromGetResponse(res.GetPolicyResponse.Policy)
		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// If we don't have an ID but we do have a displayName we can use the Search API to fetch the latest data.
	request := shared.SearchPoliciesRequest{
		DisplayName: displayName,
	}
	res, err := r.client.PolicySearch.Search(ctx, request)
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
	if res.ListPolicyResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if len(res.ListPolicyResponse.List) == 0 {
		resp.Diagnostics.AddError("unexpected response from API. Policy was not found", debugResponse(res.RawResponse))
		return
	}

	if len(res.ListPolicyResponse.List) > 2 {
		resp.Diagnostics.AddError("unexpected response from API. More than 1 policy was found", debugResponse(res.RawResponse))
		return
	}

	data.RefreshFromGetResponse(&res.ListPolicyResponse.List[0])

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

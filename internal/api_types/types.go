// Package api_types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api_types

import (
	"time"
)

// Defines values for CampaignStatusEnum.
const (
	CampaignStatusEnumCancelled CampaignStatusEnum = "Cancelled"
	CampaignStatusEnumDraft     CampaignStatusEnum = "Draft"
	CampaignStatusEnumRunning   CampaignStatusEnum = "Running"
	CampaignStatusEnumScheduled CampaignStatusEnum = "Scheduled"
	CampaignStatusEnumSent      CampaignStatusEnum = "Sent"
)

// Defines values for ContactStatusEnum.
const (
	ContactStatusEnumActive   ContactStatusEnum = "Active"
	ContactStatusEnumBlocked  ContactStatusEnum = "Blocked"
	ContactStatusEnumInactive ContactStatusEnum = "Inactive"
)

// Defines values for IntegrationStatusEnum.
const (
	IntegrationStatusEnumActive   IntegrationStatusEnum = "Active"
	IntegrationStatusEnumInactive IntegrationStatusEnum = "Inactive"
)

// Defines values for MessageDirectionEnum.
const (
	InBound  MessageDirectionEnum = "InBound"
	OutBound MessageDirectionEnum = "OutBound"
)

// Defines values for MessageStatusEnum.
const (
	MessageStatusEnumDelivered   MessageStatusEnum = "Delivered"
	MessageStatusEnumFailed      MessageStatusEnum = "Failed"
	MessageStatusEnumRead        MessageStatusEnum = "Read"
	MessageStatusEnumSent        MessageStatusEnum = "Sent"
	MessageStatusEnumUnDelivered MessageStatusEnum = "UnDelivered"
	MessageStatusEnumUnread      MessageStatusEnum = "Unread"
)

// Defines values for MessageTypeEnum.
const (
	Address  MessageTypeEnum = "Address"
	Audio    MessageTypeEnum = "Audio"
	Contacts MessageTypeEnum = "Contacts"
	Document MessageTypeEnum = "Document"
	Image    MessageTypeEnum = "Image"
	Location MessageTypeEnum = "Location"
	Reaction MessageTypeEnum = "Reaction"
	Sticker  MessageTypeEnum = "Sticker"
	Text     MessageTypeEnum = "Text"
	Video    MessageTypeEnum = "Video"
)

// Defines values for OrderEnum.
const (
	Asc  OrderEnum = "asc"
	Desc OrderEnum = "desc"
)

// Defines values for RolePermissionEnum.
const (
	GetApiKey             RolePermissionEnum = "GetApiKey"
	GetAppSettings        RolePermissionEnum = "GetAppSettings"
	GetCampaign           RolePermissionEnum = "GetCampaign"
	GetConversations      RolePermissionEnum = "getConversations"
	GetList               RolePermissionEnum = "GetList"
	GetTeam               RolePermissionEnum = "GetTeam"
	MessageInConversation RolePermissionEnum = "MessageInConversation"
	UpdateApikey          RolePermissionEnum = "UpdateApikey"
	UpdateAppSettings     RolePermissionEnum = "UpdateAppSettings"
	UpdateCampaign        RolePermissionEnum = "UpdateCampaign"
	UpdateConversation    RolePermissionEnum = "UpdateConversation"
	UpdateList            RolePermissionEnum = "UpdateList"
	UpdateTeam            RolePermissionEnum = "UpdateTeam"
)

// Defines values for UpdateOrganizationMemberRoleSchemaAction.
const (
	Add    UpdateOrganizationMemberRoleSchemaAction = "add"
	Remove UpdateOrganizationMemberRoleSchemaAction = "remove"
)

// Defines values for UserRoleEnum.
const (
	Admin  UserRoleEnum = "Admin"
	Member UserRoleEnum = "Member"
	Owner  UserRoleEnum = "Owner"
)

// Defines values for GetConversationsParamsStatus.
const (
	Resolved   GetConversationsParamsStatus = "resolved"
	Unresolved GetConversationsParamsStatus = "unresolved"
)

// Defines values for GetMessagesParamsStatus.
const (
	Failed GetMessagesParamsStatus = "failed"
	Read   GetMessagesParamsStatus = "read"
	Sent   GetMessagesParamsStatus = "sent"
	Unread GetMessagesParamsStatus = "unread"
)

// ApiKeySchema defines model for ApiKeySchema.
type ApiKeySchema struct {
	CreatedAt time.Time `json:"createdAt"`
	Key       string    `json:"key"`
	UniqueId  string    `json:"uniqueId"`
}

// CampaignSchema defines model for CampaignSchema.
type CampaignSchema struct {
	CreatedAt             time.Time           `json:"createdAt"`
	Description           *string             `json:"description,omitempty"`
	IsLinkTrackingEnabled bool                `json:"isLinkTrackingEnabled"`
	Lists                 []ContactListSchema `json:"lists"`
	Name                  string              `json:"name"`
	ScheduledAt           *time.Time          `json:"scheduledAt,omitempty"`
	SentAt                *time.Time          `json:"sentAt,omitempty"`
	Status                CampaignStatusEnum  `json:"status"`
	Tags                  []TagSchema         `json:"tags"`
	TemplateMessageId     *string             `json:"templateMessageId,omitempty"`
	UniqueId              string              `json:"uniqueId"`
}

// CampaignStatusEnum defines model for CampaignStatusEnum.
type CampaignStatusEnum string

// ContactListSchema defines model for ContactListSchema.
type ContactListSchema struct {
	CreatedAt             time.Time   `json:"createdAt"`
	Description           string      `json:"description"`
	Name                  string      `json:"name"`
	NumberOfCampaignsSent int         `json:"numberOfCampaignsSent"`
	NumberOfContacts      int         `json:"numberOfContacts"`
	Tags                  []TagSchema `json:"tags"`
	UniqueId              string      `json:"uniqueId"`
}

// ContactSchema defines model for ContactSchema.
type ContactSchema struct {
	Attributes map[string]interface{} `json:"attributes"`
	CreatedAt  time.Time              `json:"createdAt"`
	Lists      []ContactListSchema    `json:"lists"`
	Name       string                 `json:"name"`
	Phone      string                 `json:"phone"`
	UniqueId   string                 `json:"uniqueId"`
}

// ContactStatusEnum defines model for ContactStatusEnum.
type ContactStatusEnum string

// ConversationSchema defines model for ConversationSchema.
type ConversationSchema struct {
	ContactId string     `json:"contactId"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Message   *string    `json:"message,omitempty"`
	UniqueId  string     `json:"uniqueId"`
}

// CreateNewCampaignResponseSchema defines model for CreateNewCampaignResponseSchema.
type CreateNewCampaignResponseSchema struct {
	Campaign CampaignSchema `json:"campaign"`
}

// CreateNewOrganizationResponseSchema defines model for CreateNewOrganizationResponseSchema.
type CreateNewOrganizationResponseSchema struct {
	Organization OrganizationSchema `json:"organization"`
}

// CreateNewOrganizationTagResponseSchema defines model for CreateNewOrganizationTagResponseSchema.
type CreateNewOrganizationTagResponseSchema struct {
	Tag TagSchema `json:"tag"`
}

// CreateNewRoleResponseSchema defines model for CreateNewRoleResponseSchema.
type CreateNewRoleResponseSchema struct {
	Role OrganizationRoleSchema `json:"role"`
}

// CreateOrganizationMemberResponseSchema defines model for CreateOrganizationMemberResponseSchema.
type CreateOrganizationMemberResponseSchema struct {
	Member OrganizationMemberSchema `json:"member"`
}

// DeleteOrganizationMemberByIdResponseSchema defines model for DeleteOrganizationMemberByIdResponseSchema.
type DeleteOrganizationMemberByIdResponseSchema struct {
	Data bool `json:"data"`
}

// DeleteRoleByIdResponseSchema defines model for DeleteRoleByIdResponseSchema.
type DeleteRoleByIdResponseSchema struct {
	Data bool `json:"data"`
}

// GetApiKeysResponseSchema defines model for GetApiKeysResponseSchema.
type GetApiKeysResponseSchema struct {
	ApiKeys        []ApiKeySchema `json:"apiKeys"`
	PaginationMeta PaginationMeta `json:"paginationMeta"`
}

// GetCampaignByIdResponseSchema defines model for GetCampaignByIdResponseSchema.
type GetCampaignByIdResponseSchema struct {
	Campaign CampaignSchema `json:"campaign"`
}

// GetCampaignResponseSchema defines model for GetCampaignResponseSchema.
type GetCampaignResponseSchema struct {
	Campaigns      []CampaignSchema `json:"campaigns"`
	PaginationMeta PaginationMeta   `json:"paginationMeta"`
}

// GetContactByIdResponseSchema defines model for GetContactByIdResponseSchema.
type GetContactByIdResponseSchema struct {
	Contact ContactSchema `json:"contact"`
}

// GetContactListByIdSchema defines model for GetContactListByIdSchema.
type GetContactListByIdSchema struct {
	List ContactListSchema `json:"list"`
}

// GetContactListResponseSchema defines model for GetContactListResponseSchema.
type GetContactListResponseSchema struct {
	Lists          []ContactListSchema `json:"lists"`
	PaginationMeta PaginationMeta      `json:"paginationMeta"`
}

// GetContactsResponseSchema defines model for GetContactsResponseSchema.
type GetContactsResponseSchema struct {
	Contacts       []ContactSchema `json:"contacts"`
	PaginationMeta PaginationMeta  `json:"paginationMeta"`
}

// GetOrganizationByIdResponseSchema defines model for GetOrganizationByIdResponseSchema.
type GetOrganizationByIdResponseSchema struct {
	Organization OrganizationSchema `json:"organization"`
}

// GetOrganizationMemberByIdResponseSchema defines model for GetOrganizationMemberByIdResponseSchema.
type GetOrganizationMemberByIdResponseSchema struct {
	Member OrganizationMemberSchema `json:"member"`
}

// GetOrganizationMembersResponseSchema defines model for GetOrganizationMembersResponseSchema.
type GetOrganizationMembersResponseSchema struct {
	Members        []OrganizationMemberSchema `json:"members"`
	PaginationMeta PaginationMeta             `json:"paginationMeta"`
}

// GetOrganizationRolesResponseSchema defines model for GetOrganizationRolesResponseSchema.
type GetOrganizationRolesResponseSchema struct {
	PaginationMeta PaginationMeta           `json:"paginationMeta"`
	Roles          []OrganizationRoleSchema `json:"roles"`
}

// GetOrganizationSettingsResponseSchema defines model for GetOrganizationSettingsResponseSchema.
type GetOrganizationSettingsResponseSchema struct {
	Settings *[]struct {
		Key   *string `json:"key,omitempty"`
		Value *string `json:"value,omitempty"`
	} `json:"settings,omitempty"`
}

// GetOrganizationTagsResponseSchema defines model for GetOrganizationTagsResponseSchema.
type GetOrganizationTagsResponseSchema struct {
	PaginationMeta PaginationMeta `json:"paginationMeta"`
	Tags           []TagSchema    `json:"tags"`
}

// GetOrganizationsResponseSchema defines model for GetOrganizationsResponseSchema.
type GetOrganizationsResponseSchema struct {
	Organizations  []OrganizationSchema `json:"organizations"`
	PaginationMeta PaginationMeta       `json:"paginationMeta"`
}

// GetRoleByIdResponseSchema defines model for GetRoleByIdResponseSchema.
type GetRoleByIdResponseSchema struct {
	Role OrganizationRoleSchema `json:"role"`
}

// GetTemplateByIdResponseSchema defines model for GetTemplateByIdResponseSchema.
type GetTemplateByIdResponseSchema struct {
	Template TemplateSchema `json:"template"`
}

// GetUserResponseSchema defines model for GetUserResponseSchema.
type GetUserResponseSchema struct {
	User UserSchema `json:"user"`
}

// IntegrationSchema defines model for IntegrationSchema.
type IntegrationSchema struct {
	CreatedAt *time.Time             `json:"createdAt,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Slug      *string                `json:"slug,omitempty"`
	Status    *IntegrationStatusEnum `json:"status,omitempty"`
	Type      *string                `json:"type,omitempty"`
	UniqueId  *string                `json:"uniqueId,omitempty"`
}

// IntegrationStatusEnum defines model for IntegrationStatusEnum.
type IntegrationStatusEnum string

// LoginRequestBodySchema defines model for LoginRequestBodySchema.
type LoginRequestBodySchema struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginResponseBodySchema defines model for LoginResponseBodySchema.
type LoginResponseBodySchema struct {
	IsOnboardingCompleted bool   `json:"isOnboardingCompleted"`
	Token                 string `json:"token"`
}

// MessageDirectionEnum defines model for MessageDirectionEnum.
type MessageDirectionEnum string

// MessageSchema defines model for MessageSchema.
type MessageSchema struct {
	Content        *map[string]interface{} `json:"content,omitempty"`
	ConversationId *string                 `json:"conversationId,omitempty"`
	CreatedAt      *time.Time              `json:"createdAt,omitempty"`
	Direction      *MessageDirectionEnum   `json:"direction,omitempty"`
	Message        *string                 `json:"message,omitempty"`
	MessageType    *MessageTypeEnum        `json:"message_type,omitempty"`
	Status         *MessageStatusEnum      `json:"status,omitempty"`
	UniqueId       *string                 `json:"uniqueId,omitempty"`
}

// MessageStatusEnum defines model for MessageStatusEnum.
type MessageStatusEnum string

// MessageTypeEnum defines model for MessageTypeEnum.
type MessageTypeEnum string

// NewCampaignSchema defines model for NewCampaignSchema.
type NewCampaignSchema struct {
	Description           *string  `json:"description,omitempty"`
	IsLinkTrackingEnabled bool     `json:"isLinkTrackingEnabled"`
	ListIds               []string `json:"listIds"`
	Name                  string   `json:"name"`
	Tags                  []string `json:"tags"`
	TemplateMessageId     *string  `json:"templateMessageId,omitempty"`
}

// NewContactListSchema defines model for NewContactListSchema.
type NewContactListSchema struct {
	ContactIds  *[]string   `json:"contactIds,omitempty"`
	Description *string     `json:"description,omitempty"`
	Name        string      `json:"name"`
	Tags        []TagSchema `json:"tags"`
}

// NewContactSchema defines model for NewContactSchema.
type NewContactSchema struct {
	Attributes map[string]interface{} `json:"attributes"`
	ListsIds   []string               `json:"listsIds"`
	Name       string                 `json:"name"`
	Phone      string                 `json:"phone"`
	Status     ContactStatusEnum      `json:"status"`
}

// NewOrganizationMemberInviteSchemaSchema defines model for NewOrganizationMemberInviteSchemaSchema.
type NewOrganizationMemberInviteSchemaSchema struct {
	AccessLevel UserRoleEnum              `json:"accessLevel"`
	Email       string                    `json:"email"`
	Roles       *[]OrganizationRoleSchema `json:"roles,omitempty"`
}

// NewOrganizationRoleSchema defines model for NewOrganizationRoleSchema.
type NewOrganizationRoleSchema struct {
	Description *string              `json:"description,omitempty"`
	Name        string               `json:"name"`
	Permissions []RolePermissionEnum `json:"permissions"`
}

// NewOrganizationSchema defines model for NewOrganizationSchema.
type NewOrganizationSchema struct {
	Name string `json:"name"`
}

// NewOrganizationTagSchema defines model for NewOrganizationTagSchema.
type NewOrganizationTagSchema struct {
	Name string `json:"name"`
}

// OrderEnum defines model for OrderEnum.
type OrderEnum string

// OrganizationMemberSchema defines model for OrganizationMemberSchema.
type OrganizationMemberSchema struct {
	AccessLevel UserRoleEnum             `json:"accessLevel"`
	CreatedAt   time.Time                `json:"createdAt"`
	Email       string                   `json:"email"`
	Name        string                   `json:"name"`
	Roles       []OrganizationRoleSchema `json:"roles"`
	UniqueId    string                   `json:"uniqueId"`
}

// OrganizationRoleSchema defines model for OrganizationRoleSchema.
type OrganizationRoleSchema struct {
	Description *string              `json:"description,omitempty"`
	Name        string               `json:"name"`
	Permissions []RolePermissionEnum `json:"permissions"`
	UniqueId    string               `json:"uniqueId"`
}

// OrganizationSchema defines model for OrganizationSchema.
type OrganizationSchema struct {
	CreatedAt  time.Time `json:"createdAt"`
	FaviconUrl *string   `json:"faviconUrl,omitempty"`
	LogoUrl    *string   `json:"logoUrl,omitempty"`
	Name       string    `json:"name"`
	UniqueId   string    `json:"uniqueId"`
	WebsiteUrl *string   `json:"websiteUrl,omitempty"`
}

// PaginationMeta defines model for PaginationMeta.
type PaginationMeta struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
	Total   int   `json:"total"`
}

// RolePermissionEnum defines model for RolePermissionEnum.
type RolePermissionEnum string

// SwitchOrganizationResponseSchema defines model for SwitchOrganizationResponseSchema.
type SwitchOrganizationResponseSchema struct {
	Token string `json:"token"`
}

// TagSchema defines model for TagSchema.
type TagSchema struct {
	Name     string `json:"name"`
	UniqueId string `json:"uniqueId"`
}

// TemplateSchema defines model for TemplateSchema.
type TemplateSchema struct {
	BodyText  string                  `json:"bodyText"`
	Content   string                  `json:"content"`
	CreatedAt time.Time               `json:"createdAt"`
	Footer    *map[string]interface{} `json:"footer,omitempty"`
	Header    struct {
		Content    *string `json:"content,omitempty"`
		HeaderType *string `json:"headerType,omitempty"`
	} `json:"header"`
	TemplateId string `json:"templateId"`
}

// UpdateCampaignSchema defines model for UpdateCampaignSchema.
type UpdateCampaignSchema struct {
	Description        *string  `json:"description,omitempty"`
	EnableLinkTracking bool     `json:"enableLinkTracking"`
	ListIds            []string `json:"listIds"`
	Name               string   `json:"name"`
	Tags               []string `json:"tags"`
	TemplateMessageId  *string  `json:"templateMessageId,omitempty"`
}

// UpdateContactListSchema defines model for UpdateContactListSchema.
type UpdateContactListSchema struct {
	Description *string     `json:"description,omitempty"`
	Name        string      `json:"name"`
	Tags        []TagSchema `json:"tags"`
}

// UpdateContactSchema defines model for UpdateContactSchema.
type UpdateContactSchema struct {
	Attributes map[string]interface{} `json:"attributes"`
	Lists      *[]string              `json:"lists,omitempty"`
	Name       string                 `json:"name"`
	Phone      string                 `json:"phone"`
	Status     ContactStatusEnum      `json:"status"`
}

// UpdateOrganizationByIdResponseSchema defines model for UpdateOrganizationByIdResponseSchema.
type UpdateOrganizationByIdResponseSchema struct {
	Organization OrganizationSchema `json:"organization"`
}

// UpdateOrganizationMemberByIdResponseSchema defines model for UpdateOrganizationMemberByIdResponseSchema.
type UpdateOrganizationMemberByIdResponseSchema struct {
	Member OrganizationMemberSchema `json:"member"`
}

// UpdateOrganizationMemberRoleSchema defines model for UpdateOrganizationMemberRoleSchema.
type UpdateOrganizationMemberRoleSchema struct {
	Action       UpdateOrganizationMemberRoleSchemaAction `json:"action"`
	RoleUniqueId *string                                  `json:"roleUniqueId,omitempty"`
}

// UpdateOrganizationMemberRoleSchemaAction defines model for UpdateOrganizationMemberRoleSchema.Action.
type UpdateOrganizationMemberRoleSchemaAction string

// UpdateOrganizationMemberSchema defines model for UpdateOrganizationMemberSchema.
type UpdateOrganizationMemberSchema struct {
	AccessLevel *UserRoleEnum `json:"accessLevel,omitempty"`
}

// UpdateOrganizationSchema defines model for UpdateOrganizationSchema.
type UpdateOrganizationSchema struct {
	Name *string `json:"name,omitempty"`
}

// UpdateOrganizationSettingsResponseSchema defines model for UpdateOrganizationSettingsResponseSchema.
type UpdateOrganizationSettingsResponseSchema struct {
	Setting *struct {
		Key   *string `json:"key,omitempty"`
		Value *string `json:"value,omitempty"`
	} `json:"setting,omitempty"`
}

// UpdateRoleByIdResponseSchema defines model for UpdateRoleByIdResponseSchema.
type UpdateRoleByIdResponseSchema struct {
	Role OrganizationRoleSchema `json:"role"`
}

// UserRoleEnum defines model for UserRoleEnum.
type UserRoleEnum string

// UserSchema defines model for UserSchema.
type UserSchema struct {
	CreatedAt               time.Time            `json:"createdAt"`
	CurrentOrganizationRole *string              `json:"currentOrganizationRole",omitempty"`
	Email                   string               `json:"email"`
	Name                    string               `json:"name"`
	Organizations           []OrganizationSchema `json:"organizations"`
	ProfilePicture          *string              `json:"profilePicture,omitempty"`
	UniqueId                string               `json:"uniqueId"`
	Username                string               `json:"username"`
}

// SwitchOrganizationJSONBody defines parameters for SwitchOrganization.
type SwitchOrganizationJSONBody struct {
	OrganizationId *string `json:"organizationId,omitempty"`
}

// GetCampaignsParams defines parameters for GetCampaigns.
type GetCampaignsParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// Order order by asc or desc
	Order *OrderEnum `form:"order,omitempty" json:"order,omitempty"`

	// Status sort by a field
	Status *CampaignStatusEnum `form:"status,omitempty" json:"status,omitempty"`
}

// DeleteContactsByListParams defines parameters for DeleteContactsByList.
type DeleteContactsByListParams struct {
	// Id contact id/s to be deleted
	Id string `form:"id" json:"id"`
}

// GetContactsParams defines parameters for GetContacts.
type GetContactsParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// ListId query subscribers with a list id.
	ListId *string `form:"list_id,omitempty" json:"list_id,omitempty"`

	// Order order by asc or desc
	Order *OrderEnum `form:"order,omitempty" json:"order,omitempty"`

	// Status sort by a field
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}

// GetConversationsParams defines parameters for GetConversations.
type GetConversationsParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// Order order by asc or desc
	Order *OrderEnum `form:"order,omitempty" json:"order,omitempty"`

	// Status sort by a field
	Status *GetConversationsParamsStatus `form:"status,omitempty" json:"status,omitempty"`

	// ContactId query conversations with a contact id.
	ContactId *string `form:"contact_id,omitempty" json:"contact_id,omitempty"`

	// CampaignId query conversations with a campaign id.
	CampaignId *string `form:"campaign_id,omitempty" json:"campaign_id,omitempty"`

	// ListId query conversations with a list id.
	ListId *string `form:"list_id,omitempty" json:"list_id,omitempty"`

	// MessageId query conversations with a message id.
	MessageId *string `form:"message_id,omitempty" json:"message_id,omitempty"`
}

// GetConversationsParamsStatus defines parameters for GetConversations.
type GetConversationsParamsStatus string

// GetIntegrationsParams defines parameters for GetIntegrations.
type GetIntegrationsParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Order order by asc or desc
	Order *OrderEnum `form:"order,omitempty" json:"order,omitempty"`

	// Status status of the integration
	Status *IntegrationStatusEnum `form:"status,omitempty" json:"status,omitempty"`
}

// GetContactListsParams defines parameters for GetContactLists.
type GetContactListsParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// Order order by asc or desc
	Order *OrderEnum `form:"order,omitempty" json:"order,omitempty"`
}

// GetMessagesParams defines parameters for GetMessages.
type GetMessagesParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Order order by asc or desc
	Order *OrderEnum `form:"order,omitempty" json:"order,omitempty"`

	// Status status of the message
	Status *GetMessagesParamsStatus `form:"status,omitempty" json:"status,omitempty"`

	// Direction direction of the message
	Direction *MessageDirectionEnum `form:"direction,omitempty" json:"direction,omitempty"`

	// ContactId query messages with a contact id.
	ContactId *string `form:"contact_id,omitempty" json:"contact_id,omitempty"`

	// CampaignId query messages with a campaign id.
	CampaignId *string `form:"campaign_id,omitempty" json:"campaign_id,omitempty"`

	// ListId query messages with a list id.
	ListId *string `form:"list_id,omitempty" json:"list_id,omitempty"`

	// ConversationId query messages with a conversation id.
	ConversationId *string `form:"conversation_id,omitempty" json:"conversation_id,omitempty"`
}

// GetMessagesParamsStatus defines parameters for GetMessages.
type GetMessagesParamsStatus string

// GetUserOrganizationsParams defines parameters for GetUserOrganizations.
type GetUserOrganizationsParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// SortBy sorting order
	SortBy *OrderEnum `form:"sortBy,omitempty" json:"sortBy,omitempty"`
}

// GetOrganizationMembersParams defines parameters for GetOrganizationMembers.
type GetOrganizationMembersParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// SortBy sorting order
	SortBy *OrderEnum `form:"sortBy,omitempty" json:"sortBy,omitempty"`
}

// GetOrganizationRolesParams defines parameters for GetOrganizationRoles.
type GetOrganizationRolesParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// SortBy sorting order
	SortBy *OrderEnum `form:"sortBy,omitempty" json:"sortBy,omitempty"`
}

// UpdateOrganizationRoleByIdJSONBody defines parameters for UpdateOrganizationRoleById.
type UpdateOrganizationRoleByIdJSONBody struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// UpdateSettingsJSONBody defines parameters for UpdateSettings.
type UpdateSettingsJSONBody struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// GetOrganizationTagsParams defines parameters for GetOrganizationTags.
type GetOrganizationTagsParams struct {
	// Page number of records to skip
	Page int64 `form:"page" json:"page"`

	// PerPage max number of records to return per page
	PerPage int64 `form:"per_page" json:"per_page"`

	// SortBy sorting order
	SortBy *OrderEnum `form:"sortBy,omitempty" json:"sortBy,omitempty"`
}

// GetTemplatesParams defines parameters for GetTemplates.
type GetTemplatesParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Order order by asc or desc
	Order *OrderEnum `form:"order,omitempty" json:"order,omitempty"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LoginRequestBodySchema

// SwitchOrganizationJSONRequestBody defines body for SwitchOrganization for application/json ContentType.
type SwitchOrganizationJSONRequestBody SwitchOrganizationJSONBody

// CreateCampaignJSONRequestBody defines body for CreateCampaign for application/json ContentType.
type CreateCampaignJSONRequestBody = NewCampaignSchema

// UpdateCampaignByIdJSONRequestBody defines body for UpdateCampaignById for application/json ContentType.
type UpdateCampaignByIdJSONRequestBody = UpdateCampaignSchema

// CreateContactJSONRequestBody defines body for CreateContact for application/json ContentType.
type CreateContactJSONRequestBody = NewContactSchema

// UpdateContactByIdJSONRequestBody defines body for UpdateContactById for application/json ContentType.
type UpdateContactByIdJSONRequestBody = UpdateContactSchema

// CreateListJSONRequestBody defines body for CreateList for application/json ContentType.
type CreateListJSONRequestBody = NewContactListSchema

// UpdateListByIdJSONRequestBody defines body for UpdateListById for application/json ContentType.
type UpdateListByIdJSONRequestBody = UpdateContactListSchema

// CreateOrganizationJSONRequestBody defines body for CreateOrganization for application/json ContentType.
type CreateOrganizationJSONRequestBody = NewOrganizationSchema

// CreateOrganizationMemberJSONRequestBody defines body for CreateOrganizationMember for application/json ContentType.
type CreateOrganizationMemberJSONRequestBody = NewOrganizationMemberInviteSchemaSchema

// UpdateOrganizationMemberByIdJSONRequestBody defines body for UpdateOrganizationMemberById for application/json ContentType.
type UpdateOrganizationMemberByIdJSONRequestBody = UpdateOrganizationMemberSchema

// UpdateOrganizationMemberRoleByIdJSONRequestBody defines body for UpdateOrganizationMemberRoleById for application/json ContentType.
type UpdateOrganizationMemberRoleByIdJSONRequestBody = UpdateOrganizationMemberRoleSchema

// CreateOrganizationRoleJSONRequestBody defines body for CreateOrganizationRole for application/json ContentType.
type CreateOrganizationRoleJSONRequestBody = NewOrganizationRoleSchema

// UpdateOrganizationRoleByIdJSONRequestBody defines body for UpdateOrganizationRoleById for application/json ContentType.
type UpdateOrganizationRoleByIdJSONRequestBody UpdateOrganizationRoleByIdJSONBody

// UpdateSettingsJSONRequestBody defines body for UpdateSettings for application/json ContentType.
type UpdateSettingsJSONRequestBody UpdateSettingsJSONBody

// CreateOrganizationTagJSONRequestBody defines body for CreateOrganizationTag for application/json ContentType.
type CreateOrganizationTagJSONRequestBody = NewOrganizationTagSchema

// UpdateOrganizationJSONRequestBody defines body for UpdateOrganization for application/json ContentType.
type UpdateOrganizationJSONRequestBody = UpdateOrganizationSchema

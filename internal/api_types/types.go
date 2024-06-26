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

// Defines values for UserRoleEnum.
const (
	Admin  UserRoleEnum = "Admin"
	Member UserRoleEnum = "Member"
	Owner  UserRoleEnum = "Owner"
)

// Defines values for GetCampaignsParamsOrder.
const (
	GetCampaignsParamsOrderAsc  GetCampaignsParamsOrder = "asc"
	GetCampaignsParamsOrderDesc GetCampaignsParamsOrder = "desc"
)

// Defines values for GetConversationsParamsOrder.
const (
	GetConversationsParamsOrderAsc  GetConversationsParamsOrder = "asc"
	GetConversationsParamsOrderDesc GetConversationsParamsOrder = "desc"
)

// Defines values for GetConversationsParamsStatus.
const (
	Resolved   GetConversationsParamsStatus = "resolved"
	Unresolved GetConversationsParamsStatus = "unresolved"
)

// Defines values for GetContactListsParamsOrder.
const (
	GetContactListsParamsOrderAsc  GetContactListsParamsOrder = "asc"
	GetContactListsParamsOrderDesc GetContactListsParamsOrder = "desc"
)

// Defines values for GetMessagesParamsOrder.
const (
	GetMessagesParamsOrderAsc  GetMessagesParamsOrder = "asc"
	GetMessagesParamsOrderDesc GetMessagesParamsOrder = "desc"
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
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Key       *string    `json:"key,omitempty"`
	UniqueId  *string    `json:"uniqueId,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// CampaignSchema defines model for CampaignSchema.
type CampaignSchema struct {
	CreatedAt             *time.Time           `json:"createdAt,omitempty"`
	Description           *string              `json:"description,omitempty"`
	IsLinkTrackingEnabled *bool                `json:"isLinkTrackingEnabled,omitempty"`
	Lists                 *[]ContactListSchema `json:"lists,omitempty"`
	Name                  *string              `json:"name,omitempty"`
	ScheduledAt           *time.Time           `json:"scheduledAt,omitempty"`
	SentAt                *time.Time           `json:"sentAt,omitempty"`
	Status                *CampaignStatusEnum  `json:"status,omitempty"`
	Tags                  *[]TagSchema         `json:"tags,omitempty"`
	TemplateMessageId     *string              `json:"templateMessageId,omitempty"`
	UniqueId              *string              `json:"uniqueId,omitempty"`
	UpdatedAt             *time.Time           `json:"updatedAt,omitempty"`
}

// CampaignStatusEnum defines model for CampaignStatusEnum.
type CampaignStatusEnum string

// ContactListSchema defines model for ContactListSchema.
type ContactListSchema struct {
	CreatedAt             *string      `json:"created_at,omitempty"`
	Description           *string      `json:"description,omitempty"`
	Name                  *string      `json:"name,omitempty"`
	NumberOfCampaignsSent *int         `json:"numberOfCampaignsSent,omitempty"`
	NumberOfContacts      *int         `json:"numberOfContacts,omitempty"`
	Tags                  *[]TagSchema `json:"tags,omitempty"`
	UniqueId              *string      `json:"uniqueId,omitempty"`
	UpdatedAt             *string      `json:"updated_at,omitempty"`
}

// ContactSchema defines model for ContactSchema.
type ContactSchema struct {
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	CreatedAt  *string                 `json:"created_at,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Phone      *string                 `json:"phone,omitempty"`
	UniqueId   *int                    `json:"uniqueId,omitempty"`
	UpdatedAt  *string                 `json:"updated_at,omitempty"`
}

// ConversationSchema defines model for ConversationSchema.
type ConversationSchema struct {
	ContactId *string    `json:"contactId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Message   *string    `json:"message,omitempty"`
	UniqueId  *string    `json:"uniqueId,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// GetApiKeysResponseSchema defines model for GetApiKeysResponseSchema.
type GetApiKeysResponseSchema struct {
	ApiKeys *[]ApiKeySchema `json:"apiKeys,omitempty"`
}

// GetCampaignResponseSchema defines model for GetCampaignResponseSchema.
type GetCampaignResponseSchema struct {
	Campaigns      *[]CampaignSchema `json:"campaigns,omitempty"`
	PaginationMeta *PaginationMeta   `json:"paginationMeta,omitempty"`
}

// GetUserResponseSchema defines model for GetUserResponseSchema.
type GetUserResponseSchema struct {
	User *UserSchema `json:"user,omitempty"`
}

// LoginRequestBodySchema defines model for LoginRequestBodySchema.
type LoginRequestBodySchema struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginResponseBodySchema defines model for LoginResponseBodySchema.
type LoginResponseBodySchema struct {
	IsOnboardingCompleted *bool   `json:"isOnboardingCompleted,omitempty"`
	Token                 *string `json:"token,omitempty"`
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
	UpdatedAt      *time.Time              `json:"updatedAt,omitempty"`
}

// MessageStatusEnum defines model for MessageStatusEnum.
type MessageStatusEnum string

// MessageTypeEnum defines model for MessageTypeEnum.
type MessageTypeEnum string

// NewCampaignSchema defines model for NewCampaignSchema.
type NewCampaignSchema struct {
	Description        *string      `json:"description,omitempty"`
	EnableLinkTracking *bool        `json:"enableLinkTracking,omitempty"`
	ListIds            *[]string    `json:"listIds,omitempty"`
	Name               *string      `json:"name,omitempty"`
	Tags               *[]TagSchema `json:"tags,omitempty"`
	TemplateMessageId  *string      `json:"templateMessageId,omitempty"`
}

// NewContactListSchema defines model for NewContactListSchema.
type NewContactListSchema struct {
	Description *string      `json:"description,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Tags        *[]TagSchema `json:"tags,omitempty"`
}

// NewContactSchema defines model for NewContactSchema.
type NewContactSchema struct {
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Phone      *string                 `json:"phone,omitempty"`
}

// NewOrganizationMemberSchema defines model for NewOrganizationMemberSchema.
type NewOrganizationMemberSchema struct {
	Email    *string       `json:"email,omitempty"`
	Password *string       `json:"password,omitempty"`
	Role     *UserRoleEnum `json:"role,omitempty"`
	Username *string       `json:"username,omitempty"`
}

// OrganizationMemberSchema defines model for OrganizationMemberSchema.
type OrganizationMemberSchema struct {
	CreatedAt *time.Time    `json:"created_at,omitempty"`
	Role      *UserRoleEnum `json:"role,omitempty"`
	UniqueId  *string       `json:"uniqueId,omitempty"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}

// OrganizationSchema defines model for OrganizationSchema.
type OrganizationSchema struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Name      *string    `json:"name,omitempty"`
	UniqueId  *string    `json:"uniqueId,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// PaginationMeta defines model for PaginationMeta.
type PaginationMeta struct {
	Page    *int64 `json:"page,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
	Total   *int   `json:"total,omitempty"`
}

// SwitchOrganizationResponseSchema defines model for SwitchOrganizationResponseSchema.
type SwitchOrganizationResponseSchema struct {
	Token *string `json:"token,omitempty"`
}

// TagSchema defines model for TagSchema.
type TagSchema struct {
	Name     *string `json:"name,omitempty"`
	UniqueId *string `json:"uniqueId,omitempty"`
}

// UpdateCampaignSchema defines model for UpdateCampaignSchema.
type UpdateCampaignSchema struct {
	Description        *string      `json:"description,omitempty"`
	EnableLinkTracking *bool        `json:"enableLinkTracking,omitempty"`
	ListId             *string      `json:"listId,omitempty"`
	Name               *string      `json:"name,omitempty"`
	Tags               *[]TagSchema `json:"tags,omitempty"`
	TemplateMessageId  *string      `json:"templateMessageId,omitempty"`
}

// UpdateContactListSchema defines model for UpdateContactListSchema.
type UpdateContactListSchema struct {
	Description *string      `json:"description,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Tags        *[]TagSchema `json:"tags,omitempty"`
}

// UpdateContactSchema defines model for UpdateContactSchema.
type UpdateContactSchema struct {
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Phone      *string                 `json:"phone,omitempty"`
}

// UpdateOrganizationMemberSchema defines model for UpdateOrganizationMemberSchema.
type UpdateOrganizationMemberSchema struct {
	Email    *string       `json:"email,omitempty"`
	Password *string       `json:"password,omitempty"`
	Role     *UserRoleEnum `json:"role,omitempty"`
	Username *string       `json:"username,omitempty"`
}

// UserRoleEnum defines model for UserRoleEnum.
type UserRoleEnum string

// UserSchema defines model for UserSchema.
type UserSchema struct {
	CreatedAt               *time.Time            `json:"created_at,omitempty"`
	CurrentOrganizationRole *string               `json:"currentOrganizationRole",omitempty"`
	Email                   *string               `json:"email,omitempty"`
	Name                    *string               `json:"name,omitempty"`
	Organizations           *[]OrganizationSchema `json:"organizations,omitempty"`
	ProfilePicture          *string               `json:"profilePicture,omitempty"`
	UniqueId                *string               `json:"uniqueId,omitempty"`
	UpdatedAt               *time.Time            `json:"updated_at,omitempty"`
	Username                *string               `json:"username,omitempty"`
}

// SwitchOrganizationJSONBody defines parameters for SwitchOrganization.
type SwitchOrganizationJSONBody struct {
	OrganizationId *string `json:"organizationId,omitempty"`
}

// GetCampaignsParams defines parameters for GetCampaigns.
type GetCampaignsParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Order order by asc or desc
	Order *GetCampaignsParamsOrder `form:"order,omitempty" json:"order,omitempty"`

	// Status sort by a field
	Status *CampaignStatusEnum `form:"status,omitempty" json:"status,omitempty"`
}

// GetCampaignsParamsOrder defines parameters for GetCampaigns.
type GetCampaignsParamsOrder string

// DeleteContactsByListParams defines parameters for DeleteContactsByList.
type DeleteContactsByListParams struct {
	// Id contact id/s to be deleted
	Id string `form:"id" json:"id"`
}

// GetContactsParams defines parameters for GetContacts.
type GetContactsParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// ListId query subscribers with a list id.
	ListId *string `form:"list_id,omitempty" json:"list_id,omitempty"`

	// Order order by asc or desc
	Order *string `form:"order,omitempty" json:"order,omitempty"`

	// Status sort by a field
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}

// GetConversationsParams defines parameters for GetConversations.
type GetConversationsParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Order order by asc or desc
	Order *GetConversationsParamsOrder `form:"order,omitempty" json:"order,omitempty"`

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

// GetConversationsParamsOrder defines parameters for GetConversations.
type GetConversationsParamsOrder string

// GetConversationsParamsStatus defines parameters for GetConversations.
type GetConversationsParamsStatus string

// GetContactListsParams defines parameters for GetContactLists.
type GetContactListsParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Order order by asc or desc
	Order *GetContactListsParamsOrder `form:"order,omitempty" json:"order,omitempty"`
}

// GetContactListsParamsOrder defines parameters for GetContactLists.
type GetContactListsParamsOrder string

// GetOrganizationMembersParams defines parameters for GetOrganizationMembers.
type GetOrganizationMembersParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// GetMessagesParams defines parameters for GetMessages.
type GetMessagesParams struct {
	// Page number of records to skip
	Page *int64 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage max number of records to return per page
	PerPage *int64 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Order order by asc or desc
	Order *GetMessagesParamsOrder `form:"order,omitempty" json:"order,omitempty"`

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

// GetMessagesParamsOrder defines parameters for GetMessages.
type GetMessagesParamsOrder string

// GetMessagesParamsStatus defines parameters for GetMessages.
type GetMessagesParamsStatus string

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

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LoginRequestBodySchema

// UpdateOrganizationMemberByIdJSONRequestBody defines body for UpdateOrganizationMemberById for application/json ContentType.
type UpdateOrganizationMemberByIdJSONRequestBody = UpdateOrganizationMemberSchema

// CreateOrganizationMemberJSONRequestBody defines body for CreateOrganizationMember for application/json ContentType.
type CreateOrganizationMemberJSONRequestBody = NewOrganizationMemberSchema

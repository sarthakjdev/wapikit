//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	UniqueId        uuid.UUID `sql:"primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ContactId       uuid.UUID
	OrganizationId  uuid.UUID
	Status          ConversationStatus
	PhoneNumberUsed uuid.UUID
	InitiatedBy     ConversationInitiatedEnum
}
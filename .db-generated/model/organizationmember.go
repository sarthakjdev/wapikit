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

type OrganizationMember struct {
	UniqueId       uuid.UUID `sql:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Role           UserPermissionLevel
	OrganizationId uuid.UUID
	UserId         uuid.UUID
}

//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type UserPermissionLevel string

const (
	UserPermissionLevel_Owner  UserPermissionLevel = "owner"
	UserPermissionLevel_Admin  UserPermissionLevel = "admin"
	UserPermissionLevel_Member UserPermissionLevel = "member"
)

func (e *UserPermissionLevel) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "owner":
		*e = UserPermissionLevel_Owner
	case "admin":
		*e = UserPermissionLevel_Admin
	case "member":
		*e = UserPermissionLevel_Member
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for UserPermissionLevel enum")
	}

	return nil
}

func (e UserPermissionLevel) String() string {
	return string(e)
}

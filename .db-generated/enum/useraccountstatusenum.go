//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var UserAccountStatusEnum = &struct {
	Active    postgres.StringExpression
	Deleted   postgres.StringExpression
	Suspended postgres.StringExpression
}{
	Active:    postgres.NewEnumValue("Active"),
	Deleted:   postgres.NewEnumValue("Deleted"),
	Suspended: postgres.NewEnumValue("Suspended"),
}

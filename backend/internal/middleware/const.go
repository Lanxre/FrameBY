package middleware

import "github.com/lanxre/frameby/internal/types"

type contextKey string

const (
	UserIDKey    contextKey  = "userID"
	UserEmailKey contextKey  = "userEmail"
	UserRoleKey   contextKey = "userRole"
)
const (
	RoleUser      	types.RoleName   = "user"
	RoleStudent   	types.RoleName   = "student"
	RoleBRSM      	types.RoleName   = "brsm"
	RoleUniversity  types.RoleName   = "university"
	RoleCustomer  	types.RoleName   = "customer"
)

var roleWeights = map[types.RoleName]int{
	RoleUser:      1,
	RoleStudent:   2,
	RoleUniversity: 3,
	RoleCustomer:  4,
	RoleBRSM:      5,
}
package constant

type UserStatus string

const (
	UserStatusPending   UserStatus = "PENDING"   // Waiting for email verification or admin approval
	UserStatusActive    UserStatus = "ACTIVE"    // Approved and normal access
	UserStatusRejected  UserStatus = "REJECTED"  // Admin declined initial registration/onboarding
	UserStatusInactive  UserStatus = "INACTIVE"  // Dormant or self-deactivated
	UserStatusLocked    UserStatus = "LOCKED"    // Temporary security lockout (failed logins)
	UserStatusSuspended UserStatus = "SUSPENDED" // Temporarily paused by admin
	UserStatusBanned    UserStatus = "BANNED"    // Permanently blocked by admin
)

func (s UserStatus) CanLogin() bool {
	return s == UserStatusPending || s == UserStatusActive || s == UserStatusInactive
}

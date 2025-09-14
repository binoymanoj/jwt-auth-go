package models

import (
	"gorm.io/gorm"

	"time"
)

type User struct {
	gorm.Model

	// Basic Information
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name"`
	Name      string `json:"name" gorm:"not null"` // combines first_name & last_name
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `json:"-" gorm:"not null"` // Hide password in JSON responses

	// Authentication Fields
	EmailVerified     bool       `json:"email_verified" gorm:"default:false"`
	EmailVerifiedAt   *time.Time `json:"email_verified_at"`
	VerificationToken *string    `json:"-" gorm:"index"` // For email verification

	// Password Reset
	ResetToken       *string    `json:"-" gorm:"index"`
	ResetTokenExpiry *time.Time `json:"-"`

	// Account Security
	IsActive      bool       `json:"is_active" gorm:"default:true"`
	LoginAttempts int        `json:"-" gorm:"default:0"`
	LockedUntil   *time.Time `json:"-"`
	LastLoginAt   *time.Time `json:"last_login_at"`
	LastLoginIP   string     `json:"-"`

	// Two-Factor Authentication
	TwoFactorEnabled bool   `json:"two_factor_enabled" gorm:"default:false"`
	TwoFactorSecret  string `json:"-"` // TOTP secret

	// Session Management
	RefreshToken    string     `json:"-" gorm:"index"`
	RefreshTokenExp *time.Time `json:"-"`

	// Profile Information (optional)
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`

	// Roles and Permissions
	Role string `json:"role" gorm:"default:user"` // user, admin, moderator, etc.
	// Permissions string `json:"-" gorm:"type:text"`       // JSON string of permissions

	// Audit Fields
	LastPasswordChange *time.Time `json:"-"`
	CreatedByIP        string     `json:"-"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

// BeforeCreate hook to set default values
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = "user"
	}
	return nil
}

// Helper methods for the User model
func (u *User) IsAccountLocked() bool {
	return u.LockedUntil != nil && u.LockedUntil.After(time.Now())
}

func (u *User) FullName() string {
	if u.FirstName != "" && u.LastName != "" {
		return u.FirstName + " " + u.LastName
	}
	return u.Name
}

func (u *User) IncrementLoginAttempts() {
	u.LoginAttempts++
	if u.LoginAttempts >= 5 {
		lockDuration := time.Now().Add(15 * time.Minute)
		u.LockedUntil = &lockDuration
	}
}

func (u *User) ResetLoginAttempts() {
	u.LoginAttempts = 0
	u.LockedUntil = nil
	now := time.Now()
	u.LastLoginAt = &now
}

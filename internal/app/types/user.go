package types

import (
	"fmt"
	"strings"
	"time"
)

type (
	UserStatus string
	User       struct {
		ID          string     `json:"id,omitempty" bson:"_id,omitempty"`
		Email       string     `json:"email,omitempty" bson:"email,omitempty"`
		Password    string     `json:"-" bson:"password,omitempty"`
		FirstName   string     `json:"first_name,omitempty" bson:"first_name,omitempty"`
		LastName    string     `json:"last_name,omitempty" bson:"last_name,omitempty"`
		Status      UserStatus `json:"status,omitempty" bson:"status,omitempty"`
		Roles       []string   `json:"roles,omitempty" bson:"roles,omitempty"`
		Groups      []string   `json:"groups,omitempty" bson:"groups,omitempty"`
		Provider    string     `json:"provider,omitempty" bson:"provider,omitempty"`
		Name        string     `json:"name,omitempty" bson:"name,omitempty"`
		NickName    string     `json:"nickname,omitempty" bson:"nickname,omitempty"`
		Description string     `json:"description,omitempty" bson:"description,omitempty"`
		UserID      string     `json:"user_id,omitempty" bson:"user_id,omitempty"`
		AvatarURL   string     `json:"avatar_url,omitempty" bson:"avatar_url,omitempty"`
		Location    string     `json:"location,omitempty" bson:"location,omitempty"`

		CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdateAt  *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}

	RegisterRequest struct {
		Email     string `json:"email,omitempty" validate:"required,email"`
		Password  string `json:"password" validate:"required,gt=3"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	}

	UserInfo struct {
		ID          string `json:"id,omitempty"`
		Email       string `json:"email,omitempty"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Name        string `json:"name,omitempty"`
		NickName    string `json:"nickname,omitempty"`
		Description string `json:"description,omitempty"`
		UserID      string `json:"user_id,omitempty"`
		AvatarURL   string `json:"avatar_url,omitempty"`
		Location    string `json:"location,omitempty"`

		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdateAt  *time.Time `json:"updated_at,omitempty"`
	}
)

const (
	UserStatusActive UserStatus = "active"
	UserStatusLocked UserStatus = "locked"

	ProviderLocal = "local"
)

// FullName return full name in form of first name + space + last name
func (user *User) FullName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

// Strip return a new user without sensitive information
func (user *User) Strip() *User {
	stripedUser := User(*user)
	stripedUser.Password = ""
	return &stripedUser
}

func (user *User) IsLocked() bool {
	return user.Status == UserStatusLocked
}

func (user User) GetName() string {
	if user.Name != "" {
		return user.Name
	}
	if user.FullName() != "" {
		return user.FullName()
	}
	return strings.Split(user.Email, "@")[0]
}

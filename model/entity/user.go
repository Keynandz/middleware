package entity

import (
	"context"
	abstraction "go-authorization/model/base"
	"go-authorization/pkg/util/date"

	"gorm.io/gorm"
)

type MasterUserEntity struct {
	ID             uint16 `json:"id"  gorm:"primaryKey;"`
	Isactive       int32  `json:"isactive"`
	Name           string `gorm:"type:varchar" json:"name"`
	Email          string `gorm:"type:varchar" json:"email"`
	Phone          string `gorm:"type:varchar" json:"phone"`
	DepartmentID   int32  `json:"department_id"`
	RoleID         int32  `json:"role_id"`
	JobtitleID     int32  `json:"jobtitle_id"`
	Username       string `gorm:"type:varchar" json:"username"`
	Password       string `gorm:"type:varchar" json:"password"`
	OauthProvider  string `gorm:"type:varchar" json:"oauth_provider"`
	OauthUid       string `gorm:"type:varchar" json:"oauth_uid"`
	ImageProfile   string `gorm:"type:varchar" json:"image_profile"`
	FormatProfile  string `gorm:"type:varchar" json:"format_profile"`
	WorkingAreaID  int32  `json:"working_area_id"`
	Isarea         int32  `json:"isarea"`
	CategoryAreaID int32  `json:"category_area_id"`
	UserID         string `gorm:"type:varchar" json:"user_id"`
	Approver       int32  `json:"approver"`
}

type MasterUserModel struct {
	// entity
	MasterUserEntity

	// abstraction
	abstraction.Entity

	// context
	Context context.Context `json:"-" gorm:"-"`
}

func (MasterUserModel) TableName() string {
	return "master_user"
}

func (m *MasterUserModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.Created = *date.DateTodayLocal()
	return
}

func (m *MasterUserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.Updated = date.DateTodayLocal()
	return
}

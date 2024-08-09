package entity

import (
	abstraction "go-base-structure/model/base"
	"go-base-structure/pkg/util/date"
	"context"

	"gorm.io/gorm"
)

type FeatureEntity struct {
	ID              uint16  `json:"id"  gorm:"primaryKey;"`
}

type FeatureModel struct {
	// entity
	FeatureEntity

	// abstraction
	abstraction.Entity

	// context
	Context context.Context `json:"-" gorm:"-"`
}

func (FeatureModel) TableName() string {
	return "master_feature"
}

func (m *FeatureModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *FeatureModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = date.DateTodayLocal()
	return
}

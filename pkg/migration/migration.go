package migration

import (
	"fmt"

	dbDriver "go-authorization/pkg/db"
	model "go-authorization/model/entity"
	"go-authorization/pkg/constant"
	"go-authorization/pkg/util/env"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	if !env.NewEnv().GetBool(constant.MIGRATION_ENABLED) {
		return
	}

	mgConfigurations := map[string]Migration{
		constant.DB_BASE_STRUCTURE: &migration{
			DbModels: &[]interface{}{
				model.MasterRoleModel{},
				model.MasterRoleModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := dbDriver.GetConnection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}
}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}

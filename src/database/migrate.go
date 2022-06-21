package gorm

import (
	"log"

	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run database migration",
	RunE:  dbMigrate,
}

// var migUp bool
// var migDown bool

// func init() {
// 	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run miggration up")
// 	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run miggration down")
// }

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{})

	// if migUp {
	// 	if err = m.Migrate(); err != nil {
	// 		return err
	// 	}
	// 	log.Fatal("migration run successful")
	// 	return nil
	// }

	// if migDown {
	// 	if err = m.RollbackLast(); err != nil {
	// 		return err
	// 	}
	// 	log.Fatal("rollback run successful")
	// 	return nil
	// }

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
		)

		if err != nil {
			return err
		}

		return nil
	})

	if err := m.Migrate(); err != nil {
		return err
	}

	log.Fatal("ini schema succesful")
	return nil
}

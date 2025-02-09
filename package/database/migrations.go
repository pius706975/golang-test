package database

import (
	"log"

	"github.com/pius706975/golang-test/package/database/models"

	"github.com/spf13/cobra"
)

var MigrationCMD = &cobra.Command{
	Use:   "migration",
	Short: "DB migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrationCMD.Flags().BoolVarP(&migUp, "up", "u", true, "run migration up")

	MigrationCMD.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {

	db, err := NewDB()
	if err != nil {
		return err
	}

	if migDown {
		log.Println("Migration down done")
		return db.Migrator().DropTable(&models.Role{}, &models.User{}, &models.RefreshToken{}, models.Transaction{})
	}

	if migUp {
		log.Println("Migration up done")
		return db.AutoMigrate(&models.Role{}, &models.User{}, &models.RefreshToken{}, models.Transaction{})
	}

	return nil
}

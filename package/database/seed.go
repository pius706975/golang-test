package database

import (
	"fmt"
	"log"

	seeder "github.com/pius706975/golang-test/package/database/seeders"

	"github.com/pius706975/golang-test/package/database/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type seederData struct {
	name  string
	model interface{}
	size  int
}

var SeedCMD = &cobra.Command{
	Use:   "seed",
	Short: "For running db seeder",
	RunE:  Seed,
}

var seedUP bool
var seedDOWN bool

func init() {
	SeedCMD.Flags().BoolVarP(&seedUP, "seedUP", "u", true, "run seed up")

	SeedCMD.Flags().BoolVarP(&seedDOWN, "seedDOWN", "d", false, "run seed down")
}

func seedDown(db *gorm.DB) error {

	var err error

	transactionSeedModels := []seederData{
		{
			name:  models.Transaction{}.TableName(),
			model: models.Transaction{},
		},
	}

	for _, data := range transactionSeedModels {
		log.Println("Delete seeding data for", data.name)
		sql := fmt.Sprintf("DELETE FROM %v", data.name)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	refreshTokenSeedModels := []seederData{
		{
			name:  models.RefreshToken{}.TableName(),
			model: models.RefreshToken{},
		},
	}

	for _, data := range refreshTokenSeedModels {
		log.Println("Delete seeding data for", data.name)
		sql := fmt.Sprintf("DELETE FROM %v", data.name)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	userSeedModels := []seederData{
		{
			name:  models.User{}.TableName(),
			model: models.User{},
		},
	}

	for _, data := range userSeedModels {
		log.Println("Delete seeding data for", data.name)
		sql := fmt.Sprintf("DELETE FROM %v", data.name)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	var seedModel = []seederData{
		{
			name:  models.Role{}.TableName(),
			model: models.Role{},
		},
	}

	for _, data := range seedModel {
		log.Println("Delete seeding data for", data.name)
		sql := fmt.Sprintf("DELETE FROM %v", data.name)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedUp(db *gorm.DB) error {

	var err error

	var seedModel = []seederData{
		{
			name:  "roles",
			model: seeder.RoleSeed,
			size:  cap(seeder.RoleSeed),
		},

		{
			name:  "users",
			model: seeder.UserSeed,
			size:  cap(seeder.UserSeed),
		},

		{
			name:  "transactions",
			model: seeder.TransactionSeed,
			size:  cap(seeder.TransactionSeed),
		},
	}

	for _, data := range seedModel {
		log.Println("create seeding data for", data.name)
		err = db.CreateInBatches(data.model, data.size).Error
	}

	return err
}

func Seed(cmd *cobra.Command, args []string) error {

	var err error

	db, err := NewDB()
	if err != nil {
		return err
	}

	if seedDOWN {
		err = seedDown(db)
		return err
	}

	if seedUP {
		err = seedUp(db)
	}

	return err
}

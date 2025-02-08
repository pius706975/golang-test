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
	Use: "seed",
	Short: "For running db seeder",
	RunE: Seed,
}

var seedUP bool
var seedDOWN bool

func init()  {
	SeedCMD.Flags().BoolVarP(&seedUP, "seedUP", "u", true, "run seed up")

	SeedCMD.Flags().BoolVarP(&seedDOWN, "seedDOWN", "d", false, "run seed down")
}

func seedDown(db *gorm.DB) error {
	
	var err error

	var seedModel = []seederData{
		{
			name: models.Role{}.TableName(),
			model: models.Role{},
		},
	}

	for _, data := range seedModel {
		log.Println("Delete seeding data for", data.name)
		sql := fmt.Sprintf("DELETE FROM %v ", data.name)
		err = db.Exec(sql).Error
	}

	return err
}

func seedUp(db *gorm.DB) error {
	
	var err error

	var seedModel = []seederData{
		{
			name: "roles",
			model: seeder.RoleSeed,
			size: cap(seeder.RoleSeed),
		},
	}

	for _, data := range seedModel {

		log.Println("create seeding data for ", data.name)
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
package seeders

import (
	"time"

	"github.com/pius706975/golang-test/package/database/models"
)

var RoleSeed = models.Roles{
	{
		ID:        "73b1d7d7-b319-47f8-ab47-8c92627b13bd",
		Name:      "admin",
		CreatedAt: time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),
	},
	{
		ID:        "f4e1855f-80a2-4ee5-a1ec-e80a9a3d3648",
		Name:      "user",
		UpdatedAt: time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),
	},
}

package seeders

import (
	"time"

	"github.com/pius706975/golang-test/package/database/models"
)

var UserSeed = models.Users{
	// User password is User@123
	{
		RoleID: "73b1d7d7-b319-47f8-ab47-8c92627b13bd",
		Name: "Jagon Mouse",
		Username: "jagon9953220",
		Email: "jagon@gmail.com",
		Password: "$2a$10$JokQO.zw7EEVPf/w9Hbm.eouMpxkI.lv.mnVCnWLJhipL6NYIl16G",
		PhoneNumber: "6285153639797",
		IsVerified: true,
		CreatedAt:     time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),

	},

	{
		ID: "2956d89d-51eb-488a-a865-f4ec8b198c62",
		RoleID: "f4e1855f-80a2-4ee5-a1ec-e80a9a3d3648",
		Name: "Gerald Jinx Mouse",
		Username: "gerald9953220",
		Email: "jerry@gmail.com",
		Password: "$2a$10$JokQO.zw7EEVPf/w9Hbm.eouMpxkI.lv.mnVCnWLJhipL6NYIl16G",
		PhoneNumber: "6285153639797",
		IsVerified: true,
		CreatedAt:     time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),

	},

	{
		ID: "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		RoleID: "f4e1855f-80a2-4ee5-a1ec-e80a9a3d3648",
		Name: "Thomas Jasper Cat",
		Username: "thomas9953220",
		Email: "tom@gmail.com",
		Password: "$2a$10$JokQO.zw7EEVPf/w9Hbm.eouMpxkI.lv.mnVCnWLJhipL6NYIl16G",
		PhoneNumber: "6285153639797",
		IsVerified: true,
		CreatedAt:     time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 10, 54, 11, 446308000, time.FixedZone("UTC+8", 8*60*60)),

	},
}
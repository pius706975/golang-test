package seeders

import "github.com/pius706975/golang-test/package/database/models"

var RoleSeed = models.Roles{
	{
		ID: "73b1d7d7-b319-47f8-ab47-8c92627b13bd",
		Name: "admin",
	},
	{
		ID: "f4e1855f-80a2-4ee5-a1ec-e80a9a3d3648",
		Name: "user",
	},
}
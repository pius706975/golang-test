package seeders

import (
	"time"

	"github.com/pius706975/golang-test/package/database/models"
	"gorm.io/datatypes"
)

var TransactionSeed = models.Transactions{
	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        15000,
		Status:        "pending",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:42:07.150957637+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 11, 42, 7, 150956745, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 42, 7, 151937663, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		Amount:        10000,
		Status:        "pending",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:42:07.150957637+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 11, 42, 7, 150956745, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 42, 7, 151937663, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        20000,
		Status:        "failed",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:40:00+08:00"}, {"status":"failed","timestamp":"2025-02-09T11:45:00+08:00"}]`),
		PaymentMethod: "ovo",
		CreatedAt:     time.Date(2025, 2, 9, 11, 40, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 45, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		Amount:        25000,
		Status:        "failed",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:30:00+08:00"}, {"status":"failed","timestamp":"2025-02-09T11:35:00+08:00"}]`),
		PaymentMethod: "dana",
		CreatedAt:     time.Date(2025, 2, 9, 11, 30, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 35, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        18000,
		Status:        "failed",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:20:00+08:00"}, {"status":"failed","timestamp":"2025-02-09T11:25:00+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 11, 20, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 25, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		Amount:        15000,
		Status:        "success",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:10:00+08:00"}, {"status":"success","timestamp":"2025-02-09T11:15:00+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 11, 10, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 15, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        20000,
		Status:        "success",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:22:07+08:00"},{"status":"success","timestamp":"2025-02-09T11:25:07+08:00"}]`),
		PaymentMethod: "dana",
		CreatedAt:     time.Date(2025, 2, 9, 11, 22, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 25, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		Amount:        50000,
		Status:        "success",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:12:07+08:00"},{"status":"success","timestamp":"2025-02-09T11:15:07+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 11, 12, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 15, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        15000,
		Status:        "pending",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:42:07+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 11, 42, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 42, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		Amount:        10000,
		Status:        "pending",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:32:07+08:00"}]`),
		PaymentMethod: "ovo",
		CreatedAt:     time.Date(2025, 2, 9, 11, 32, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 32, 7, 0, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        20000,
		Status:        "pending",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:22:07.150957637+08:00"}]`),
		PaymentMethod: "ovo",
		CreatedAt:     time.Date(2025, 2, 9, 11, 22, 7, 150956745, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 22, 7, 151937663, time.FixedZone("UTC+8", 8*60*60)),
	},
	
	{
		UserID:        "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		Amount:        50000,
		Status:        "success",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:12:07.150957637+08:00"}, {"status":"success","timestamp":"2025-02-09T11:14:07.150957637+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 11, 12, 7, 150956745, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 14, 7, 151937663, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        75000,
		Status:        "success",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T11:02:07.150957637+08:00"}, {"status":"success","timestamp":"2025-02-09T11:04:07.150957637+08:00"}]`),
		PaymentMethod: "dana",
		CreatedAt:     time.Date(2025, 2, 9, 11, 2, 7, 150956745, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 11, 4, 7, 151937663, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "5244cf96-67f0-49d8-9e0a-31c02f244f24",
		Amount:        12000,
		Status:        "pending",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T10:52:07.150957637+08:00"}]`),
		PaymentMethod: "ovo",
		CreatedAt:     time.Date(2025, 2, 9, 10, 52, 7, 150956745, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 10, 52, 7, 151937663, time.FixedZone("UTC+8", 8*60*60)),
	},

	{
		UserID:        "2956d89d-51eb-488a-a865-f4ec8b198c62",
		Amount:        30000,
		Status:        "pending",
		StatusHistory: datatypes.JSON(`[{"status":"pending","timestamp":"2025-02-09T10:42:07.150957637+08:00"}]`),
		PaymentMethod: "gopay",
		CreatedAt:     time.Date(2025, 2, 9, 10, 42, 7, 150956745, time.FixedZone("UTC+8", 8*60*60)),
		UpdatedAt:     time.Date(2025, 2, 9, 10, 42, 7, 151937663, time.FixedZone("UTC+8", 8*60*60)),
	},
}

package mocks

import "timesheet-be/pkg/models"

var Timesheets = []models.Timesheet{
	{
		ID:            1,
		Date:          "2022-01-25",
		WorkingStart:  "15:00",
		WorkingEnd:    "16:00",
		OvertimeStart: "00:00",
		OvertimeEnd:   "00:00",
		Activity:      "Membuat dashboard utama",
		ProjectID:     1,
		StatusID:      1,
	},
}

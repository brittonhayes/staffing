package inmem

import "github.com/brittonhayes/staffing"

var (
	sara = &staffing.Employee{
		ID:   "9d12784e-81d1-4782-9bb3-4170533fb484",
		Name: "Sara",
	}

	john = &staffing.Employee{
		ID:   "b5b5b5b5-b5b5-b5b5-b5b5-b5b5b5b5b5b5",
		Name: "29439877-92b9-41af-8fab-de8776d4b05e",
	}

	trish = &staffing.Employee{
		ID:   "52928ed0-8122-4a72-a5eb-15d87dc759f8",
		Name: "Trish",
	}
)

var (
	marketing = &staffing.Project{
		ID:   "bbbeffe6-d0ee-46d2-961e-f20c6f053b89",
		Name: "Marketing Campaign",
	}

	cloudMigration = &staffing.Project{
		ID:   "7b67e136-5c96-43fc-b1b7-6f2d15c6f5b7",
		Name: "Cloud Migration",
	}
)

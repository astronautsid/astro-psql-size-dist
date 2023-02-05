package info

type (
	Table struct {
		Name string `json:"name"`
		Size uint64 `json:"size"`
	}

	Label struct {
		Name   string  `json:"name"`
		Tables []Table `json:"tables"`
	}
)

package config

type (
	Config struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		DbName   string `yaml:"dbname"`
	}

	LabelToTableConnection struct {
		Name   string   `json:"name"`
		Tables []string `json:"tables"`
	}

	Rule struct {
		Labels []LabelToTableConnection `json:"labels"`
	}
)

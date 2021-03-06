package model

type Config struct {
	Server struct {
		Port       string `yaml:"port"`
		Host       string `yaml:"host"`
		JwrExpires int    `yaml:"jwtExprires"`
	}

	Database struct {
		DbName     string `yaml:"dbName"`
		DbHost     string `yaml:"dbHost"`
		DbPort     string `yaml:"dbPort"`
		DbUserName string `yaml:"dbUserName"`
		DbPassword string `yaml:"dbPassword"`
	}
}

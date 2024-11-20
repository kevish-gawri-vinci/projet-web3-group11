package database

type DbVariables struct {
	Username string
	DbName   string
	Password string
	Hostname string
	Port     string
}

func GetVariables() *DbVariables {
	dbVariable := DbVariables{
		Username: "",
		DbName:   "web3_project_db",
		Password: "",
		Hostname: "localhost",
		Port:     "5432",
	}
	return &dbVariable
}

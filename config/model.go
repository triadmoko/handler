package config

// connect database for postgresql
type Database struct {
	// host database
	//
	// example : 127.0.0.1
	Host string
	/*
		user database example : postgre
	*/
	User string
	// name database
	Name string
	// password database
	Password string
	Port     string
}

package main

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "workdb"
)

func main() {

	a := App{}
	a.Initialize(user, password, dbname)
	a.Run(":8010")
}

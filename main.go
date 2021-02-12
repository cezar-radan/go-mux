package main

var (
	user, password, dbname string = "postgres", "postgres", "workdb"
)

func main() {
	a := App{}
	a.Initialize(user, password, dbname)

	a.Run(":8010")
}

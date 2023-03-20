package main

func main() {
	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ride-kaki")
	// if err != nil {
	//     panic(err.Error())
	// }
	// defer db.Close()

	// fmt.Println("Hello, Modules!")

	InitRoutes()
}

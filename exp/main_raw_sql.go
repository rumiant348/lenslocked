package main

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//)
//
//const (
//	host     = "localhost"
//	port     = 5432
//	userName = "aru"
//	password = ""
//	dbname   = "lenslocked_dev"
//)
//
//var db *sql.DB
//
//func main() {
//
//	defer tearDown()
//	createTables()
//}
//
//func dropTables() {
//	_, err := db.Exec(`DROP TABLE users`)
//	if err != nil {
//		panic(err)
//	}
//	_, err = db.Exec(`DROP TABLE orders`)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func createTables() {
//	_, err := db.Exec(`
//	CREATE TABLE users (
// id SERIAL PRIMARY KEY,
// name TEXT,
// email TEXT NOT NULL
//)`)
//	if err != nil {
//		panic(err)
//	}
//
//	_, err = db.Exec(`
//	CREATE TABLE orders (
// id SERIAL PRIMARY KEY,
// user_id INT NOT NULL,
// amount INT,
// description TEXT
//)`)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func populateRelatedData() {
//	var id int
//	for i := 1; i < 6; i++ {
//		// Create some fake data
//		userId := 17
//		if i > 3 {
//			userId = 19
//		}
//		amount := 1000 * i
//		description := fmt.Sprintf("USB-C Adapter x%d", i)
//
//		err := db.QueryRow(`
//     		INSERT INTO orders (user_id, amount, description)
//   		VALUES ($1, $2, $3)
//    		RETURNING id`,
//			userId, amount, description).Scan(&id)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("Created an order with the ID:", id)
//	}
//}
//
//func getRelated() {
//	rows, err := db.Query(
//		`SELECT users.id, users.email, users.name,
//		orders.id AS order_id,
//		orders.amount AS order_amount,
//		orders.description AS order_description
//	FROM users
//	INNER JOIN orders
//	ON users.id = orders.user_id;`)
//	if err != nil {
//		panic(err)
//	}
//	for rows.Next() {
//		var user user
//		var order order
//
//		err := rows.Scan(&user.id, &user.email, &user.name,
//			&order.id, &order.amount, &order.description)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("%+v %+v\n", user, order)
//	}
//}
//
//type user struct {
//	id    int
//	name  string
//	email string
//}
//
//type order struct {
//	id          int
//	amount      int
//	description string
//}
//
//func init() {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"dbname=%s sslmode=disable password=%s ",
//		host, port, userName, dbname, password)
//
//	var err error
//	db, err = sql.Open("postgres", psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Successfully connected!")
//}
//
//func tearDown() {
//	err := db.Close()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Tear down successful")
//}
//
//// insert a row
//func insertRow(name, email string) int {
//	var id int
//	row := db.QueryRow(`
// 	INSERT INTO users(name, email)
// 	VALUES($1, $2) RETURNING id`,
//		name, email)
//	err := row.Scan(&id)
//	if err != nil {
//		panic(err)
//	}
//	return id
//}
//
//// get Row gets a single record of a &User
//func getRow(id int) int {
//	var name, email string
//	row := db.QueryRow(`
//		SELECT id, name, email
//		FROM users
//		WHERE id=$1`, id)
//
//	err := row.Scan(&id, &name, &email)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("ID: %d Name: %s Email: %s\n", id, name, email)
//
//	return id
//}
//
//func getRows() []user {
//	rows, err := db.Query(`
//		SELECT id, name, email
//		FROM users
//		WHERE email=$1`, "jon@calhoun.io")
//	if err != nil {
//		panic(err)
//	}
//	var users []user
//	for rows.Next() {
//		var user user
//		err := rows.Scan(&user.id, &user.name, &user.email)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("ID:", user.id, "Name:", user.name, "Email:", user.email)
//		users = append(users, user)
//	}
//	return users
//}
//
//func deleteRow(id int) {
//	var deletedId int
//	err := db.QueryRow(`
//		DELETE FROM users
//		WHERE id=$1
//		RETURNING id`, id).Scan(&deletedId)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Deleted with id: %d\n", deletedId)
//}

package main

//
//import (
//	"bufio"
//	"fmt"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/postgres"
//	"os"
//	"strings"
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
//var db *gorm.DB
//
//func init() {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, userName, password, dbname)
//
//	var err error
//	db, err = gorm.Open("postgres", psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//
//	db.LogMode(true)
//
//	//err = db..Ping()
//	//if err != nil {
//	//	panic(err)
//	//}
//	fmt.Println("Successfully connected!")
//
//	db.AutoMigrate(&user{}, &order{})
//}
//
//type user struct {
//	gorm.Model
//	Name   string
//	Email  string `gorm:"not null;unique_index"`
//	Orders []order
//}
//
//type order struct {
//	gorm.Model
//	UserID      uint
//	Amount      int
//	Description string 1
//}
//
//func main() {
//	defer db.Close()
//	findOrdersForUser()
//}
//
//func findOrdersForUser() {
//	var user user
//	// ищет пользователя а потом все его заказы, ыыы
//	db.Preload("Orders").First(&user)
//	if db.Error != nil {
//		panic(db.Error)
//	}
//	fmt.Println("Email:", user.Email)
//	fmt.Println("Number of orders:", len(user.Email))
//	fmt.Println("Orders:", user.Orders)
//}
//
//func populateOrders() {
//	user := getFirstUser()
//	createOrder(db, user, 1001, "Fake Description #1")
//	createOrder(db, user, 9999, "Fake Description #2")
//	createOrder(db, user, 8800, "Fake Description #3")
//}
//
//func createOrder(db *gorm.DB, user user, amount int, desc string) {
//	order := order{
//		UserID:      user.ID,
//		Amount:      amount,
//		Description: desc,
//	}
//	db.Create(&order)
//	if db.Error != nil {
//		panic(db.Error)
//	}
//}
//
//func findAll() {
//	var users []user
//	db.Find(&users)
//	if db.Error != nil {
//		panic(db.Error)
//	}
//	fmt.Println("Retrieved", len(users), "users.")
//	fmt.Println(users)
//}
//
//func where() {
//	var u user
//	maxId := 3
//	db.Where("id <= ?", maxId).First(&u)
//	if db.Error != nil {
//		panic(db.Error)
//	}
//	fmt.Println(u)
//}
//
//func userQuery() {
//	var u user
//	u.Email = "jon@calhoun.com"
//	db.Where(u).First(&u)
//	if db.Error != nil {
//		panic(db.Error)
//	}
//	fmt.Println(u)
//}
//
//func getFirstUser() user {
//	var u user
//	db.First(&u)
//	if db.Error != nil {
//		panic(db.Error)
//	}
//	fmt.Println(u)
//	return u
//}
//
//func create(name, email string) {
//	u := &user{
//		Name:  name,
//		Email: email,
//	}
//	if err := db.Create(u).Error; err != nil {
//		panic(err)
//	}
//	fmt.Printf("%+v\n", u)
//}
//
//func getInfo() (name, email string) {
//	reader := bufio.NewReader(os.Stdin)
//
//	name = getFromUser("What is your name?", reader)
//	email = getFromUser("What is your email?", reader)
//
//	return name, email
//}
//
//func getFromUser(question string, reader *bufio.Reader) string {
//	fmt.Println(question)
//	res, _ := reader.ReadString('\n')
//	res = strings.TrimSpace(res)
//	return res
//}

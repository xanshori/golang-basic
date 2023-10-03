package database

import "fmt"

var conn string
// init dipanggil setiap kali package digunakan
func init() {
	fmt.Println("init dipanggil")
	conn = "Mysql"
}

func GetDatabase() string {
	return conn
}
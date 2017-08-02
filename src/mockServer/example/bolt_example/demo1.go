package main

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
)

// 校验错误
func checkerr(e error) {
	if e != nil {
		fmt.Println("error:", e)
		panic(e)
	}
}

type User struct {
	ID    int    // primary key
	Group string `storm:"index"`  // this field will be indexed
	Email string `storm:"unique"` // this field will be indexed with a unique constraint
	Name  string // this field will not be indexed
	Age   int    `storm:"index"`
	// CreatedAt
}

func main() {
	// 链接数据库
	db, err := storm.Open("temp/temp.db", storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
	// 数据库关闭
	defer db.Close()
	checkerr(err)

	user := User{
		ID:    10,
		Group: "staff",
		Email: "john@provider.com",
		Name:  "John",
		Age:   21,
		// CreatedAt: time.Now().UnixNano(),
	}

	err1 := db.Save(&user)
	checkerr(err1)
	// err == nil

	// user.ID++
	// err = db.Save(&user)
	// checkerr(err)

	var user1 User
	err2 := db.One("Email", "john@provider.com", &user1)
	checkerr(err2)
	// fmt.Println(err2)
}

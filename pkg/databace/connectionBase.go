package databace

import (
	"database/sql"
	"fmt"
	  "github.com/Tayduro/registration-web-server/pkg/models"
	_ "github.com/lib/pq"
)

func DataB(users * models.User) {
	user := models.User{
		FirstName: users.FirstName,
		LastName: users.LastName,
		Email: users.Email,
		Password: users.Password,
	}

	psqlconn := "postgres://postgres:12345@localhost:6080/users?sslmode=disable"

	db, err := sql.Open("postgres", psqlconn)
	if err != nil{
		panic(err)
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (first_name,last_name,email,password) VALUES($1, $2, $3, $4)", &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil{
		panic(err)
	}
	defer insert.Close()

	fmt.Println("connect to server...")
}

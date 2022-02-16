package databace

import (
	"fmt"
	"github.com/Tayduro/registration-web-server/pkg/config"
	jwtToken "github.com/Tayduro/registration-web-server/pkg/jwt"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DataBaseRegistration(users *models.User) {
	currentUser := models.User{
		FirstName: users.FirstName,
		LastName:  users.LastName,
		Email:     users.Email,
		Password:  users.Password,
	}

	connstring := config.ConfigServer()

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	salt := service.RandomString()

	hash := service.CreatingHash(salt, currentUser.Password)

	db, err = sqlx.Connect("postgres", connstring)

	defer db.Close()

	insert, err := db.Queryx("INSERT INTO users (first_name,last_name,email) VALUES($1, $2, $3)", &currentUser.FirstName, &currentUser.LastName, &currentUser.Email)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	var dbUserId string
	err = db.QueryRow("SELECT user_id FROM users WHERE email= $1", currentUser.Email).Scan(&dbUserId)
	if err != nil {
		panic(err)
	}

	insert, err = db.Queryx("INSERT INTO credentials (user_id,salt ,hash) VALUES($1, $2, $3)", dbUserId, salt, hash)
	if err != nil {
		panic(err)
	}

	defer insert.Close()

	fmt.Println("connect to server...")
}

func Login(user *models.User) string {

	connstring := config.ConfigServer()

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var dbUserId string
	err = db.QueryRow("SELECT user_id FROM users WHERE email= $1", user.Email).Scan(&dbUserId)
	if err != nil {
		panic(err)
	}

	var dbSalt string
	err = db.QueryRow("SELECT salt FROM credentials WHERE user_id= $1", dbUserId).Scan(&dbSalt)
	if err != nil {
		panic(err)
	}

	hash := service.CreatingHash(dbSalt, user.Password)

	var dbHash string
	err = db.QueryRow("SELECT hash FROM credentials WHERE user_id= $1", dbUserId).Scan(&dbHash)
	if err != nil {
		panic(err)
	}
	var token string
	if dbHash == hash {

		var hmacSampleSecret = []byte(config.GetKey())
		//token := jwtToken.NewHmac(dbUserId, hmacSampleSecret)
		token = jwtToken.NewJWT(dbUserId, hmacSampleSecret)

		insert, err := db.Queryx("INSERT INTO access_token (user_id,token) VALUES($1, $2)", dbUserId, token)
		if err != nil {
			panic(err)
		}
		defer insert.Close()

	}
	return token
}

func DeleteToken(token string) {
	connstring := config.ConfigServer()

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	drop, err := db.Query("delete from access_token where token = $1", token)
	if err != nil {
		panic(err)
	}
	defer drop.Close()

}

func databaseData(token string) []models.UserData {

	userData := make([]models.UserData, 0, 0)

	if len(token) == 0 {
		return userData
	}

	connstring := config.ConfigServer()

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var dbUserId string
	err = db.QueryRow("SELECT user_id FROM access_token WHERE token= $1", token).Scan(&dbUserId)
	if err != nil {
		panic(err)
	}

	var dbFirstName string
	err = db.QueryRow("SELECT first_name FROM users WHERE user_id= $1", dbUserId).Scan(&dbFirstName)
	if err != nil {
		panic(err)
	}

	var dbLastName string
	err = db.QueryRow("SELECT last_name FROM users WHERE user_id= $1", dbUserId).Scan(&dbLastName)
	if err != nil {
		panic(err)
	}

	userData = append(userData, models.UserData{
		Field:      "FirstName",
		FieldValue: dbFirstName,
	})
	userData = append(userData, models.UserData{
		Field:      "LastName",
		FieldValue: dbLastName,
	})
	return userData
}


//func GettingUserData(token string) []models.UserData {
//
//	hmacSampleSecret := []byte(config.GetKey())
//	if jwtToken.ParseHmac(token, hmacSampleSecret) == nil {
//
//		return DatabaseData(token)
//	}
//
//	return DatabaseData("")
//}

func GettingUserData(token string) []models.UserData {

	var hmacSampleSecret = []byte(config.GetKey())

	//jwtToken.ParseHmac(token, hmacSampleSecret)
	if jwtToken.ParseHmac(token, hmacSampleSecret) == nil {
		return databaseData(token)
	}

	return databaseData("")
}

func GetUserId(token string) string {

	connstring := config.ConfigServer()
	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var dbUserId string
	err = db.QueryRow("SELECT user_id FROM access_token WHERE token= $1", token).Scan(&dbUserId)
	if err != nil {
		panic(err)
	}
  return dbUserId
}





func GetUserFirstName(UserId string) string {
	connstring := config.ConfigServer()
	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var dbFirstName string
	err = db.QueryRow("SELECT first_name FROM users WHERE user_id= $1", UserId).Scan(&dbFirstName)
	if err != nil {
		panic(err)
	}

	return dbFirstName

}

func GetUserLastName(UserId string) string {
	connstring := config.ConfigServer()
	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var dbLastName string
	err = db.QueryRow("SELECT last_name FROM users WHERE user_id= $1", UserId).Scan(&dbLastName)
	if err != nil {
		panic(err)
	}

	return dbLastName

}




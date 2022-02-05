package databace

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/Tayduro/registration-web-server/pkg/config"
	"github.com/Tayduro/registration-web-server/pkg/models"
)

const (
	Host     = "localhost"
	Port     = 6080
	User     = "postgres"
	Password = "12345"
	Dbname   = "users"
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

	insert, err := db.Queryx("INSERT INTO users (first_name,last_name,email) VALUES($1, $2, $3)", &currentUser.FirstName, &currentUser.LastName, &currentUser.Email)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	salt := RandStringRunes(5)
	newPassword := fmt.Sprintf("%s%s", salt, currentUser.Password)
	hashBits := sha256.Sum256([]byte(newPassword))
	hash := fmt.Sprintf("%x", hashBits)

	db, err = sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	defer db.Close()

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

	db, err = sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//var dbFirstName string
	//err = db.QueryRow("SELECT first_name FROM users WHERE email= $1", currentUser.Email).Scan(&dbFirstName)
	//if err != nil {
	//	panic(err)
	//}
	//
	//var dbLastName string
	//err = db.QueryRow("SELECT last_name FROM users WHERE email= $1", currentUser.Email).Scan(&dbLastName)
	//if err != nil {
	//	panic(err)
	//}


	var hmacSampleSecret = []byte(GetKey())
	tocken := ExamplenewHmac(dbUserId, hmacSampleSecret)

	insert, err = db.Queryx("INSERT INTO access_token (user_id,token) VALUES($1, $2)", dbUserId, tocken)
	if err != nil {
		panic(err)
	}

	defer insert.Close()




	fmt.Println("connect to server...")
}


func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


func ExamplenewHmac(userId string, hmacSampleSecret []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"userId ": userId,
		//"nbf": time.Date(2022, 1, 31, 23, 0, 0, 0, time.UTC).Unix(),
		"iat": 1516239022,
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)

	// Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU <nil>
	return tokenString
}

func ExampleParse_hmac(tokenString string, hmacSampleSecret []byte) error {
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})


	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		fmt.Println(err)
		return err
	}

}


func GetKey() string {
	yfile, err := ioutil.ReadFile("./cmd/signup-server/config.yaml")

	if err != nil {

		log.Fatal(err)
	}

	conf := *&config.Config{}

	err = yaml.Unmarshal(yfile, &conf)

	if err != nil {

		log.Fatal(err)
	}
	return conf.Key
}

func GettingUserData(tocken string) string {

	var hmacSampleSecret = []byte(GetKey())


	ExampleParse_hmac(tocken, hmacSampleSecret)
	if(ExampleParse_hmac(tocken, hmacSampleSecret) == nil){
		return databaseData(tocken)
		fmt.Println("ololo")
	}

	return "errorGettingUserData"
}

func databaseData(token string) string {
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

	fmt.Println(dbUserId)
	fmt.Println(dbFirstName)
	fmt.Println(dbLastName)

	result := dbFirstName + " " + dbLastName
	return result
}

func DeleteToken(token string)  {
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
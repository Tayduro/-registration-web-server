package databace

import (
	"crypto/sha256"
	"fmt"
<<<<<<< HEAD
	jwtToken "github.com/Tayduro/registration-web-server/pkg/jwt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strings"
	"time"

	"github.com/Tayduro/registration-web-server/pkg/config"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"math/rand"
=======
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/Tayduro/registration-web-server/pkg/config"
	"github.com/Tayduro/registration-web-server/pkg/models"
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
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

<<<<<<< HEAD
	//defer db.Close()

	_, err = db.Queryx("INSERT INTO users (first_name,last_name,email) VALUES($1, $2, $3)", &currentUser.FirstName, &currentUser.LastName, &currentUser.Email)
	if err != nil {
		panic(err)
	}
	//defer insert.Close()

	salt := RandomString()

	hash := creatingHash (salt,currentUser.Password)

	db, err = sqlx.Connect("postgres", connstring)

=======
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
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	//defer db.Close()

	var dbUserId string
	err = db.QueryRow("SELECT user_id FROM users WHERE email= $1", currentUser.Email).Scan(&dbUserId)
	if err != nil {
		panic(err)
	}

	_, err = db.Queryx("INSERT INTO credentials (user_id,salt ,hash) VALUES($1, $2, $3)", dbUserId, salt, hash)
	if err != nil {
		panic(err)
	}

	//defer insert.Close()
=======
	defer insert.Close()



>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3

	fmt.Println("connect to server...")
}


<<<<<<< HEAD

func Login(user *models.User) string {

	connstring :=config.ConfigServer()

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	//defer db.Close()

	var dbUserId string
	err = db.QueryRow("SELECT user_id FROM users WHERE email= $1", user.Email).Scan(&dbUserId)
	if err != nil{
		panic(err)
	}

	var dbSalt string
	err = db.QueryRow("SELECT salt FROM credentials WHERE user_id= $1", dbUserId).Scan(&dbSalt)
	if err != nil{
		panic(err)
	}

	hash := creatingHash (dbSalt,user.Password)

	var dbHash string
	err = db.QueryRow("SELECT hash FROM credentials WHERE user_id= $1", dbUserId).Scan(&dbHash)
	if err != nil{
		panic(err)
	}
	var token string
	if dbHash == hash {

		var hmacSampleSecret = []byte(config.GetKey())
		//token := jwtToken.NewHmac(dbUserId, hmacSampleSecret)
		token = jwtToken.NewJWT(dbUserId, hmacSampleSecret)

		_, err = db.Queryx("INSERT INTO access_token (user_id,token) VALUES($1, $2)", dbUserId, token)
		if err != nil {
			panic(err)
		}
		//defer insert.Close()

	}
	return token
}





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func creatingHash(dbSalt string, userPassword string) string {
	password := dbSalt + userPassword
	hashBits := sha256.Sum256([]byte(password))
	hash := fmt.Sprintf("%x", hashBits)

	return hash
}

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()

	return str
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////



func GettingUserData(token string) []models.UserData {

	var hmacSampleSecret = []byte(config.GetKey())


	//jwtToken.ParseHmac(token, hmacSampleSecret)
	if(jwtToken.ParseHmac(token, hmacSampleSecret) == nil){
		return databaseData(token)
	}


	return databaseData("")
}

func databaseData(token string) []models.UserData {
	userData := make([]models.UserData, 0, 0)
	if len(token) == 0{
		return userData
	}
		connstring := config.ConfigServer()
=======
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
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	//defer db.Close()
=======
	defer db.Close()
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3

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

<<<<<<< HEAD
	userData = append(userData, models.UserData{
		Field: "FirstName",
		FieldValue: dbFirstName,

	})
	userData = append(userData, models.UserData{
		Field: "LastName",
		FieldValue: dbLastName,

	})
	return userData
=======
	fmt.Println(dbUserId)
	fmt.Println(dbFirstName)
	fmt.Println(dbLastName)

	result := dbFirstName + " " + dbLastName
	return result
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
}

func DeleteToken(token string)  {
	connstring := config.ConfigServer()

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	//defer db.Close()

	_, err = db.Query("delete from access_token where token = $1", token)
	if err != nil {
		panic(err)
	}
	//defer drop.Close()

}


=======
	defer db.Close()

	drop, err := db.Query("delete from access_token where token = $1", token)
	if err != nil {
		panic(err)
	}
	defer drop.Close()

}
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3

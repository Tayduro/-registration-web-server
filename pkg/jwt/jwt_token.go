package jwtToken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)


type Claims struct {
	UserId string
	jwt.StandardClaims
}


func NewJWT(userId string, hmacSampleSecret []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		},
		UserId: userId,

	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return tokenString
}



func ParseHmac(tokenString string, hmacSampleSecret []byte) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})


	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["userId"])
		//result := fmt.Sprintln(claims["ExpiresAt"])
		//checkTimeValidToken(result)

		//fmt.Println(result)
		return nil
	} else {
		fmt.Println(err)
		return err
	}

}

func NewHmac(userId string, hmacSampleSecret []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"userId": userId,
		"ExpiresAt": fmt.Sprintln(time.Now().Add(time.Second * time.Duration(12)).Unix()),
		"iat": fmt.Sprintln(time.Now().Unix()),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}


	return tokenString
}

//func checkTimeValidToken(timeToken string) string {
//
//	timeNow := fmt.Sprintln(time.Now().Unix())
//
//	if timeNow > timeToken {
//		fmt.Println("ok")
//		return "ok"
//	} else {
//		fmt.Println("Not ok")
//	}
//
//	fmt.Println(timeToken, "timeToken")
//	fmt.Println(timeNow, "timeNow")
//
//	return ""
//}







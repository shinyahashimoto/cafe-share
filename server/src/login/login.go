package login

import (
	"errors"
	"log"
	"net/http"
	"server/auth"
	"server/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Name     string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Param struct {
	Id       int    `json:"id"`
	Name     string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// 要修正
	Token string `json:"token"`
}

func Login(c echo.Context) error {
	param := new(User)
	if err := c.Bind(param); err != nil {
		return err
	}

	db := db.DbConnect()
	defer db.Close()
	//パスワードのハッシュ化
	// hashPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), 12)

	// メールアドレスの重複がないか確認

	// shopをIDを元に検索
	user := User{}
	db.Where("name = ?", "jinzhu").First(&user)
	result := db.Debug().Where("email = ?", param.Email).First(&user)
	err := errors.Is(result.Error, gorm.ErrRecordNotFound)
	if err == true {
		return c.String(http.StatusInternalServerError, "メールアドレスが違います!")
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))

	if passwordErr != nil {
		//パスワードが合っているか
		return c.String(http.StatusInternalServerError, "パスワードが違います!")
	}

	response := Param{}

	response.Id = user.Id
	response.Name = user.Name
	response.Email = user.Email
	response.Password = user.Password

	token := auth.GetTokenHandler()

	response.Token = token
	log.Print(response)

	return c.JSON(http.StatusOK, response)
}

// func UserToParam(user *User){
// 	Para
// }

func 退避() {
	// if err := db.QueryRow("SELECT password FROM users WHERE email = $1", param.Email).Scan(&password); err != nil {
	// 	log.Print(err)
	// 	return c.String(http.StatusOK, "メールアドレスが違います!")
	// }
}

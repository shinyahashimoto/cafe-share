package user

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Name     string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	param := new(User)
	if err := c.Bind(param); err != nil {
		return err
	}

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=Niko1004 dbname=cafe-share sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	// メールアドレスの重複がないか確認
	rows := db.QueryRow("SELECT * FROM users WHERE email = $1 ", param.Email)

	if rows != nil {
		// メールアドレスが重複していた場合、登録はしない
		return c.String(http.StatusOK, "既にメールアドレスが使用されています")
	}

	//パスワードのハッシュ化
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), 12)

	_, err = db.Exec("INSERT INTO users(name, email, password) VALUES($1, $2, $3)", param.Name, param.Email, hashPassword)

	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, "登録完了!")
}

func dbAccess() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=Niko1004 dbname=cafe-share sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	// INSERT
	// var empID string
	// id := 3
	// number := 4444
	// err = db.QueryRow("INSERT INTO employee(emp_id, emp_number) VALUES($1,$2) RETURNING emp_id", id, number).Scan(&empID)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(empID)

	// SELECT
	rows, err := db.Query("SELECT * FROM users	")

	if err != nil {
		fmt.Println(err)
	}

	var es []User
	for rows.Next() {
		var e User
		rows.Scan(&e.Id, &e.Name)
		es = append(es, e)
	}
	fmt.Printf("%v", es)
}

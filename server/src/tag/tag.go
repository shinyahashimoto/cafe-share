package tag

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetAll(c echo.Context) error {
	param := new(Tag)
	if err := c.Bind(param); err != nil {
		return err
	}

	// connectDB.ConnectDB()

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=Niko1004 dbname=cafe-share sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(reflect.TypeOf(db))

	rows, err := db.Query("SELECT id, name FROM tags")

	if err != nil {
		fmt.Println(err)
	}

	var es []Tag
	// var profJson []uint8
	for rows.Next() {
		var e Tag
		rows.Scan(&e.Id, &e.Name)
		es = append(es, e)

	}
	fmt.Print(es)
	return c.JSON(http.StatusOK, es)
}

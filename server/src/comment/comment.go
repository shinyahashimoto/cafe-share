package like

import (
	"log"
	"net/http"
	"server/db"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Comment struct {
	Id        int       `gorm:"primary_key" json:"id"`
	ShopId    int       `json:"shopId"`
	UserId    int       `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdateAt"`
	DeletedAt *time.Time
}

func Create(c echo.Context) error {

	comment := bindComment(c)
	db := db.DbConnect()
	defer db.Close()

	if comment.Id == 0 {
		log.Print("新規登録")
		db.Debug().Create(&comment)
		// TODO エラーハンドラー
		return c.JSON(http.StatusCreated, comment)

	} else {
		log.Print("更新登録")
		db.Debug().Update(&comment)
		// TODO エラーハンドラー
		return c.JSON(http.StatusOK, comment)
	}
}

func Delete(c echo.Context) error {
	commentId := c.QueryParam("commentId")
	db := db.DbConnect()
	defer db.Close()
	comment := Comment{}
	comment.Id, _ = strconv.Atoi(commentId)
	// TODO エラーハンドラー
	db.Delete(&comment)
	return c.String(http.StatusOK, "正常に削除しました。")
}

// shop構造体にマッピング
func bindComment(c echo.Context) *Comment {
	comment := new(Comment)
	if err := c.Bind(comment); err != nil {
		logrus.Error(err)
	}
	return comment
}

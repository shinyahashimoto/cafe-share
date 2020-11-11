package like

import (
	"log"
	"net/http"
	"server/db"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Like struct {
	Id        int       `gorm:"primary_key" json:"id"`
	ShopId    int       `json:"shopId"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdateAt"`
	IsDeleted bool
}

func Create(c echo.Context) error {

	like := bindLike(c)

	db := db.DbConnect()
	defer db.Close()

	if like.Id == 0 {
		logrus.Info("新規登録")

		db.Debug().Create(&like)
		// if info := db.Debug().Create(&like); info. != nil {
		// 	logrus.WithFields(logrus.Fields{
		// 		"error": err,
		// 	}).Error("Failed to Create")
		// 	return c.String(http.StatusInternalServerError, "登録が失敗しました。")
		// }

	} else {
		log.Print("更新登録")
		like.IsDeleted = !like.IsDeleted
		log.Print(like.IsDeleted)
		log.Print(!like.IsDeleted)

		db.Debug().Save(&like)
		// if err := db.Debug().Save(&shop); err != nil {
		// 	logrus.WithFields(logrus.Fields{
		// 		"error": err,
		// 	}).Fatal("Failed to Create")
		// 	return c.String(http.StatusInternalServerError, "登録が失敗しました。")
		// }

	}
	return c.JSON(http.StatusCreated, like)
}

// shop構造体にマッピング
func bindLike(c echo.Context) *Like {
	like := new(Like)
	if err := c.Bind(like); err != nil {
		logrus.Error(err)
	}
	log.Print(like)
	return like
}

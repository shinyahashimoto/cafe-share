package favorite

import (
	"log"
	"net/http"
	"server/db"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Favorite struct {
	Id        int       `gorm:"primary_key" json:"id"`
	ShopId    int       `json:"shopId"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdateAt"`
	IsDeleted bool      `json:"isDeleted"`
}

func Create(c echo.Context) error {

	favorite := bindFavorite(c)

	log.Print(&favorite)

	db := db.DbConnect()
	defer db.Close()

	if favorite.Id == 0 {
		log.Print("新規登録")
		db.Debug().Create(&favorite)

		//TODO エラーハンドリング
		return c.JSON(http.StatusCreated, favorite)

	} else {
		log.Print("更新登録")
		favorite.IsDeleted = !favorite.IsDeleted
		db.Debug().Save(&favorite)
		//TODO エラーハンドリング
		return c.JSON(http.StatusOK, favorite)

	}
}

// shop構造体にマッピング
func bindFavorite(c echo.Context) *Favorite {
	favorite := new(Favorite)
	if err := c.Bind(favorite); err != nil {
		logrus.Error(err)
	}
	log.Print(favorite)
	return favorite
}

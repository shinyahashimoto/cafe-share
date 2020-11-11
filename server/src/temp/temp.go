import (
	"log"
	"net/http"
	"server/db"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func Create(c echo.Context) error {
	param := bindParam(c)
	shop := bindShop(param)

	db := db.DbConnect()
	defer db.Close()

	if param.ShopId == 0 {
		logrus.Info("新規登録")

		db.Debug().Create(&shop)
		// if info := db.Debug().Create(&shop); info. != nil {
		// logrus.WithFields(logrus.Fields{
		// 	"error": err,
		// }).Error("Failed to Create")
		// return c.String(http.StatusInternalServerError, "登録が失敗しました。")
		// }

	} else {
		log.Print("更新登録")

		db.Debug().Save(&shop)
		// if err := db.Debug().Save(&shop); err != nil {
		// 	logrus.WithFields(logrus.Fields{
		// 		"error": err,
		// 	}).Fatal("Failed to Create")
		// 	return c.String(http.StatusInternalServerError, "登録が失敗しました。")
		// }

	}
	return c.String(http.StatusCreated, "登録完了!")
}

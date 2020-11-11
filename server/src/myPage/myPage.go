package myPage

import (
	"log"
	"net/http"
	comment "server/comment"
	"server/db"
	"server/favorite"
	"server/shop"
	"server/tag"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Shop shop.Shop

// type Param shop.Param
type Tag tag.Tag
type Like shop.Like
type Comment comment.Comment
type Favorite favorite.Favorite

type Param struct {
	ShopId           int      `json:"id"`
	ShopName         string   `json:"shopName"`
	Address          string   `json:"address"`
	ShopUrl          string   `json:"shopUrl"`
	ShopImage        string   `json:"shopImage"`
	OpenTime         string   `json:"openTime"`
	CloseTime        string   `json:"closeTime"`
	Impression       string   `json:"impression"`
	SelectedTags     []int    `json:"selectedTags"`
	SelectedTagNames []string `json:"selectedTagNames"`

	UserId    int   `json:"userId"`
	Like      Like  `json:"like"`
	LikeCount int64 `gorm:"-" json:"likeCount"`

	Favorite      Favorite `json:"favorite"`
	FavoriteCount int64    `gorm:"-" json:"favoriteCount"`

	Comments []Comment `json:"comments"`
}

var count int64

func Get(c echo.Context) error {

	userId := c.QueryParam("userId")

	log.Print(userId)

	request := new(Param)
	if err := c.Bind(request); err != nil {
		logrus.Error(err)
	}
	db := db.DbConnect()
	defer db.Close()
	// shopをIDを元に検索
	shops := []Shop{}
	// 指定したIDを元にレコードを１つ引っ張ってくる
	// result := db.Debug().Joins("left join likes on likes.shop_id = shops.id and likes.user_id = ?", userId).Where("shops.id = ?", id).First(&shop)
	db.Debug().Joins("inner join favorites on shops.id = favorites.shop_id").Where("favorites.user_id= ?", userId).Find(&shops)

	return c.JSON(http.StatusOK, shops)
}

func FormatToParam(shop *Shop, tagList *[]Tag) Param {
	param := Param{}

	param.ShopId = shop.Id
	param.ShopName = shop.Name
	param.Address = shop.Address
	param.ShopUrl = shop.ShopUrl
	param.OpenTime = shop.OpenTime
	param.CloseTime = shop.CloseTime
	param.Impression = shop.Impression
	param.SelectedTags = []int{}
	param.SelectedTagNames = []string{}

	param.UserId = shop.UserId

	for _, tag := range *tagList {
		param.SelectedTags = append(param.SelectedTags, tag.Id)
		param.SelectedTagNames = append(param.SelectedTagNames, tag.Name)
	}

	// param.SelectedTags = selectedTags

	return param
}

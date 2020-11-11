package shop

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"server/db"
	"server/favorite"
	"strconv"

	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/guregu/null"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
)

type Shop struct {
	Id             int    `gorm:"primary_key" json:"id"`
	Name           string `json:"shopName"`
	Address        string `json:"address"`
	ShopUrl        string `json:"shopUrl"`
	OpenTime       string `json:"openTime"`
	CloseTime      string `json:"closeTime"`
	Impression     string `json:"impression"`
	UserId         int    `json:"userId"`
	ExistShopImage bool
	Like           Like  `json:"like"`
	LikeCount      int64 `gorm:"-" json:"likeCount"`
	FavoriteCount  int64 `gorm:"-" json:"favoriteCount"`
	CommentCount   int64 `gorm:"-" json:"commentCount"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	SelectedTags   []int       `gorm:"-" json:"selectedTags"`
	ImageNames     []string    `gorm:"-" json:"imageNames"`
	ImageName      null.String `gorm:"-"`
	TagId          null.Int    `gorm:"-" `
}

type ShopTag struct {
	Id     int `gorm:"primary_key" json:"id"`
	ShopId int
	TagId  int
}

type Like struct {
	Id        int `json:"id"`
	ShopId    int
	UserId    int  `json:"userId"`
	IsDeleted bool `json:"isDeleted"`
}

type Favorite favorite.Favorite

type Comment struct {
	Id        int       `gorm:"primary_key" json:"id"`
	ShopId    int       `json:"shopId"`
	UserId    int       `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdateAt"`
	DeletedAt *time.Time

	UserName string `json:"userName"`
}

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

	Comments     []Comment `json:"comments"`
	CommentCount int64     `gorm:"-" json:"commentCount"`
}

type Tag struct {
	Id        int
	Name      string
	ImageName string
}

var likeCount int64
var favoriteCount int64
var commentCount int64

//パラメータからパラメータ受け取り構造体にマッピング
func bindParam(c echo.Context) *Param {

	param := new(Param)
	if err := c.Bind(param); err != nil {
		logrus.Error(err)
	}

	log.Print("param")
	log.Print(param)

	// param := new(Param)
	// if err := c.Bind(param); err != nil {
	// 	logrus.WithFields(logrus.Fields{
	// 		"error": err,
	// 	}).Error("Failed to Bind Param")
	// 	return c.String(http.StatusInternalServerError, "登録が失敗しました。")
	// }
	return param
}

// shop構造体にマッピング
func bindShop(param *Param) *Shop {
	// 構造体のインスタンス化
	shop := Shop{}
	// 挿入したい情報を構造体に与える
	shop.Id = param.ShopId
	shop.Name = param.ShopName
	shop.Address = param.Address
	shop.ShopUrl = param.ShopUrl
	shop.OpenTime = param.OpenTime
	shop.CloseTime = param.CloseTime
	shop.Impression = param.Impression
	shop.UserId = param.UserId

	if param.ShopImage != "" {
		shop.ExistShopImage = true
	}
	return &shop
}

// 画像をAWSs3にアップロード処理
func s3Upload(fileData string, shopName string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	bucketName := "cafe-share"
	objectKey := shopName

	file := strings.NewReader(fileData)

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")

}

// shopを登録します
func Create(c echo.Context) error {
	param := bindParam(c)
	shop := bindShop(param)

	db := db.DbConnect()
	defer db.Close()

	// 新規か更新登録かの条件分岐
	if param.ShopId == 0 {
		logrus.Info("新規登録")

		result := db.Debug().Create(&shop)
		// err := errors.Is(result.Error)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		if len(param.SelectedTags) != 0 {
			for _, tagId := range param.SelectedTags {
				shopTag := ShopTag{}
				shopTag.ShopId = shop.Id
				shopTag.TagId = tagId
				db.Debug().Create(&shopTag)
			}
		}
	} else {
		log.Print("更新登録")

		// db.Debug().Save(&shop)
		db.Debug().Model(&shop).Updates(Shop{Name: shop.Name, Address: shop.Address, ShopUrl: shop.ShopUrl, OpenTime: shop.OpenTime, CloseTime: shop.CloseTime, Impression: shop.Impression})

		// if err := db.Debug().Save(&shop); err != nil {
		// 	logrus.WithFields(logrus.Fields{
		// 		"error": err,
		// 	}).Fatal("Failed to Create")
		// 	return c.String(http.StatusInternalServerError, "登録が失敗しました。")
		// }

		// shop_tagsは削除してから登録
		db.Debug().Delete(&ShopTag{}, "shop_id = ?", param.ShopId)
		// tagを1件ずつ登録
		if len(param.SelectedTags) != 0 {
			for _, tagId := range param.SelectedTags {
				shopTag := ShopTag{}
				shopTag.ShopId = shop.Id
				shopTag.TagId = tagId
				db.Debug().Create(&shopTag)
			}
		}
	}

	if shop.ExistShopImage {
		log.Print("s3にアップロードします")
		// s3Upload(param.ShopImage, param.ShopName)
	}

	dec, _ := base64.StdEncoding.DecodeString(param.ShopImage)
	fmt.Println(string(dec))

	return c.String(http.StatusCreated, "登録完了!")
}

func GetAll(c echo.Context) error {
	// db := db.DbConnect()
	// defer db.Close()

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=Niko1004 dbname=cafe-share sslmode=disable")
	defer db.Close()

	shops := []Shop{}

	rows, err := db.Query("SELECT shops.id, shops.name, shops.address, shops.shop_url, shops.open_time, shops.close_time, shops.exist_shop_image, A.likeCount, B.favoriteCount, C.commentCount, tags.id, tags.image_name FROM shops LEFT JOIN ( SELECT shops.id, COUNT(likes.id) as likeCount FROM shops LEFT JOIN likes on  shops.id = likes.shop_id GROUP BY shops.id ) as A on  shops.id = A.id LEFT JOIN ( SELECT shops.id, COUNT(favorites.id) as favoriteCount FROM shops LEFT JOIN favorites on  shops.id = favorites.shop_id GROUP BY shops.id ) AS B on  shops.id = B.id LEFT JOIN ( SELECT shops.id, COUNT(comments.id) as commentCount FROM shops LEFT JOIN comments on  shops.id = comments.shop_id GROUP BY shops.id ) AS C on  shops.id = C.id LEFT JOIN shop_tags on  shop_tags.shop_id = shops.id left join tags on  tags.id = shop_tags.tag_id order by shops.id")

	if err != nil {
		panic(err.Error())
	}

	currentShopId := 0

	// tagId := 0

	for rows.Next() {
		shop := Shop{}
		if err := rows.Scan(&shop.Id, &shop.Name, &shop.Address, &shop.ShopUrl, &shop.OpenTime, &shop.CloseTime, &shop.ExistShopImage, &shop.LikeCount, &shop.FavoriteCount, &shop.CommentCount, &shop.TagId, &shop.ImageName); err != nil {
			log.Fatal(err)
		}

		jsonTagId, _ := json.Marshal(shop.TagId)
		tagId, _ := strconv.Atoi(string(jsonTagId))
		jsonImageName, _ := json.Marshal(shop.ImageName)
		imageName := string(jsonImageName)
		fmt.Println(&shop.ImageName)

		if shop.Id != currentShopId {
			shops = append(shops, shop)
			currentShopId = shop.Id
		}

		if tagId != 0 {
			replaced1 := strings.Replace(imageName, "\"", "", -1)
			shops[len(shops)-1].ImageNames = append(shops[len(shops)-1].ImageNames, replaced1)
		}

	}
	// file := s3Read("TODO")

	return c.JSON(http.StatusOK, shops)

}

// shops.idを元に検索します
func GetById(c echo.Context) error {
	shopId := c.QueryParam("id")
	userId := c.QueryParam("userId")

	request := new(Param)
	if err := c.Bind(request); err != nil {
		logrus.Error(err)
	}
	db := db.DbConnect()
	defer db.Close()
	// shopをIDを元に検索
	shop := Shop{}
	// 指定したIDを元にレコードを１つ引っ張ってくる
	// result := db.Debug().Joins("left join likes on likes.shop_id = shops.id and likes.user_id = ?", userId).Where("shops.id = ?", id).First(&shop)
	result := db.Debug().Where("id = ?", shopId).First(&shop)

	err := errors.Is(result.Error, gorm.ErrRecordNotFound)
	if err == true {
		return c.String(http.StatusInternalServerError, "お探しのお店は削除されているか、お時間を置いて再度アクセスしてみだくさい。")
	}

	// shop.idを元にshop_tagsを検索

	tagList := []Tag{}
	db.Debug().Model(&Tag{}).Select("tags.id, tags.name").Joins("left join shop_tags st on tags.id = st.tag_id left join shops s on st.shop_id = s.id").Where("s.id = ?", shopId).Scan(&tagList)

	param := FormatToParam(&shop, &tagList)
	if shop.ExistShopImage {
		//お金がかかってしまうから、コメントアウト
		// param.ShopImage = s3Read(param.ShopName)
	}

	like := Like{}
	result = db.Debug().Where("shop_id = ? AND user_id = ? and is_deleted is false", shopId, userId).First(&like)
	if result.RowsAffected == 1 {
		param.Like = like
		log.Print("like")
	}

	db.Debug().Model(&Like{}).Where("shop_id = ? and is_deleted is false", shopId).Count(&likeCount)

	favorite := Favorite{}
	result = db.Debug().Where("shop_id = ? AND user_id = ?", shopId, userId).First(&favorite)
	if result.RowsAffected == 1 {
		param.Favorite = favorite
	}

	db.Debug().Model(&Favorite{}).Where("shop_id = ? and is_deleted = false", shopId).Count(&favoriteCount)

	comments := []Comment{}
	// db.Debug().Where("shop_id = ?", shopId).Find(&comments)
	commentResult := db.Debug().Select("users.name as user_name, comments.*").Joins("left join users on users.id = comments.user_id").Where("shop_id = ?", shopId).Find(&comments)

	param.FavoriteCount = favoriteCount
	param.LikeCount = likeCount
	param.CommentCount = commentResult.RowsAffected
	param.Comments = comments

	return c.JSON(http.StatusOK, param)
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

func s3Read(shopName string) string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	bucketName := "cafe-share"
	objectKey := shopName

	// S3 clientを作成
	svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatal(err)
	}

	rc := obj.Body
	defer rc.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(obj.Body)
	rslt := buf.String()
	return rslt
}

func Delete(c echo.Context) error {
	paramShopId := c.QueryParam("shopId")
	shop := Shop{}
	shopId, _ := strconv.Atoi(paramShopId)
	shop.Id = shopId

	db := db.DbConnect()
	defer db.Close()
	db.Debug().Delete(&shop)

	return c.String(http.StatusOK, "削除が完了しました")

}

// db.Debug().Select("shops.*,COUNT(likes.id) AS like_count, count(favorites.id) as favorite_count , count(comments.id) as comment_count").Joins("left join likes on likes.shop_id = shops.id left join favorites on favorites.shop_id = shops.id left join comments on comments.shop_id = shops.id").Where("shops.deleted_at is null").Group("shops.id , shops.name, shops.address, shops.shop_url, shops.open_time, shops.close_time, shops.impression, shops.exist_shop_image, shops.user_id, shops.created_at, shops.updated_at, shops.deleted_at").Order("shops.id").Find(&shops)

// 	select A.id, A.name, A.address, A.open_time, A.close_time, A.exist_shop_image, A.like_count, B.favorite_count, C.comment_count
// from
// 	(select shops.id, shops.name, shops.address, shops.open_time, shops.close_time, shops.exist_shop_image, COUNT(likes.id) as like_count
// 		from shops
// 		left join likes
// 	 	on shops.id = likes.shop_id
// 	   GROUP BY shops.id,
// 	 shops.name,shops.address,shops.open_time,shops.close_time,shops.exist_shop_image

// 		)as A
// left join
// (select shops.id, COUNT(favorites.id) as favorite_count
// 		from shops
// 		left join favorites
// 	 	on shops.id = favorites.shop_id
// 	   GROUP BY shops.id

// 		)as B USING(id)
// left join
// (select shops.id, COUNT(comments.id) as comment_count
// 		from shops
// 		left join comments
// 	 	on shops.id = comments.shop_id
// 	   GROUP BY shops.id
// 		)as C USING(id)
// order by id;

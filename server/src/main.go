package main

import (
	"bytes"
	"fmt"
	"log"
	comment "server/comment"
	"server/favorite"
	"server/like"
	"server/login"
	"server/myPage"
	"server/shop"
	"server/tag"
	"server/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// // サーバー起動

	e.Use(middleware.CORS())

	e.POST("/login", login.Login)
	e.POST("/users", user.Register)

	// shop
	e.POST("/shop", shop.Create)
	e.DELETE("/shop", shop.Delete)
	// e.PUT("/shop", shop.Update)
	// param含めて
	e.GET("/shop", shop.GetById)
	e.GET("/shopAll", shop.GetAll)
	e.GET("/shopA", tag.GetAll)
	e.POST("/like", like.Create)
	e.POST("/favorite", favorite.Create)

	// s3Upload()

	//comment
	e.POST("/comment", comment.Create)
	e.DELETE("/comment", comment.Delete)

	//myPage
	e.GET("/myPage", myPage.Get)

	e.Start(":8000") //ポート番号指定してね

}

func s3Read() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	bucketName := "cafe-share"
	objectKey := "testsss"

	// S3 clientを作成
	svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("aaa")

	rc := obj.Body
	defer rc.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(obj.Body)
	rslt := buf.String()
	fmt.Println(rslt)
}

// func s3Upload() {
// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))

// 	// ファイルを開く
// 	targetFilePath := "img/みにおん.jpeg"
// 	file, err := os.Open(targetFilePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	opt := jpeg.Options{
// 		Quality: 90,
// 	}
// 	err = jpeg.Encode(file, &opt)
// 	if err != nil {
// 		// Handle error
// 	}

// 	bucketName := "cafe-share"
// 	objectKey := "test"

// 	// Uploaderを作成し、ローカルファイルをアップロード
// 	uploader := s3manager.NewUploader(sess)
// 	_, err = uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(objectKey),
// 		Body:   file,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("done")

// }

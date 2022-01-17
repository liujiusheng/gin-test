package main

import (
	"image/png"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	// "os"
	// "fmt"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		cap := captcha.New()
		cap.SetFont("comic.ttf")
		img, str := cap.Create(4, captcha.NUM)
		// fmt.Printf(c,str,img)
		// f, err := os.Create("image.png")
		png.Encode(c.Writer, img)
		println(str)
		// c.SaveUploadedFile(png, "./test.png")
		// os.WriteFile("image.png", img, 0600)
		// reader := img
		// contentLength := img.ContentLength
		// contentType := img.Header.Get("Content-Type")
		// extraHeaders := map[string]string{
		// 	"Content-Disposition": `attachment; filename="gopher.png"`,
		// }
		// c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		// c.JSON(200, gin.H{
		// 	"message":str,
		// })
	})
	r.GET("/pg", func(c *gin.Context) {

	})
	r.Run(":8080")
}

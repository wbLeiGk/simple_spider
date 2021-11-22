package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func gin_invoke() {

	r := gin.Default()
	r.POST("/url", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		var json_data map[string]interface{}
		c.BindJSON(&json_data)
		url := json_data["url"].(string)
		// url := c.Param("url")
		content := invoke_spider(url)
		fmt.Println("the url : ", url)
		c.JSON(200, gin.H{
			"message": content,
		})
	})
	r.Run(":8030") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func invoke_spider(url string) string {
	var wa sync.WaitGroup
	var result string

	wa.Add(1)

	c := colly.NewCollector()
	c.OnHTML("div[id]", func(e *colly.HTMLElement) {
		if e.Attr("id") == "content" {
			defer wa.Done()
			r_br := regexp.MustCompile("(<br/>)+")
			r_0a := regexp.MustCompile("(\u00a0)")
			ret, _ := e.DOM.Html()
			b2 := r_br.ReplaceAll([]byte(ret), []byte("\n"))
			b2 = r_0a.ReplaceAll(b2, []byte(" "))

			b, _ := GbkToUtf8([]byte(b2))
			result = string(b)
			fmt.Printf("content:", string(b))
		}
		// e.Request.Visit(e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// c.Visit("https://www.777zw.net/2/2142/751765.html")
	c.Visit(url)

	wa.Wait()
	return result
}

func main() {

	gin_invoke()

}

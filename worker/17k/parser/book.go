package parser

import (
	"encoding/json"
	"log"
	"regexp"
	"worker/engine"
	"worker/model"
)

/*
<meta itemprop="name"
      content="《人格的忏悔》百诺酱著-17K小说网">
<meta itemprop="description" content="分享各种下饭操作。">
*/

func ParseBook(contents []byte) (pr engine.SerializationParseResult) {
	const (
		bookRe = `<meta name="Description"
	      content="(.+)由作家(.+)所著的(.+)小说，小说，您可以在17k小说网免费试读" />`
		infoRe = `<meta itemprop="description" content="([^"]+)">`
		urlRe  = `<meta property="og:url" content="(https://www.17k.com/book/[0-9]+.html)" />`
	)
	bookre := regexp.MustCompile(bookRe)
	infore := regexp.MustCompile(infoRe)
	urlre := regexp.MustCompile(urlRe)
	bookSubmatch := bookre.FindSubmatch(contents)
	infoSubmatch := infore.FindSubmatch(contents)
	urlSubmatch := urlre.FindSubmatch(contents)
	var book model.Book
	if len(bookSubmatch) > 0 {
		book.Platform = "17k"
		book.Url = string(urlSubmatch[1])
		book.Name = string(bookSubmatch[1])
		book.Author = string(bookSubmatch[2])
		book.Kind = string(bookSubmatch[3])
		if len(infoSubmatch) > 0 {
			book.Info = string(infoSubmatch[1])
		}
		marshal, err := json.Marshal(&book)
		if err != nil {
			log.Printf("Faild to marshal object: %v, err: %s", book, err)
			return
		}
		pr.Items = append(pr.Items, string(marshal))
	}
	//fmt.Println(book)
	/*utils.Save(fmt.Sprintf("[name: %s]\n[author: %s]\n[kind: %s]\n[info: %s]\n\n", book.Name, book.Author, book.Kind, book.Info))
	utils.Flush()*/
	return
}

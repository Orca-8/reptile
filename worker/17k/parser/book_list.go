package parser

import (
	"regexp"
	"worker/engine"
)

func ParseBookList(contents []byte) (pr engine.SerializationParseResult) {
	const (
		bookListRe     = `<strong><a target="_blank" href="(//www.17k.com/book/[0-9]+.html)" >+([^<]+)</a></strong>`
		prefix         = "https:"
		nextPage       = `<a href="(/all/book/[0-9_]+_0_0_0_0_0_0_[0-9]+.html)"><span>下一页</span></a>`
		nextPagePrefix = "https://www.17k.com"
	)
	re := regexp.MustCompile(bookListRe)
	nextPagere := regexp.MustCompile(nextPage)
	nextUrl := nextPagere.FindSubmatch(contents)
	if len(nextUrl) > 0 {
		// 下一页
		pr.Requests = append(pr.Requests, engine.SerializationRequest{
			Url:        nextPagePrefix + string(nextUrl[1]),
			ParserFunc: "ParseBookList",
		})
	}
	// 当前页小说
	allString := re.FindAllSubmatch(contents, -1)
	for _, s := range allString {
		//pr.Items = append(pr.Items, string(s[2]))
		pr.Requests = append(pr.Requests, engine.SerializationRequest{
			Url:        prefix + string(s[1]),
			ParserFunc: "ParseBook",
		})
	}
	return
}

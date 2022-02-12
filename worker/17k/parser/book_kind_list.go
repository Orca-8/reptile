package parser

import (
	"regexp"
	"worker/engine"
)

func ParseBookKindList(contents []byte) (pr engine.SerializationParseResult) {
	const (
		bookKindListRe = `<a href="(/all/book/[0-9_]+_0_0_0_0_0_0_1.html)" >+([^<]+)</a>`
		prefix         = "https://www.17k.com"
	)
	re := regexp.MustCompile(bookKindListRe)
	allString := re.FindAllSubmatch(contents, -1)
	for _, s := range allString {
		//pr.Items = append(pr.Items, string(s[2]))
		pr.Requests = append(pr.Requests, engine.SerializationRequest{
			Url:        prefix + string(s[1]),
			ParserFunc: "ParseBookList",
		})
	}
	return
}

package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type SerializationRequest struct {
	Url        string
	ParserFunc string
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type SerializationParseResult struct {
	Requests []SerializationRequest
	Items    []string
}

package model

type FizzBuzzQuery struct {
	Int1  int    `form:"int1"`
	Int2  int    `form:"int2"`
	Limit int    `form:"limit"`
	Str1  string `form:"str1"`
	Str2  string `form:"str2"`
}

func (fbq FizzBuzzQuery) IsEmpty() bool {
	return fbq == FizzBuzzQuery{}
}

type MostUsedQuery struct {
	Params string `json:"params"`
	Hits   int64  `json:"hits"`
}

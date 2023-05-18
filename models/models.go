package models

type Url struct {
	LongUrl  string
	ShortUrl string
}

func (u Url) TableName() string {
	return "converted_urls"
}

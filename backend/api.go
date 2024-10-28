package backend

import (
	"gorm.io/gorm"
)

// Model for storing shortened urls
type LittleLink struct {
	gorm.Model
	ShortUrl string
	LongUrl  string
}

// Function to create tables in database based on provided struct
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&LittleLink{})
}

// Function to generate a random short url path
func ShortenUrl() string {
	return String(14)
}

// Function to store a url and return a shortened url
func LogUrl(db *gorm.DB, longurl string) string {
	shorturl := ShortenUrl()
	db.Create(&LittleLink{ShortUrl: shorturl, LongUrl: longurl})
	return shorturl
}

// Function to retrieve a shortened url
func RetrieveUrl(db *gorm.DB, key string) string {
	var result LittleLink
	db.Model(LittleLink{ShortUrl: key}).First(&result)

	return result.LongUrl
}

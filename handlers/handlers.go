package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"urlshortener/models"
	"urlshortener/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Store map[string]string

type UrlWrapper struct {
	Url string `json:"url" binding:"required"`
}

func Shorten(c *gin.Context) {
	var longUrl, shortUrl UrlWrapper
	if err := c.ShouldBind(&longUrl); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// check if it's present in the db
	var urlRetrieved models.Url
	gormDb := c.MustGet(utils.DbKey).(*gorm.DB)
	if err := gormDb.Model(&models.Url{}).First(&urlRetrieved, "long_url = ?", longUrl.Url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			shortUrl.Url = fmt.Sprintf("%v%v", utils.ShortPrefix, utils.GenerateRandomCharacters(utils.ChararcterSet, 6))
			gormDb.Create(models.Url{LongUrl: longUrl.Url, ShortUrl: shortUrl.Url})
			c.JSON(http.StatusOK, shortUrl)
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	shortUrl.Url = urlRetrieved.ShortUrl
	c.JSON(http.StatusOK, shortUrl)
}

func Unshorten(c *gin.Context) {
	var shortUrl UrlWrapper
	if err := c.ShouldBind(&shortUrl); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// check if this short url has been already converted
	var urlRetrieved models.Url
	gormDb := c.MustGet(utils.DbKey).(*gorm.DB)
	if err := gormDb.Model(&models.Url{}).First(&urlRetrieved, "short_url = ?", shortUrl.Url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.String(http.StatusNotFound, fmt.Sprintf("the url %q is not known", shortUrl.Url))
			return
		}
	}

	c.JSON(http.StatusOK, UrlWrapper{Url: urlRetrieved.LongUrl})
}

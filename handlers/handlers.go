package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"urlshortener/utils"

	"github.com/gin-gonic/gin"
)

var Store sync.Map

type UrlWrapper struct {
	Url string `json:"url" binding:"required"`
}

func Shorten(c *gin.Context) {
	var longUrl, shortUrl UrlWrapper
	if err := c.ShouldBind(&longUrl); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// check if it's already present in the map
	value, isFound := Store.Load(longUrl.Url)
	if isFound {
		shortUrl.Url = value.(string)
		c.JSON(http.StatusOK, shortUrl)
		return
	}

	// generate new url
	shortUrl.Url = fmt.Sprintf("%v%v", utils.ShortPrefix, utils.GenerateRandomCharacters(utils.ChararcterSet, 6))
	Store.Store(longUrl.Url, shortUrl.Url)
	c.JSON(http.StatusOK, shortUrl)
}

func Unshorten(c *gin.Context) {
	var shortUrl UrlWrapper
	if err := c.ShouldBind(&shortUrl); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// check if this short url has been already converted
	var longUrl UrlWrapper
	var isFound bool
	Store.Range(func(key, value any) bool {
		if value.(string) == shortUrl.Url {
			longUrl.Url = key.(string)
			isFound = true
			return false
		}
		return true
	})

	if !isFound {
		c.String(http.StatusNotFound, fmt.Sprintf("the url %q is not known", shortUrl.Url))
		return
	}

	c.JSON(http.StatusOK, longUrl)
}

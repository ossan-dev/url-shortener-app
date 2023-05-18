package handlers

import (
	"fmt"
	"net/http"

	"urlshortener/utils"

	"github.com/gin-gonic/gin"
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

	// check if it's already present in the map
	value, isFound := Store[longUrl.Url]
	if isFound {
		shortUrl.Url = value
		c.JSON(http.StatusOK, shortUrl)
		return
	}

	// generate new url
	shortUrl.Url = fmt.Sprintf("%v%v", utils.ShortPrefix, utils.GenerateRandomCharacters(utils.ChararcterSet, 6))
	Store[longUrl.Url] = shortUrl.Url
	c.JSON(http.StatusOK, shortUrl)
}

func Unshorten(c *gin.Context) {
	var shortUrl UrlWrapper
	if err := c.ShouldBind(&shortUrl); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// check if this short url has been already converted
	for k, v := range Store {
		if v == shortUrl.Url {
			c.JSON(http.StatusOK, UrlWrapper{Url: k})
			return
		}
	}

	c.String(http.StatusNotFound, fmt.Sprintf("the url %q is not known", shortUrl.Url))

}

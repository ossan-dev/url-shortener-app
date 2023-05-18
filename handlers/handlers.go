package handlers

import (
	"fmt"
	"math/rand"
	"net/http"

	"urlshortener/utils"

	"github.com/gin-gonic/gin"
)

var Store map[string]string

type UrlWrapper struct {
	Url string `json:"url" binding:"required"`
}

func generateRandomCharacters(chararcterSet string, numberOfDigits int) string {
	bytes := make([]byte, numberOfDigits)
	for k := range bytes {
		bytes[k] = chararcterSet[rand.Intn(len(chararcterSet))]
	}
	return string(bytes)
}

func Shorten(c *gin.Context) {
	var urlWrapper, res UrlWrapper
	if err := c.ShouldBind(&urlWrapper); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	// check if it's already present in the map
	value, isFound := Store[urlWrapper.Url]
	if isFound {
		res.Url = value
		c.JSON(http.StatusOK, res)
		return
	}

	// generate new url
	res.Url = fmt.Sprintf("%v%v", utils.ShortPrefix, generateRandomCharacters(utils.ChararcterSet, 6))
	Store[urlWrapper.Url] = res.Url
	c.JSON(http.StatusOK, res)
}

func Unshorten(c *gin.Context) {
	c.String(200, "Unshorten")
}

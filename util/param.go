package util

import (
	"dbo-test/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageParam(c *gin.Context) (page int, err error) {
	pg := c.Query("page")
	if pg == "" {
		err = fmt.Errorf("%s: page mandatory", config.Configure.Source.Name)
		return
	}
	page, err = strconv.Atoi(pg)
	if err != nil {
		err = fmt.Errorf("%s: page must number", config.Configure.Source.Name)
		return
	}
	if page < 1 {
		err = fmt.Errorf("%s: page must greater than 0", config.Configure.Source.Name)
		return
	}
	return
}

package tokenService

import (
	"dbo-test/config"
	"dbo-test/dao/userDao"
	"dbo-test/model/dto"
	"dbo-test/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context, _ dto.DataIN) (result interface{}, err error) {
	var dtIn dto.UserIn
	err = c.ShouldBindJSON(&dtIn)
	if err != nil {
		return
	}

	user, err := userDao.GetUser(dtIn.UserName, util.GenerateHmacSHA256(dtIn.Password))
	if err != nil {
		return
	}

	if user.Id.Int64 == 0 {
		err = fmt.Errorf("%s: unknown user and password", config.Configure.Source.Name)
		return
	}

	result = util.GenerateToken(user.UseName.String, user.Id.Int64)
	return
}

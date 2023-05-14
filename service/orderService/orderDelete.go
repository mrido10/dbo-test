package orderService

import (
	"dbo-test/config"
	"dbo-test/dao/orderDao"
	"dbo-test/model/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DeleteOrder(c *gin.Context, dtIn dto.DataIN) (_ interface{}, err error) {
	var od dto.OrderIn
	err = c.ShouldBindJSON(&od)
	if err != nil {
		return
	}

	count, err := orderDao.CountOrderById(od.Id)
	if err != nil {
		return
	}
	if count == 0 {
		err = fmt.Errorf("%s: unknown with this Customer ID", config.Configure.Source.Name)
		return
	}
	err = orderDao.DeleteOrder(dtIn.Tx, od.Id)
	return
}

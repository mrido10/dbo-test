package orderService

import (
	"dbo-test/config"
	"dbo-test/dao/orderDao"
	"dbo-test/model/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpdateOrder(c *gin.Context, dtIn dto.DataIN) (_ interface{}, err error) {
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
	order, err := validateInsertOrUpdate(od)
	if err != nil {
		return
	}
	order.Id.Int64 = od.Id
	err = orderDao.UpdateOrder(dtIn.Tx, order)
	return
}

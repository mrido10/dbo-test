package orderService

import (
	"dbo-test/config"
	"dbo-test/dao/orderDao"
	"dbo-test/model"
	"dbo-test/model/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetDetail(c *gin.Context, _ dto.DataIN) (result interface{}, err error) {
	conf := config.Configure
	idStr := c.Param("id")
	if idStr == "" {
		err = fmt.Errorf("%s: id mandatory", conf.Source.Name)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("%s: id must number", config.Configure.Source.Name)
		return
	}
	od, err := orderDao.GetOrderById(id)
	if err != nil {
		return
	}
	if od.Id.Int64 == 0 {
		err = fmt.Errorf("%s: unknown with this ID", conf.Source.Name)
		return
	}
	result = setDataDetail(od)
	return
}

func setDataDetail(od model.OrderModel) dto.OrderOut {
	return dto.OrderOut{
		Id: od.Id.Int64,
		CustomerId: od.CustomerId.Int64,
		Customername: od.CustomerModel.Name.String,
		ProductId: od.ProductId.Int64,
		ProductName: od.ProductModel.Name.String,
		TotalOrder: od.TotalOrder.Int32,
		Amount: od.Amount.Float64,
	}
}

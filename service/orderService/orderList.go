package orderService

import (
	"dbo-test/dao/orderDao"
	"dbo-test/model"
	"dbo-test/model/dto"
	"dbo-test/util"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context, _ dto.DataIN) (listData interface{}, err error) {
	page, err := util.GetPageParam(c)
	if err != nil {
		return
	}

	list, err := orderDao.GetListOrder(page, c.Query("searchCustomer"), c.Query("searchProduct"))
	if err != nil {
		return
	}
	listData = setDataList(list)
	return
}

func setDataList(list []model.OrderModel) (result []dto.OrderOut) {
	for _, od := range list {
		temp := dto.OrderOut{
			Id:           od.Id.Int64,
			CustomerId:   od.CustomerId.Int64,
			Customername: od.CustomerModel.Name.String,
			ProductId:    od.ProductId.Int64,
			ProductName:  od.ProductModel.Name.String,
			TotalOrder:   od.TotalOrder.Int32,
			Amount:       od.Amount.Float64,
		}
		result = append(result, temp)
	}
	return
}

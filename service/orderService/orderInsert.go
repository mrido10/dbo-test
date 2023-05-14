package orderService

import (
	"database/sql"
	"dbo-test/config"
	"dbo-test/dao/customerDao"
	"dbo-test/dao/orderDao"
	"dbo-test/dao/productDao"
	"dbo-test/model"
	"dbo-test/model/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InsertOrder(c *gin.Context, dtIn dto.DataIN) (_ interface{}, err error) {
	var od dto.OrderIn
	err = c.ShouldBindJSON(&od)
	if err != nil {
		return
	}

	conf := config.Configure
	cust, err := customerDao.GetCustomerById(od.CustomerId, true)
	if err != nil {
		return
	}
	if cust.Id.Int64 == 0 {
		err = fmt.Errorf("%s: unknown with this Customer ID", conf.Source.Name)
		return
	}

	prod, err := productDao.GetProductById(od.ProductId)
	if prod.Id.Int64 == 0 {
		err = fmt.Errorf("%s: unknown with this Product ID", conf.Source.Name)
		return
	}

	amount := prod.Price.Float64 * float64(od.TotalOrder)
	order := model.OrderModel{
		CustomerId: cust.Id,
		ProductId:  prod.Id,
		TotalOrder: sql.NullInt32{Int32: od.TotalOrder},
		Amount:     sql.NullFloat64{Float64: amount},
	}

	err = orderDao.InsertOrder(dtIn.Tx, order)
	return
}

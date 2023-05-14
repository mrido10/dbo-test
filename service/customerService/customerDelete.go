package customerService

import (
	"dbo-test/config"
	"dbo-test/dao/customerDao"
	"dbo-test/model/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DeleteCustomer(c *gin.Context, dtIn dto.DataIN) (_ interface{}, err error) {
	var cust dto.CustomerIn
	err = c.ShouldBindJSON(&cust)
	if err != nil {
		return
	}

	customer, err := customerDao.GetCustomerById(cust.Id, true)
	if err != nil {
		return
	}
	conf := config.Configure
	if customer.Id.Int64 == 0 {
		err = fmt.Errorf("%s: unknown with this ID", conf.Source.Name)
		return
	}

	err = customerDao.DeleteCustomer(dtIn.Tx, cust.Id)
	return
}

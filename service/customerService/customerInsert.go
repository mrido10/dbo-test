package customerService

import (
	"database/sql"
	"dbo-test/config"
	"dbo-test/dao/customerDao"
	"dbo-test/model"
	"dbo-test/model/dto"
	"dbo-test/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InsertCustomer(c *gin.Context, dtIn dto.DataIN) (_ interface{}, err error) {
	var cust dto.CustomerIn
	err = c.ShouldBindJSON(&cust)
	if err != nil {
		return
	}

	err = util.ValidateEmail(cust.Email)
	if err != nil {
		return
	}

	customer, err := customerDao.GetCustomerByUnique(cust.Name, cust.Email, true)
	if err != nil {
		return
	}
	if customer.Id.Int64 > 0 {
		err = fmt.Errorf("%s: name and email already exist", config.Configure.Source.Name)
		return
	}

	err = customerDao.UpsertCustomer(dtIn.Tx, model.CustomerModel{
		Name:  sql.NullString{String: cust.Name},
		Email: sql.NullString{String: cust.Email},
	})
	return
}

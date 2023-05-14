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

func UpdateCustomer(c *gin.Context, dtIn dto.DataIN) (_ interface{}, err error) {
	var cust dto.CustomerIn
	err = c.ShouldBindJSON(&cust)
	if err != nil {
		return
	}

	err = util.ValidateEmail(cust.Email)
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

	customer, err = customerDao.GetCustomerByUnique(cust.Name, cust.Email, true)
	if err != nil {
		return
	}
	if customer.Id.Int64 > 0 && cust.Id != customer.Id.Int64 {
		err = fmt.Errorf("%s: ame and email already exist", conf.Source.Name)
		return
	}

	err = customerDao.UpdateCustomer(dtIn.Tx, model.CustomerModel{
		Id:    sql.NullInt64{Int64: cust.Id},
		Name:  sql.NullString{String: cust.Name},
		Email: sql.NullString{String: cust.Email},
	})
	return
}

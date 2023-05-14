package customerService

import (
	"dbo-test/config"
	"dbo-test/dao/customerDao"
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
	customer, err := customerDao.GetCustomerById(id, true)
	if err != nil {
		return
	}
	if customer.Id.Int64 == 0 {
		err = fmt.Errorf("%s: unknown with this ID", conf.Source.Name)
		return
	}
	result = setDataDetail(customer)
	return
}

func setDataDetail(cust model.CustomerModel) (result dto.CustomerOut) {
	return dto.CustomerOut{
		Id:    cust.Id.Int64,
		Name:  cust.Name.String,
		Email: cust.Email.String,
	}
}

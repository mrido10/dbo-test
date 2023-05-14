package customerService

import (
	"dbo-test/dao/customerDao"
	"dbo-test/model"
	"dbo-test/model/dto"
	"dbo-test/util"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context, _ dto.DataIN) (listData interface{}, err error) {
	page, err := util.GetPageLimitParam(c)
	if err != nil {
		return
	}

	list, err := customerDao.GetListCustomer(page, c.Query("search"))
	if err != nil {
		return
	}
	listData = setDataList(list)
	return
}

func setDataList(list []model.CustomerModel) (result []dto.CustomerOut) {
	for _, cust := range list {
		temp := dto.CustomerOut{
			Id:    cust.Id.Int64,
			Name:  cust.Name.String,
			Email: cust.Email.String,
		}
		result = append(result, temp)

	}
	return
}

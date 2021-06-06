package handlers

import (
	"jstock/src/data/getter"

	"github.com/gin-gonic/gin"
)

func PurChaseTypeList(c *gin.Context) {
	Return(c)("success", "10001", getter.PurchaseTypeGetter.GetPurchaseTypeList())(OK)
}

package network

import (
	"github.com/gin-gonic/gin"
	"mygrpcp_project/types"
	"net/http"
)

func (r *Network) login(c *gin.Context) {
	var req types.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else if res, err := r.service.CreateAuth(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, res)
	}

}

func (r *Network) verify(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}

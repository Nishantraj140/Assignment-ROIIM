package pay

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Nishantraj140/Assignment-ROIIM/internal/user"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Pay(c *gin.Context) {
	session := sessions.Default(c)
	emaili := session.Get("user-id")
	email := emaili.(string)
	u := &user.User{
		Email: email,
	}
	err := u.Get()
	if err != nil {
		logger.ErrorLogger.Printf("error in find user, err:%v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"data":     nil,
			"debugMsg": "Unauthorized provided",
		})
		return
	}
	p := &PayReq{}
	err = c.ShouldBindJSON(p)
	if err != nil {
		logger.ErrorLogger.Printf("invalid pay req struct provided, err:%v", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "Invalid json provided",
		})
		return
	}
	//p.MerchantRefNum = uuid.New().String()
	rand.Seed(time.Now().UnixNano())
	p.MerchantRefNum = fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(1000000000))
	p.CustomerIp = "192.168.1.69"
	p.Description = "pay"
	resp, err := ProcessPay(p)
	if err != nil {
		logger.ErrorLogger.Printf("error in process pay api, err:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "error in process pay api",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":     resp,
		"debugMsg": "success",
	})
}

package profile

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Nishantraj140/Assignment-ROIIM/internal/address"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/singleToken"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Nishantraj140/Assignment-ROIIM/internal/user"
)

func Show(c *gin.Context) {
	session := sessions.Default(c)
	emaili := session.Get("user-id")
	email := emaili.(string)
	u := &user.User{
		Email: email,
	}
	err := u.Get()
	if err != nil {
		log.Printf("error in find user, err:%v", err)
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Unauthorized"})
		return
	}

	if u.CustomerId == "" {
		c.HTML(http.StatusOK, "profile.html", gin.H{"Profile": true, "FirstName": u.FirstName})
		return
	}
	rand.Seed(time.Now().UnixNano())
	stReq := &singleToken.SingleTokenReq{
		MerchantRefNum: fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(1000000000)),
		PaymentTypes:   []string{"CARD"},
	}
	stResp, err := singleToken.CreateSingleUseToken(stReq, u.CustomerId)
	if err != nil {
		log.Printf("error in creating single use token, error:%v", err)
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Single token api error"})
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{"Profile": false, "FirstName": u.FirstName /*"Cards": cards,*/, "SToken": stResp, "User": u, "Addr": stResp.Addresses[0]})
}

func CreateProfileHandler(c *gin.Context) {
	session := sessions.Default(c)
	emaili := session.Get("user-id")
	email := emaili.(string)
	u := &user.User{
		Email: email,
	}
	err := u.Get()
	if err != nil {
		logger.ErrorLogger.Printf("error in find user, err:%v", err)
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Unauthorized"})
		return
	}
	year, _ := strconv.Atoi(c.PostForm("year"))
	month, _ := strconv.Atoi(c.PostForm("month"))
	day, _ := strconv.Atoi(c.PostForm("day"))

	cp := CreateProfile{
		MerchantCustomerId: u.MerchantCustomerId,
		Locale:             "en_US",
		FirstName:          u.FirstName,
		MiddleName:         u.MiddleName,
		LastName:           u.LastName,
		DateOfBirth: DateOfBirth{
			Year:  year,
			Month: month,
			Day:   day,
		},
		Email:       u.Email,
		Phone:       u.Phone,
		Ip:          "192.168.1.69",
		Gender:      c.PostForm("gender"),
		Nationality: "",
		CellPhone:   "",
	}
	res, err := CreateProfileService(cp)
	if err != nil || res == nil || res.Id == "" {
		logger.ErrorLogger.Printf("error in create profile api call, res:%v err:%v", res, err)
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Create profile service failed", "error": err})
		return
	}
	ap := address.Address{
		NickName: c.PostForm("nickName"),
		Street:   c.PostForm("street"),
		City:     c.PostForm("city"),
		Zip:      c.PostForm("zip"),
		State:    c.PostForm("state"),
		Country:  c.PostForm("country"),
		Phone:    cp.Phone,
	}

	resp, err := address.CreateAddressService(ap, res.Id)
	if err != nil || resp == nil || resp.Id == "" {
		logger.ErrorLogger.Printf("error in create address api call, resp:%v err:%v", resp, err)
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Create address api service failed", "error": err})
		return
	}
	ap.Id = resp.Id
	err = ap.Create()
	if err != nil {
		logger.ErrorLogger.Printf("error in create address db, err:%v", err)
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "address create failed db error", "error": err})
		return
	}

	err = u.Update(res.Id, resp.Id)
	if err != nil {
		logger.ErrorLogger.Printf("error in update user db, err:%v", err)
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "User update failed db error", "error": err})
		return
	}
	c.Redirect(http.StatusFound, "/client")
}

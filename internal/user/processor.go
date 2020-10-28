package user

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	h, err := HashPassword(c.PostForm("password"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Internal Server Error"})
		return
	}
	u := &User{
		Email:      c.PostForm("email"),
		FirstName:  c.PostForm("firstName"),
		MiddleName: c.PostForm("middleName"),
		LastName:   c.PostForm("lastName"),
		Phone:      c.PostForm("phone"),
		Password:   h,
	}
	rand.Seed(time.Now().UnixNano())
	u.MerchantCustomerId = fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(1000000000))
	err = u.Create()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Internal Server Error"})
		return
	}
	c.Redirect(http.StatusFound, "/login")
}

func VerifyUser(c *gin.Context) {
	u := &User{
		Email: c.PostForm("email"),
	}
	err := u.Get()
	if err != nil && err.Error() != "record not found" {
		log.Printf("error in user form db, errorL%v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Internal Server Error"})
		return
	}
	if CheckPasswordHash(c.PostForm("password"), u.Password) {
		session := sessions.Default(c)
		session.Set("user-id", u.Email)
		err = session.Save()
		if err != nil {
			log.Printf("error in saving session, error:%v", err)
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Internal Server Error"})
			return
		}
		c.Redirect(http.StatusFound, "/client")
	} else {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"message": "Unauthorized"})
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

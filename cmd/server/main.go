package main

import (
	"flag"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/config"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/pay"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/profile"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/user"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/utils"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/logger"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
)

var configFile = flag.String("config", "conf/config.json","config file")

func main() {
	flag.Parse()
	config.ReadConfig(*configFile)
	err := sql.DBConn(&config.C.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer sql.DB.Close()

	router := gin.Default()
	gin.DefaultWriter = logger.File
	sql.DB.SetLogger(logger.DBLogger)

	token, err := utils.RandToken(64)
	if err != nil {
		log.Fatal("Unable to generate random tokens ", err)
	}

	store := cookie.NewStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})

	router.Use(sessions.Sessions("assign", store))

	router.GET("/register", register)
	router.GET("/login", login)
	router.POST("/register", user.CreateUser)
	router.POST("/login",user.VerifyUser)
	authorized := router.Group("/client")
	authorized.Use(AuthorizeRequest())
	{
		authorized.GET("/",profile.Show)
		authorized.POST("/profile", profile.CreateProfileHandler)
		authorized.POST("/pay", pay.Pay)
		authorized.GET("/logout", logout)
	}

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "error.html", gin.H{})
	})

	router.LoadHTMLGlob("templates/*")
	router.Static("/static/", "./static/")
	err = router.Run(":" + config.C.Common.ServerPort)
	if err != nil {
		log.Fatalf("error in starting server , err:%v", err)
	}
}

func AuthorizeRequest() gin.HandlerFunc {
	return func(g *gin.Context) {
		session := sessions.Default(g)
		v := session.Get("user-id")
		if v == nil {
			log.Println("client unauthorized")
			g.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Please login."})
			g.Abort()
		}
		g.Next()
	}
}

func register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func logout(g *gin.Context) {
	session := sessions.Default(g)
	session.Clear()
	session.Save()
	g.Redirect(http.StatusFound, "/login")
}
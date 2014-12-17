package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go/gongfu/api/spec"
	"go/gongfu/api/v1"
	"go/gongfu/cfg"
	"go/gongfu/util"
	"log"
	"net/http"
	"time"
)

// redirHttps is handlerFunc to redirect http traffic to https, while preserving the url
func redirHttps(w http.ResponseWriter, req *http.Request) {
	log.Printf("redirecting to https from %s via %s to host: %s", req.RemoteAddr, req.Method, req.Host)
	http.Redirect(w, req, "https://"+req.Host+req.RequestURI, http.StatusMovedPermanently)
}

// RequireAuth is a middleware to read the cookie and see if user has ACCESS_OK
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		hasAccess, email, userid := util.CookieHasAccess(c)
		if hasAccess == false {
			log.Printf("... denied access ... for %v", email)
			//c.JSON(401, gin.H{"message": "Required user authentication.", "status": sys.FAIL })
			c.Fail(401, errors.New("Unauthorized"))
		}
		log.Print("RequireAuth userid : " + userid)
		c.Set("email", email)
		c.Set("userid", userid)
		c.Next()
		// after request
		latency := time.Since(t)
		log.Print(latency)
	}
}

func main() {
	//SETUP
	cfg.Setup("cfg.json")
	util.SetupDB()
	util.SetupKeebler()

	g := gin.Default()

	// ROUTES
	g.GET("/", v1.GetApp)
	g.GET("/index", v1.GetIndex)
	g.GET("/logout", v1.GetLogout)

	g.POST("/login", v1.PostLogin)
	g.POST("/signup", v1.PostSignup)

	sp := g.Group("/spec")
	{
		//		sp.DELETE("/chirp/:id", spec.DeleteChirp)
		//		sp.GET("/chirp/:id", spec.GetChirpById)
		//		sp.GET("/chirps", spec.GetChirps)
		sp.GET("/user", spec.GetUser)
		//		sp.GET("/chirps-search/:term", spec.GetChirpsSearch)
		//		sp.PATCH("/chirp/:id", spec.PatchChirp)
		//		sp.POST("/chirp", spec.PostChirp)
	}

	api := g.Group("/api")
	api.Use(RequireAuth())
	{
		api.DELETE("/chirp/:id", v1.DeleteChirp)

		api.GET("/chirp/:id", v1.GetChirpById)
		api.GET("/chirps", v1.GetChirps)
		api.GET("/chirps-by-user/:id", v1.GetChirpsByUser)
		api.GET("/chirps-search/:term", v1.GetChirpsSearch)

		api.GET("/user", v1.GetUser)
		api.GET("/user/:id", v1.GetUserById)

		api.PATCH("/chirp", v1.PatchChirp)
		api.PATCH("/user", v1.PatchUser)

		api.POST("/chirp", v1.PostChirp)
	}

	// SERVICES
	go http.ListenAndServe(cfg.HTTP_PORT, http.HandlerFunc(redirHttps))
	log.Printf("... started " + cfg.SERVER + " redirects to Https ...")

	g.ServeFiles("/webapp/*filepath", http.Dir("webapp"))
	g.ServeFiles("/html/*filepath", http.Dir("webapp/html"))

	log.Printf("... started ServeFile ...")
	log.Printf("... serving %v TLS on port %v ", cfg.APPNAME, cfg.HTTP_PORT)
	g.RunTLS(cfg.HTTP_PORT, "certs/cert.pem", "certs/key.pem")

}

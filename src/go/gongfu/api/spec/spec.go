package spec

import (
	"../../model"
	"github.com/gin-gonic/gin"
	. "strconv"
)

type Msg struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(c *gin.Context) {
	m := Msg{
		Message: "This is GetUser",
		Status:  200,
	}
	c.JSON(200, m)
}

func GetUsers(c *gin.Context) {
	m := Msg{
		Message: "This is GetUsers",
		Status:  200,
	}
	c.JSON(200, m)
}

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong", "status": 200})
}

func PostLogin(c *gin.Context) {
	data := User{Name: "kyle", Email: "kyle@redbridge.io"}
	c.JSON(200, gin.H{"data": data, "status": "OK"})
}

func GetChirp(c *gin.Context) {
	id, err := Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(200, gin.H{"msg": "Cannot parse id param.", "status": "FAILED"})
	}
	payload := model.Chirp{Id: id, Userid: 102, Message: "Tweet tweet, said the mouse."}
	c.JSON(200, gin.H{"payload": payload, "status": "OK"})
}

func GetChirps(c *gin.Context) {
	chirp1 := model.Chirp{Id: 91, Userid: 101, Message: "One Chirp"}
	chirp2 := model.Chirp{Id: 92, Userid: 102, Message: "Another chirp."}
	payload := []model.Chirp{chirp1, chirp2}
	c.JSON(200, gin.H{"payload": payload, "status": "OK"})
}

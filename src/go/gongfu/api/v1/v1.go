package v1

import (
	"../../cfg"
	"../../model"
	"../../sys"
	"../../util"
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	. "strconv"
	//"strings"
	//"time"
)

func GetApp(c *gin.Context) {
	log.Printf("... GetApp ... ")
	http.Redirect(c.Writer, c.Request, "webapp/app.html#/", http.StatusMovedPermanently)
}

func GetIndex(c *gin.Context) {
	log.Printf("... GetIndex ... ")
	http.Redirect(c.Writer, c.Request, "webapp/index.html#/", http.StatusMovedPermanently)
}

func GetLogout(c *gin.Context) {
	log.Printf("... GetLogout ... ")
	SetCookieHandlerAccessRevoked(c.Writer, c.Request)
	http.Redirect(c.Writer, c.Request, "webapp/app.html#/", http.StatusMovedPermanently)
}

// App API

func DeleteChirp(c *gin.Context) {
	var id = c.Params.ByName("id")
	var chirp model.Chirp

	uid, _ := c.Get("userid")
	userid, _ := Atoi(uid.(string))
	chirp.Id, _ = Atoi(id)

	row := util.DB.QueryRow(chirp.StmtGetById())
	err := row.Scan(&chirp.Id, &chirp.Userid, &chirp.Type, &chirp.Message, &chirp.Create_date, &chirp.Active)
	if err != nil {
		c.JSON(400, sys.JSON_NOACCESS)
		return
	}

	// Check if chirp belongs to requester's account
	if chirp.Userid != userid {
		c.JSON(403, sys.JSON_NOACCESS)
		return
	}

	row = util.DB.QueryRow(chirp.StmtSetActive(false))
	err = row.Scan(&chirp.Id)
	if err != nil {
		c.JSON(500, gin.H{"message": sys.ERR_DB_DELETE, "status": sys.FAIL, "payload": id})
		return
	}
	c.JSON(200, gin.H{"payload": "Deleted Chirp " + Itoa(chirp.Id), "status": sys.SUCCESS})

}

func GetChirpById(c *gin.Context) {
	var id = c.Params.ByName("id")
	var chirp model.Chirp

	uid, _ := c.Get("userid")
	userid, _ := Atoi(uid.(string))
	chirp.Id, _ = Atoi(id)

	row := util.DB.QueryRow(chirp.StmtGetById())
	err := row.Scan(&chirp.Id, &chirp.Userid, &chirp.Type, &chirp.Message, &chirp.Create_date, &chirp.Active)
	if err != nil {
		c.JSON(400, sys.JSON_NOACCESS)
		return
	}

	// Check if chirp belongs to requester's account
	if chirp.Userid != userid {
		c.JSON(403, sys.JSON_NOACCESS)
		return
	}

	c.JSON(200, gin.H{"payload": chirp, "status": sys.SUCCESS})
}

func GetChirps(c *gin.Context) {
	var arr []model.Chirp
	var chirp model.Chirp

	rows, err := util.DB.Query(chirp.StmtSelectActive())
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_DB_NO_MATCH, "status": sys.FAIL})
		return
	}

	//util.PanicIf(err)
	for rows.Next() {
		var cp model.Chirp
		if err := rows.Scan(&cp.Id, &cp.Userid, &cp.Type, &cp.Message, &cp.Create_date, &cp.Active); err != nil {
			c.JSON(400, gin.H{"message": sys.ERR_SCAN_ROW, "status": sys.FAIL})
			return
		} else {
			arr = append(arr, cp)
		}
	}

	c.JSON(200, gin.H{"payload": arr, "status": sys.SUCCESS})
}

func GetChirpsByUser(c *gin.Context) {
	var arr []model.Chirp
	var chirp model.Chirp

	uid, _ := c.Get("userid")
	chirp.Id, _ = Atoi(uid.(string))

	log.Printf("... GetUser ... ")
	rows, err := util.DB.Query(chirp.StmtSelectByUserId())
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_DB_NO_MATCH, "status": sys.FAIL})
		return
	}

	//util.PanicIf(err)
	for rows.Next() {
		var cp model.Chirp
		if err := rows.Scan(&cp.Id, &cp.Type, &cp.Message, &cp.Create_date, &cp.Active); err != nil {
			c.JSON(400, gin.H{"message": sys.ERR_SCAN_ROW, "status": sys.FAIL})
			return
		} else {
			arr = append(arr, cp)
		}
	}

	c.JSON(200, arr)
}

//TODO: Implement
func GetChirpsSearch(c *gin.Context) {
	c.JSON(400, gin.H{"message": sys.ERR_DB_NO_MATCH, "status": sys.FAIL})
}

func GetUser(c *gin.Context) {
	var user model.User
	uid, _ := c.Get("userid")
	user.Id, _ = Atoi(uid.(string))
	log.Printf("... GetUser ... " + uid.(string))

	err := util.DB.QueryRow("select name, email from users where id=$1", user.Id).Scan(&user.Name, &user.Email)

	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_DB_NO_MATCH, "status": sys.FAIL})
	} else {
		c.JSON(200, gin.H{"payload": user, "status": sys.SUCCESS})
	}
}

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	log.Printf("... GetUserById ... " + id)

	var user model.User
	err := util.DB.QueryRow("select name, email from users where id=$1", id).Scan(&user.Name, &user.Email)
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_DB_NO_MATCH, "status": sys.FAIL})
	} else {
		c.JSON(200, gin.H{"payload": user, "status": sys.SUCCESS})
	}
}

func PatchChirp(c *gin.Context) {
	var chirp model.Chirp
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_READ_BODY, "status": sys.FAIL})
		return
	}
	err = json.Unmarshal(body, &chirp)
	if err != nil {
		c.JSON(500, gin.H{"message": sys.ERR_UNMARSHAL_BODY, "status": sys.FAIL, "payload": body})
		return
	}

	// TODO : Check if this Userid can access this payment
	row := util.DB.QueryRow(chirp.StmtUpdate())
	err = row.Scan(&chirp.Id)
	if err != nil {
		c.JSON(500, gin.H{"message": sys.ERR_DB_UPDATE, "status": sys.FAIL, "payload": body})
		return
	}
	log.Printf("======== type: " + chirp.Type)
	c.JSON(200, gin.H{"payload": "Updated Chirp " + Itoa(chirp.Id), "status": sys.SUCCESS})
}

//TODO: Implement
func PatchUser(c *gin.Context) {
	c.JSON(400, gin.H{"message": sys.ERR_DB_NO_MATCH, "status": sys.FAIL})
}

func PostChirp(c *gin.Context) {
	var chirp model.Chirp
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"message": sys.ERR_READ_BODY, "status": sys.FAIL})
		return
	}

	err = json.Unmarshal(body, &chirp)
	if err != nil {
		c.JSON(500, gin.H{"message": sys.ERR_UNMARSHAL_BODY, "status": sys.FAIL, "payload": body})
		return
	}

	uid, _ := c.Get("userid")
	chirp.Userid, _ = Atoi(uid.(string))

	row := util.DB.QueryRow(chirp.StmtInsert())
	err = row.Scan(&chirp.Id)
	if err != nil {
		c.JSON(500, gin.H{"message": sys.ERR_DB_INSERT, "status": sys.FAIL, "payload": body})
		return
	}
	c.JSON(200, gin.H{"payload": chirp.Id, "status": sys.SUCCESS})
}

func PostLogin(c *gin.Context) {
	var u model.User
	u.Email = c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	row := util.DB.QueryRow(u.StmtGetByEmail())
	err := row.Scan(&u.Id, &u.Accountid, &u.Name, &u.Email, &u.Password, &u.Active)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		log.Printf(".... auth fail from db " + u.Password)
		log.Printf(".... from form         " + password)
		http.Redirect(c.Writer, c.Request, "webapp/app.html?msg=loginfailed", http.StatusMovedPermanently)
	} else {
		SetCookieHandlerAccessOK(c, u.Email, u.Id)
		http.Redirect(c.Writer, c.Request, "webapp/app.html#/loggedin/"+u.Email, http.StatusMovedPermanently)
	}
}

func PostSignup(c *gin.Context) {
	var user model.User
	user.Name = c.Request.FormValue("name")
	user.Email = c.Request.FormValue("email")
	formpass := c.Request.FormValue("password")
	_ = user.HashPassword(formpass)

	stmt, err := util.DB.Prepare(user.StmtInsert())
	CatchErrorInsertDB(err, c)

	_, err = stmt.Exec()
	CatchErrorInsertDB(err, c)
	newId := 0

	defer stmt.Close()
	http.Redirect(c.Writer, c.Request, "webapp/app.html#/?msg=new+user+"+Itoa(newId), http.StatusMovedPermanently)
}

// SetCookieHandlerAccessOK sets the user's cookie with ACCESS_OK and email
func SetCookieHandlerAccessOK(c *gin.Context, email string, userid int) {
	value := map[string]string{
		sys.COOKIE_APP_ACCESS: sys.ACCESS_OK,
		sys.COOKIE_EMAIL:      email,
		sys.COOKIE_USERID:     Itoa(userid),
	}
	if encoded, err := util.Keebler.Encode(cfg.COOKIE_NAME, value); err == nil {
		cookie := &http.Cookie{
			Name:  cfg.COOKIE_NAME,
			Value: encoded,
			Path:  "/",
		}
		log.Printf("Setting cookie: %v", cfg.COOKIE_NAME)
		http.SetCookie(c.Writer, cookie)
		log.Printf("cookie set %v with ACCESS_OK : ", cfg.COOKIE_NAME)
	} else {
		log.Printf("FAILED TO set cookie!! for %v", email)
	}
}

// SetCookieHandlerAccessRevoked resets the user's cookie with blanks
func SetCookieHandlerAccessRevoked(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		sys.COOKIE_APP_ACCESS: sys.BLANK,
		sys.COOKIE_EMAIL:      sys.BLANK,
		sys.COOKIE_USERID:     sys.BLANK,
	}
	if encoded, err := util.Keebler.Encode(cfg.COOKIE_NAME, value); err == nil {
		cookie := &http.Cookie{
			Name:  cfg.COOKIE_NAME,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

//// Error Handlers ////
func CatchErrorInsertDB(err error, c *gin.Context) {
	if err != nil {
		log.Fatal(err)
		http.Redirect(c.Writer, c.Request, "webapp/signup.html?msg=creationfailed", http.StatusConflict)
	}
}

func CatchErrorRedirect(err error, c *gin.Context, path string, code int) {
	if err != nil {
		http.Redirect(c.Writer, c.Request, path, code)
	}
}

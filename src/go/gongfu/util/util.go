package util

import (
	"go/gongfu/cfg"
	"go/gongfu/sys"
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	_ "github.com/lib/pq"
	"io"
	"log"
	"math"
	"time"
)

var DB *sql.DB
var Keebler *securecookie.SecureCookie

func SetupDB() {
	settings := "user=" + cfg.DB_USER + " password=" + cfg.DB_PASSWORD + " dbname=" + cfg.DB_NAME + " sslmode=disable"
	db, err := sql.Open("postgres", settings)

	PanicIf(err)
	DB = db
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
	log.Printf("DB has been pinged")
}

func SetupKeebler() {
	Keebler = securecookie.New([]byte(cfg.HASHKEY), []byte(cfg.BLOCKKEY))
	log.Printf("Keebler has been created.... ready to make cookies...")
}

func GenHash(c int) string {
	//c := 8
	b := make([]byte, c)
	n, err := io.ReadFull(rand.Reader, b)
	if n != len(b) || err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func CookieHasAccess(c *gin.Context) (bool, string, string) {
	if cookie, err := c.Request.Cookie(cfg.COOKIE_NAME); err == nil {
		value := make(map[string]string)
		if err = Keebler.Decode(cfg.COOKIE_NAME, cookie.Value, &value); err == nil {
			log.Printf("The value of access is " + value[sys.COOKIE_APP_ACCESS])
			log.Printf("The value of email is " + value[sys.COOKIE_EMAIL])
			log.Printf("Access value of [COOKIE_APP_ACCESS]: %v", value[sys.COOKIE_APP_ACCESS])
			if value[sys.COOKIE_APP_ACCESS] == sys.ACCESS_OK {
				return true, value[sys.COOKIE_EMAIL], value[sys.COOKIE_USERID]
			}
		}
	}
	log.Printf("Where is the cookie? %v", cfg.COOKIE_NAME)
	return false, "", ""
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// Schedule is a generic timer func that will run a func after a delay
func Schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
	// stop is the channel that will stop if you do: stop <- true
}

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

func Round(val float64, prec int) float64 {
	var rounder float64
	intermed := val * math.Pow(10, float64(prec))

	if val >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / math.Pow(10, float64(prec))
}

// Btc_Sat converts BTC float to Satochi Int; 100 million sat = 1BTC
// int64 value +/- 9223372036854775807
func Btc_Sat(btc float64) (sat int64, err error) {
	num := btc * 100000000
	if (num > 9223372036854775807) || (num < -9223372036854775807) {
		num = -1
		err = errors.New("Out of scale for int64")
	}
	return int64(num), err
}

// TODO write check for float64 size
func Sat_Btc(sat int64) (btc float64, err error) {
	num := float64(sat)
	if (num > 9223372036854775807) || (num < -9223372036854775807) {
		num = -1
		err = errors.New("Out of scale for float64")
	}
	btc = num / 100000000
	return btc, err
}

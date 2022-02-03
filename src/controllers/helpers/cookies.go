package helpers

import (
	"net/http"
	"time"

	"github.com/paulobezerra/goblog/src/models"
)

var CookieNameWithAuthToken string = "auth_token"

func SaveCookie(name string, value string, w http.ResponseWriter) {
	cookie := http.Cookie{Name: name, Value: value, Expires: time.Now().Add(30 * 24 * time.Hour), Secure: true}
	http.SetCookie(w, &cookie)
}

func DropCookie(name string, w http.ResponseWriter) {
	cookie := http.Cookie{Name: name, Value: "", Expires: time.Unix(0, 0), Secure: true}
	http.SetCookie(w, &cookie)
}

func GetUserByCookieName(cookieName string, r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	user, err := GetUserByJWT(cookie.Value)

	if err != nil {
		return nil, err
	}

	return user, err
}

package cookiestore

import (
	"net/http"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// New func
func New() CookieStore {
	return make(CookieStore)
}

// CookieName const
const CookieName = "session-id"

func newCookie(value string) *http.Cookie {
	c := &http.Cookie{
		Name:  CookieName,
		Value: value,
		Path:  "/",
		// 48 hours
		MaxAge:   60 * 60 * 48,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return c
}

// UserData type
type UserData struct {
	Guest bool
	User  *forum.User
}

// CookieStore type
type CookieStore map[uuid.UUID]UserData

// SetNewGuestCookie func
func (c CookieStore) SetNewGuestCookie() *http.Cookie {
	id := uuid.Must(uuid.NewV4())
	c.Set(id, UserData{Guest: true, User: nil})
	cookie := newCookie(id.String())
	return cookie
}

// SetNewCookie func
func (c CookieStore) SetNewCookie(user *forum.User) *http.Cookie {
	if sid, ok := c.userSID(user.ID); ok {
		c.clear(sid)
	}

	id := uuid.Must(uuid.NewV4())
	c.Set(id, UserData{Guest: false, User: user})
	cookie := newCookie(id.String())

	return cookie
}

func (c CookieStore) clear(sid uuid.UUID) {
	delete(c, sid)
}

func (c CookieStore) userSID(userID uuid.UUID) (uuid.UUID, bool) {
	for sid, data := range c {
		if data.User != nil && data.User.ID == userID {
			return sid, true
		}
	}

	return uuid.UUID{}, false
}

// Get func
func (c CookieStore) Get(cookieID uuid.UUID) (UserData, bool) {
	data, ok := c[cookieID]
	return data, ok
}

// Set func
func (c CookieStore) Set(cookieID uuid.UUID, data UserData) {
	c[cookieID] = data
}

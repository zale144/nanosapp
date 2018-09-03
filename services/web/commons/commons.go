package commons

import "github.com/dgrijalva/jwt-go"

const (
	CookieName          = "Nanocorp.Challenge"
	SECRET              = "$P$Bd2WdVjaRR/De58OX2qVu3XA6aiPaf."
	HEADER_AUTH_USER_ID = "Auth-User-Id"
	CRYPT_SETTING = "$P$BwQZDcQaNU9zAOF.6MOUdEhz9X68fL1"
)

var (
	ApiURL string
)

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
package secure

import (
	"errors"
	"time"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
	"github.com/dgrijalva/jwt-go"

	"github.com/rg-km/final-project-engineering-11/backend/config"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blowfish"
)

func VerifyPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateToken(username string, id int, role string) (string, error) {
	timer := time.Now().Add(config.Configuration.JWT_EXPIRATION_DURATION).Unix()
	secret := config.Configuration.JWT_SECRET
	claims := jwt.MapClaims{}

	claims["role"] = role
	claims["id"] = id
	claims["username"] = username
	claims["exp"] = int(timer)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyCookie(hash string, req string) (string, error) {

	if err := VerifyPassword(hash, req); err != nil {
		return "", errors.New("User Unauthorized")
	}

	return hash, nil

}

func Verifyid(data int, req int) error {
	if data != req {
		return errors.New("User Unauthorized")
	}
	return nil
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Configuration.JWT_SECRET), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		timer := time.Now().Add(config.Configuration.JWT_EXPIRATION_DURATION).Unix()
		claims["exp"] = int(timer)
		tokenString, err := token.SignedString([]byte(config.Configuration.JWT_SECRET))
		if err != nil {
			return "", err
		}

		return tokenString, nil
	}
	return "", errors.New("Token not valid")
}

func Encrypt(pt, key []byte) string {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return "error"
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt)
	if err != nil {
		return "error"
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return string(ct)
}

func Decrypt(ct, key []byte) (string, error) {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return "", errors.New("UnAuthorized")
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	if len(pt)%8 != 0 {
		return "", errors.New("UnAuthorized")
	}
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		return "", errors.New("UnAuthorized")
	}
	return string(pt), nil
}

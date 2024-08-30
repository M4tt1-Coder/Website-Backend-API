package authenticator

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Auth_key"] != nil {
			//get the api key from the environment variable
			envs, err := godotenv.Read(".env")
			if err != nil {
				log.Printf("Error loading.env file")
			}
			//check if the api key is the same as the one in the header
			//the auth_key is encrypted
			if decrypt(r.Header["Auth_key"][0], envs["DECRYPTION_KEY"]) == envs["API_KEY"] {
				log.Printf("authenticated")
				//call the next handler
				next.ServeHTTP(w, r)
			} else {
				log.Printf("not authenticated")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not authenticated"))
			}
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

// decrytion
// source: https://gist.github.com/manishtpatel/8222606

func decrypt(cryptoText string, key string) string {
	Key := []byte(key)

	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(Key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

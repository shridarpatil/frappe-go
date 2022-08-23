package frappe


import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/fernet/fernet-go"
)


func Authorize(r *http.Request) error {
	token := r.Header.Get("Authorization")

	if token == "" {
		err := fmt.Errorf("Header Authorization not found!",)
		return err
	}

	authorization := strings.Split(token, " ")
	switch authorization[0]{
		case "token":
			api := strings.Split(authorization[1], ":")
			var user = session{}

			Frappe.Db.Get(
				&user,
				`SELECT u.name as user, u.email, a.password
					FROM "tabUser" u
					INNER JOIN "__Auth" as a on u.name = a.name
						AND a.doctype = 'User' AND a.fieldname = 'api_secret'
					WHERE u.api_key = $1  limit 1`, api[0])

			Frappe.Session = user

			if api[1] != Decrypt(user.Password){
				err := fmt.Errorf("Authorization failed!",)
				return err
			}
	}

	return nil
}


func Decrypt(token string) string {
	key, _ := base64.URLEncoding.DecodeString(Frappe.Config.EncryptionKey)

	k, _ := fernet.DecodeKey(hex.EncodeToString(key))

	tok := []byte(token)
			// require.NoError(t, err)


	msg := fernet.VerifyAndDecrypt(tok, 0, []*fernet.Key{k})
	return string(msg)
}
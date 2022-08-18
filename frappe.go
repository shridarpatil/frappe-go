package frappe

import (
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/jmoiron/sqlx"

	"net/http"

	"fmt"
	"strings"

	"github.com/kr/fernet"
	"time"
	// "encoding/base64"
)

var Frappe *frappe

type frappe struct {
	Server 	*rpc.Server
	Db 		*sqlx.DB
}


func init() {
	// Initialize frappe
	fmt.Println("\033[33m Initializing frappe: \033[0m")
	Frappe = &frappe{}
	Frappe.Server = rpc.NewServer()
	Frappe.Server.RegisterCodec(json.NewCodec(), "application/json")
	Frappe.Db = InitDB("postgres", "host=172.17.0.1 port=5432 user=crm password=Gorv71YDDqYaW0kl dbname=crm sslmode=disable")

	fmt.Println("\033[;32m Frappe initialized: \033[0m")
}

func (s *frappe) Ping() string{

	return "Pong"
}


func Authorize(r *http.Request) error{
	token := r.Header.Get("Authorization")

	if token == "" {
		err := fmt.Errorf("Header Authorization not found!",)
		return err
	}

	authorization := strings.Split(token, " ")
	switch authorization[0]{
		case "token":
			fmt.Println("%v ---- ", authorization[0])
			api := strings.Split(authorization[1], ":")

	// 		api_secret, _ := Decrypt("gAAAAABiBLP3FdpPKsoBGf4dRkFIZnRIbNgvW-Y7CE3gtcKhfWTJODWk7Am9xhijW082hzG81HKHnr9Nc7JLAGItRGX2eto02g==")
	// // 		if err1 == false{
	// // 			err1 := fmt.Errorf("Api secret match!",)
				// // return err1

	// // 		}

	// 		fmt.Println("%s ---- ", api_secret)

			if api[0] != "123" || api[1] != "456"{
				err := fmt.Errorf("Token miss match!",)
				return err
			}
	}

	return nil
}



func Decrypt(token string) (string, bool) {

	// btok, err := base64.URLEncoding.DecodeString(token)

	// if err != nil {
	// 	fmt.Println(" +++++++++++++++++++++++++++++++++++")
	// 	return "", false

	// }
	// fmt.Println(" +++++++++++++++++++++++++++++++++++ %v", btok)
	k := fernet.MustDecodeKeys("sEgBb3h1KKIlGayaGUem65aowNkGQp_3WgWqYnONMa4=")

	email := fernet.VerifyAndDecrypt([]byte(token), 60*time.Second, k)

	fmt.Println(" +++++++++++++++++++++++++++++++++++ %v", email)
	return string(email), true
}


func New(driver, dsn string) *frappe {
	var frappe = &frappe{}
	frappe.Server = rpc.NewServer()
	frappe.Server.RegisterCodec(json.NewCodec(), "application/json")

	frappe.Db = InitDB(driver, dsn)

	return frappe
}

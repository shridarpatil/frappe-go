package main

import (
	"github.com/shridarpatil/frappe"
	"github.com/shridarpatil/frappe/example/api"
	"net/http"
	"fmt"
)

type HelloService struct{
	*frappe.Frappe
}


type HelloReply struct {
	Message string
}


type HelloArgs struct {
	Who string
}


type Place struct {
	Name 	string
	Owner	string
}


type Opt struct {
	api      *api.Api
}


func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	err := h.Authorize(r)

	if err != nil{
		return err
	}

	reply.Message = "Hello, " + args.Who + "!"
	h.Log.NoticeF("args: %v\nreply: %v, \n", "", h)

	fmt.Println(h.Session.User)
	var jason = Place{}
	h.Db.Get(&jason, `SELECT name, owner FROM "tabUser" limit 1 `)


	return nil
}


func (h *HelloService) Validate(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	// Frappe.Authorize(r)
	return nil
}




func main() {

	var config = &frappe.SiteConfig{
		Driver: "postgres",
		DSN: "host=172.17.0.1 port=5432 user=crm password=Gorv71YDDqYaW0kl dbname=crm sslmode=disable",
		EncryptionKey: "sEgBb3h1KKIlGayaGUem65aowNkGQp_3WgWqYnONMa4=",
		SetLogLevel: "INFO",
	}
	var app = frappe.New(
		config,
	)
	// Register methods for rpc
	var opt = &Opt{}
	opt.api = api.New(app)

	app.RegisterService(&HelloService{app}, "")
	app.RegisterService(opt.api, "")

	http.Handle("/rpc", app.GetServer())
	http.ListenAndServe("localhost:10000", nil)

}
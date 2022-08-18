package main

import "github.com/user/frappe"

import "net/http"
import "fmt"
import "log"



type HelloService struct{}

type HelloReply struct {
	Message string
}

type HelloArgs struct {
	Who string
}


type Place struct {
    Name	 string
    Owner     string
}


func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	err := frappe.Authorize(r)

	if err != nil{
		return err
	}
	reply.Message = "Hello, " + args.Who + "!"
	log.Printf("args: %v\nreply: %v, \n %v", r, r.Header.Get("Authorization"), frappe.Frappe)

	// frappe.InitDB("mysql", "root:r00t@(172.17.0.1:3306)/crux")
	fmt.Println(frappe.Frappe.Ping())
	var jason = Place{}
    frappe.Frappe.Db.Get(&jason, `SELECT name, owner FROM "tabUser" limit 1 `)
    fmt.Printf("%#v\n", jason)


	return nil
}



func main() {

	// Register methods for rpc
	frappe.Frappe.Server.RegisterService(new(HelloService), "")

	http.Handle("/rpc", frappe.Frappe.Server)
	http.ListenAndServe("localhost:10000", nil)

}
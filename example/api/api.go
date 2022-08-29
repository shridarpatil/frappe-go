package api
import "github.com/user/frappe"
import (
	"log"
	"net/http"
)


type Api struct{
	*frappe.Frappe
}


// Api args
type Args struct {
	Who string
	Me string
}


// Api response
type Response struct {
	Message string
}


func (h *Api) Who(r *http.Request, args *Args, reply *Response) error {
	err := h.Authorize(r)

	if err != nil{
		return err
	}

	h.Log.Debug(h.Session.User)

	reply.Message = "Hello, " + args.Who + "!"
	log.Printf("request: %v\nargs: %v\nreply: %v", r, args, reply)
	return nil
}


func (h *Api) Ping(r *http.Request, args *Args, reply *Response) error {
	reply.Message = "Pong"
	log.Printf("request: %v\nargs: %v\nreply: %v", r, args, reply)
	return nil
}


// New API object.
func New(f  *frappe.Frappe) *Api {

	return &Api{f}
}

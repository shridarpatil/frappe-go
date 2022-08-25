# Frappe-go


Frappe go ,
Write apis for frappe framework in go

## Features

- Token based authetication
- RPC

## Tech

Frappe-go uses a number of open source projects to work properly:

- [Gorilla RPC](https://github.com/gorilla/rpc) - A golang foundation for RPC over HTTP services.
- [Sqlx](https://github.com/jmoiron/sqlx) - General purpose extensions to golang's database/sql.
- [koanf](https://github.com/knadh/koanf) - Simple, lightweight, extensible, configuration management library for Go.
- [Logger](https://github.com/apsdehal/go-logger) - Simple logger for Go programs. Allows custom formats for messages.

And of course Frappe-go itself is open source with a [public repository][dill]
 on GitHub.

## Installation

Frappe-go requires [Go](https://go.dev/doc/install) to run.

Install the dependencies and devDependencies and start the server.

```go
package main

import "github.com/shridarpatil/frappe-go"
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


type User struct {
	Name 	string
	Owner	string
}


func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	err := frappe.Authorize(r)

	if err != nil{
		return err
	}
	reply.Message = "Hello, " + args.Who + "!"
	log.Printf("args: %v\nreply: %v, \n %v", r, r.Header.Get("Authorization"), frappe.Frappe)

	fmt.Println(frappe.Frappe.Ping())
	var user = User{}
	frappe.Frappe.Db.Get(&user, `SELECT name, owner FROM "tabUser" limit 1 `)
	fmt.Printf("%#v\n", user)


	return nil
}

func main() {

	// Register methods for rpc
	frappe.Frappe.Server.RegisterService(new(HelloService), "")

	http.Handle("/rpc", frappe.Frappe.Server)
	http.ListenAndServe("localhost:10000", nil)

}
```
Save above code in main.go and run 

```sh
go run main.go
```

```sh
localhost:10000/rpc
```

```sh
curl --location --request POST 'http://localhost:10000/rpc' \
--header 'Authorization: token <api_key>:<api_secret>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "method":"HelloService.Say",
    "params":[{"Who":"Shri"}],
    "id":"1"
}'
```

```json
{
    "result": {
        "Message": "Hello, Shri!"
    },
    "error": null,
    "id": "1"
}
```

### Stats

Python based api took around 100ms on an average 
![image](https://user-images.githubusercontent.com/11792643/186417332-e9db9270-032b-4adb-bdc3-95add5fe1c84.png)


Golang based api took around less then 10ms on an average
![image](https://user-images.githubusercontent.com/11792643/186417563-21cd8c20-f379-4cc6-b608-794ee7ec746c.png)

## License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/shridarpatil/frappe-go>
   [git-repo-url]: <https://github.com/shridarpatil/frappe-go.git>

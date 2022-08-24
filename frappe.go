package frappe

import (
	// "os"
	// "fmt"
	// "strings"
	// "net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/jmoiron/sqlx"
	"github.com/apsdehal/go-logger"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	// "github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"

	// "encoding/base64"
)

var Frappe *frappe

type frappe struct {
	Server 	*rpc.Server
	Db 		*sqlx.DB
	Log 	*logger.Logger
	Session  session
	Config   SiteConfig
}

type config struct {
	Site     SiteConfig
}

type session struct {
	User 	    string
	// FullName    string
	Email	    string
	Password    string
}


func init() {
	// Initialize frappe
	log, err := logger.New("frappe", 1)
	log.SetLogLevel(6)
	log.SetFormat("#%{id} [%{time}] [%{level}] [%{filename}:%{line}] ▶ %{message}")

	ko := koanf.New(".")
	var conf config

	if err := ko.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.ErrorF("error reading config: %v", err)
		panic("error reading config:")
	}

	if err := ko.Unmarshal("app", &conf); err != nil {
		log.ErrorF("error while parsing config: %v", err)
	}

	if err != nil {
		panic(err) // Check for error
	}


	Frappe = &frappe{}

	Frappe.Log = log
	Frappe.Config = conf.Site

	Frappe.Log.Info("Initializing frappe:")

	Frappe.Server = rpc.NewServer()
	Frappe.Server.RegisterCodec(json.NewCodec(), "application/json")


	Frappe.Db = InitDB(Frappe.Config.Driver, Frappe.Config.DSN)
	Frappe.Log.Notice("Frappe initialized: ")
}

func (s *frappe) Ping() string{

	return "Pong"
}


func (s *frappe) HasRole(role string, user ...string) bool{
	// Retrun true or false if user has the role
	// Default user is the session user

	query := `SELECT name
		FROM "tabHas Role"
		WHERE parent = $1
			and parenttype = 'User'
			and role = $2`
	var id string
	_user := Frappe.Session.User

	Frappe.Log.DebugF("%v -- ", user)

	if len(user) != 0{
		_user = user[0]
	}

	err := Frappe.Db.Get(&id, query, _user, role)
	if err != nil{
		return false
	}
	return true
}


func(s *frappe) GetRoles(user ...string) []string{
	// Retrun list of roles for specified user
	// Default user is the session user

	query := `SELECT role
		FROM "tabHas Role"
		WHERE parent = $1
			and parenttype = 'User'`
	_user := Frappe.Session.User

	Frappe.Log.DebugF("%v -- ", user)

	if len(user) != 0{
		_user = user[0]
	}

	var role []string
	err := Frappe.Db.Select(&role, query, _user)

	Frappe.Log.DebugF("%v -- ", role)
	if err != nil{
		return nil
	}

	return role
}


func New(driver, dsn string) *frappe {
	var frappe = &frappe{}
	frappe.Server = rpc.NewServer()
	frappe.Server.RegisterCodec(json.NewCodec(), "application/json")

	frappe.Db = InitDB(driver, dsn)

	return frappe
}

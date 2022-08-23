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



func New(driver, dsn string) *frappe {
	var frappe = &frappe{}
	frappe.Server = rpc.NewServer()
	frappe.Server.RegisterCodec(json.NewCodec(), "application/json")

	frappe.Db = InitDB(driver, dsn)

	return frappe
}

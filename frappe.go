package frappe

import (


	"github.com/shridarpatil/rpc"
	"github.com/shridarpatil/rpc/json"
	"github.com/jmoiron/sqlx"
	"github.com/apsdehal/go-logger"

)


type Frappe struct {
	server 	*rpc.Server
	Db		*sqlx.DB
	Log		*logger.Logger
	Session	session
	Config	*SiteConfig
}


type config struct {
	Site	SiteConfig
}


type session struct {
	User 	    string
	// FullName    string
	Email	    string
	Password    string
}


func (f *Frappe) GetServer() *rpc.Server{
	return f.server
}


func (f *Frappe) Ping() string{
	return "Pong"
}


func (f *Frappe) RegisterService(receiver interface{}, name string) error {
	return f.server.RegisterService(receiver, name)
}


func (f *Frappe) HasRole(role string, user ...string) bool{
	// Retrun true or false if user has the role
	// Default user is the session user

	query := `SELECT name
		FROM "tabHas Role"
		WHERE parent = $1
			and parenttype = 'User'
			and role = $2`
	var id string
	_user := f.Session.User

	if len(user) != 0{
		_user = user[0]
	}

	err := f.Db.Get(&id, query, _user, role)
	if err != nil{
		return false
	}
	return true
}


func(f *Frappe) GetRoles(user ...string) []string{
	// Retrun list of roles for specified user
	// Default user is the session user

	query := `SELECT role
		FROM "tabHas Role"
		WHERE parent = $1
			and parenttype = 'User'`
	_user := f.Session.User

	if len(user) != 0{
		_user = user[0]
	}

	var role []string
	err := f.Db.Select(&role, query, _user)

	if err != nil{
		return nil
	}

	return role
}



func New(config *SiteConfig) *Frappe {
	var frappe = &Frappe{
		Config: config,
	}

	log, err := logger.New("frappe", 1)


	if config.SetFormat == ""{
		config.SetFormat = "#%{id} [%{time}] [%{level}] [%{filename}:%{line}] ▶ %{message}"
	}

	switch config.SetLogLevel{
		case "CRITICAL":
		log.SetLogLevel(logger.CriticalLevel)
		case "ERROR":
		log.SetLogLevel(logger.ErrorLevel)
		case "WARNING":
		log.SetLogLevel(logger.WarningLevel)
		case "NOTICE":
		log.SetLogLevel(logger.NoticeLevel)
		case "INFO":
		log.SetLogLevel(logger.InfoLevel)
		case "DEBUG":
		log.SetLogLevel(logger.DebugLevel)
		default:
		log.SetLogLevel(logger.DebugLevel)
	}

	log.SetFormat(config.SetFormat)

	if err != nil {
		panic(err) // Check for error
	}


	frappe.Log = log
	frappe.Log.Info("Initializing frappe:")
	frappe.server = rpc.NewServer()
	frappe.server.RegisterCodec(json.NewCodec(), "application/json")
	frappe.Config.EncryptionKey = config.EncryptionKey
	frappe.initDB()

	return frappe
}

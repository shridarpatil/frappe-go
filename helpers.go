package frappe

import (
	"context"
	"errors"
	"fmt"
)

type SiteConfig struct {
	DSN 		      	string
	Driver    			string
	EncryptionKey    	string
	SetFormat			string
	SetLogLevel         string
}


func (f *Frappe) BeforeFind(ctx context.Context) (err error) {
  if ctx == nil {
    err = errors.New("can't save invalid data")
  }
  fmt.Println("BeforeFind %v", f)
  return
}
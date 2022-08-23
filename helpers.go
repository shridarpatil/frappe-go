package frappe


type SiteConfig struct {
	DSN 		      	string `koanf:"db_dsn"`
	Driver    			string `koanf:"db_driver"`
	EncryptionKey    	string `koanf:"encryption_key"`
}
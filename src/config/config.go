package config

// Config struct
type Config struct {
	WebServer struct {
		Port int
	}
	Database struct {
		Host string
		Port int
	}
}

// func GetConfig() {
// 	os.
// }

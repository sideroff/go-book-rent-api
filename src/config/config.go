package config

// Roles - Enum that specifies the access rights of the app, admin > guest
var Roles = struct{
    Guest int
	User int
	Admin int
}{
    Guest: 0,
	User: 1,
	Admin: 2,
}

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
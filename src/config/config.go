package config


// Role - role in the app, admin > guest
type Role int

// Roles - Enum that specifies the acciss rights of the app
var Roles = struct{
    Guest Role
	User Role
	Admin Role
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

// func GetConfig() {
// 	os.
// }

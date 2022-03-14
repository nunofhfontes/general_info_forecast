package datamodels

type AuthInfo struct {
	User     string
	Password string
	Name     string
	Role     string
	Admin    bool

	// TODO - add more info
}

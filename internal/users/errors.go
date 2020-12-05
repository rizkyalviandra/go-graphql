package users

// WrongUsernameOrPasswordError struct
type WrongUsernameOrPasswordError struct {}


func (m *WrongUsernameOrPasswordError) Error() string {
	return "Incorrect Username or Password"
}
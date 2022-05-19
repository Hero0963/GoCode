package p15

var UserLogin = map[string]string{
	"hero": "abc$123",
}

func addNewUser(key string, value string) {
	UserLogin[key] = value
}

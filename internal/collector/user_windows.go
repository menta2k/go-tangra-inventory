package collector

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func GetUserInfo() (string, error) {
	var token windows.Token
	err := windows.OpenProcessToken(
		windows.CurrentProcess(),
		windows.TOKEN_QUERY,
		&token,
	)
	if err != nil {
		return "unknown", err
	}
	defer token.Close()

	user, err := token.GetTokenUser()
	if err != nil {
		return "unknown", err
	}

	account, domain, _, err := user.User.Sid.LookupAccount("")
	if err != nil {
		return "unknown", err
	}

	fmt.Printf("User: %s\\%s\n", domain, account)

	return account, nil

}

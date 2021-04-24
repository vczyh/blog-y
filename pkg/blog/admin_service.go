package blog

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func LoginService(username, password string) (string, error) {
	usernameVal, err := GetConfigValueService("username")
	if err != nil {
		l.Error("failed  to get config value", "name", username, "error", err)
		return "", err
	}

	passwordVal, err := GetConfigValueService("password")
	if err != nil {
		l.Error("failed  to get config value", "name", password, "error", err)
		return "", err
	}

	if username != usernameVal || password != passwordVal {
		return "", fmt.Errorf("username or password incorrect")
	}

	token := base64.URLEncoding.EncodeToString([]byte(username + ":" + password))

	return token, nil
}

func AuthService(token string) (bool, error) {
	b, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		l.Error("failed to base64 decode", "string", token, "error", err)
		return false, nil
	}

	arr := strings.Split(string(b), ":")
	if len(arr) != 2 {
		return false, nil
	}
	username := arr[0]
	password := arr[1]

	usernameVal, err := GetConfigValueService("username")
	if err != nil {
		l.Error("failed  to get config value", "name", username, "error", err)
		return false, err
	}
	passwordVal, err := GetConfigValueService("password")
	if err != nil {
		l.Error("failed  to get config value", "name", password, "error", err)
		return false, err
	}
	if username == usernameVal && password == passwordVal {
		return true, nil
	}

	return false, nil
}

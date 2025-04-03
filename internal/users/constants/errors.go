package constants

import (
	"errors"
)

var (
	ErrSignIn       = errors.New("ログインに失敗しました。ユーザー名とパスワードを確認してください。")
	ErrSignInServer = errors.New("ログインに失敗しました。しばらくしてからもう一度お試しください。")
)

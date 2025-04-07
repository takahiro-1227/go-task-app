package constants

import (
	"errors"
	"strconv"
)

const (
	MaxPasswordLength = 32
	MinPasswordLength = 10
)

var (
	ErrSignIn                    = errors.New("ログインに失敗しました。ユーザー名とパスワードを確認してください。")
	ErrSignInServer              = errors.New("ログインに失敗しました。しばらくしてからもう一度お試しください。")
	ErrOverPasswordLength        = errors.New("パスワードは" + strconv.Itoa(MaxPasswordLength) + "文字以下で入力してください。")
	ErrLessPasswordLength        = errors.New("パスワードは" + strconv.Itoa(MinPasswordLength) + "文字以上で入力してください。")
	ErrDuplicatedUserName        = errors.New("入力されたユーザー名は既に存在します。")
	ErrPasswordCharacterCategory = errors.New("パスワードは英数字記号を含める必要があります。")
)

const (
	ErrSuffixRequiredInput = "の入力は必須です。"
)

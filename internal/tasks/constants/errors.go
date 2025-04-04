package constants

import "errors"

var (
	ErrGetTasks        = errors.New("タスクの取得に失敗しました。")
	ErrTitleIsEmpty    = errors.New("タスクのタイトルを入力してください")
	ErrDuplicatedTitle = errors.New("タスクのタイトルが重複しています")
	ErrCreateFailed    = errors.New("タスクの作成に失敗しました")
)

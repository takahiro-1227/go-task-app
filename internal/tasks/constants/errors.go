package constants

import "errors"

var (
	ErrGetTasks      = errors.New("タスクの取得に失敗しました。")
	ErrTitleIsEmpty  = errors.New("タスクのタイトルを入力してください。")
	ErrCreateFailed  = errors.New("タスクの作成に失敗しました。")
	ErrUpdateFailed  = errors.New("タスクの更新に失敗しました。")
	ErrDeleteFailed  = errors.New("タスクの削除に失敗しました。")
	ErrInvalidUpdate = errors.New("不正なタスク変更です。")
	ErrInvalidDelete = errors.New("不正なタスク削除です。")
)

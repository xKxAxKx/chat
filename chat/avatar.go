package main

import (
	"errors"
)

// ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatarはユーザーのプロフィール画像を表す型です。
type Avatar interface {
	/*
	   指定されたクライアントのアバターのURLを返す
	   問題が発生した場合にはエラーを返す
	   特に、URLを取得できなかった場合にはErrNoAvatarURLを返す
	*/
	GetAvatarURL(c *client) (string, error)
}

package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っているユーザを指す
type client struct {
	// このユーザのためのWebSocket
	socket *websocket.Conn

	// メッセージが送られるチャネル
	send chan []byte

	// クライアントが参加しているチャットルーム
	room *room
}

/*
  clientがWebSocketからReadMessageを使ってデータを読み込むために使う
  受け取ったメッセージはすぐにroomのforwardチャネルに送られる
  何かしらのエラーが起こった場合、ループから抜けてWebSocketを閉じる
*/
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

/*
  継続的にsendチャネルからメッセージを受け取る
  WebSocketのWriteMessageメソッドを使って書き出す
  書き込みが失敗するとbreakでforから抜けて、WebSocketを閉じる
*/
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// clientはチャットを行っているユーザを指す
type client struct {
	socket   *websocket.Conn        // このユーザのためのWebSocket
	send     chan *message          // メッセージが送られるチャネル
	room     *room                  // クライアントが参加しているチャットルーム
	userData map[string]interface{} // クライアントのユーザデータ
}

/*
  clientがWebSocketからReadMessageを使ってデータを読み込むために使う
  受け取ったメッセージはすぐにroomのforwardチャネルに送られる
  何かしらのエラーが起こった場合、ループから抜けてWebSocketを閉じる
*/
func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
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
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

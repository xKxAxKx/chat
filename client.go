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

package main

type room struct {
  // 他のクライアントに転送するためのメッセージを保持するチャネル
  forward chan [] byte

  // チャットルームに参加しようとしているクライアントのためのチャネル
  join chan *client

  // チャットルームから退室しようとしているクライアントのためのチャネル
  leave chan *client

  // 在室している全てのクライアントを保持
  clients map[*client]bool
}

func (r *room) run() {
  for {
    select {
    case client := <-r.join:
      // 参加
      r.clients[client] = true
    case client := <- r.leave:
      // 退室
      delete(r.clients, client)
      close(client.send)
    case msg := <- r.forward:
      // 全てのクライアントにメッセージを転送
      for client := range r.clients {
        select {
        case client.send <- msg:
          // メッセージを送信
        default:
          // 送信に失敗
          delete(r.clients, client)
          close(client.send)
        }
      }
    }
  }
}

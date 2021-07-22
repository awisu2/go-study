// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"log"
	"time"
)

type BroadCast struct {
	message []byte
	outClient *Client
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *BroadCast

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *BroadCast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case broadcast := <-h.broadcast:
			for client := range h.clients {
				if broadcast.outClient != nil && broadcast.outClient == client {
					continue
				}

				select {
				case client.send <- broadcast.message:
				default:
					delete(h.clients, client)
					close(client.send)
				}
			}

		// 一定時間ごとに何らかの処理を行う
		case <-ticker.C:
			log.Println("interval")
			// 再起動
			// ticker.Stop()
			// ticker.Reset(time.Second * 1)
		}
	}
}
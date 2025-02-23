
package services

import (
    "encoding/json"
    "log"
    "sync"
    "time"
    "github.com/gorilla/websocket"
)

const (
    writeWait      = 10 * time.Second
    pongWait       = 60 * time.Second
    pingPeriod     = (pongWait * 9) / 10
    maxMessageSize = 512
)

type Client struct {
    Hub  *Hub
    ID   string
    Conn *websocket.Conn
    Send chan []byte
}

type Hub struct {
    Clients    map[*Client]bool
    Broadcast  chan []byte
    Register   chan *Client
    Unregister chan *Client
    mutex      sync.RWMutex
}

var WebsocketHub = NewHub()

func NewHub() *Hub {
    return &Hub{
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan []byte),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.mutex.Lock()
            h.Clients[client] = true
            h.mutex.Unlock()
            log.Printf("Client registered: %s", client.ID)

        case client := <-h.Unregister:
            h.mutex.Lock()
            if _, ok := h.Clients[client]; ok {
                delete(h.Clients, client)
                close(client.Send)
            }
            h.mutex.Unlock()
            log.Printf("Client unregistered: %s", client.ID)

        case message := <-h.Broadcast:
            h.mutex.RLock()
            for client := range h.Clients {
                select {
                case client.Send <- message:
                default:
                    close(client.Send)
                    delete(h.Clients, client)
                }
            }
            h.mutex.RUnlock()
        }
    }
}

func (c *Client) ReadPump() {
    defer func() {
        c.Hub.Unregister <- c
        c.Conn.Close()
    }()

    c.Conn.SetReadLimit(maxMessageSize)
    c.Conn.SetReadDeadline(time.Now().Add(pongWait))
    c.Conn.SetPongHandler(func(string) error {
        c.Conn.SetReadDeadline(time.Now().Add(pongWait))
        return nil
    })

    for {
        _, message, err := c.Conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            break
        }
        
       
        var msg map[string]interface{}
        if err := json.Unmarshal(message, &msg); err != nil {
            log.Printf("Error parsing message: %v", err)
            continue
        }

        
        if msgType, ok := msg["type"].(string); ok {
            switch msgType {
            case "ping":
                
                c.Send <- []byte(`{"type": "pong"}`)
            default:
                log.Printf("Received message of type %s from client %s", msgType, c.ID)
            }
        }
    }
}

func (c *Client) WritePump() {
    ticker := time.NewTicker(pingPeriod)
    defer func() {
        ticker.Stop()
        c.Conn.Close()
    }()

    for {
        select {
        case message, ok := <-c.Send:
            c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
            if !ok {
                c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            w, err := c.Conn.NextWriter(websocket.TextMessage)
            if err != nil {
                return
            }
            w.Write(message)

            if err := w.Close(); err != nil {
                return
            }

        case <-ticker.C:
            c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                return
            }
        }
    }
}

func BroadcastMessage(messageType string, data interface{}) {
    message := map[string]interface{}{
        "type":      messageType,
        "data":      data,
        "timestamp": time.Now(),
    }

    jsonMessage, err := json.Marshal(message)
    if err != nil {
        log.Printf("Error marshaling message: %v", err)
        return
    }

    WebsocketHub.Broadcast <- jsonMessage
}
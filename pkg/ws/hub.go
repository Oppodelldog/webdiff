package ws

var hub *Hub

func StartHub() *Hub {
	hub = NewHub()
	go hub.Run()

	return hub
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Broadcast(message []byte) {
	h.broadcast <- message
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true

			c.OpenConnection()

		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				if c.CancelSession != nil {
					c.CancelSession()
				}
				closeClientConnection(c, h)
			}

		case message := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.ToClient() <- message:
				default:
					closeClientConnection(c, h)
				}
			}
		}
	}
}

func closeClientConnection(c *Client, h *Hub) {
	c.CloseConnection()
	delete(h.clients, c)
}

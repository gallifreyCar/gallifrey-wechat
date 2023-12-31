package server

var MyServer = NewServer()

type Server struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client // 注册
	Unregister chan *Client // 注销
}

func NewServer() *Server {
	return &Server{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

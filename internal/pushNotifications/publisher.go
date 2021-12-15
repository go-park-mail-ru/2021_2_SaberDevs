package pushNotifications

type Publisher struct {
	subscribers map[*Subscriber]bool
	broadcast   chan string
	register    chan *Subscriber
	unregister  chan *Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{
		broadcast:   make(chan string),
		register:    make(chan *Subscriber),
		unregister:  make(chan *Subscriber),
		subscribers: make(map[*Subscriber]bool),
	}
}

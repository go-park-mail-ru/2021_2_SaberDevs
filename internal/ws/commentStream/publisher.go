package commentStream

import cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"

type Publisher struct {
	subscribers map[*Subscriber]bool
	// канал для получения последних коментариев
	broadcast  chan []cmodels.StreamComment // // todo поменять на коменты
	register   chan *Subscriber
	unregister chan *Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{
		broadcast:   make(chan []cmodels.StreamComment),
		register:    make(chan *Subscriber),
		unregister:  make(chan *Subscriber),
		subscribers: make(map[*Subscriber]bool),
	}
}

func (pub *Publisher) Run() {
	for {
		select {
		case subscriber := <-pub.register:
			// todo если активных сабов слишком много то отказать
			pub.subscribers[subscriber] = true

		case subscriber := <-pub.unregister:
			if _, ok := pub.subscribers[subscriber]; ok {
				delete(pub.subscribers, subscriber)
				close(subscriber.send)
			}

		case message := <-pub.broadcast:
			for subscriber := range pub.subscribers {
				select {

				case subscriber.send <- message:

				default:
					close(subscriber.send)
					delete(pub.subscribers, subscriber)
				}
			}
		}
	}
}

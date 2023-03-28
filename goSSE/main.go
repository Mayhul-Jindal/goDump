package main

import (
	"fmt"
	"log"
	"net/http"
)

type Broker struct {
	Messages chan string
	NewClients chan chan string // ekh hee chan mein multiple  channel kaam ata jab concurrenly ekh hee baare mein multiple jaga braodcast karna hota
	ClosingClients chan chan string
	Clients map[chan string]bool
}

func NewBrocker() *Broker{
	return &Broker{
		Messages:       make(chan string),
		NewClients:     make(chan chan string),
		ClosingClients: make(chan chan string),
		Clients:        make(map[chan string]bool),
	}
}

func (b *Broker) Listen(){
	for {
		select {
		case newClientch := <- b.NewClients:
			b.Clients[newClientch] = true
			log.Printf("Client added. Total %d registered clients", len(b.Clients))

		case closedClientch := <- b.ClosingClients:
			delete(b.Clients, closedClientch)
			log.Printf("Client removed. %d registered clients left", len(b.Clients))

		case msg := <- b.Messages:
			// TODO: later optimize this by making it concurrent
			for clientMsgch := range b.Clients {
				clientMsgch <- msg
			}
			log.Printf("Broadcasted messages to %d registered clients", len(b.Clients))
		}
	}
}

/*
Yeh karna padra because in main I have used drect the brocker object.
Toh voh cheeze kaam karne ke liye you have implement ServeHTTP function apn struct ke andar

After creating the broker instance, we should listen and wait for three actions to happen. 
Those are: 
1. if a new client comes
2. if a client disconnected from our service
3. if a notification message has arrived from our second endpoint.
*/
func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request){
	/*
	http.buffer() ka reason:
	When data is written to an http.ResponseWriter in a Go HTTP server, it is not necessary that the data be immediately sent to the client.
	Instead, the data may be buffered on the server until a certain amount of data has been accumulated or until the request is complete. 
	However, when streaming data, it is necessary that the data be sent to the client as soon as possible.
	*/
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// creating channel over which brocker can send this client, messages from the message source
	msgch := make(chan string)

	// Add this client msg channel to the channel of those that should receive updates
	b.NewClients <- msgch

	/*
	By checking for client disconnect twice in the code, 
	the server can ensure that it always detects when the client has disconnected, 
	regardless of how the client disconnects.
	*/ 
	defer func() {
		b.ClosingClients <- msgch
	}()

	closingctx := r.Context()
	go func(){
		<- closingctx.Done() // blocking line isiliye go routine mein daal diya main ko na rok deh
		b.ClosingClients <- msgch
		log.Printf("Http connection for %v just closed because of context condition", r.RemoteAddr)
	}()

	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	
	// not closing connection with the client. Possible because of http1.1 or http2 architecture 
	b.ReadLoop(w, r, flusher, msgch)
}

func (b *Broker)ReadLoop(w http.ResponseWriter, r *http.Request, flusher http.Flusher, msgch chan string){
	for {
		msg, open := <- msgch
		if !open{
			break
		}

		fmt.Fprintf(w, "Message: %s\n\n", msg)

		// Flush the response as soon as possible. This is only possible if the repsonse supports streaming.
		flusher.Flush()
	}

	log.Printf("Http connection for %v just closed because of end of handler reached", r.RemoteAddr)
}

func (b *Broker) PublishMessages(){
	for i := 0; i <= 100; i++ {
		b.Messages <- fmt.Sprintf("%d%%", i)
	}
}

func main(){
	brocker := NewBrocker()
	go brocker.Listen()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello wolrd")
	})

	http.HandleFunc("/getLogs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello wolrd")
	})

	http.Handle("/liveData", brocker)

	http.HandleFunc("/testLiveData", func(w http.ResponseWriter, r *http.Request) {
		brocker.PublishMessages()
	})

	// TODO: implementing graceful shutdown
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
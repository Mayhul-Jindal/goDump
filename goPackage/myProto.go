package myProto

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

var defaultDialInterval = 1 * time.Second 

type Sender[T any] struct{
	Chan chan T
	outBoundConn net.Conn
	addr string
	dialInterval time.Duration
}

func NewSender[T any](addr string)(*Sender[T], error){
	sen := &Sender[T]{
		addr: addr,
		Chan: make(chan T, 10),
		dialInterval: defaultDialInterval,
	}

	go sen.dialRemote()

	return sen, nil
}

func (s *Sender[T]) dialRemote(){
	for i := 0; i < 5; i++ {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Println("Dial error: ", err)
			time.Sleep(s.dialInterval)
		}else{
			s.outBoundConn = conn
			break
		}
	}

	go s.sendMessageFromChanDaemon()
}

func (s *Sender[T]) sendMessageFromChanDaemon(){
	for{
		if err := gob.NewEncoder(s.outBoundConn).Encode(<-s.Chan); err != nil{
			log.Println(err)
			continue
		} 
	}
}

type Reciever[T any] struct{
	Chan chan T
	listenAddr string
	listener net.Listener
}

func NewReciever[T any](listenAddr string)(*Reciever[T], error){
	recv := &Reciever[T]{
		listenAddr: listenAddr,
		Chan: make(chan T),
	}

	ln, err := net.Listen("tcp",listenAddr)
	if err != nil {
		return nil, err
	}
	recv.listener = ln

	// yeh go routine mein daal diya jisse object mil jaye aur voh peeche process mein chalte rahe 
	go recv.acceptLoop()
	return recv, nil
}

func (r *Reciever[T]) acceptLoop(){
	defer r.listener.Close()

	// considering there are multiple sending parties
	for{
		conn, err := r.listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go r.handleConn(conn)
	}
}

func (r *Reciever[T]) handleConn(conn net.Conn){
	var msg T

	for{
		if err := gob.NewDecoder(conn).Decode(&msg); err != nil{
			log.Println(err)
			continue // basically yeh wali iteration skip karne ke liye
		}

		r.Chan <- msg
	}
}
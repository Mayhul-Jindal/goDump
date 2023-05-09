package myProto_test

import (
	"fmt"
	"myProto"
	"testing"
)

func TestReciever(t *testing.T) {
	sender, err := myProto.NewSender[string](":3000")
	if err != nil {
		t.Error(err)
	}
	
	reciever, err := myProto.NewReciever[string](":3000")
	if err != nil {
		t.Error(err)
	}
	
	sender.Chan <- "hello1"
	msg := <- reciever.Chan
	
	fmt.Printf("Message: %v\n", msg)
}
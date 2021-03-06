/*
 * CS3210 - Principles of Programming Languages - Fall 2020
 * Instructor: Thyago Mota
 * Description: Prg04 - Publish Subscribe Simulation
 * Student(s) Name(s):  Larsen Close and Matt Hurt
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	stopCh    chan struct{}
	publishCh chan interface{}
	subCh     chan chan interface{}
	unsubCh   chan chan interface{}
	mutex     sync.Mutex
	topics    map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		stopCh:    make(chan struct{}),
		publishCh: make(chan interface{}, 1),
		subCh:     make(chan chan interface{}, 1),
		unsubCh:   make(chan chan interface{}, 1),
		topics:    make(map[string][]chan string),
	}
}

var wg sync.WaitGroup

// TODO: writes the given message on all the channels associated with the given topic.
func (ps *PubSub) Start() {
	//defer wg.Done()
	subs := map[chan interface{}]struct{}{}
	for {
		select {
		case <-ps.stopCh:
			return
		case msgCh := <-ps.subCh:
			subs[msgCh] = struct{}{}
		case msgCh := <-ps.unsubCh:
			delete(subs, msgCh)
		case msg := <-ps.publishCh:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the PubSub:
				select {
				case msgCh <- msg:
				default:
				}
			} // end case
		} // end select
	} // end for
} // end Start

func (ps *PubSub) Stop() {
	close(ps.stopCh)
}

// TODO: creates and returns a new channel on a given topic, updating the PubSub struct
func (ps *PubSub) Subscribe() chan interface{} {
	//defer wg.Done()
	msgCh := make(chan interface{}, 5) // Create slice
	ps.subCh <- msgCh
	return msgCh
}

//func (ps *PubSub) Unsubscribe(msgCh chan interface{}) {
//	ps.unsubCh <- msgCh
//}

func (ps *PubSub) Publish(msg interface{}) {
	//defer wg.Done()
	ps.publishCh <- msg
	//wg.Wait()
}

func main() {
	// TODO: create the ps struct
	// Create and start a PubSub:
	ps := NewPubSub()

	// TODO: create the subscriber goroutines
	go ps.Start()

	// TODO: create the arrays of messages to be sent on each topic
	// Create and subscribe 3 clients:
	subscriber := func(name string) {
		msgCh := ps.Subscribe()
		for {
			fmt.Printf("* %s got message: %v\n", name, <-msgCh)

		}

	}
	//sub2 := func(name string) {
	//	msgCh := ps.Subscribe()
	//	for {
	//		//ps.mutex.Lock()
	//		fmt.Printf("* %s got message: %v\n", name, <-msgCh)
	//		//ps.mutex.Unlock()
	//	}
	//}
	//go sub2("Marry")
	go subscriber("Tom")

	//go publish(beefacts)

	// TODO: set wait group to 2 (# of publishers)
	// TODO: create the publisher goroutines
	// Start publishing messages:
	go func() {
		for publisher := 0; ; publisher++ {
			defer wg.Done()
			//ps.mutex.Lock()
			//time.Sleep(200 * time.Millisecond)
			ps.Publish("bees are pollinators.")
			time.Sleep(200 * time.Millisecond)
			ps.Publish("bees produce honey")
			time.Sleep(200 * time.Millisecond)
			ps.Publish("all worker bees are female")
			time.Sleep(200 * time.Millisecond)
			ps.Publish("bees have 5 eyes,")
			time.Sleep(200 * time.Millisecond)
			ps.Publish("bees fly about 20mph")
			time.Sleep(200 * time.Millisecond)
			ps.Publish("Streets are cool")
			time.Sleep(200 * time.Millisecond)
			//ps.mutex.Unlock()
			//wg.Wait()
		}
	}()
	//go func() {
	//	for pub2 := 0; ; pub2++ {
	//		//ps.mutex.Unlock()
	//		ps.Publish("Streets are cool")
	//		time.Sleep(200 * time.Millisecond)
	//		//ps.mutex.Unlock()
	//		wg.Wait()
	//	}
	//}()

	// TODO: wait for all publishers to be done
	time.Sleep(200 * time.Second)
}

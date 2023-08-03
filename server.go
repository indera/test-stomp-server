package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/go-stomp/stomp/v3"
	"github.com/go-stomp/stomp/v3/server"
	"github.com/sirupsen/logrus"
)

const defaultPort = ":61613"

var serverAddr = flag.String("server", "localhost:61613", "STOMP server endpoint")
var messageCount = flag.Int("count", 10, "Number of messages to send/receive")
var queueName = flag.String("queue", "/queue/test", "Destination queue")
var helpFlag = flag.Bool("help", false, "Print help text")

// these are the default options that work with RabbitMQ
var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("guest", "guest"),
	stomp.ConnOpt.Host("/"),
}

func client(log *logrus.Logger) {
	log.Info("Start client")
	var stop = make(chan bool)

	flag.Parse()

	if *helpFlag {
		_, err := fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		if err != nil {
			fmt.Printf("error %w", err)
		}

		flag.PrintDefaults()
		os.Exit(1)
	}

	subscribed := make(chan bool)
	go recvMessages(subscribed, stop)

	// wait until we know the receiver has subscribed
	x := <-subscribed
	log.Infof("from subscribed chan got x %v", x)

	go sendMessages(stop)

	y := <-stop
	log.Infof("from stop chan got y %v", y)

	z := <-stop
	log.Infof("from stop chan got z %v", z)
}

func main() {
	//var logger := logrus.WithFields(logrus.Fields{"a": "b"})
	var log = logrus.New()
	var wg = sync.WaitGroup{}

	wg.Add(1)
	go client(log)

	// Create a STOMP server
	srv := server.Server{
		//Addr: server.DefaultAddr,
		Addr: *serverAddr,
	}

	log.Infof("srv %+v", srv)

	go func() {
		log.Info("Start server...")
		_ = srv.ListenAndServe()

		//if err != nil {
		//	fmt.Println("Error creating STOMP server:", err)
		//	return
		//}
	}()

	wg.Wait()

	//// Subscribe to a destination
	//sub := srv.Sub("/queue/example")
	//
	//// Handle incoming messages
	//go func() {
	//	for msg := range sub.C {
	//		fmt.Println("Received message:", string(msg.Body))
	//		sub.Ack(msg)
	//	}
	//}()
	//
	//fmt.Println("STOMP server started on localhost:61613")
	//select {}
}

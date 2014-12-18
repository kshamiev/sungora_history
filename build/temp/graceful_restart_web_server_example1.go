package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"
	_ "expvar"

	"github.com/webdeskltd/gograce"
)

func main() {
	var srv1, srv2 *gograce.WebServer
	var cnf1, cnf2 gograce.Configuration
	var err error
	var fhlog *os.File

	// Loging
	fhlog, err = os.Create("log.log")
	if err != nil {
		log.Panicf("Can't open log file 'log.log': %v", err)
	}
	defer fhlog.Close()
	log.SetOutput(fhlog)
	log.Println()
	log.Printf("Programm started: %s", os.Args)
	log.Printf("ENV: GOGRACE_ENV=%s", os.Getenv(`GOGRACE_ENV`))

	// HTTP Handler
	http.HandleFunc("/", DefaultHandler)

	// Server one
	cnf1.Host = "localhost"
	cnf1.Port = 2080
	cnf1.Mode = gograce.ModeTCP
	cnf1.MaxHeaderBytes = 1 << 20
	cnf1.ReadTimeout = time.Second * 3
	cnf1.WriteTimeout = time.Second * 15
	cnf1.KeepAlive = 3

	// Server two
	cnf2.Mode = gograce.ModeUnixSocket
	cnf2.Socket = `/Users/kallisto/domains/gograce_tests/application.sock`
	cnf2.MaxHeaderBytes = 1 << 20
	cnf2.ReadTimeout = time.Second * 3
	cnf2.WriteTimeout = time.Second * 15
	cnf2.KeepAlive = 0

	srv1, err = gograce.NewServer(cnf1)
	if err != nil {
		log.Fatalf("Error create server object for configuration 1: %v\n", err)
	}
	srv2, err = gograce.NewServer(cnf2)
	if err != nil {
		log.Fatalf("Error create server object for configuration 2: %v\n", err)
	}

	// Register event on change status for server 1
	srv1.OnServerStatus(onStatusChange)

	// Register event on change status for server 2
	srv2.OnServerStatus(onStatusChange)

	// Register event on connection count change for server 1
	srv1.OnConnectionCount(onConnectionCount)

	// Register event on connection count change for server 2
	srv2.OnConnectionCount(onConnectionCount)

	err = gograce.GracefulStart(srv1, srv2)
	if err != nil {
		log.Fatalf("Error starting servers: %v\n", err)
	}

	log.Printf("Wait five seconds")
	time.Sleep(time.Second * 5)

	log.Printf("GracefulRestart")
	err = gograce.GracefulRestart(func() {
		log.Printf("GracefulRestart done")
		time.Sleep(time.Second)
		os.Exit(0)
	}, srv1, srv2)
	if err != nil {
		log.Fatalf("GracefulRestart error: %v\n", err)
	}

	select {}
}

func onConnectionCount(srv *gograce.WebServer) {
	var pid = syscall.Getpid()
	log.Printf("Pid: %d onConnectionCount: %v", pid, srv.ConcurrentConnections())
}

func onStatusChange(srv *gograce.WebServer) {
	var pid = syscall.Getpid()

	log.Printf("Pid: %d onStatusChange: %v", pid, srv.Status())
}

func onIdle(srv *gograce.WebServer) {
	var pid = syscall.Getpid()

	log.Printf("Pid: %d onIDLE: %v", pid, srv.Status())
	log.Printf("Pid: %d, application exit", pid)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	var pid = syscall.Getpid()

	fmt.Fprintf(w, "PID: %d, Default Handler, host %v, time is %s", pid, r.Host, time.Now().String())
}

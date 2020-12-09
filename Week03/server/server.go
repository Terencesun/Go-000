package server

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var closeCh = make(chan struct{})

func s1() error {
	mux := http.ServeMux{}
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test1"))
	})

	server := http.Server{
		Addr:":8000",
		Handler: &mux,
	}

	go func() {
		<- closeCh
		log.Println("s1 received close signal")
		timeout, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
		defer cancel()
		server.Shutdown(timeout)
		log.Println("s1 shutdown")
	}()

	log.Println("s1: listening on port :8000....")
	return server.ListenAndServe()
}

func s2() error {
	mux := http.ServeMux{}
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test2"))
	})

	server := http.Server{
		Addr:":8001",
		Handler: &mux,
	}

	go func() {
		<- closeCh
		log.Println("s2 received close signal")
		timeout, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		server.Shutdown(timeout)
		log.Println("s2 shutdown")
	}()
	log.Println("s2: listening on port :8001....")
	return server.ListenAndServe()
}

func sigRegister() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
			}
		}()

		si := <- sig
		log.Println("Get signal :", si)
		if closeCh != nil {
			closeCh <- struct{}{}
		}
	}()
}


func Run() {
	fmt.Println("pid = ", os.Getpid())
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(s1)
	group.Go(s2)

	sigRegister()

	<-ctx.Done()
	if err := ctx.Err(); err != nil {
		log.Println("select received: ", err)
		if closeCh != nil {
			closeCh <- struct{}{}
		}
	}
	close(closeCh)
	time.Sleep(time.Second * 2)
	log.Println("All server run")
}

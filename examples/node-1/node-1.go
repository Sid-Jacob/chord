package main

import (
	"log"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/arriqaaq/chord"
	"github.com/arriqaaq/chord/models"
)

func createNode(id string, addr string, sister *models.Node) (*chord.Node, error) {

	cnf := chord.DefaultConfig()
	cnf.Id = id
	cnf.Addr = addr
	cnf.Timeout = 10 * time.Millisecond
	cnf.MaxIdle = 100 * time.Millisecond

	n, err := chord.NewNode(cnf, sister)
	return n, err
}

func createID(id string) []byte {
	val := big.NewInt(0)
	val.SetString(id, 10)
	return val.Bytes()
}

func main() {

	h, err := createNode("1", "127.0.0.1:8001", nil)
	if err != nil {
		log.Fatalln(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-time.After(10 * time.Second)
	<-c
	h.Stop()

}

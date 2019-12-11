package main

import (
	"bufio"
	"cryptocrouse/src/go/FiatShamirProtocol"
	"cryptocrouse/src/go/Fingerprints"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	data *ClientData
}

type ClientData struct {
	S *big.Int
	V *big.Int
	N *big.Int
}

func (c *Client) ConnectToServer() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	connect := arguments[1]
	conn, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.conn = conn
}

func (c* Client) setupConnections() {
	c.reader = bufio.NewReader(c.conn)
	c.writer = bufio.NewWriter(c.conn)
}

func (c *Client) Round() {
	c.receiveN()
}

func (c *Client) receiveN() {
	_, _ = c.writer.WriteString(FiatShamirProtocol.COMMAND_GET_N)
	_ = c.writer.Flush()

	msg, _ := c.reader.ReadString('\n')
	log.Printf("received N: %s\n", msg)

	c.data.N, _ = big.NewInt(0).SetString(msg, 10)
}

func (c *Client) generateS() {
	for {
		c.data.S = Fingerprints.GetBigRandomWithLimit(MAX_P)
		GCD := big.NewInt(0).GCD(
			nil,
			nil,
			c.data.S,
			c.data.N)
		if GCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}

//for {
//fmt.Print(">> ")
//text, _ := rStdin.ReadString('\n')
//_, _ = fmt.Fprintf(conn, text)
//
//message, _ := rConn.ReadString('\n')
//fmt.Print("->: " + message)
//if strings.TrimSpace(text) == "STOP" {
//fmt.Println("TCP client exiting...")
//return
//}
//}
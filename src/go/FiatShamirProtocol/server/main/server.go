package main

import (
	"bufio"
	"cryptocrouse/src/go/Fingerprints"
	"log"
	"math/big"
	"net"
	"strings"
)

type Server struct {
	data ServerData
}

type ServerData struct {
	p *big.Int
	q *big.Int
	N *big.Int
}

func (s *Server) Run() {
	s.serverPrepare()
	log.Printf("Server up with\n")
	log.Printf("N=%s\n", s.data.N.Text(10))
	s.serverListen()
}

func (s *Server) serverPrepare() {
	s.data.generateP()
	s.data.computeN()
}

func (s *Server) serverListen() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		log.Fatalf("Error listening:", err.Error())
	}
	defer l.Close()

	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error accepting: ", err.Error())
		}

		log.Printf("New connection: %s\n", conn.RemoteAddr())

		go startRound(conn)
	}
}

func startRound(conn net.Conn) {
	defer func() {
		log.Printf("closing connection from %v\n", conn.RemoteAddr())
		conn.Close()
	}()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	scanr := bufio.NewScanner(r)

	for {
		scanned := scanr.Scan()
		if !scanned {
			if err := scanr.Err(); err != nil {
				log.Printf("%v(%v)\n", err, conn.RemoteAddr())
				return
			}
			break
		}
		msg := scanr.Text()
		log.Printf("Reveived [%s]: %s\n", conn.RemoteAddr(), msg)
		_, _ = w.WriteString(strings.ToUpper(msg) + "\n")
		_ = w.Flush()
	}
}

func (data *ServerData) generateQ() *big.Int {
	return Fingerprints.GenerateBigPrimeNumberWithLimit(MIN_P)
}

func (data *ServerData) generateP() {
	data.p = big.NewInt(0)

	for {
		q := data.generateQ()
		data.p.Add(
			big.NewInt(0).Mul(
				big.NewInt(2),
				q),
			big.NewInt(1))
		if Fingerprints.IsPrimeRef(data.p) {
			if data.p.Cmp(MIN_P) > 0 && data.p.Cmp(MAX_P) < 0 {
				break
			}
		}
	}
}

func (data *ServerData) computeN() {
	data.N = big.NewInt(0).Mul(data.p, data.q)
}

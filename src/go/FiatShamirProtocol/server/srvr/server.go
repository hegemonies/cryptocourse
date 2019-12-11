package srvr

import (
	"bufio"
	"cryptocrouse/src/go/FiatShamirProtocol"
	"cryptocrouse/src/go/Fingerprints"
	"log"
	"math/big"
	"net"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	data ServerData
}

type ServerData struct {
	p *big.Int
	q *big.Int
	N *big.Int
}

func ServerInit() *Server {
	return &Server{}
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
	l, err := net.Listen(CONN_TYPE, CONN_HOST+ ":" +CONN_PORT)
	if err != nil {
		log.Fatalf("Error listening:", err.Error())
	}
	defer l.Close()

	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s\n", err.Error())
		}

		log.Printf("New connection: %s\n", conn.RemoteAddr())

		go s.startRound(conn)
	}
}

func (s *Server) startRound(conn net.Conn) {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	scanr := bufio.NewScanner(r)

	var x *big.Int
	var y *big.Int
	var v *big.Int
	e := generateE()

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
		msg = strings.TrimSuffix(msg, "\n")
		log.Printf("Reveived [%s]: ~%s~\n", conn.RemoteAddr(), msg)

		switch msg {
		case FiatShamirProtocol.COMMAND_GET_N:
			s.sendN(w)
		case FiatShamirProtocol.COMMAND_GET_X:
			x = s.receiveX(r)
		case FiatShamirProtocol.COMMAND_GET_Y:
			y = s.receiveY(r)
			statusCode := s.computeY(y, x, v, w, e)
			s.sendAnswerCode(w, statusCode)
		case FiatShamirProtocol.COMMAND_GET_V:
			v = s.receiveV(r)
		case FiatShamirProtocol.COMMAND_GET_E:
			s.sendE(w, e)
		}
	}
}

func (s *Server) sendN(w *bufio.Writer) {
	log.Printf("Send N: %s\n", s.data.N.Text(10))
	_, _ = w.WriteString(s.data.N.Text(10) + "\n")
	_ = w.Flush()
}

func (s *Server) receiveX(r *bufio.Reader) *big.Int {
	msg, _ := r.ReadString('\n')
	msg = strings.TrimSuffix(msg, "\n")
	x, _ := big.NewInt(0).SetString(msg, 10)
	log.Printf("Receive x: %s\n", msg)
	return x
}

func (s *Server) receiveV(r *bufio.Reader) *big.Int {
	msg, _ := r.ReadString('\n')
	msg = strings.TrimSuffix(msg, "\n")
	v, _ := big.NewInt(0).SetString(msg, 10)
	log.Printf("Receive v: %s\n", msg)
	return v
}

func (s *Server) sendE(w *bufio.Writer, e int) {
	_, _ = w.WriteString(strconv.Itoa(e) + "\n")
	_ = w.Flush()
	log.Println("Send E: " + strconv.Itoa(e))
}

func (s *Server) receiveY(r *bufio.Reader) *big.Int {
	msg, _ := r.ReadString('\n')
	msg = strings.TrimSuffix(msg, "\n")
	y, _ := big.NewInt(0).SetString(msg, 10)
	log.Printf("Receive y: %s\n", msg)
	return y
}

func (s *Server) computeY(y *big.Int, x *big.Int, v *big.Int, w *bufio.Writer, e int) string {
	if y.Cmp(big.NewInt(0)) == 0 {
		_, _ = w.WriteString(FiatShamirProtocol.COMMAND_ANSWER_CODE_ERROR)
		_ = w.Flush()
		return FiatShamirProtocol.COMMAND_ANSWER_CODE_ERROR
	}

	l := big.NewInt(0).Exp(y, big.NewInt(2), s.data.N)
	var r *big.Int

	switch e {
	case 0:
		r = x
	case 1:
		r = big.NewInt(0).Mod(
			big.NewInt(0).Mul(
				x,
				v),
			s.data.N)
	}

	log.Printf("l = %s r = %s\n", l.Text(10), r.Text(10))

	code := ""

	if l.Cmp(r) == 0 {
		code = FiatShamirProtocol.COMMAND_ANSWER_CODE_SUCCESS
	} else {
		code = FiatShamirProtocol.COMMAND_ANSWER_CODE_ERROR
	}

	return code
}

func (s *Server) sendAnswerCode(w *bufio.Writer, statusCode string) {
	_, err := w.WriteString(statusCode + "\n")
	if err != nil {
		log.Fatal(err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Send status code: " + statusCode)
	time.Sleep(50 * time.Millisecond)
}

func generateE() int {
	rand := Fingerprints.GetBigRandom()
	answer, _ := strconv.Atoi(big.NewInt(0).Mod(rand, big.NewInt(2)).Text(10))
	return answer
}

func (data *ServerData) generateQ() {
	data.q = Fingerprints.GenerateBigPrimeNumberWithLimit(MIN_P)
}

func (data *ServerData) generateP() {
	data.p = big.NewInt(0)

	for {
		data.generateQ()
		data.p.Add(
			big.NewInt(0).Mul(
				big.NewInt(2),
				data.q),
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

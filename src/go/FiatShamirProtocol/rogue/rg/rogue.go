package rg

import (
	"bufio"
	"cryptocrouse/src/go/FiatShamirProtocol"
	"cryptocrouse/src/go/Fingerprints"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Rogue struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	data   *RogueData
}

type RogueData struct {
	singleV *big.Int
	V       []*big.Int
	N       *big.Int
	R       *big.Int
	X       *big.Int
	E       int
	Y       *big.Int
	S       *big.Int
}

func InitRogue() *Rogue {
	data := &RogueData{
		V: make([]*big.Int, 0),
		N: big.NewInt(0),
	}
	return &Rogue{
		data: data,
	}
}

func (r *Rogue) ConnectToServer() {
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
	log.Printf("Connect to %s\n", connect)
	r.conn = conn

	r.data = &RogueData{}
	r.setupConnections()
}

func (r *Rogue) setupConnections() {
	r.reader = bufio.NewReader(r.conn)
	r.writer = bufio.NewWriter(r.conn)
}

func (r *Rogue) TryToAcceptSecret() {
	r.receiveOpenKeys()

	r.generateS()

	var v *big.Int
	for i := 0; i < len(r.data.V); i++ {
		v = r.data.V[i]
		if v != nil {
			r.data.singleV = v
			break
		}
	}

	flag := r.hackSecret(v)
	if flag {
		log.Println("Hack is SUCCESSFUL")
		return
	}
	log.Println("Hack is BAD")
}

func (r *Rogue) receiveOpenKeys() {
	r.receiveN()
	r.receiveV()
}

func (r *Rogue) receiveN() {
	time.Sleep(50 * time.Millisecond)

	msg, err := r.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	msg = strings.TrimSuffix(msg, "\n")
	log.Printf("Received N: ~%s~\n", msg)

	var flag bool
	r.data.N, flag = big.NewInt(0).SetString(msg, 10)
	if flag == false {
		log.Fatal("Received N is bad")
	}
}

func (r *Rogue) receiveV() {
	time.Sleep(50 * time.Millisecond)

	msg, err := r.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	msg = strings.TrimSuffix(msg, "\n")
	log.Printf("received array of V: ~%s~\n", msg)

	vString := strings.Split(msg, ",")
	for _, vstr := range vString {
		var flag bool
		v, flag := big.NewInt(0).SetString(vstr, 10)
		if flag == false {
			log.Fatal("Received N is bad")
		}
		r.data.V = append(r.data.V, v)
	}
}

func (r *Rogue) hackSecret(v *big.Int) bool {
	r.sendV()
	for i := 0; i < 5; i++ {
		if !r.round() {
			log.Printf("Can not proof with v=%s on %d iteration\n", v.Text(10), i)
			return false
		}
	}
	log.Printf("Secret accepted successfuly")
	r.sendEnd()
	return true
}

func (r *Rogue) round() bool {
	r.generateR()
	r.computeX()
	r.sendX()

	r.receiveE()
	r.computeY()
	r.sendY()
	return r.getAnswer()
}

func (r *Rogue) generateR() {
	for {
		r.data.R = Fingerprints.GetBigRandomWithLimit(r.data.N)
		if r.data.R.Cmp(big.NewInt(1)) > 0 && r.data.R.Cmp(r.data.N) < 0 {
			break
		}
	}
}

func (r *Rogue) computeX() {
	r.data.X = big.NewInt(0).Exp(r.data.R, big.NewInt(2), r.data.N)
}

func (r *Rogue) sendX() {
	time.Sleep(50 * time.Millisecond)

	_, _ = r.writer.WriteString(r.data.X.Text(10) + "\n")
	_ = r.writer.Flush()
	log.Printf("Send X %s\n", r.data.X.Text(10))

	time.Sleep(50 * time.Millisecond)
}

func (r *Rogue) receiveE() {
	time.Sleep(50 * time.Millisecond)

	msg, _ := r.reader.ReadString('\n')
	msg = strings.TrimSuffix(msg, "\n")
	log.Printf("Received E: %s\n", msg)

	r.data.E, _ = strconv.Atoi(msg)
}

func (r *Rogue) computeY() {
	switch r.data.E {
	case 0:
		r.data.Y = r.data.R
	case 1:
		r.data.Y = big.NewInt(0).Mod(
			big.NewInt(0).Mul(
				r.data.S,
				r.data.R),
			r.data.N)
	}
}

func (r *Rogue) generateS() {
	for {
		r.data.S = Fingerprints.GetBigRandomWithLimit(r.data.N)
		if r.data.S.Cmp(big.NewInt(1)) == 0 {
			continue
		}
		GCD := big.NewInt(0).GCD(
			nil,
			nil,
			r.data.S,
			r.data.N)
		if GCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}

func (r *Rogue) sendY() {
	time.Sleep(50 * time.Millisecond)

	_, _ = r.writer.WriteString(r.data.Y.Text(10) + "\n")
	_ = r.writer.Flush()
	log.Printf("Send Y %s\n", r.data.Y.Text(10))

	time.Sleep(50 * time.Millisecond)
}

func (r *Rogue) getAnswer() bool {
	msg, err := r.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	msg = strings.TrimSuffix(msg, "\n")
	log.Printf("Received answer: ~%s~\n", msg)

	switch msg {
	case FiatShamirProtocol.COMMAND_ANSWER_CODE_SUCCESS:
		return true
	case FiatShamirProtocol.COMMAND_ANSWER_CODE_ERROR:
		return false
	default:
		return false
	}
}

func (r *Rogue) sendEnd() {
	_, _ = r.writer.WriteString(FiatShamirProtocol.COMMAND_END + "\n")
	_ = r.writer.Flush()
}

func (r *Rogue) computeV() {
	r.data.singleV = big.NewInt(0).Exp(r.data.S, big.NewInt(2), r.data.N)
}

func (r *Rogue) sendV() {
	time.Sleep(50 * time.Millisecond)
	_, _ = r.writer.WriteString(r.data.singleV.Text(10) + "\n")
	_ = r.writer.Flush()
	log.Printf("Send V %s\n", r.data.singleV.Text(10))
}

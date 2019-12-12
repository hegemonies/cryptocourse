package rg

import (
	"bufio"
	"cryptocrouse/src/go/FiatShamirProtocol"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

type Rogue struct {
	conn net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	data *RogueData
}

type RogueData struct {
	V []*big.Int
	N *big.Int
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
	for i := 0; i < len(r.data.V); i++ {
		flag := r.hackSecret(r.data.V[i])
		if flag {
			log.Println("Hack is SUCCESSFUL")
		}
	}
	log.Println("Hack is BAD")
}

func (r *Rogue) receiveOpenKeys() {
	r.receiveN()
	r.receiveV()
}

func (r *Rogue) receiveN() {
	_, err := r.writer.WriteString(FiatShamirProtocol.COMMAND_GET_N + "\n")
	if err != nil {
		log.Fatal(err)
	}
	err = r.writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(50 * time.Millisecond)

	msg, err := r.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	msg = strings.TrimSuffix(msg, "\n")
	log.Printf("received N: ~%s~\n", msg)

	var flag bool
	r.data.N, flag = big.NewInt(0).SetString(msg, 10)
	if flag == false {
		log.Fatal("Received N is bad")
	}
}

func (r *Rogue) receiveV() {
	_, err := r.writer.WriteString(FiatShamirProtocol.COMMAND_SERVER_GET_V + "\n")
	if err != nil {
		log.Fatal(err)
	}
	err = r.writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

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

}

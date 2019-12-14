package clnt

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

type Client struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	data   *ClientData
}

type ClientData struct {
	S      *big.Int
	V      *big.Int
	N      *big.Int
	E      int
	Y      *big.Int
	R      *big.Int
	X      *big.Int
	arrayV []*big.Int
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
	log.Printf("Connect to %s\n", connect)
	c.conn = conn

	c.data = &ClientData{}
	c.setupConnections()
}

func (c* Client) setupConnections() {
	c.reader = bufio.NewReader(c.conn)
	c.writer = bufio.NewWriter(c.conn)
}

func (c *Client) StartProof() {
	c.receiveN()
	c.receiveV()
	c.generateS()
	c.computeV()
	c.sendV()

	for i := 0; i < 5; i++ {
		answerCode := c.round()
		if answerCode == false {
			log.Fatalf("Can not proof on %d iteration\n", i)
			return
		}
	}
	log.Printf("Secret accepted successfuly")
	c.sendEnd()
}

func (c *Client) round() bool {
	c.generateR()
	c.computeX()
	c.sendX()

	c.receiveE()

	c.computeY()
	c.sendY()

	return c.getAnswer()
}

func (c *Client) receiveN() {
	time.Sleep(50 * time.Millisecond)

	msg, err := c.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	msg = strings.TrimSuffix(msg, "\n")
	log.Printf("received N: ~%s~\n", msg)

	var flag bool
	c.data.N, flag = big.NewInt(0).SetString(msg, 10)
	if flag == false {
		log.Fatal("Received N is bad")
	}
}

func (c *Client) generateS() {
	for {
		c.data.S = Fingerprints.GetBigRandomWithLimit(c.data.N)
		if c.data.S.Cmp(big.NewInt(1)) == 0 {
			continue
		}
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

func (c *Client) computeV() {
	c.data.V = big.NewInt(0).Exp(c.data.S, big.NewInt(2), c.data.N)
}

func (c *Client) receiveE() {
	time.Sleep(50 * time.Millisecond)

	msg, _ := c.reader.ReadString('\n')
	msg = strings.TrimSuffix(msg, "\n")
	log.Printf("Received E: %s\n", msg)

	c.data.E, _ = strconv.Atoi(msg)
}

func (c *Client) computeY() {
	switch c.data.E {
	case 0:
		c.data.Y = big.NewInt(0).Mod(c.data.R, c.data.N)
	case 1:
		c.data.Y = big.NewInt(0).Mod(
			big.NewInt(0).Mul(
				c.data.S,
				c.data.R),
			c.data.N)
	}
}

func (c *Client) generateR() {
	for {
		c.data.R = Fingerprints.GetBigRandomWithLimit(c.data.N)
		if c.data.R.Cmp(big.NewInt(1)) > 0 && c.data.R.Cmp(c.data.N) < 0 {
			break
		}
	}
}

func (c *Client) computeX() {
	c.data.X = big.NewInt(0).Exp(c.data.R, big.NewInt(2), c.data.N)
}

func (c *Client) sendX() {
	time.Sleep(50 * time.Millisecond)
	_, _ = c.writer.WriteString(c.data.X.Text(10) + "\n")
	_ = c.writer.Flush()
	log.Printf("Send X %s\n", c.data.X.Text(10))
}

func (c *Client) sendY() {
	time.Sleep(50 * time.Millisecond)
	_, _ = c.writer.WriteString(c.data.Y.Text(10) + "\n")
	_ = c.writer.Flush()
	log.Printf("Send Y %s\n", c.data.Y.Text(10))
}

func (c *Client) sendV() {
	time.Sleep(50 * time.Millisecond)
	_, _ = c.writer.WriteString(c.data.V.Text(10) + "\n")
	_ = c.writer.Flush()
	log.Printf("Send V %s\n", c.data.V.Text(10))
}

func (c *Client) getAnswer() bool {
	msg, err := c.reader.ReadString('\n')
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

func (c *Client) sendEnd() {
	_, _ = c.writer.WriteString(FiatShamirProtocol.COMMAND_END + "\n")
	_ = c.writer.Flush()
}

func (c *Client) receiveV() {
	time.Sleep(50 * time.Millisecond)

	_, err := c.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
}

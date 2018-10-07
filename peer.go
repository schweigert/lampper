package lampper

import (
	"net"

	"github.com/schweigert/lampper/internal"
)

type Peer struct {
	conn net.Conn
}

func New(conn net.Conn) *Peer {
	return &Peer{conn: conn}
}

func Dial(network, address string) (*Peer, error) {
	conn, err := net.Dial(network, address)
	return New(conn), err
}

func NewPipe() (*Peer, *Peer) {
	connOne, connTwo := net.Pipe()
	return New(connOne), New(connTwo)
}

func (peer *Peer) Close() error {
	return peer.conn.Close()
}

func (peer *Peer) WriteBytes(data []byte) error {
	_, err := peer.conn.Write(data)
	return err
}

func (peer *Peer) ReadBytes(size uint32) ([]byte, error) {
	ret := make([]byte, size)
	_, err := peer.conn.Read(ret)
	return ret, err
}

func (peer *Peer) WriteUint32(num uint32) error {
	data := bytes.FromUint32(num)
	return peer.WriteBytes(data)
}

func (peer *Peer) ReadUint32() (uint32, error) {
	data, err := peer.ReadBytes(4)
	if err != nil {
		return 0, err
	}
	return bytes.ToUint32(data)
}

func (peer *Peer) WriteString(str string) error {
	err := peer.WriteUint32(uint32(len(str)))
	if err != nil {
		return err
	}
	return peer.WriteBytes([]byte(str))
}

func (peer *Peer) ReadString() (string, error) {
	size, err := peer.ReadUint32()
	if err != nil {
		return "", err
	}

	str, err := peer.ReadBytes(size)
	return string(str), err
}

func (peer *Peer) WriteLambda(lambda *Lambda) error {
	err := peer.WriteString(lambda.Name)
	if err != nil {
		return err
	}

	return peer.WriteString(lambda.Body)
}

func (peer *Peer) ReadLambda() (*Lambda, error) {
	name, err := peer.ReadString()
	if err != nil {
		return nil, err
	}
	body, err := peer.ReadString()

	return NewLambda(name, body), err
}

func (peer *Peer) Handle(lambdaSet *LambdaSet) {
	defer func() {
		_ = peer.Close()
	}()

	for {
		lambda, err := peer.ReadLambda()
		if err != nil {
			return
		}

		lambdaFunction := lambdaSet.Get(lambda.Name)
		if lambdaFunction != nil && lambdaFunction.F != nil {
			lambdaFunction.F(lambda, peer)
		}
	}
}

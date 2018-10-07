package lampper

import "net"

type Service struct {
	listener net.Listener
}

func NewService(listener net.Listener) *Service {
	return &Service{listener: listener}
}

func Listen(network, address string) (*Service, error) {
	listener, err := net.Listen(network, address)
	return NewService(listener), err
}

func (service *Service) Accept() *Peer {
	for {
		conn, err := service.listener.Accept()
		if err == nil {
			return New(conn)
		}
	}
}

func (service *Service) Handle(lambdaSet *LambdaSet) {
	for {
		peer := service.Accept()
		go peer.Handle(lambdaSet)
	}
}

func (service *Service) Close() error {
	return service.listener.Close()
}

package lampper

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	service, err := Listen("tcp", "localhost:3030")
	assert.NoError(t, err)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		peerOne, errT := Dial("tcp", "localhost:3030")
		assert.NoError(t, errT)

		errT = peerOne.Close()
		assert.NoError(t, errT)
		wg.Done()
	}()

	peer := service.Accept()

	err = peer.Close()
	assert.NoError(t, err)

	wg.Wait()

	err = service.Close()
	assert.NoError(t, err)
}

func TestServiceHandle(t *testing.T) {
	listener, err := Listen("tcp", "localhost:3030")
	assert.NoError(t, err)
	var eventOne, eventTwo, eventThree int

	go func() {
		listener.Handle(NewLambdaSet().Add("event one", func(lambda *Lambda, peer *Peer) {
			eventOne++
			peer.WriteLambda(NewLambda("event one", "event one"))
		}).Add("event two", func(lambda *Lambda, peer *Peer) {
			eventTwo++
			peer.WriteLambda(NewLambda("event two", "event two"))
		}).Add("event three", func(lambda *Lambda, peer *Peer) {
			eventThree++
			peer.WriteLambda(NewLambda("event three", "event three"))
		}))
	}()

	peer, err := Dial("tcp", "localhost:3030")
	assert.NoError(t, err)

	assert.Equal(t, 0, eventOne)
	assert.Equal(t, 0, eventTwo)
	assert.Equal(t, 0, eventThree)

	err = peer.WriteLambda(NewLambda("event one", "nothong here"))
	assert.NoError(t, err)

	err = peer.WriteLambda(NewLambda("event two", "nothong here"))
	assert.NoError(t, err)

	err = peer.WriteLambda(NewLambda("event three", "nothong here"))
	assert.NoError(t, err)

	lambda, err := peer.ReadLambda()
	assert.NoError(t, err)
	assert.Equal(t, lambda.Name, lambda.Body)
	assert.Equal(t, 1, eventOne)

	lambda, err = peer.ReadLambda()
	assert.NoError(t, err)
	assert.Equal(t, lambda.Name, lambda.Body)
	assert.Equal(t, 1, eventTwo)

	lambda, err = peer.ReadLambda()
	assert.NoError(t, err)
	assert.Equal(t, lambda.Name, lambda.Body)
	assert.Equal(t, 1, eventThree)

	err = peer.Close()
	assert.NoError(t, err)

	err = listener.Close()
	assert.NoError(t, err)
}

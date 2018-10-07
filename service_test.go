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

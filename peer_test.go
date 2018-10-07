package lampper

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeer(t *testing.T) {
	p1, p2 := NewPipe()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		p1.WriteLambda(NewLambda("name", "body"))
		wg.Done()
	}()

	lambda, err := p2.ReadLambda()
	assert.NoError(t, err)
	assert.Equal(t, "name", lambda.Name)
	assert.Equal(t, "body", lambda.Body)

	wg.Wait()
	err = p1.Close()
	assert.NoError(t, err)

	err = p2.Close()
	assert.NoError(t, err)
}

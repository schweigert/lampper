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

func TestWriteAndReadOnClosedConnection(t *testing.T) {
	p1, p2 := NewPipe()

	err := p1.Close()
	assert.NoError(t, err)

	err = p2.Close()
	assert.NoError(t, err)

	err = p1.WriteBytes([]byte{})
	assert.Error(t, err)

	err = p1.WriteUint32(0)
	assert.Error(t, err)

	err = p1.WriteString("testing")
	assert.Error(t, err)

	err = p1.WriteLambda(NewLambda("name", "body"))
	assert.Error(t, err)

	data, err := p2.ReadBytes(32)
	assert.Error(t, err)
	assert.Equal(t, make([]byte, 32), data)

	num, err := p2.ReadUint32()
	assert.Error(t, err)
	assert.Equal(t, uint32(0), num)

	str, err := p2.ReadString()
	assert.Error(t, err)
	assert.Equal(t, "", str)

	lambda, err := p2.ReadLambda()
	assert.Error(t, err)
	assert.Nil(t, lambda)
}

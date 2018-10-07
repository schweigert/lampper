package lampper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaSet(t *testing.T) {
	set := NewLambdaSet().Add("name one", func(lambda *Lambda, peer *Peer) {}).Add("name two", func(lambda *Lambda, peer *Peer) {})

	assert.Equal(t, "name one", set.Get("name one").Name)
	assert.Equal(t, "name two", set.Get("name two").Name)
}

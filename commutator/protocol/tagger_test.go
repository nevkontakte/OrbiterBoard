package protocol

import (
    "testing"
    "github.com/stretchr/testify/require"
)


func TestTagger(t *testing.T) {
    assert := require.New(t)
    tagger := NewTagger(TAGGER_PC)
    assert.Equal(MsgTag(1), tagger.Next())
    assert.Equal(MsgTag(2), tagger.Next())

    tagger = NewTagger(TAGGER_BOARD)
    assert.Equal(MsgTag(0x8001), tagger.Next())
    assert.Equal(MsgTag(0x8002), tagger.Next())
}
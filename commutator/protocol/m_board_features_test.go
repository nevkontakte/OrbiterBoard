package protocol

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBoardFeaturesRequest(t *testing.T) {
	assert := require.New(t)
	assert.Implements((*Message)(nil), new(MessageBoardFeatures))

	tagger := NewTagger(TAGGER_PC)
	tagger.Next()

	// Create message
	msg := NewMessageBoardFeaturesRequest(&tagger)
	assert.Equal(MSG_BOARD_FEATURES, msg.GetMessageId())
	assert.Equal(MsgTag(2), msg.GetMessageTag())
	assert.Equal(make(PayloadData, 0), msg.GetPayloadData())

	// Send message
	buf := new(bytes.Buffer)
	err := SendMessage(buf, &msg)
	assert.Nil(err)
	assert.Equal([]byte{0x02, 0x00, 0x02, 0x00, 0x00}, buf.Bytes())

	// Receive packet
	packet, err := ReceivePacket(buf)
	assert.Nil(err)
	assert.Equal(MSG_BOARD_FEATURES, packet.Meta.MessageId)
	assert.Equal(msg.GetMessageTag(), packet.Meta.MessageTag)
	assert.Equal(make(PayloadData, 0), packet.Payload)
}

func TestBoardFeaturesResponse(t *testing.T) {
	assert := require.New(t)
	assert.Implements((*Message)(nil), new(MessageBoardFeatures))

	tagger := NewTagger(TAGGER_PC)
	tagger.Next()

	// Create message
	buttons := []bool{true, false, true, true}
	switches := []bool{false, true, true}
	sliders := []uint8{127}
	msg := NewMessageBoardFeaturesResponse(&tagger, buttons, switches, sliders)
	assert.Equal(MSG_BOARD_FEATURES, msg.GetMessageId())
	assert.Equal(MsgTag(2), msg.GetMessageTag())
	assert.Equal(PayloadData{4, 1 | (0 << 1) | (1 << 2) | (1 << 3), 3, 0 | (1 << 1) | (1 << 2), 1, 127}, msg.GetPayloadData())

	// Send message
	buf := new(bytes.Buffer)
	err := SendMessage(buf, &msg)
	assert.Nil(err)
	assert.Equal([]byte{0x02, 0x00, 0x02, 0x00, 0x6, 0x4, 0xd, 0x3, 0x6, 0x1, 0x7f}, buf.Bytes())

	// Receive packet
	packet, err := ReceivePacket(buf)
	assert.Nil(err)
	assert.Equal(MSG_BOARD_FEATURES, packet.Meta.MessageId)
	assert.Equal(msg.GetMessageTag(), packet.Meta.MessageTag)
	assert.Equal(PayloadData{0x4, 0xd, 0x3, 0x6, 0x1, 0x7f}, packet.Payload)
}

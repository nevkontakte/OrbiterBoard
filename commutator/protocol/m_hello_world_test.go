package protocol

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	assert := require.New(t)

	assert.Implements((*Message)(nil), new(MessageHelloWorld))

	tagger := NewTagger(TAGGER_PC)
	tagger.Next() // Skip one tag

	// Create message
	msg := NewHelloWorld(&tagger)
	assert.Equal(MSG_HELLO_WORLD, msg.GetMessageId())
	assert.Equal(MsgTag(2), msg.GetMessageTag())
	assert.Equal(make(PayloadData, 0), msg.GetPayloadData())

	// Send message
	buf := new(bytes.Buffer)
	err := SendMessage(buf, &msg)
	assert.Nil(err)
	assert.Equal([]byte{0x01, 0x00, 0x02, 0x00, 0x00}, buf.Bytes())

	// Receive packet
	packet, err := ReceivePacket(buf)
	assert.Nil(err)
	assert.Equal(MSG_HELLO_WORLD, packet.Meta.MessageId)
	assert.Equal(msg.GetMessageTag(), packet.Meta.MessageTag)
	assert.Equal(make(PayloadData, 0), packet.Payload)
}

func TestPacket2HelloWorld(t *testing.T) {
	assert := require.New(t)

	packet := Packet{
		Meta: PacketMeta{
			MessageId:  MSG_RESERVED,
			MessageTag: 1,
		},
		Payload: make(PayloadData, 0),
	}

	_, err := Packet2HelloWorld(packet)
	assert.NotNil(err)

	packet = Packet{
		Meta: PacketMeta{
			MessageId:  MSG_HELLO_WORLD,
			MessageTag: 1,
		},
		Payload: make(PayloadData, 1),
	}

	_, err = Packet2HelloWorld(packet)
	assert.NotNil(err)

	packet = Packet{
		Meta: PacketMeta{
			MessageId:  MSG_HELLO_WORLD,
			MessageTag: 1,
		},
		Payload: make(PayloadData, 0),
	}

	msg, err := Packet2HelloWorld(packet)
	assert.Nil(err)
	assert.Equal(MsgTag(1), msg.GetMessageTag())
}

package protocol

import (
	"fmt"
)

type MessageHelloWorld struct {
	message_tag MsgTag
}

func NewHelloWorld(tagger *Tagger) MessageHelloWorld {
	self := MessageHelloWorld{
		message_tag: tagger.Next(),
	}

	return self
}

func Packet2HelloWorld(packet Packet) (MessageHelloWorld, error) {
	if packet.Meta.MessageId != MSG_HELLO_WORLD {
		return MessageHelloWorld{},
			fmt.Errorf("Packet2HelloWorld: Unexpected message id: %d.", packet.Meta.MessageId)
	}
	if len(packet.Payload) != 0 {
		return MessageHelloWorld{},
			fmt.Errorf("Packet2HelloWorld: Unexpected payload.")
	}

	return MessageHelloWorld{
		message_tag: packet.Meta.MessageTag,
	}, nil
}

func (self *MessageHelloWorld) GetMessageId() MsgId {
	return MSG_HELLO_WORLD
}

func (self *MessageHelloWorld) GetMessageTag() MsgTag {
	return self.message_tag
}

func (self *MessageHelloWorld) GetPayloadData() PayloadData {
	return PayloadData{}
}

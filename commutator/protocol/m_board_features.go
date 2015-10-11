package protocol

import "bytes"

type MessageBoardFeatures struct {
	message_tag MsgTag
	is_request  bool
	buttons     []bool
	switches    []bool
	sliders     []uint8
}

func NewMessageBoardFeaturesRequest(tagger *Tagger) MessageBoardFeatures {
	return MessageBoardFeatures{
		message_tag: tagger.Next(),
		is_request:  true,
		buttons:     make([]bool, 0),
		switches:    make([]bool, 0),
		sliders:     make([]uint8, 0),
	}
}

func NewMessageBoardFeaturesResponse(tagger *Tagger, buttons []bool, switches []bool, sliders []uint8) MessageBoardFeatures {
	return MessageBoardFeatures{
		message_tag: tagger.Next(),
		is_request:  false,
		buttons:     buttons,
		switches:    switches,
		sliders:     sliders,
	}
}

func (self *MessageBoardFeatures) GetMessageId() MsgId {
	return MSG_BOARD_FEATURES
}

func (self *MessageBoardFeatures) GetMessageTag() MsgTag {
	return self.message_tag
}

func ToBitArray(bits []bool) []uint8 {
	size := len(bits) / 8

	if len(bits)%8 != 0 {
		size += 1
	}

	bit_array := make([]uint8, size)

	for i, bit := range bits {
		if bit {
			chunk_id := i / 8
			chunk_offset := uint(i % 8)
			bit_array[chunk_id] |= uint8(1 << chunk_offset)
		}
	}

	return bit_array
}

func (self *MessageBoardFeatures) GetPayloadData() PayloadData {
	if self.is_request {
		return make([]uint8, 0)
	} else {
		payload := new(bytes.Buffer)
		payload.WriteByte(uint8(len(self.buttons)))
		payload.Write(ToBitArray(self.buttons))
		payload.WriteByte(uint8(len(self.switches)))
		payload.Write(ToBitArray(self.switches))
		payload.WriteByte(uint8(len(self.sliders)))
		payload.Write(self.sliders)
		return payload.Bytes()
	}
}

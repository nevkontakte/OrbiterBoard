package protocol

const (
	TAGGER_BOARD = iota
	TAGGER_PC    = iota
)

type Tagger struct {
	counter_mask     MsgTag
	highest_bit_mask MsgTag
	current_tag      MsgTag
}

func NewTagger(mode int) Tagger {
	self := Tagger{}

	self.counter_mask = ^MsgTag(0) >> 1

	if mode == TAGGER_PC {
		self.highest_bit_mask = 0
	} else if mode == TAGGER_BOARD {
		self.highest_bit_mask = ^self.counter_mask
	} else {
		panic("Invalid tagger mode")
	}

	return self
}

func (self *Tagger) Next() MsgTag {
	self.current_tag = (self.current_tag + 1) & self.counter_mask
	return self.current_tag | self.highest_bit_mask
}

package main

import (
	"github.com/nevkontakte/OrbiterBoard/commutator/protocol"
	"log"
)

func main() {
	//    log.Print(protocol.NewPacket(protocol.MSG_HELLO_WORLD, 0))
	log.Print(protocol.NewTagger(protocol.TAGGER_PC))
	log.Print(protocol.NewTagger(protocol.TAGGER_BOARD))
}

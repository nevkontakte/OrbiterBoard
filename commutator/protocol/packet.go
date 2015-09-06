package protocol
import (
    "io"
    "encoding/binary"
)

const (
    MSG_RESERVED MsgId = 0
    MSG_HELLO_WORLD MsgId = 1
    MSG_BOARD_FEATURES MsgId = 2
    MSG_EVENT MsgId = 3
    MSG_SYNC_REQUIRED MsgId = 4
)

type MsgId uint8
type MsgTag uint16
type PayloadDataLength uint16
type PayloadData []uint8

type PacketMeta struct {
    message_id MsgId
    message_tag MsgTag
}

type Packet struct {
    meta PacketMeta
    payload PayloadData
}

type Message interface {
    GetMessageId() MsgId
    GetMessageTag() MsgTag
    GetPayloadData() PayloadData
}

func Message2Packet(msg Message) Packet {
    return Packet{
        meta: PacketMeta{
            message_id: msg.GetMessageId(),
            message_tag: msg.GetMessageTag(),
        },
        payload: msg.GetPayloadData(),
    }
}

func SendPacket(writer io.Writer, packet Packet) error {
    if err := binary.Write(writer, binary.BigEndian, packet.meta.message_id); err != nil {
        return err
    }
    if err := binary.Write(writer, binary.BigEndian, packet.meta.message_tag); err != nil {
        return err
    }
    if err := binary.Write(writer, binary.BigEndian, uint16(len(packet.payload))); err != nil {
        return err
    }
    if err := binary.Write(writer, binary.BigEndian, packet.payload); err != nil {
        return err
    }
    return nil
}

func ReceivePacket(reader io.Reader) (Packet, error) {
    var packet Packet;

    if err := binary.Read(reader, binary.BigEndian, &packet.meta.message_id); err != nil {
        return Packet{}, err
    }
    if err := binary.Read(reader, binary.BigEndian, &packet.meta.message_tag); err != nil {
        return Packet{}, err
    }

    var payload_length uint16;
    if err := binary.Read(reader, binary.BigEndian, &payload_length); err != nil {
        return Packet{}, err
    }

    packet.payload = make(PayloadData, payload_length)
    if err := binary.Read(reader, binary.BigEndian, &packet.payload); err != nil {
        return Packet{}, err
    }

    return packet, nil
}

func SendMessage(writer io.Writer, msg Message) error {
    return SendPacket(writer, Message2Packet(msg))
}


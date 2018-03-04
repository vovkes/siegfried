// Copyright (c) 2014-2015 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package wire

import (
	"io"
)

// MsgFilterClear implements the Message interface and represents a wbcoin
// filterclear message which is used to reset a Bloom filter.
//
// This message was not added until protocol version BIP0037Version and has
// no payload.
type MsgFilterClear struct{}

// BtcDecode decodes r using the wbcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgFilterClear) BtcDecode(r io.Reader, pver uint32) error {
	return nil
}

// BtcEncode encodes the receiver to w using the wbcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgFilterClear) BtcEncode(w io.Writer, pver uint32) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgFilterClear) Command() string {
	return CmdFilterClear
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgFilterClear) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

// NewMsgFilterClear returns a new wbcoin filterclear message that conforms to the Message
// interface.  See MsgFilterClear for details.
func NewMsgFilterClear() *MsgFilterClear {
	return &MsgFilterClear{}
}

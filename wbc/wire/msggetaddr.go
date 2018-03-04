// Copyright (c) 2013-2015 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package wire

import (
	"io"
)

// MsgGetAddr implements the Message interface and represents a wbcoin
// getaddr message.  It is used to request a list of known active peers on the
// network from a peer to help identify potential nodes.  The list is returned
// via one or more addr messages (MsgAddr).
//
// This message has no payload.
type MsgGetAddr struct{}

// BtcDecode decodes r using the wbcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgGetAddr) BtcDecode(r io.Reader, pver uint32) error {
	return nil
}

// BtcEncode encodes the receiver to w using the wbcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgGetAddr) BtcEncode(w io.Writer, pver uint32) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgGetAddr) Command() string {
	return CmdGetAddr
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgGetAddr) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

// NewMsgGetAddr returns a new wbcoin getaddr message that conforms to the
// Message interface.  See MsgGetAddr for details.
func NewMsgGetAddr() *MsgGetAddr {
	return &MsgGetAddr{}
}

// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleAcceptFriendship(packet nex.PacketInterface) {
	if protocol.AcceptFriendship == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::AcceptFriendship not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var uiPlayer types.UInt32

	err := uiPlayer.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AcceptFriendship(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, uiPlayer)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AcceptFriendship(nil, packet, callID, uiPlayer)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

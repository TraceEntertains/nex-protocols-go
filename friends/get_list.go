// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetList(packet nex.PacketInterface) {
	var err error

	if protocol.GetList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::GetList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	byRelationship := types.NewPrimitiveU8(0)
	err = byRelationship.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetList(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	bReversed := types.NewPrimitiveBool(false)
	err = bReversed.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetList(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetList(nil, packet, callID, byRelationship, bReversed)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByID(packet nex.PacketInterface) {
	var err error

	if protocol.FindByID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::FindByID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstID := types.NewList[*types.PrimitiveU32]()
	lstID.Type = types.NewPrimitiveU32(0)
	err = lstID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByID(fmt.Errorf("Failed to read lstID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByID(nil, packet, callID, lstID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

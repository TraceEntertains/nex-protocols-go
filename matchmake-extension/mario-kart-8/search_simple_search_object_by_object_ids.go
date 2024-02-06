// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSearchSimpleSearchObjectByObjectIDs(packet nex.PacketInterface) {
	var err error

	if protocol.SearchSimpleSearchObjectByObjectIDs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionMarioKart8::SearchSimpleSearchObjectByObjectIDs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	objectIDs := types.NewList[*types.PrimitiveU32]()
	objectIDs.Type = types.NewPrimitiveU32(0)
	err = objectIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchSimpleSearchObjectByObjectIDs(fmt.Errorf("Failed to read objectIDs from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchSimpleSearchObjectByObjectIDs(nil, packet, callID, objectIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

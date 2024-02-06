// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCommonData(packet nex.PacketInterface) {
	var err error

	if protocol.DeleteCommonData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::DeleteCommonData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uniqueID := types.NewPrimitiveU64(0)
	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteCommonData(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteCommonData(nil, packet, callID, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUnregisterApplication(packet nex.PacketInterface) {
	var err error

	if protocol.UnregisterApplication == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AAUser::UnregisterApplication not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	titleID := types.NewPrimitiveU64(0)
	err = titleID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UnregisterApplication(fmt.Errorf("Failed to read titleID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UnregisterApplication(nil, packet, callID, titleID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	var err error

	if protocol.GetCommonData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::GetCommonData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	optionFlags := types.NewPrimitiveU32(0)
	err = optionFlags.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCommonData(fmt.Errorf("Failed to read optionFlags from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	principalID := types.NewPID(0)
	err = principalID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCommonData(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	nexUniqueID := types.NewPrimitiveU64(0)
	err = nexUniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCommonData(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCommonData(nil, packet, callID, optionFlags, principalID, nexUniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

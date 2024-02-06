// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleResetRateCustomRankingCounter(packet nex.PacketInterface) {
	var err error

	if protocol.ResetRateCustomRankingCounter == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::ResetRateCustomRankingCounter not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	applicationID := types.NewPrimitiveU32(0)
	err = applicationID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ResetRateCustomRankingCounter(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ResetRateCustomRankingCounter(nil, packet, callID, applicationID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

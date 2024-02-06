// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetRound(packet nex.PacketInterface) {
	var err error

	if protocol.GetRound == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeReferee::GetRound not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	roundID := types.NewPrimitiveU64(0)
	err = roundID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRound(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRound(nil, packet, callID, roundID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

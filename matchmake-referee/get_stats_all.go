// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-referee/types"
)

func (protocol *Protocol) handleGetStatsAll(packet nex.PacketInterface) {
	if protocol.GetStatsAll == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeReferee::GetStatsAll not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	target := matchmake_referee_types.NewMatchmakeRefereeStatsTarget()

	err := target.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStatsAll(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, target)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetStatsAll(nil, packet, callID, target)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

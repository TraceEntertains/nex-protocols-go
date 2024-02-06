// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDebugEndCommunityCompetitionParticipation(packet nex.PacketInterface) {
	if protocol.DebugEndCommunityCompetitionParticipation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, rmcError := protocol.DebugEndCommunityCompetitionParticipation(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

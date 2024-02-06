// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAutoTournamentMatchmake(packet nex.PacketInterface) {
	if protocol.AutoTournamentMatchmake == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionSuperSmashBros4::AutoTournamentMatchmake not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::AutoTournamentMatchmake STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, rmcError := protocol.AutoTournamentMatchmake(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

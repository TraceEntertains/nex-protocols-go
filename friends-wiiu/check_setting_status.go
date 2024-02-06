// Package protocol implements the Friends WiiU protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCheckSettingStatus(packet nex.PacketInterface) {
	if protocol.CheckSettingStatus == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::CheckSettingStatus not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, rmcError := protocol.CheckSettingStatus(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

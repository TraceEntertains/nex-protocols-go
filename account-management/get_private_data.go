// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPrivateData(packet nex.PacketInterface) {
	if protocol.GetPrivateData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::GetPrivateData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, rmcError := protocol.GetPrivateData(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

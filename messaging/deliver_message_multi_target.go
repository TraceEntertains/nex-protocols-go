// Package protocol implements the Messaging protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeliverMessageMultiTarget(packet nex.PacketInterface) {
	if protocol.DeliverMessageMultiTarget == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Messaging::DeliverMessageMultiTarget not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("Messaging::DeliverMessageMultiTarget STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	rmcMessage, rmcError := protocol.DeliverMessageMultiTarget(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

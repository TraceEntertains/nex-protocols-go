// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

func (protocol *Protocol) handleGetMessagesHeaders(packet nex.PacketInterface) {
	var err error

	if protocol.GetMessagesHeaders == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Messaging::GetMessagesHeaders not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	recipient := messaging_types.NewMessageRecipient()
	err = recipient.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMessagesHeaders(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	resultRange := types.NewResultRange()
	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMessagesHeaders(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMessagesHeaders(nil, packet, callID, recipient, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

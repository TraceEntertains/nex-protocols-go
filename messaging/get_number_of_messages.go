// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
)

func (protocol *Protocol) handleGetNumberOfMessages(packet nex.PacketInterface) {
	if protocol.GetNumberOfMessages == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Messaging::GetNumberOfMessages not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	recipient := messaging_types.NewMessageRecipient()

	err := recipient.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetNumberOfMessages(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, recipient)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetNumberOfMessages(nil, packet, callID, recipient)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
)

func (protocol *Protocol) handleDeleteMessages(packet nex.PacketInterface) {
	if protocol.DeleteMessages == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Messaging::DeleteMessages not implemented")

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
	var lstMessagesToDelete types.List[types.UInt32]

	var err error

	err = recipient.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteMessages(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, recipient, lstMessagesToDelete)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstMessagesToDelete.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteMessages(fmt.Errorf("Failed to read lstMessagesToDelete from parameters. %s", err.Error()), packet, callID, recipient, lstMessagesToDelete)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteMessages(nil, packet, callID, recipient, lstMessagesToDelete)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

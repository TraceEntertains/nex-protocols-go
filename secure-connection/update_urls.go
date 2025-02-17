// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateURLs(packet nex.PacketInterface) {
	if protocol.UpdateURLs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SecureConnection::UpdateURLs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var vecMyURLs types.List[types.StationURL]

	err := vecMyURLs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateURLs(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), packet, callID, vecMyURLs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateURLs(nil, packet, callID, vecMyURLs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

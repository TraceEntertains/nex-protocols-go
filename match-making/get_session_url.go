// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetSessionURL(packet nex.PacketInterface) {
	if protocol.GetSessionURL == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::GetSessionURL not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var idGathering types.UInt32

	err := idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSessionURL(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, idGathering)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSessionURL(nil, packet, callID, idGathering)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

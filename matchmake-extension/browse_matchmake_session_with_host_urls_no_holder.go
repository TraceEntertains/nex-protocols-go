// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleBrowseMatchmakeSessionWithHostURLsNoHolder(packet nex.PacketInterface) {
	if protocol.BrowseMatchmakeSessionWithHostURLsNoHolder == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::BrowseMatchmakeSessionWithHostURLsNoHolder not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	searchCriteria := match_making_types.NewMatchmakeSessionSearchCriteria()
	resultRange := types.NewResultRange()

	var err error

	err = searchCriteria.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BrowseMatchmakeSessionWithHostURLsNoHolder(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BrowseMatchmakeSessionWithHostURLsNoHolder(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BrowseMatchmakeSessionWithHostURLsNoHolder(nil, packet, callID, searchCriteria, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	var err error

	if protocol.GetSimplePlayingSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::GetSimplePlayingSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	listPID := types.NewList[*types.PID]()
	listPID.Type = types.NewPID(0)
	err = listPID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read listPID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	includeLoginUser := types.NewPrimitiveBool(false)
	err = includeLoginUser.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read includeLoginUser from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSimplePlayingSession(nil, packet, callID, listPID, includeLoginUser)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

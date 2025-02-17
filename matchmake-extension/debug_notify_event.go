// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDebugNotifyEvent(packet nex.PacketInterface) {
	if protocol.DebugNotifyEvent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::DebugNotifyEvent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var pid types.PID
	var mainType types.UInt32
	var subType types.UInt32
	var param1 types.UInt64
	var param2 types.UInt64
	var stringParam types.String

	var err error

	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, pid, mainType, subType, param1, param2, stringParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = mainType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read mainType from parameters. %s", err.Error()), packet, callID, pid, mainType, subType, param1, param2, stringParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = subType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read subType from parameters. %s", err.Error()), packet, callID, pid, mainType, subType, param1, param2, stringParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = param1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read param1 from parameters. %s", err.Error()), packet, callID, pid, mainType, subType, param1, param2, stringParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = param2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read param2 from parameters. %s", err.Error()), packet, callID, pid, mainType, subType, param1, param2, stringParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = stringParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read stringParam from parameters. %s", err.Error()), packet, callID, pid, mainType, subType, param1, param2, stringParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DebugNotifyEvent(nil, packet, callID, pid, mainType, subType, param1, param2, stringParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

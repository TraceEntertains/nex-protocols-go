// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetUserStatuses(packet nex.PacketInterface) {
	var err error

	if protocol.GetUserStatuses == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::GetUserStatuses not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pids := types.NewList[*types.PID]()
	pids.Type = types.NewPID(0)
	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetUserStatuses(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	unknown := types.NewList[*types.PrimitiveU8]()
	unknown.Type = types.NewPrimitiveU8(0)
	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetUserStatuses(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetUserStatuses(nil, packet, callID, pids, unknown)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetProfiles(packet nex.PacketInterface) {
	var err error

	if protocol.GetProfiles == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperSmashBros4::GetProfiles not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pidList := types.NewList[*types.PID]()
	pidList.Type = types.NewPID(0)
	err = pidList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetProfiles(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetProfiles(nil, packet, callID, pidList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

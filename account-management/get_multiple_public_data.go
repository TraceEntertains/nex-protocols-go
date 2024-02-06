// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetMultiplePublicData(packet nex.PacketInterface) {
	var err error

	if protocol.GetMultiplePublicData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::GetMultiplePublicData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstPrincipals := types.NewList[*types.PID]()
	lstPrincipals.Type = types.NewPID(0)
	err = lstPrincipals.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMultiplePublicData(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMultiplePublicData(nil, packet, callID, lstPrincipals)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

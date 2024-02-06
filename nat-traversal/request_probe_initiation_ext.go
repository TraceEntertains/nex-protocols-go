// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	var err error

	if protocol.ReportNATProperties == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::RequestProbeInitiationExt not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	targetList := types.NewList[*types.String]()
	targetList.Type = types.NewString("")
	err = targetList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestProbeInitiationExt(fmt.Errorf("Failed to read targetList from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	stationToProbe := types.NewString("")
	err = stationToProbe.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestProbeInitiationExt(fmt.Errorf("Failed to read stationToProbe from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RequestProbeInitiationExt(nil, packet, callID, targetList, stationToProbe)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

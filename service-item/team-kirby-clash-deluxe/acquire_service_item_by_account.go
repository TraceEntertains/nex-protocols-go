// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleAcquireServiceItemByAccount(packet nex.PacketInterface) {
	if protocol.AcquireServiceItemByAccount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ServiceItemTeamKirbyClashDeluxe::AcquireServiceItemByAccount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	acquireServiceItemByAccountParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemAcquireServiceItemByAccountParam()

	err := acquireServiceItemByAccountParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AcquireServiceItemByAccount(fmt.Errorf("Failed to read acquireServiceItemByAccountParam from parameters. %s", err.Error()), packet, callID, acquireServiceItemByAccountParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AcquireServiceItemByAccount(nil, packet, callID, acquireServiceItemByAccountParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

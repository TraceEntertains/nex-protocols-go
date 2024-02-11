// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetRankingByPIDList(packet nex.PacketInterface) {
	if protocol.GetRankingByPIDList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetRankingByPIDList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	principalIDList := types.NewList[*types.PID]()
	principalIDList.Type = types.NewPID(0)
	rankingMode := types.NewPrimitiveU8(0)
	category := types.NewPrimitiveU32(0)
	orderParam := ranking_types.NewRankingOrderParam()
	uniqueID := types.NewPrimitiveU64(0)

	var err error

	err = principalIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = rankingMode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRankingByPIDList(nil, packet, callID, principalIDList, rankingMode, category, orderParam, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

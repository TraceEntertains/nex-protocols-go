// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/ranking/mario-kart-8/types"
)

func (protocol *Protocol) handleUploadCompetitionRankingScore(packet nex.PacketInterface) {
	var err error

	if protocol.UploadCompetitionRankingScore == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "RankingMarioKart8::UploadCompetitionRankingScore not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	param := ranking_mario_kart8_types.NewCompetitionRankingUploadScoreParam()
	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadCompetitionRankingScore(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UploadCompetitionRankingScore(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

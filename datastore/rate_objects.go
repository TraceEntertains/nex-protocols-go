// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRateObjects(packet nex.PacketInterface) {
	var err error

	if protocol.RateObjects == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::RateObjects not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	targets := types.NewList[*datastore_types.DataStoreRatingTarget]()
	targets.Type = datastore_types.NewDataStoreRatingTarget()
	err = targets.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	params := types.NewList[*datastore_types.DataStoreRateObjectParam]()
	params.Type = datastore_types.NewDataStoreRateObjectParam()
	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	transactional := types.NewPrimitiveBool(false)
	err = transactional.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	fetchRatings := types.NewPrimitiveBool(false)
	err = fetchRatings.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RateObjects(nil, packet, callID, targets, params, transactional, fetchRatings)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

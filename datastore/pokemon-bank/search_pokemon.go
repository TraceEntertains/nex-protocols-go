// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSearchPokemon(packet nex.PacketInterface) {
	if protocol.SearchPokemon == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::SearchPokemon not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := datastore_pokemon_bank_types.NewGlobalTradeStationSearchPokemonParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchPokemon(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, param)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchPokemon(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

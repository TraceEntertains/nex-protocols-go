package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostFightingPowerScore sets the PostFightingPowerScore handler function
func (protocol *DataStoreSuperSmashBros4Protocol) PostFightingPowerScore(handler func(err error, client *nex.Client, callID uint32, params []*DataStorePostFightingPowerScoreParam)) {
	protocol.PostFightingPowerScoreHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandlePostFightingPowerScore(packet nex.PacketInterface) {
	if protocol.PostFightingPowerScoreHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::PostFightingPowerScore not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(NewDataStorePostFightingPowerScoreParam())
	if err != nil {
		go protocol.PostFightingPowerScoreHandler(err, client, callID, nil)
		return
	}

	go protocol.PostFightingPowerScoreHandler(nil, client, callID, params.([]*DataStorePostFightingPowerScoreParam))
}

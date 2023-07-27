// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetProfiles sets the GetProfiles handler function
func (protocol *Protocol) GetProfiles(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	protocol.getProfilesHandler = handler
}

func (protocol *Protocol) handleGetProfiles(packet nex.PacketInterface) {
	if protocol.getProfilesHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetProfiles not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pidList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getProfilesHandler(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getProfilesHandler(nil, client, callID, pidList)
}

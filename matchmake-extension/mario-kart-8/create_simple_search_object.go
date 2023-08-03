// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

// CreateSimpleSearchObject sets the CreateSimpleSearchObject handler function
func (protocol *Protocol) CreateSimpleSearchObject(handler func(err error, client *nex.Client, callID uint32, object *matchmake_extension_mario_kart8_types.SimpleSearchObject)) {
	protocol.createSimpleSearchObjectHandler = handler
}

func (protocol *Protocol) handleCreateSimpleSearchObject(packet nex.PacketInterface) {
	if protocol.createSimpleSearchObjectHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::CreateSimpleSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	object, err := parametersStream.ReadStructure(matchmake_extension_mario_kart8_types.NewSimpleSearchObject())
	if err != nil {
		go protocol.createSimpleSearchObjectHandler(fmt.Errorf("Failed to read object from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.createSimpleSearchObjectHandler(nil, client, callID, object.(*matchmake_extension_mario_kart8_types.SimpleSearchObject))
}

// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrincipalIDByLocalFriendCode sets the GetPrincipalIDByLocalFriendCode handler function
func (protocol *Friends3DSProtocol) GetPrincipalIDByLocalFriendCode(handler func(err error, client *nex.Client, callID uint32, lfc uint64, lfcList []uint64)) {
	protocol.getPrincipalIDByLocalFriendCodeHandler = handler
}

func (protocol *Friends3DSProtocol) handleGetPrincipalIDByLocalFriendCode(packet nex.PacketInterface) {
	if protocol.getPrincipalIDByLocalFriendCodeHandler == nil {
		globals.Logger.Warning("Friends3DS::GetPrincipalIDByLocalFriendCode not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getPrincipalIDByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	lfcList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.getPrincipalIDByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.getPrincipalIDByLocalFriendCodeHandler(nil, client, callID, lfc, lfcList)
}

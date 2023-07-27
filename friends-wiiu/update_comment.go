// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateComment sets the UpdateComment handler function
func (protocol *Protocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment *friends_wiiu_types.Comment)) {
	protocol.updateCommentHandler = handler
}

func (protocol *Protocol) handleUpdateComment(packet nex.PacketInterface) {
	if protocol.updateCommentHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateComment not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	comment, err := parametersStream.ReadStructure(friends_wiiu_types.NewComment())
	if err != nil {
		go protocol.updateCommentHandler(fmt.Errorf("Failed to read comment from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateCommentHandler(nil, client, callID, comment.(*friends_wiiu_types.Comment))
}

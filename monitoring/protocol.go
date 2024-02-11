// Package protocol implements the Monitoring protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Monitoring protocol
	ProtocolID = 0x13

	// MethodPingDaemon is the method ID for the method PingDaemon
	MethodPingDaemon = 0x1

	// MethodGetClusterMembers is the method ID for the method GetClusterMembers
	MethodGetClusterMembers = 0x2
)

// Protocol handles the Monitoring protocol
type Protocol struct {
	endpoint          nex.EndpointInterface
	PingDaemon        func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetClusterMembers func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
}

// Interface implements the methods present on the Monitoring protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerPingDaemon(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetClusterMembers(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerPingDaemon sets the handler for the PingDaemon method
func (protocol *Protocol) SetHandlerPingDaemon(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.PingDaemon = handler
}

// SetHandlerGetClusterMembers sets the handler for the GetClusterMembers method
func (protocol *Protocol) SetHandlerGetClusterMembers(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetClusterMembers = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
	case MethodPingDaemon:
		protocol.handlePingDaemon(packet)
	case MethodGetClusterMembers:
		protocol.handleGetClusterMembers(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Monitoring method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Monitoring protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	return &Protocol{endpoint: endpoint}
}

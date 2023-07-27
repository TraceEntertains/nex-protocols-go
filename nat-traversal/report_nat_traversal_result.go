// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATTraversalResult sets the ReportNATTraversalResult handler function
func (protocol *Protocol) ReportNATTraversalResult(handler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32)) {
	protocol.reportNATTraversalResultHandler = handler
}

func (protocol *Protocol) handleReportNATTraversalResult(packet nex.PacketInterface) {
	if protocol.reportNATTraversalResultHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResult not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	natTraversalVersion := protocol.Server.NATTraversalProtocolVersion()

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.reportNATTraversalResultHandler(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), client, callID, 0, false, 0)
		return
	}

	result, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.reportNATTraversalResultHandler(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), client, callID, 0, false, 0)
		return
	}

	var rtt uint32 = 0

	// TODO - Is this the right version?
	if natTraversalVersion.Major >= 3 && natTraversalVersion.Minor >= 0 {
		rtt, err = parametersStream.ReadUInt32LE()
		if err != nil {
			go protocol.reportNATTraversalResultHandler(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), client, callID, 0, false, 0)
			return
		}
	}

	go protocol.reportNATTraversalResultHandler(nil, client, callID, cid, result, rtt)
}

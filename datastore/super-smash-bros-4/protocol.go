// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the DataStore (Super Smash Bros 4) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodPostProfile is the method ID for the method PostProfile
	MethodPostProfile = 0x2D

	// MethodGetProfiles is the method ID for the method GetProfiles
	MethodGetProfiles = 0x2E

	// MethodSendPlayReport is the method ID for the method SendPlayReport
	MethodSendPlayReport = 0x2F

	// MethodGetWorldPlayReport is the method ID for the method GetWorldPlayReport
	MethodGetWorldPlayReport = 0x30

	// MethodGetReplayMeta is the method ID for the method GetReplayMeta
	MethodGetReplayMeta = 0x31

	// MethodPrepareGetReplay is the method ID for the method PrepareGetReplay
	MethodPrepareGetReplay = 0x32

	// MethodPreparePostReplay is the method ID for the method PreparePostReplay
	MethodPreparePostReplay = 0x33

	// MethodCompletePostReplay is the method ID for the method CompletePostReplay
	MethodCompletePostReplay = 0x34

	// MethodCheckPostReplay is the method ID for the method CheckPostReplay
	MethodCheckPostReplay = 0x35

	// MethodGetNextReplay is the method ID for the method GetNextReplay
	MethodGetNextReplay = 0x36

	// MethodPreparePostSharedData is the method ID for the method PreparePostSharedData
	MethodPreparePostSharedData = 0x37

	// MethodCompletePostSharedData is the method ID for the method CompletePostSharedData
	MethodCompletePostSharedData = 0x38

	// MethodSearchSharedData is the method ID for the method SearchSharedData
	MethodSearchSharedData = 0x39

	// MethodGetApplicationConfig is the method ID for the method GetApplicationConfig
	MethodGetApplicationConfig = 0x3A

	// MethodSearchReplay is the method ID for the method SearchReplay
	MethodSearchReplay = 0x3B

	// MethodPostFightingPowerScore is the method ID for the method PostFightingPowerScore
	MethodPostFightingPowerScore = 0x3C

	// MethodGetFightingPowerChart is the method ID for the method GetFightingPowerChart
	MethodGetFightingPowerChart = 0x3D

	// MethodGetFightingPowerChartAll is the method ID for the method GetFightingPowerChartAll
	MethodGetFightingPowerChartAll = 0x3E

	// MethodReportSharedData is the method ID for the method ReportSharedData
	MethodReportSharedData = 0x3F
)

var patchedMethods = []uint32{
	MethodPostProfile,
	MethodGetProfiles,
	MethodSendPlayReport,
	MethodGetWorldPlayReport,
	MethodGetReplayMeta,
	MethodPrepareGetReplay,
	MethodPreparePostReplay,
	MethodCompletePostReplay,
	MethodCheckPostReplay,
	MethodGetNextReplay,
	MethodPreparePostSharedData,
	MethodCompletePostSharedData,
	MethodSearchSharedData,
	MethodGetApplicationConfig,
	MethodSearchReplay,
	MethodPostFightingPowerScore,
	MethodGetFightingPowerChart,
	MethodGetFightingPowerChartAll,
	MethodReportSharedData,
}

type datastoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Super Smash Bros 4) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	Server *nex.Server
	datastoreProtocol
	postProfileHandler              func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePostProfileParam)
	getProfilesHandler              func(err error, client *nex.Client, callID uint32, pidList []uint32)
	sendPlayReportHandler           func(err error, client *nex.Client, callID uint32, playReport []int32)
	getWorldPlayReportHandler       func(err error, client *nex.Client, callID uint32)
	getReplayMetaHandler            func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreGetReplayMetaParam)
	prepareGetReplayHandler         func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePrepareGetReplayParam)
	preparePostReplayHandler        func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostReplayParam)
	completePostReplayHandler       func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreCompletePostReplayParam)
	checkPostReplayHandler          func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostReplayParam)
	getNextReplayHandler            func(err error, client *nex.Client, callID uint32)
	preparePostSharedDataHandler    func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostSharedDataParam)
	completePostSharedDataHandler   func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreCompletePostSharedDataParam)
	searchSharedDataHandler         func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreSearchSharedDataParam)
	getApplicationConfigHandler     func(err error, client *nex.Client, callID uint32, applicationID uint32)
	searchReplayHandler             func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreSearchReplayParam)
	postFightingPowerScoreHandler   func(err error, client *nex.Client, callID uint32, params []*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam)
	getFightingPowerChartHandler    func(err error, client *nex.Client, callID uint32, mode uint8)
	getFightingPowerChartAllHandler func(err error, client *nex.Client, callID uint32)
	reportSharedDataHandler         func(err error, client *nex.Client, callID uint32, dataID uint64)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.datastoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodPostProfile:
		go protocol.handlePostProfile(packet)
	case MethodGetProfiles:
		go protocol.handleGetProfiles(packet)
	case MethodSendPlayReport:
		go protocol.handleSendPlayReport(packet)
	case MethodGetWorldPlayReport:
		go protocol.handleGetWorldPlayReport(packet)
	case MethodGetReplayMeta:
		go protocol.handleGetReplayMeta(packet)
	case MethodPrepareGetReplay:
		go protocol.handlePrepareGetReplay(packet)
	case MethodPreparePostReplay:
		go protocol.handlePreparePostReplay(packet)
	case MethodCompletePostReplay:
		go protocol.handleCompletePostReplay(packet)
	case MethodCheckPostReplay:
		go protocol.handleCheckPostReplay(packet)
	case MethodGetNextReplay:
		go protocol.handleGetNextReplay(packet)
	case MethodPreparePostSharedData:
		go protocol.handlePreparePostSharedData(packet)
	case MethodCompletePostSharedData:
		go protocol.handleCompletePostSharedData(packet)
	case MethodSearchSharedData:
		go protocol.handleSearchSharedData(packet)
	case MethodGetApplicationConfig:
		go protocol.handleGetApplicationConfig(packet)
	case MethodSearchReplay:
		go protocol.handleSearchReplay(packet)
	case MethodPostFightingPowerScore:
		go protocol.handlePostFightingPowerScore(packet)
	case MethodGetFightingPowerChart:
		go protocol.handleGetFightingPowerChart(packet)
	case MethodGetFightingPowerChartAll:
		go protocol.handleGetFightingPowerChartAll(packet)
	case MethodReportSharedData:
		go protocol.handleReportSharedData(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStoreSuperSmashBros4 method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new DataStore (Super Smash Bros 4) protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.datastoreProtocol.Server = server

	protocol.Setup()

	return protocol
}

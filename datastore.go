package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreProtocolID is the protocol ID for the DataStore protocol
	DataStoreProtocolID = 0x73

	// DataStoreMethodPrepareGetObjectV1 is the method ID for the method PrepareGetObjectV1
	DataStoreMethodPrepareGetObjectV1 = 0x1

	// DataStoreMethodPreparePostObjectV1 is the method ID for the method PreparePostObjectV1
	DataStoreMethodPreparePostObjectV1 = 0x2

	// DataStoreMethodCompletePostObjectV1 is the method ID for the method CompletePostObjectV1
	DataStoreMethodCompletePostObjectV1 = 0x3

	// DataStoreMethodDeleteObject is the method ID for the method DeleteObject
	DataStoreMethodDeleteObject = 0x4

	// DataStoreMethodDeleteObjects is the method ID for the method DeleteObjects
	DataStoreMethodDeleteObjects = 0x5

	// DataStoreMethodChangeMetaV1 is the method ID for the method ChangeMetaV1
	DataStoreMethodChangeMetaV1 = 0x6

	// DataStoreMethodChangeMetasV1 is the method ID for the method ChangeMetasV1
	DataStoreMethodChangeMetasV1 = 0x7

	// DataStoreMethodGetMeta is the method ID for the method GetMeta
	DataStoreMethodGetMeta = 0x8

	// DataStoreMethodGetMetas is the method ID for the method GetMetas
	DataStoreMethodGetMetas = 0x9

	// DataStoreMethodPrepareUpdateObject is the method ID for the method PrepareUpdateObject
	DataStoreMethodPrepareUpdateObject = 0xA

	// DataStoreMethodCompleteUpdateObject is the method ID for the method CompleteUpdateObject
	DataStoreMethodCompleteUpdateObject = 0xB

	// DataStoreMethodSearchObject is the method ID for the method SearchObject
	DataStoreMethodSearchObject = 0xC

	// DataStoreMethodGetNotificationURL is the method ID for the method GetNotificationUrl
	DataStoreMethodGetNotificationURL = 0xD

	// DataStoreMethodGetNewArrivedNotificationsV1 is the method ID for the method GetNewArrivedNotificationsV1
	DataStoreMethodGetNewArrivedNotificationsV1 = 0xE

	// DataStoreMethodRateObject is the method ID for the method RateObject
	DataStoreMethodRateObject = 0xF

	// DataStoreMethodGetRating is the method ID for the method GetRating
	DataStoreMethodGetRating = 0x10

	// DataStoreMethodGetRatings is the method ID for the method GetRatings
	DataStoreMethodGetRatings = 0x11

	// DataStoreMethodResetRating is the method ID for the method ResetRating
	DataStoreMethodResetRating = 0x12

	// DataStoreMethodResetRatings is the method ID for the method ResetRatings
	DataStoreMethodResetRatings = 0x13

	// DataStoreMethodGetSpecificMetaV1 is the method ID for the method GetSpecificMetaV1
	DataStoreMethodGetSpecificMetaV1 = 0x14

	// DataStoreMethodPostMetaBinary is the method ID for the method PostMetaBinary
	DataStoreMethodPostMetaBinary = 0x15

	// DataStoreMethodTouchObject is the method ID for the method TouchObject
	DataStoreMethodTouchObject = 0x16

	// DataStoreMethodGetRatingWithLog is the method ID for the method GetRatingWithLog
	DataStoreMethodGetRatingWithLog = 0x17

	// DataStoreMethodPreparePostObject is the method ID for the method PreparePostObject
	DataStoreMethodPreparePostObject = 0x18

	// DataStoreMethodPrepareGetObject is the method ID for the method PrepareGetObject
	DataStoreMethodPrepareGetObject = 0x19

	// DataStoreMethodCompletePostObject is the method ID for the method CompletePostObject
	DataStoreMethodCompletePostObject = 0x1A

	// DataStoreMethodGetNewArrivedNotifications is the method ID for the method GetNewArrivedNotifications
	DataStoreMethodGetNewArrivedNotifications = 0x1B

	// DataStoreMethodGetSpecificMeta is the method ID for the method GetSpecificMeta
	DataStoreMethodGetSpecificMeta = 0x1C

	// DataStoreMethodGetPersistenceInfo is the method ID for the method GetPersistenceInfo
	DataStoreMethodGetPersistenceInfo = 0x1D

	// DataStoreMethodGetPersistenceInfos is the method ID for the method GetPersistenceInfos
	DataStoreMethodGetPersistenceInfos = 0x1E

	// DataStoreMethodPerpetuateObject is the method ID for the method PerpetuateObject
	DataStoreMethodPerpetuateObject = 0x1F

	// DataStoreMethodUnperpetuateObject is the method ID for the method UnperpetuateObject
	DataStoreMethodUnperpetuateObject = 0x20

	// DataStoreMethodPrepareGetObjectOrMetaBinary is the method ID for the method PrepareGetObjectOrMetaBinary
	DataStoreMethodPrepareGetObjectOrMetaBinary = 0x21

	// DataStoreMethodGetPasswordInfo is the method ID for the method GetPasswordInfo
	DataStoreMethodGetPasswordInfo = 0x22

	// DataStoreMethodGetPasswordInfos is the method ID for the method GetPasswordInfos
	DataStoreMethodGetPasswordInfos = 0x23

	// DataStoreMethodGetMetasMultipleParam is the method ID for the method GetMetasMultipleParam
	DataStoreMethodGetMetasMultipleParam = 0x24

	// DataStoreMethodCompletePostObjects is the method ID for the method CompletePostObjects
	DataStoreMethodCompletePostObjects = 0x25

	// DataStoreMethodChangeMeta is the method ID for the method ChangeMeta
	DataStoreMethodChangeMeta = 0x26

	// DataStoreMethodChangeMetas is the method ID for the method ChangeMetas
	DataStoreMethodChangeMetas = 0x27

	// DataStoreMethodRateObjects is the method ID for the method RateObjects
	DataStoreMethodRateObjects = 0x28

	// DataStoreMethodPostMetaBinaryWithDataID is the method ID for the method PostMetaBinaryWithDataId
	DataStoreMethodPostMetaBinaryWithDataID = 0x29

	// DataStoreMethodPostMetaBinariesWithDataID is the method ID for the method PostMetaBinariesWithDataId
	DataStoreMethodPostMetaBinariesWithDataID = 0x2A

	// DataStoreMethodRateObjectWithPosting is the method ID for the method RateObjectWithPosting
	DataStoreMethodRateObjectWithPosting = 0x2B

	// DataStoreMethodRateObjectsWithPosting is the method ID for the method RateObjectsWithPosting
	DataStoreMethodRateObjectsWithPosting = 0x2C

	// DataStoreMethodGetObjectInfos is the method ID for the method GetObjectInfos
	DataStoreMethodGetObjectInfos = 0x2D

	// DataStoreMethodSearchObjectLight is the method ID for the method SearchObjectLight
	DataStoreMethodSearchObjectLight = 0x2E
)

// DataStoreProtocol handles the DataStore nex protocol
type DataStoreProtocol struct {
	server                       *nex.Server
	GetMetaHandler               func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)
	PrepareGetObjectHandler      func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *DataStorePrepareGetParam)
	GetMetasMultipleParamHandler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParams []*DataStoreGetMetaParam)
	ChangeMetaHandler            func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *DataStoreChangeMetaParam)
}

// DataStoreSearchParam is sent in DataStore search methods
type DataStoreSearchParam struct {
	SearchTarget           uint8
	OwnerIds               []uint32
	OwnerType              uint8
	DestinationIds         []uint64
	DataType               uint16
	CreatedAfter           *nex.DateTime
	CreatedBefore          *nex.DateTime
	UpdatedAfter           *nex.DateTime
	UpdatedBefore          *nex.DateTime
	ReferDataId            uint32
	Tags                   []string
	ResultOrderColumn      uint8
	ResultOrder            uint8
	ResultRange            *nex.ResultRange
	ResultOption           uint8
	MinimalRatingFrequency uint32
	UseCache               bool
	nex.Structure
}

// ExtractFromStream extracts a DataStoreSearchParam structure from a stream
func (dataStoreSearchParam *DataStoreSearchParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSearchParam.SearchTarget = stream.ReadUInt8()
	dataStoreSearchParam.OwnerIds = stream.ReadListUInt32LE()
	dataStoreSearchParam.OwnerType = stream.ReadUInt8()
	dataStoreSearchParam.DestinationIds = stream.ReadListUInt64LE()
	dataStoreSearchParam.DataType = stream.ReadUInt16LE()
	dataStoreSearchParam.CreatedAfter = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.CreatedBefore = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.UpdatedAfter = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.UpdatedBefore = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.ReferDataId = stream.ReadUInt32LE()
	dataStoreSearchParam.Tags = stream.ReadListString()
	dataStoreSearchParam.ResultOrderColumn = stream.ReadUInt8()
	dataStoreSearchParam.ResultOrder = stream.ReadUInt8()

	resultRange, err := stream.ReadStructure(nex.NewResultRange())

	if err != nil {
		return err
	}

	dataStoreSearchParam.ResultRange = resultRange.(*nex.ResultRange)
	dataStoreSearchParam.ResultOption = stream.ReadUInt8()
	dataStoreSearchParam.MinimalRatingFrequency = stream.ReadUInt32LE()
	dataStoreSearchParam.UseCache = (stream.ReadUInt8() == 1)

	return nil
}

// NewDataStoreSearchParam returns a new DataStoreSearchParam
func NewDataStoreSearchParam() *DataStoreSearchParam {
	return &DataStoreSearchParam{}
}

// DataStoreGetMetaParam is sent in the GetMeta method
type DataStoreGetMetaParam struct {
	nex.Structure
	DataID            uint64
	PersistenceTarget *DataStorePersistenceTarget
	ResultOption      uint8
	AccessPassword    uint64
}

// ExtractFromStream extracts a DataStoreGetMetaParam structure from a stream
func (dataStoreGetMetaParam *DataStoreGetMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 23 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStoreGetMetaParam::ExtractFromStream] Data size too small")
	}

	dataID := stream.ReadUInt64LE()
	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())

	if err != nil {
		return err
	}

	dataStoreGetMetaParam.DataID = dataID
	dataStoreGetMetaParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStoreGetMetaParam.ResultOption = stream.ReadUInt8()
	dataStoreGetMetaParam.AccessPassword = stream.ReadUInt64LE()

	return nil
}

// NewDataStoreGetMetaParam returns a new DataStoreGetMetaParam
func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	return &DataStoreGetMetaParam{}
}

// DataStoreChangeMetaParam is sent in the ChangeMeta method
type DataStoreChangeMetaParam struct {
	nex.Structure
	DataID         uint64
	ModifiesFlag   uint32
	Name           string
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         uint16
	MetaBinary     []byte
	Tags           []string
	UpdatePassword uint64
	ReferredCnt    uint32
	DataType       uint16
	Status         uint8
	CompareParam   *DataStoreChangeMetaCompareParam
	//PersistenceTarget *DataStorePersistenceTarget (not seen in SMM1??)
}

// ExtractFromStream extracts a DataStoreChangeMetaParam structure from a stream
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO: Check size

	dataStoreChangeMetaParam.DataID = stream.ReadUInt64LE()
	dataStoreChangeMetaParam.ModifiesFlag = stream.ReadUInt32LE()

	name, err := stream.ReadString()

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.Name = name

	permission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaParam.Period = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.MetaBinary = metaBinary
	dataStoreChangeMetaParam.Tags = stream.ReadListString()
	dataStoreChangeMetaParam.UpdatePassword = stream.ReadUInt64LE()
	dataStoreChangeMetaParam.ReferredCnt = stream.ReadUInt32LE()
	dataStoreChangeMetaParam.DataType = stream.ReadUInt16LE()
	dataStoreChangeMetaParam.Status = stream.ReadUInt8()

	compareParam, err := stream.ReadStructure(NewDataStoreChangeMetaCompareParam())

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.CompareParam = compareParam.(*DataStoreChangeMetaCompareParam)

	/*
		persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())

		if err != nil {
			return err
		}

		dataStoreChangeMetaParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	*/

	return nil
}

// NewDataStoreChangeMetaParam returns a new DataStoreChangeMetaParam
func NewDataStoreChangeMetaParam() *DataStoreChangeMetaParam {
	return &DataStoreChangeMetaParam{}
}

// DataStoreChangeMetaCompareParam is sent in the ChangeMeta method
type DataStoreChangeMetaCompareParam struct {
	nex.Structure
	ComparisonFlag uint32
	Name           string
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         uint16
	MetaBinary     []byte
	Tags           []string
	ReferredCnt    uint32
	DataType       uint16
	Status         uint8
}

// ExtractFromStream extracts a DataStoreChangeMetaCompareParam structure from a stream
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO: Check size

	dataStoreChangeMetaCompareParam.ComparisonFlag = stream.ReadUInt32LE()

	name, err := stream.ReadString()

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.Name = name

	permission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaCompareParam.Period = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.MetaBinary = metaBinary
	dataStoreChangeMetaCompareParam.Tags = stream.ReadListString()
	dataStoreChangeMetaCompareParam.ReferredCnt = stream.ReadUInt32LE()
	dataStoreChangeMetaCompareParam.DataType = stream.ReadUInt16LE()
	dataStoreChangeMetaCompareParam.Status = stream.ReadUInt8()

	return nil
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaParam
func NewDataStoreChangeMetaCompareParam() *DataStoreChangeMetaCompareParam {
	return &DataStoreChangeMetaCompareParam{}
}

// DataStorePermission contains information about a permission for a DataStore object
type DataStorePermission struct {
	nex.Structure
	Permission   uint8
	RecipientIds []uint32
}

// ExtractFromStream extracts a DataStorePermission structure from a stream
func (dataStorePermission *DataStorePermission) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 9 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStorePermission::ExtractFromStream] Data size too small")
	}

	dataStorePermission.Permission = stream.ReadUInt8()
	dataStorePermission.RecipientIds = stream.ReadListUInt32LE()

	return nil
}

// NewDataStorePermission returns a new DataStorePermission
func NewDataStorePermission() *DataStorePermission {
	return &DataStorePermission{}
}

// DataStorePersistenceTarget contains information about a DataStore target
type DataStorePersistenceTarget struct {
	nex.Structure
	OwnerID           uint32
	PersistenceSlotID uint16
}

// ExtractFromStream extracts a DataStorePersistenceTarget structure from a stream
func (dataStorePersistenceTarget *DataStorePersistenceTarget) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 9 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStorePersistenceTarget::ExtractFromStream] Data size too small")
	}

	dataStorePersistenceTarget.OwnerID = stream.ReadUInt32LE()
	dataStorePersistenceTarget.PersistenceSlotID = stream.ReadUInt16LE()

	return nil
}

// NewDataStorePersistenceTarget returns a new DataStorePersistenceTarget
func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	return &DataStorePersistenceTarget{}
}

type DataStoreRatingInfo struct {
	nex.Structure
	TotalValue   int64
	Count        uint32
	InitialValue int64
}

// ExtractFromStream extracts a DataStoreRatingInfo structure from a stream
func (dataStoreRatingInfo *DataStoreRatingInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingInfo.TotalValue = int64(stream.ReadUInt64LE())
	dataStoreRatingInfo.Count = stream.ReadUInt32LE()
	dataStoreRatingInfo.InitialValue = int64(stream.ReadUInt64LE())

	return nil
}

// NewDataStoreRatingInfo returns a new DataStoreRatingInfo
func NewDataStoreRatingInfo() *DataStoreRatingInfo {
	return &DataStoreRatingInfo{}
}

type DataStoreRatingInfoWithSlot struct {
	nex.Structure
	Slot   int8
	Rating *DataStoreRatingInfo
}

// ExtractFromStream extracts a DataStoreRatingInfoWithSlot structure from a stream
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingInfoWithSlot.Slot = int8(stream.ReadUInt8())

	rating, err := stream.ReadStructure(NewDataStoreRatingInfo())

	if err != nil {
		return err
	}

	dataStoreRatingInfoWithSlot.Rating = rating.(*DataStoreRatingInfo)

	return nil
}

// NewDataStoreRatingInfoWithSlot returns a new DataStoreRatingInfoWithSlot
func NewDataStoreRatingInfoWithSlot() *DataStoreRatingInfoWithSlot {
	return &DataStoreRatingInfoWithSlot{}
}

// DataStoreMetaInfo contains DataStore meta information
type DataStoreMetaInfo struct {
	nex.Structure
	DataID        uint64
	OwnerID       uint32
	Size          uint32
	Name          string
	MetaBinary    []byte
	Permission    *DataStorePermission
	DelPermission *DataStorePermission
	CreatedTime   *nex.DateTime
	UpdatedTime   *nex.DateTime
	Period        uint16
	Status        uint8
	ReferredCnt   uint32
	ReferDataID   uint32
	Flag          uint32
	ReferredTime  *nex.DateTime
	ExpireTime    *nex.DateTime
	Tags          []string
	Ratings       []*DataStoreRatingInfoWithSlot
}

// ExtractFromStream extracts a DataStoreMetaInfo structure from a stream
func (dataStoreMetaInfo *DataStoreMetaInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreMetaInfo.DataID = stream.ReadUInt64LE()
	dataStoreMetaInfo.OwnerID = stream.ReadUInt32LE()
	dataStoreMetaInfo.Size = stream.ReadUInt32LE()

	name, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreMetaInfo.Name = name

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStoreMetaInfo.MetaBinary = metaBinary

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStoreMetaInfo.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStoreMetaInfo.DelPermission = delPermission.(*DataStorePermission)
	dataStoreMetaInfo.CreatedTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.UpdatedTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.Period = stream.ReadUInt16LE()
	dataStoreMetaInfo.Status = stream.ReadUInt8()
	dataStoreMetaInfo.ReferredCnt = stream.ReadUInt32LE()
	dataStoreMetaInfo.ReferDataID = stream.ReadUInt32LE()
	dataStoreMetaInfo.Flag = stream.ReadUInt32LE()
	dataStoreMetaInfo.ReferredTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.ExpireTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.Tags = stream.ReadListString()

	// TODO: Refactor this, it's disgusting
	nexProtoStream := NewStreamIn(stream.Bytes()[stream.ByteOffset():], stream.Server)

	ratings, err := nexProtoStream.ReadListDataStoreRatingInfoWithSlot()
	if err != nil {
		return err
	}

	dataStoreMetaInfo.Ratings = ratings

	return nil
}

// NewDataStoreMetaInfo returns a new DataStoreMetaInfo
func NewDataStoreMetaInfo() *DataStoreMetaInfo {
	return &DataStoreMetaInfo{}
}

// DataStorePrepareGetParam is sent in the PrepareGetObject method
type DataStorePrepareGetParam struct {
	nex.Structure
	DataID            uint64
	LockID            uint32
	PersistenceTarget *DataStorePersistenceTarget
	AccessPassword    uint64
	ExtraData         []string
}

// ExtractFromStream extracts a DataStorePrepareGetParam structure from a stream
func (dataStorePrepareGetParam *DataStorePrepareGetParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePrepareGetParam.DataID = stream.ReadUInt64LE()
	dataStorePrepareGetParam.LockID = stream.ReadUInt32LE()

	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())
	if err != nil {
		return err
	}

	dataStorePrepareGetParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStorePrepareGetParam.AccessPassword = stream.ReadUInt64LE()
	dataStorePrepareGetParam.ExtraData = stream.ReadListString()

	return nil
}

// NewDataStorePrepareGetParam returns a new DataStorePrepareGetParam
func NewDataStorePrepareGetParam() *DataStorePrepareGetParam {
	return &DataStorePrepareGetParam{}
}

// DataStoreKeyValue is sent in the PrepareGetObject method
type DataStoreKeyValue struct {
	nex.Structure
	Key   string
	Value string
}

// Bytes encodes the DataStoreKeyValue and returns a byte array
func (dataStoreKeyValue *DataStoreKeyValue) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreKeyValue.Key)
	stream.WriteString(dataStoreKeyValue.Value)

	return stream.Bytes()
}

// NewDataStoreKeyValue returns a new DataStoreKeyValue
func NewDataStoreKeyValue() *DataStoreKeyValue {
	return &DataStoreKeyValue{}
}

// DataStoreReqGetInfo is sent in the PrepareGetObject method
type DataStoreReqGetInfo struct {
	nex.Structure
	URL            string
	RequestHeaders []*DataStoreKeyValue
	Size           uint32
	RootCA         []byte
	DataID         uint64
}

// Bytes encodes the DataStoreReqGetInfo and returns a byte array
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreReqGetInfo.URL)
	stream.WriteListStructure(dataStoreReqGetInfo.RequestHeaders)
	stream.WriteUInt32LE(dataStoreReqGetInfo.Size)
	stream.WriteBuffer(dataStoreReqGetInfo.RootCA)
	stream.WriteUInt64LE(dataStoreReqGetInfo.DataID)

	return stream.Bytes()
}

// NewDataStoreReqGetInfo returns a new DataStoreReqGetInfo
func NewDataStoreReqGetInfo() *DataStoreReqGetInfo {
	return &DataStoreReqGetInfo{}
}

// Setup initializes the protocol
func (dataStoreProtocol *DataStoreProtocol) Setup() {
	nexServer := dataStoreProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if DataStoreProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case DataStoreMethodGetMeta:
				go dataStoreProtocol.handleGetMeta(packet)
			case DataStoreMethodPrepareGetObject:
				go dataStoreProtocol.handlePrepareGetObject(packet)
			case DataStoreMethodGetMetasMultipleParam:
				go dataStoreProtocol.handleGetMetasMultipleParam(packet)
			case DataStoreMethodChangeMeta:
				go dataStoreProtocol.handleChangeMeta(packet)
			default:
				fmt.Printf("Unsupported DataStore method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// GetMeta sets the GetMeta handler function
func (dataStoreProtocol *DataStoreProtocol) GetMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)) {
	dataStoreProtocol.GetMetaHandler = handler
}

// GetMeta sets the GetMeta handler function
func (dataStoreProtocol *DataStoreProtocol) PrepareGetObject(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *DataStorePrepareGetParam)) {
	dataStoreProtocol.PrepareGetObjectHandler = handler
}

// GetMetasMultipleParam sets the GetMetasMultipleParam handler function
func (dataStoreProtocol *DataStoreProtocol) GetMetasMultipleParam(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParams []*DataStoreGetMetaParam)) {
	dataStoreProtocol.GetMetasMultipleParamHandler = handler
}

// ChangeMeta sets the ChangeMeta handler function
func (dataStoreProtocol *DataStoreProtocol) ChangeMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *DataStoreChangeMetaParam)) {
	dataStoreProtocol.ChangeMetaHandler = handler
}

func (dataStoreProtocol *DataStoreProtocol) handleGetMeta(packet nex.PacketInterface) {
	if dataStoreProtocol.GetMetaHandler == nil {
		fmt.Println("[Warning] DataStoreProtocol::GetMeta not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreGetMetaParam, err := parametersStream.ReadStructure(NewDataStoreGetMetaParam())

	if err != nil {
		go dataStoreProtocol.GetMetaHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.GetMetaHandler(nil, client, callID, dataStoreGetMetaParam.(*DataStoreGetMetaParam))
}

func (dataStoreProtocol *DataStoreProtocol) handlePrepareGetObject(packet nex.PacketInterface) {
	if dataStoreProtocol.PrepareGetObjectHandler == nil {
		fmt.Println("[Warning] DataStoreProtocol::PrepareGetObject not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStorePrepareGetParam, err := parametersStream.ReadStructure(NewDataStorePrepareGetParam())

	if err != nil {
		go dataStoreProtocol.PrepareGetObjectHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.PrepareGetObjectHandler(nil, client, callID, dataStorePrepareGetParam.(*DataStorePrepareGetParam))
}

func (dataStoreProtocol *DataStoreProtocol) handleGetMetasMultipleParam(packet nex.PacketInterface) {
	if dataStoreProtocol.GetMetasMultipleParamHandler == nil {
		fmt.Println("[Warning] DataStoreProtocol::GetMetasMultipleParam not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreGetMetaParams, err := parametersStream.ReaListDataStoreGetMetaParam()

	if err != nil {
		go dataStoreProtocol.GetMetasMultipleParamHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.GetMetasMultipleParamHandler(nil, client, callID, dataStoreGetMetaParams)
}

func (dataStoreProtocol *DataStoreProtocol) handleChangeMeta(packet nex.PacketInterface) {
	if dataStoreProtocol.ChangeMetaHandler == nil {
		fmt.Println("[Warning] DataStoreProtocol::ChangeMeta not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreChangeMetaParam, err := parametersStream.ReadStructure(NewDataStoreChangeMetaParam())

	if err != nil {
		go dataStoreProtocol.ChangeMetaHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.ChangeMetaHandler(nil, client, callID, dataStoreChangeMetaParam.(*DataStoreChangeMetaParam))
}

// NewDataStoreProtocol returns a new DataStoreProtocol
func NewDataStoreProtocol(server *nex.Server) *DataStoreProtocol {
	dataStoreProtocol := &DataStoreProtocol{server: server}

	dataStoreProtocol.Setup()

	return dataStoreProtocol
}

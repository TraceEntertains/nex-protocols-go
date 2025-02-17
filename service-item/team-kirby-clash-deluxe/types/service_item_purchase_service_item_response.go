// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPurchaseServiceItemResponse is a type within the ServiceItem protocol
type ServiceItemPurchaseServiceItemResponse struct {
	types.Structure
	ServiceItemEShopResponse
	NullablePurchaseInfo types.List[ServiceItemPurchaseInfo]
}

// WriteTo writes the ServiceItemPurchaseServiceItemResponse to the given writable
func (sipsir ServiceItemPurchaseServiceItemResponse) WriteTo(writable types.Writable) {
	sipsir.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sipsir.NullablePurchaseInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sipsir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPurchaseServiceItemResponse from the given readable
func (sipsir *ServiceItemPurchaseServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = sipsir.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = sipsir.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse header. %s", err.Error())
	}

	err = sipsir.NullablePurchaseInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse.NullablePurchaseInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemResponse
func (sipsir ServiceItemPurchaseServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemPurchaseServiceItemResponse()

	copied.StructureVersion = sipsir.StructureVersion
	copied.ServiceItemEShopResponse = sipsir.ServiceItemEShopResponse.Copy().(ServiceItemEShopResponse)
	copied.NullablePurchaseInfo = sipsir.NullablePurchaseInfo.Copy().(types.List[ServiceItemPurchaseInfo])

	return copied
}

// Equals checks if the given ServiceItemPurchaseServiceItemResponse contains the same data as the current ServiceItemPurchaseServiceItemResponse
func (sipsir ServiceItemPurchaseServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemPurchaseServiceItemResponse); !ok {
		return false
	}

	other := o.(ServiceItemPurchaseServiceItemResponse)

	if sipsir.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipsir.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return sipsir.NullablePurchaseInfo.Equals(other.NullablePurchaseInfo)
}

// CopyRef copies the current value of the ServiceItemPurchaseServiceItemResponse
// and returns a pointer to the new copy
func (sipsir ServiceItemPurchaseServiceItemResponse) CopyRef() types.RVTypePtr {
	copied := sipsir.Copy().(ServiceItemPurchaseServiceItemResponse)
	return &copied
}

// Deref takes a pointer to the ServiceItemPurchaseServiceItemResponse
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sipsir *ServiceItemPurchaseServiceItemResponse) Deref() types.RVType {
	return *sipsir
}

// String returns the string representation of the ServiceItemPurchaseServiceItemResponse
func (sipsir ServiceItemPurchaseServiceItemResponse) String() string {
	return sipsir.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPurchaseServiceItemResponse using the provided indentation level
func (sipsir ServiceItemPurchaseServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, sipsir.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullablePurchaseInfo: %s,\n", indentationValues, sipsir.NullablePurchaseInfo))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaseServiceItemResponse returns a new ServiceItemPurchaseServiceItemResponse
func NewServiceItemPurchaseServiceItemResponse() ServiceItemPurchaseServiceItemResponse {
	return ServiceItemPurchaseServiceItemResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullablePurchaseInfo:     types.NewList[ServiceItemPurchaseInfo](),
	}

}

// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2GetByListParam is a type within the Ranking2 protocol
type Ranking2GetByListParam struct {
	types.Structure
	Category           types.UInt32
	Offset             types.UInt32
	Length             types.UInt32
	SortFlags          types.UInt32
	OptionFlags        types.UInt32
	NumSeasonsToGoBack types.UInt8
}

// WriteTo writes the Ranking2GetByListParam to the given writable
func (rgblp Ranking2GetByListParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rgblp.Category.WriteTo(contentWritable)
	rgblp.Offset.WriteTo(contentWritable)
	rgblp.Length.WriteTo(contentWritable)
	rgblp.SortFlags.WriteTo(contentWritable)
	rgblp.OptionFlags.WriteTo(contentWritable)
	rgblp.NumSeasonsToGoBack.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rgblp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2GetByListParam from the given readable
func (rgblp *Ranking2GetByListParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = rgblp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam header. %s", err.Error())
	}

	err = rgblp.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Category. %s", err.Error())
	}

	err = rgblp.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Offset. %s", err.Error())
	}

	err = rgblp.Length.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Length. %s", err.Error())
	}

	err = rgblp.SortFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.SortFlags. %s", err.Error())
	}

	err = rgblp.OptionFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.OptionFlags. %s", err.Error())
	}

	err = rgblp.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.NumSeasonsToGoBack. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2GetByListParam
func (rgblp Ranking2GetByListParam) Copy() types.RVType {
	copied := NewRanking2GetByListParam()

	copied.StructureVersion = rgblp.StructureVersion
	copied.Category = rgblp.Category.Copy().(types.UInt32)
	copied.Offset = rgblp.Offset.Copy().(types.UInt32)
	copied.Length = rgblp.Length.Copy().(types.UInt32)
	copied.SortFlags = rgblp.SortFlags.Copy().(types.UInt32)
	copied.OptionFlags = rgblp.OptionFlags.Copy().(types.UInt32)
	copied.NumSeasonsToGoBack = rgblp.NumSeasonsToGoBack.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given Ranking2GetByListParam contains the same data as the current Ranking2GetByListParam
func (rgblp Ranking2GetByListParam) Equals(o types.RVType) bool {
	if _, ok := o.(Ranking2GetByListParam); !ok {
		return false
	}

	other := o.(Ranking2GetByListParam)

	if rgblp.StructureVersion != other.StructureVersion {
		return false
	}

	if !rgblp.Category.Equals(other.Category) {
		return false
	}

	if !rgblp.Offset.Equals(other.Offset) {
		return false
	}

	if !rgblp.Length.Equals(other.Length) {
		return false
	}

	if !rgblp.SortFlags.Equals(other.SortFlags) {
		return false
	}

	if !rgblp.OptionFlags.Equals(other.OptionFlags) {
		return false
	}

	return rgblp.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack)
}

// CopyRef copies the current value of the Ranking2GetByListParam
// and returns a pointer to the new copy
func (rgblp Ranking2GetByListParam) CopyRef() types.RVTypePtr {
	copied := rgblp.Copy().(Ranking2GetByListParam)
	return &copied
}

// Deref takes a pointer to the Ranking2GetByListParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rgblp *Ranking2GetByListParam) Deref() types.RVType {
	return *rgblp
}

// String returns the string representation of the Ranking2GetByListParam
func (rgblp Ranking2GetByListParam) String() string {
	return rgblp.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2GetByListParam using the provided indentation level
func (rgblp Ranking2GetByListParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2GetByListParam{\n")
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rgblp.Category))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, rgblp.Offset))
	b.WriteString(fmt.Sprintf("%sLength: %s,\n", indentationValues, rgblp.Length))
	b.WriteString(fmt.Sprintf("%sSortFlags: %s,\n", indentationValues, rgblp.SortFlags))
	b.WriteString(fmt.Sprintf("%sOptionFlags: %s,\n", indentationValues, rgblp.OptionFlags))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %s,\n", indentationValues, rgblp.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2GetByListParam returns a new Ranking2GetByListParam
func NewRanking2GetByListParam() Ranking2GetByListParam {
	return Ranking2GetByListParam{
		Category:           types.NewUInt32(0),
		Offset:             types.NewUInt32(0),
		Length:             types.NewUInt32(0),
		SortFlags:          types.NewUInt32(0),
		OptionFlags:        types.NewUInt32(0),
		NumSeasonsToGoBack: types.NewUInt8(0),
	}

}

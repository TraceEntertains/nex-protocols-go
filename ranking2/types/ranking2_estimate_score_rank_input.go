// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2EstimateScoreRankInput is a type within the Ranking2 protocol
type Ranking2EstimateScoreRankInput struct {
	types.Structure
	Category           types.UInt32
	NumSeasonsToGoBack types.UInt8
	Score              types.UInt32
}

// WriteTo writes the Ranking2EstimateScoreRankInput to the given writable
func (resri Ranking2EstimateScoreRankInput) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	resri.Category.WriteTo(contentWritable)
	resri.NumSeasonsToGoBack.WriteTo(contentWritable)
	resri.Score.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	resri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2EstimateScoreRankInput from the given readable
func (resri *Ranking2EstimateScoreRankInput) ExtractFrom(readable types.Readable) error {
	var err error

	err = resri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput header. %s", err.Error())
	}

	err = resri.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.Category. %s", err.Error())
	}

	err = resri.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.NumSeasonsToGoBack. %s", err.Error())
	}

	err = resri.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.Score. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2EstimateScoreRankInput
func (resri Ranking2EstimateScoreRankInput) Copy() types.RVType {
	copied := NewRanking2EstimateScoreRankInput()

	copied.StructureVersion = resri.StructureVersion
	copied.Category = resri.Category.Copy().(types.UInt32)
	copied.NumSeasonsToGoBack = resri.NumSeasonsToGoBack.Copy().(types.UInt8)
	copied.Score = resri.Score.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given Ranking2EstimateScoreRankInput contains the same data as the current Ranking2EstimateScoreRankInput
func (resri Ranking2EstimateScoreRankInput) Equals(o types.RVType) bool {
	if _, ok := o.(Ranking2EstimateScoreRankInput); !ok {
		return false
	}

	other := o.(Ranking2EstimateScoreRankInput)

	if resri.StructureVersion != other.StructureVersion {
		return false
	}

	if !resri.Category.Equals(other.Category) {
		return false
	}

	if !resri.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack) {
		return false
	}

	return resri.Score.Equals(other.Score)
}

// CopyRef copies the current value of the Ranking2EstimateScoreRankInput
// and returns a pointer to the new copy
func (resri Ranking2EstimateScoreRankInput) CopyRef() types.RVTypePtr {
	copied := resri.Copy().(Ranking2EstimateScoreRankInput)
	return &copied
}

// Deref takes a pointer to the Ranking2EstimateScoreRankInput
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (resri *Ranking2EstimateScoreRankInput) Deref() types.RVType {
	return *resri
}

// String returns the string representation of the Ranking2EstimateScoreRankInput
func (resri Ranking2EstimateScoreRankInput) String() string {
	return resri.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2EstimateScoreRankInput using the provided indentation level
func (resri Ranking2EstimateScoreRankInput) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2EstimateScoreRankInput{\n")
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, resri.Category))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %s,\n", indentationValues, resri.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, resri.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2EstimateScoreRankInput returns a new Ranking2EstimateScoreRankInput
func NewRanking2EstimateScoreRankInput() Ranking2EstimateScoreRankInput {
	return Ranking2EstimateScoreRankInput{
		Category:           types.NewUInt32(0),
		NumSeasonsToGoBack: types.NewUInt8(0),
		Score:              types.NewUInt32(0),
	}

}

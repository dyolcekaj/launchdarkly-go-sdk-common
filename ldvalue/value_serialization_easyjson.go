// +build launchdarkly_easyjson

package ldvalue

import (
	"gopkg.in/launchdarkly/go-jsonstream.v1/jreader"
	"gopkg.in/launchdarkly/go-jsonstream.v1/jwriter"

	"github.com/mailru/easyjson/jlexer"
	ej_jwriter "github.com/mailru/easyjson/jwriter"
)

// This conditionally-compiled file provides custom marshal/unmarshal functions for ldvalue types
// in EasyJSON.
//
// EasyJSON's code generator does recognize the same MarshalJSON and UnmarshalJSON methods used by
// encoding/json, and will call them if present. But this mechanism is inefficient: when marshaling
// it requires the allocation of intermediate byte slices, and when unmarshaling it causes the
// JSON object to be parsed twice. It is preferable to have our marshal/unmarshal methods write to
// and read from the EasyJSON Writer/Lexer directly. Our go-jsonstream library provides methods for
// doing this, if the launchdarkly_easyjson build tag is set.package ldmodel
//
// For more information, see: https://gopkg.in/launchdarkly/go-jsonstream.v1

func (v Value) MarshalEasyJSON(writer *ej_jwriter.Writer) {
	wrappedWriter := jwriter.NewWriterFromEasyJSONWriter(writer)
	v.WriteToJSONWriter(&wrappedWriter)
}

func (v *Value) UnmarshalEasyJSON(lexer *jlexer.Lexer) {
	wrappedReader := jreader.NewReaderFromEasyJSONLexer(lexer)
	v.ReadFromJSONReader(&wrappedReader)
	lexer.AddError(wrappedReader.Error())
}

func (v OptionalBool) MarshalEasyJSON(writer *ej_jwriter.Writer) {
	wrappedWriter := jwriter.NewWriterFromEasyJSONWriter(writer)
	v.WriteToJSONWriter(&wrappedWriter)
}

func (v *OptionalBool) UnmarshalEasyJSON(lexer *jlexer.Lexer) {
	wrappedReader := jreader.NewReaderFromEasyJSONLexer(lexer)
	v.ReadFromJSONReader(&wrappedReader)
	lexer.AddError(wrappedReader.Error())
}

func (v OptionalInt) MarshalEasyJSON(writer *ej_jwriter.Writer) {
	wrappedWriter := jwriter.NewWriterFromEasyJSONWriter(writer)
	v.WriteToJSONWriter(&wrappedWriter)
}

func (v *OptionalInt) UnmarshalEasyJSON(lexer *jlexer.Lexer) {
	wrappedReader := jreader.NewReaderFromEasyJSONLexer(lexer)
	v.ReadFromJSONReader(&wrappedReader)
	lexer.AddError(wrappedReader.Error())
}

func (v OptionalString) MarshalEasyJSON(writer *ej_jwriter.Writer) {
	wrappedWriter := jwriter.NewWriterFromEasyJSONWriter(writer)
	v.WriteToJSONWriter(&wrappedWriter)
}

func (v *OptionalString) UnmarshalEasyJSON(lexer *jlexer.Lexer) {
	wrappedReader := jreader.NewReaderFromEasyJSONLexer(lexer)
	v.ReadFromJSONReader(&wrappedReader)
	lexer.AddError(wrappedReader.Error())
}

func (v ValueArray) MarshalEasyJSON(writer *ej_jwriter.Writer) {
	wrappedWriter := jwriter.NewWriterFromEasyJSONWriter(writer)
	v.WriteToJSONWriter(&wrappedWriter)
}

func (v *ValueArray) UnmarshalEasyJSON(lexer *jlexer.Lexer) {
	wrappedReader := jreader.NewReaderFromEasyJSONLexer(lexer)
	v.ReadFromJSONReader(&wrappedReader)
	lexer.AddError(wrappedReader.Error())
}

func (v ValueMap) MarshalEasyJSON(writer *ej_jwriter.Writer) {
	wrappedWriter := jwriter.NewWriterFromEasyJSONWriter(writer)
	v.WriteToJSONWriter(&wrappedWriter)
}

func (v *ValueMap) UnmarshalEasyJSON(lexer *jlexer.Lexer) {
	wrappedReader := jreader.NewReaderFromEasyJSONLexer(lexer)
	v.ReadFromJSONReader(&wrappedReader)
	lexer.AddError(wrappedReader.Error())
}

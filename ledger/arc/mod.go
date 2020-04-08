// Package arc defines the interfaces for the Access Rights Control.
package arc

import (
	"strings"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/fabric/encoding"
)

// Identity is an abstraction to uniquely identify a signer.
type Identity interface {
	encoding.Packable
	encoding.TextMarshaler
}

// AccessControl is an abstraction to verify if an identity has access to a
// specific rule.
type AccessControl interface {
	encoding.Packable

	Match(rule string, idents ...Identity) error
}

// AccessControlFactory is an abstraction to decode access controls from
// protobuf messages.
type AccessControlFactory interface {
	FromProto(proto.Message) (AccessControl, error)
}

// Compile returns a compacted rule from the string segments.
func Compile(segments ...string) string {
	return strings.Join(segments, ":")
}
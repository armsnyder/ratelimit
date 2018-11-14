// This file is a manually-maintained and minimal implementation of
// https://github.com/envoyproxy/data-plane-api/blob/master/envoy/api/v2/core/base.proto
// since I couldn't figure out how to make protoc compile it because of its gogoproto options.
// TODO Make this file auto-generate like the other vendor proto files
package core

// Header name/value pair.
type HeaderValue struct {
	// Header name.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Header value.
	//
	// The same :ref:`format specifier <config_access_log_format>` as used for
	// :ref:`HTTP access logging <config_access_log>` applies here, however
	// unknown header values are replaced with the empty string instead of `-`.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

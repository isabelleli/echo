// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log/v1/HttpRequest.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// [START messages]
type HttpRequest struct {
	// method + url are the basics of http
	Method string `protobuf:"bytes,10,opt,name=method" json:"method,omitempty"`
	// url is the full / raw url
	Url string `protobuf:"bytes,11,opt,name=url" json:"url,omitempty"`
	// the following unpack the url into useful sections.
	Scheme      string `protobuf:"bytes,12,opt,name=scheme" json:"scheme,omitempty"`
	Host        string `protobuf:"bytes,13,opt,name=host" json:"host,omitempty"`
	Path        string `protobuf:"bytes,14,opt,name=path" json:"path,omitempty"`
	QueryString string `protobuf:"bytes,15,opt,name=queryString" json:"queryString,omitempty"`
	// route is kind of implementation specific;
	// it depends on the app having a muxer
	// route should be the path less actual variables
	Route string `protobuf:"bytes,16,opt,name=route" json:"route,omitempty"`
	// the following describe the client
	UserAgent string `protobuf:"bytes,17,opt,name=userAgent" json:"userAgent,omitempty"`
	// remoteAddr is the reported remote addr for the request
	// very often will be a load balancer or similar
	RemoteAddr string `protobuf:"bytes,18,opt,name=remoteAddr" json:"remoteAddr,omitempty"`
	// remoteIP should be forwarded addr headers, or the ip component of remoteAddr
	RemoteIP string `protobuf:"bytes,19,opt,name=remoteIP" json:"remoteIP,omitempty"`
	// referrer isn't useful in backend contexts
	// but could be for frontend
	Referrer string `protobuf:"bytes,20,opt,name=referrer" json:"referrer,omitempty"`
	// contentLength generally only applies to post requests
	ContentLength int64 `protobuf:"varint,30,opt,name=contentLength" json:"contentLength,omitempty"`
	// contentType generally only applies to post requests
	ContentType string `protobuf:"bytes,31,opt,name=contentType" json:"contentType,omitempty"`
	// contentEncoding generally only applies to post requests
	ContentEncoding string `protobuf:"bytes,32,opt,name=contentEncoding" json:"contentEncoding,omitempty"`
	// response fields
	Elapsed                 *google_protobuf.Duration `protobuf:"bytes,40,opt,name=elapsed" json:"elapsed,omitempty"`
	StatusCode              int32                     `protobuf:"varint,41,opt,name=statusCode" json:"statusCode,omitempty"`
	ResponseContentLength   int64                     `protobuf:"varint,42,opt,name=responseContentLength" json:"responseContentLength,omitempty"`
	ResponseContentEncoding string                    `protobuf:"bytes,43,opt,name=responseContentEncoding" json:"responseContentEncoding,omitempty"`
	ResponseContentType     string                    `protobuf:"bytes,44,opt,name=responseContentType" json:"responseContentType,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *HttpRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HttpRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *HttpRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HttpRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HttpRequest) GetQueryString() string {
	if m != nil {
		return m.QueryString
	}
	return ""
}

func (m *HttpRequest) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *HttpRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *HttpRequest) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *HttpRequest) GetRemoteIP() string {
	if m != nil {
		return m.RemoteIP
	}
	return ""
}

func (m *HttpRequest) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *HttpRequest) GetContentLength() int64 {
	if m != nil {
		return m.ContentLength
	}
	return 0
}

func (m *HttpRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HttpRequest) GetContentEncoding() string {
	if m != nil {
		return m.ContentEncoding
	}
	return ""
}

func (m *HttpRequest) GetElapsed() *google_protobuf.Duration {
	if m != nil {
		return m.Elapsed
	}
	return nil
}

func (m *HttpRequest) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *HttpRequest) GetResponseContentLength() int64 {
	if m != nil {
		return m.ResponseContentLength
	}
	return 0
}

func (m *HttpRequest) GetResponseContentEncoding() string {
	if m != nil {
		return m.ResponseContentEncoding
	}
	return ""
}

func (m *HttpRequest) GetResponseContentType() string {
	if m != nil {
		return m.ResponseContentType
	}
	return ""
}

func init() {
	proto.RegisterType((*HttpRequest)(nil), "logv1.HttpRequest")
}

func init() { proto.RegisterFile("log/v1/HttpRequest.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x15, 0x95, 0x14, 0x3a, 0xa1, 0xb4, 0x4c, 0x0b, 0x2c, 0x15, 0x0a, 0x16, 0x02, 0x61,
	0x3e, 0x64, 0x53, 0xca, 0x81, 0x6b, 0x29, 0x48, 0x20, 0x71, 0x40, 0x86, 0x13, 0x37, 0x27, 0x9e,
	0xda, 0x91, 0x9c, 0x1d, 0x77, 0x77, 0x5c, 0xa9, 0x7f, 0x82, 0xdf, 0x8c, 0x3c, 0x1b, 0x17, 0x37,
	0x0a, 0xb7, 0x7d, 0x9f, 0x67, 0x2c, 0xed, 0xeb, 0x59, 0x30, 0x35, 0x97, 0xe9, 0xe5, 0x71, 0xfa,
	0x55, 0xa4, 0xc9, 0xe8, 0xa2, 0x25, 0x2f, 0x49, 0xe3, 0x58, 0x18, 0xc7, 0x35, 0x97, 0x97, 0xc7,
	0x47, 0xd3, 0x92, 0xb9, 0xac, 0x29, 0x55, 0x38, 0x6b, 0xcf, 0xd3, 0xcf, 0xad, 0xcb, 0x65, 0xc1,
	0x36, 0x8c, 0x3d, 0xfb, 0x33, 0x86, 0xc9, 0xe0, 0x63, 0x7c, 0x08, 0xdb, 0x4b, 0x92, 0x8a, 0x0b,
	0x03, 0xd1, 0x28, 0xde, 0xc9, 0x56, 0x09, 0xf7, 0x61, 0xab, 0x75, 0xb5, 0x99, 0x28, 0xec, 0x8e,
	0xdd, 0xa4, 0x9f, 0x57, 0xb4, 0x24, 0x73, 0x37, 0x4c, 0x86, 0x84, 0x08, 0xb7, 0x2a, 0xf6, 0x62,
	0x76, 0x95, 0xea, 0xb9, 0x63, 0x4d, 0x2e, 0x95, 0xb9, 0x17, 0x58, 0x77, 0xc6, 0x08, 0x26, 0x17,
	0x2d, 0xb9, 0xab, 0x9f, 0xe2, 0x16, 0xb6, 0x34, 0x7b, 0xaa, 0x86, 0x08, 0x0f, 0x61, 0xec, 0xb8,
	0x15, 0x32, 0xfb, 0xea, 0x42, 0xc0, 0x27, 0xb0, 0xd3, 0x7a, 0x72, 0xa7, 0x25, 0x59, 0x31, 0xf7,
	0xd5, 0xfc, 0x03, 0x38, 0x05, 0x70, 0xb4, 0x64, 0xa1, 0xd3, 0xa2, 0x70, 0x06, 0x55, 0x0f, 0x08,
	0x1e, 0xc1, 0x9d, 0x90, 0xbe, 0xfd, 0x30, 0x07, 0x6a, 0xaf, 0x73, 0x70, 0xe7, 0xe4, 0x1c, 0x39,
	0x73, 0xd8, 0xbb, 0x90, 0xf1, 0x39, 0xec, 0xce, 0xd9, 0x0a, 0x59, 0xf9, 0x4e, 0xb6, 0x94, 0xca,
	0x4c, 0xa3, 0x51, 0xbc, 0x95, 0xdd, 0x84, 0x5d, 0xa7, 0x15, 0xf8, 0x75, 0xd5, 0x90, 0x79, 0x1a,
	0x3a, 0x0d, 0x10, 0xc6, 0xb0, 0xb7, 0x8a, 0x5f, 0xec, 0x9c, 0x8b, 0xae, 0x79, 0xa4, 0x53, 0xeb,
	0x18, 0x4f, 0xe0, 0x36, 0xd5, 0x79, 0xe3, 0xa9, 0x30, 0x71, 0x34, 0x8a, 0x27, 0xef, 0x1f, 0x27,
	0x61, 0x97, 0x49, 0xbf, 0xcb, 0xa4, 0xdf, 0x65, 0xd6, 0x4f, 0x76, 0xf5, 0xbd, 0xe4, 0xd2, 0xfa,
	0x33, 0x2e, 0xc8, 0xbc, 0x8a, 0x46, 0xf1, 0x38, 0x1b, 0x10, 0xfc, 0x00, 0x0f, 0x1c, 0xf9, 0x86,
	0xad, 0xa7, 0xb3, 0x1b, 0x75, 0x5e, 0x6b, 0x9d, 0xcd, 0x12, 0x3f, 0xc2, 0xa3, 0x35, 0x71, 0x7d,
	0xf9, 0x37, 0x7a, 0xf9, 0xff, 0x69, 0x7c, 0x07, 0x07, 0x6b, 0x4a, 0x7f, 0xcc, 0x5b, 0xfd, 0x6a,
	0x93, 0xfa, 0xf4, 0xf2, 0xf7, 0x8b, 0x72, 0x21, 0xc9, 0xac, 0x26, 0x5b, 0xd4, 0xf9, 0xcc, 0x27,
	0x73, 0x5e, 0xa6, 0x9a, 0xc2, 0x1b, 0xf6, 0x69, 0x78, 0xf2, 0xb3, 0x6d, 0x8d, 0x27, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0x61, 0x4d, 0xb2, 0x03, 0x03, 0x00, 0x00,
}

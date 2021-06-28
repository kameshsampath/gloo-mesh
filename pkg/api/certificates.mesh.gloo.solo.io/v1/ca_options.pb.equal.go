// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/certificates/v1/ca_options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *CommonCertOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CommonCertOptions)
	if !ok {
		that2, ok := that.(CommonCertOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetTtlDays() != target.GetTtlDays() {
		return false
	}

	if m.GetRsaKeySizeBytes() != target.GetRsaKeySizeBytes() {
		return false
	}

	if strings.Compare(m.GetOrgName(), target.GetOrgName()) != 0 {
		return false
	}

	if m.GetSecretRotationGracePeriodRatio() != target.GetSecretRotationGracePeriodRatio() {
		return false
	}

	return true
}

// Equal function
func (m *IntermediateCertificateAuthority) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IntermediateCertificateAuthority)
	if !ok {
		that2, ok := that.(IntermediateCertificateAuthority)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.CaSource.(type) {

	case *IntermediateCertificateAuthority_Vault:
		if _, ok := target.CaSource.(*IntermediateCertificateAuthority_Vault); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVault()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVault()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVault(), target.GetVault()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CaSource != target.CaSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *CertificateRotationCondition) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CertificateRotationCondition)
	if !ok {
		that2, ok := that.(CertificateRotationCondition)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetTimestamp(), target.GetTimestamp()) != 0 {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}

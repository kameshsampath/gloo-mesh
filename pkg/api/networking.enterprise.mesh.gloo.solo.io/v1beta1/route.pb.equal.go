// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/route.proto

package v1beta1

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
func (m *Route) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if len(m.GetMatchers()) != len(target.GetMatchers()) {
		return false
	}
	for idx, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchers()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	switch m.Action.(type) {

	case *Route_RouteAction_:
		if _, ok := target.Action.(*Route_RouteAction_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRouteAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRouteAction(), target.GetRouteAction()) {
				return false
			}
		}

	case *Route_RedirectAction:
		if _, ok := target.Action.(*Route_RedirectAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRedirectAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRedirectAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRedirectAction(), target.GetRedirectAction()) {
				return false
			}
		}

	case *Route_DirectResponseAction:
		if _, ok := target.Action.(*Route_DirectResponseAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDirectResponseAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDirectResponseAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDirectResponseAction(), target.GetDirectResponseAction()) {
				return false
			}
		}

	case *Route_DelegateAction:
		if _, ok := target.Action.(*Route_DelegateAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDelegateAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDelegateAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDelegateAction(), target.GetDelegateAction()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Action != target.Action {
			return false
		}
	}

	return true
}

// Equal function
func (m *RedirectAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RedirectAction)
	if !ok {
		that2, ok := that.(RedirectAction)
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

	if strings.Compare(m.GetHostRedirect(), target.GetHostRedirect()) != 0 {
		return false
	}

	if m.GetResponseCode() != target.GetResponseCode() {
		return false
	}

	if m.GetHttpsRedirect() != target.GetHttpsRedirect() {
		return false
	}

	if m.GetStripQuery() != target.GetStripQuery() {
		return false
	}

	switch m.PathRewriteSpecifier.(type) {

	case *RedirectAction_PathRedirect:
		if _, ok := target.PathRewriteSpecifier.(*RedirectAction_PathRedirect); !ok {
			return false
		}

		if strings.Compare(m.GetPathRedirect(), target.GetPathRedirect()) != 0 {
			return false
		}

	case *RedirectAction_PrefixRewrite:
		if _, ok := target.PathRewriteSpecifier.(*RedirectAction_PrefixRewrite); !ok {
			return false
		}

		if strings.Compare(m.GetPrefixRewrite(), target.GetPrefixRewrite()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.PathRewriteSpecifier != target.PathRewriteSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *DirectResponseAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DirectResponseAction)
	if !ok {
		that2, ok := that.(DirectResponseAction)
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

	if m.GetStatus() != target.GetStatus() {
		return false
	}

	if strings.Compare(m.GetBody(), target.GetBody()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *DelegateAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DelegateAction)
	if !ok {
		that2, ok := that.(DelegateAction)
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

	if len(m.GetRefs()) != len(target.GetRefs()) {
		return false
	}
	for idx, v := range m.GetRefs() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRefs()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRefs()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetSelector()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSelector()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSelector(), target.GetSelector()) {
			return false
		}
	}

	if m.GetSortRoutes() != target.GetSortRoutes() {
		return false
	}

	return true
}

// Equal function
func (m *Route_RouteAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Route_RouteAction)
	if !ok {
		that2, ok := that.(Route_RouteAction)
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

	if len(m.GetDestinations()) != len(target.GetDestinations()) {
		return false
	}
	for idx, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDestinations()[idx]) {
				return false
			}
		}

	}

	return true
}

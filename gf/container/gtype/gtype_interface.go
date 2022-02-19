// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtype

import (
	"sync/atomic"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// Interface is a struct for concurrent-safe operation for type interface{}.
// Interface 实际上就是对 atomic.Value 的封装
type Interface struct {
	value atomic.Value
}

// NewInterface creates and returns a concurrent-safe object for interface{} type,
// with given initial value `value`.
func NewInterface(value ...interface{}) *Interface {
	t := &Interface{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}

// Clone clones and returns a new concurrent-safe object for interface{} type.
func (v *Interface) Clone() *Interface {
	return NewInterface(v.Val())
}

// Set atomically stores `value` into t.value and returns the previous value of t.value.
// Note: The parameter `value` cannot be nil.
func (v *Interface) Set(value interface{}) (old interface{}) {
	old = v.Val()
	v.value.Store(value)
	return
}

// Val atomically loads and returns t.value.
func (v *Interface) Val() interface{} {
	return v.value.Load()
}

// String implements String interface for string printing.
func (v *Interface) String() string {
	return gconv.String(v.Val())
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (v *Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (v *Interface) UnmarshalJSON(b []byte) error {
	var i interface{}
	// 避免 float64 精度缺失的问题，详细解答可见：https://juejin.cn/post/7064954912810991653
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.Set(i)
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for `v`.
// TODO： 未看懂设计
func (v *Interface) UnmarshalValue(value interface{}) error {
	v.Set(value)
	return nil
}

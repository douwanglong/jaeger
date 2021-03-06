// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package model

import "io"

// Process describes an instance of an application or service that emits tracing data.
type Process struct {
	ServiceName string    `json:"serviceName"`
	Tags        KeyValues `json:"tags,omitempty"`
}

// NewProcess creates a new Process for given serviceName and tags.
// The tags are sorted in place and kept in the the same array/slice,
// in order to store the Process in a canonical form that is relied
// upon by the Equal and Hash functions.
func NewProcess(serviceName string, tags []KeyValue) *Process {
	typedTags := KeyValues(tags)
	typedTags.Sort()
	return &Process{ServiceName: serviceName, Tags: typedTags}
}

// Equal compares Process object with another Process.
func (p *Process) Equal(other *Process) bool {
	if p.ServiceName != other.ServiceName {
		return false
	}
	return p.Tags.Equal(other.Tags)
}

// Hash implements Hash from Hashable.
func (p *Process) Hash(w io.Writer) (err error) {
	if _, err := w.Write([]byte(p.ServiceName)); err != nil {
		return err
	}
	return p.Tags.Hash(w)
}

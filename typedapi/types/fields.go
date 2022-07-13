// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// Fields type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/common.ts#L113-L113
type Fields []Field

// FieldsBuilder holds Fields struct and provides a builder API.
type FieldsBuilder struct {
	v Fields
}

// NewFields provides a builder for the Fields struct.
func NewFieldsBuilder() *FieldsBuilder {
	return &FieldsBuilder{}
}

// Build finalize the chain and returns the Fields struct
func (b *FieldsBuilder) Build() Fields {
	return b.v
}

func (b *FieldsBuilder) Fields(value Fields) *FieldsBuilder {
	b.v = value
	return b
}
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

// UnmappedSamplerAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L488-L489
type UnmappedSamplerAggregate struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Meta         *Metadata                   `json:"meta,omitempty"`
}

// UnmappedSamplerAggregateBuilder holds UnmappedSamplerAggregate struct and provides a builder API.
type UnmappedSamplerAggregateBuilder struct {
	v *UnmappedSamplerAggregate
}

// NewUnmappedSamplerAggregate provides a builder for the UnmappedSamplerAggregate struct.
func NewUnmappedSamplerAggregateBuilder() *UnmappedSamplerAggregateBuilder {
	r := UnmappedSamplerAggregateBuilder{
		&UnmappedSamplerAggregate{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the UnmappedSamplerAggregate struct
func (rb *UnmappedSamplerAggregateBuilder) Build() UnmappedSamplerAggregate {
	return *rb.v
}

func (rb *UnmappedSamplerAggregateBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *UnmappedSamplerAggregateBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *UnmappedSamplerAggregateBuilder) DocCount(doccount int64) *UnmappedSamplerAggregateBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *UnmappedSamplerAggregateBuilder) Meta(meta *MetadataBuilder) *UnmappedSamplerAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
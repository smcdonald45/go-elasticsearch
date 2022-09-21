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
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/smcdonald45/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L396>
//
// --------------------------------------------------------------------------------
// PUT _ingest/pipeline/set-foo
// {
//   "description" : "sets foo",
//   "processors" : [ {
//       "set" : {
//         "field": "foo",
//         "value": "bar"
//       }
//   } ]
// }
// POST twitter/_update_by_query?pipeline=set-foo
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_c4b278ba293abd0d02a0b5ad1a99f84a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:c4b278ba293abd0d02a0b5ad1a99f84a[]
	{
		res, err := es.Ingest.PutPipeline(
			"set-foo",
			strings.NewReader(`{
		  "description": "sets foo",
		  "processors": [
		    {
		      "set": {
		        "field": "foo",
		        "value": "bar"
		      }
		    }
		  ]
		}`),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.UpdateByQuery(
			[]string{"twitter"},
			es.UpdateByQuery.WithPipeline("set-foo"),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:c4b278ba293abd0d02a0b5ad1a99f84a[]
}

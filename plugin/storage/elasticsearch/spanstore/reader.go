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

package spanstore

import (
	"time"

	// "github.com/pkg/errors"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"github.com/uber/jaeger/model"
	"github.com/uber/jaeger/storage/spanstore"
	elastic "gopkg.in/olivere/elastic.v5"
)

func NewSpanReader(
	client *elastic.Client,
	metricsFactory metrics.Factory,
	logger *zap.Logger,
) *SpanReader {
	return &SpanReader{}
}

type SpanReader struct {
}

// GetTrace takes a traceID and returns a Trace associated with that traceID
func (s *SpanReader) GetTrace(traceID model.TraceID) (*model.Trace, error) {
	var testingSpan = model.Span{
		TraceID: model.TraceID{
			Low:  1,
			High: 2,
		},
		SpanID: model.SpanID(1),
		Process: &model.Process{
			ServiceName: "serviceName",
			Tags:        model.KeyValues{},
		},
		OperationName: "operationName",
		Tags: model.KeyValues{
			model.String("tagKey", "tagValue"),
		},
		Logs: []model.Log{
			{
				Timestamp: time.Now(),
				Fields: []model.KeyValue{
					model.String("logKey", "logValue"),
				},
			},
		},
		Duration:  time.Second * 5,
		StartTime: time.Unix(300, 0),
	}

	testingTrace := &model.Trace{}
	testingTrace.Spans = append(testingTrace.Spans, &testingSpan)

	return testingTrace, nil
}

// GetServices returns all services traced by Jaeger
func (s *SpanReader) GetServices() ([]string, error) {
	var services = []string{"serviceName"}
	return services, nil
}

// GetOperations returns all operations for a specific service traced by Jaeger
func (s *SpanReader) GetOperations(service string) ([]string, error) {
	var operations = []string{}
	return operations, nil
}

// FindTraces retrieves traces that match the traceQuery
func (s *SpanReader) FindTraces(traceQuery *spanstore.TraceQueryParameters) ([]*model.Trace, error) {
	var testingSpan = model.Span{
		TraceID: model.TraceID{
			Low:  1,
			High: 2,
		},
		SpanID: model.SpanID(1),
		Process: &model.Process{
			ServiceName: "serviceName",
			Tags:        model.KeyValues{},
		},
		OperationName: "operationName",
		Tags: model.KeyValues{
			model.String("tagKey", "tagValue"),
		},
		Logs: []model.Log{
			{
				Timestamp: time.Now(),
				Fields: []model.KeyValue{
					model.String("logKey", "logValue"),
				},
			},
		},
		Duration:  time.Second * 5,
		StartTime: time.Unix(300, 0),
	}

	testingTrace := &model.Trace{}
	testingTrace.Spans = append(testingTrace.Spans, &testingSpan)

	var traces []*model.Trace
	traces = append(traces, testingTrace)

	return traces, nil
}

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
	"context"
	"fmt"

	"github.com/uber/jaeger-lib/metrics"
	"github.com/uber/jaeger/model"
	"go.uber.org/zap"
	elastic "gopkg.in/olivere/elastic.v5"
)

type SpanWriter struct {
	client         *elastic.Client
	metricsFactory metrics.Factory
	logger         *zap.Logger
}

func NewSpanWriter(client *elastic.Client, metricsFactory metrics.Factory, logger *zap.Logger) *SpanWriter {
	return &SpanWriter{
		client:         client,
		metricsFactory: metricsFactory,
		logger:         logger,
	}
}

// WriteSpan saves the span into ElasticSearch
func (s *SpanWriter) WriteSpan(span *model.Span) error {
	_, err := s.client.Index().
		Index("jaeger").
		Type("span").
		BodyJson(span).
		Do(context.TODO())

	return err
}

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

package builder

import (
	"go.uber.org/zap"

	"github.com/uber/jaeger-lib/metrics"
	esDependencyStore "github.com/uber/jaeger/plugin/storage/elasticsearch/dependencystore"
	esSpanStore "github.com/uber/jaeger/plugin/storage/elasticsearch/spanstore"
	"github.com/uber/jaeger/storage/dependencystore"
	"github.com/uber/jaeger/storage/spanstore"
	elastic "gopkg.in/olivere/elastic.v5"
)

type elasticsearchBuilder struct {
	client         *elastic.Client
	logger         *zap.Logger
	metricsFactory metrics.Factory
}

func newElasticSearchBuilder(logger *zap.Logger, metricsFactory metrics.Factory) *elasticsearchBuilder {
	esBuilder := &elasticsearchBuilder{
		logger:         logger,
		metricsFactory: metricsFactory,
	}
	return esBuilder
}

func (es *elasticsearchBuilder) getClient() (*elastic.Client, error) {
	if es.client == nil {
		client, err := elastic.NewSimpleClient()
		es.client = client
		return es.client, err
	}
	return es.client, nil
}

func (es *elasticsearchBuilder) NewSpanReader() (spanstore.Reader, error) {
	client, err := es.getClient()
	if err != nil {
		return nil, err
	}
	return esSpanStore.NewSpanReader(client, es.metricsFactory, es.logger), nil
}

func (es *elasticsearchBuilder) NewDependencyReader() (dependencystore.Reader, error) {
	client, err := es.getClient()
	if err != nil {
		return nil, err
	}
	return esDependencyStore.NewDependencyStore(client, es.metricsFactory, es.logger), nil
}

// Copyright 2016 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package report

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetTimeseries(t *testing.T) {
	sp := newSecondPoints()
	now := time.Now()
	sp.Add(now, time.Second)
	sp.Add(now.Add(5*time.Second), time.Second)
	n := sp.getTimeSeries().Len()
	require.GreaterOrEqualf(t, n, 3, "expected at 6 points of time series, got %s", sp.getTimeSeries())

	// add a point with duplicate timestamp
	sp.Add(now, 3*time.Second)
	ts := sp.getTimeSeries()
	require.Equalf(t, time.Second, ts[0].MinLatency, "ts[0] min latency expected %v, got %s", time.Second, ts[0].MinLatency)
	require.Equalf(t, 2*time.Second, ts[0].AvgLatency, "ts[0] average latency expected %v, got %s", 2*time.Second, ts[0].AvgLatency)
	require.Equalf(t, 3*time.Second, ts[0].MaxLatency, "ts[0] max latency expected %v, got %s", 3*time.Second, ts[0].MaxLatency)
}

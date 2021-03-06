// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package cases

import (
	"github.com/pingcap/kvproto/pkg/metapb"
	"github.com/pingcap/pd/server/core"
)

// Store is the config to simulate tikv.
type Store struct {
	ID           uint64
	Status       metapb.StoreState
	Labels       []metapb.StoreLabel
	Capacity     uint64
	Available    uint64
	LeaderWeight float32
	RegionWeight float32
}

// Region is the config to simulate a region.
type Region struct {
	ID     uint64
	Peers  []*metapb.Peer
	Leader *metapb.Peer
	Size   int64
}

// Conf represents a test suite for simulator.
type Conf struct {
	Stores  []Store
	Regions []Region
	MaxID   uint64

	Checker func(*core.RegionsInfo) bool // To check the schedule is finished.
}

const (
	kb = 1024
	mb = kb * 1024
	gb = mb * 1024
)

type idAllocator struct {
	maxID uint64
}

func (a *idAllocator) nextID() uint64 {
	a.maxID++
	return a.maxID
}

func (a *idAllocator) setMaxID(id uint64) {
	a.maxID = id
}

var confMap = map[string]func() *Conf{
	"balance-leader": newBalanceLeader,
}

// NewConf creates a config to initialize simulator cluster.
func NewConf(name string) *Conf {
	if f, ok := confMap[name]; ok {
		return f()
	}
	return nil
}

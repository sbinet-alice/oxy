// Copyright 2018 The oxy Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simdfmt

// MCTrack holds MonteCarlo tracks processed by the Stack.
//
// An MCTrack can be a primary track put into the simulation or
// a secondary one produced by the transport through decay or interaction.
type MCTrack struct {
	StartVtxMomX float32 // px-component at start vertex [GeV]
	StartVtxMomY float32 // py-component at start vertex [GeV]
	StartVtxMomZ float32 // pz-component at start vertex [GeV]

	StartVtxX float32 // x-coordinate of start vertex [cm, ns]
	StartVtxY float32 // y-coordinate of start vertex [cm, ns]
	StartVtxZ float32 // z-coordinate of start vertex [cm, ns]
	StartVtxT float32 // t-coordinate of start vertex [cm, ns]

	PDGID int32 // PDG particle code

	MotherTrkID int // index of mother track. -1 for primary particles.

	Prop propEncoding // hitmask. if bit i is set, it means this track left a trace in detector i.
}

type propEncoding int32

// Storage returns whether to store this trock to the output
func (p propEncoding) Storage() int {
	panic("not implemented")
}

// ProcessID returns the process that created this track.
func (p propEncoding) ProcessID() int {
	panic("not implemented")
}

// Hitmask returns the hits per detector.
func (p propEncoding) Hitmask() int32 {
	panic("not implemented")
}

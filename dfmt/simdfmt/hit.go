// Copyright 2018 The oxy Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simdfmt

// BaseHit holds the track ID for all AliceO2 hits.
type BaseHit struct {
	TrackID int64
}

// BasicXYZEHit is a basic hit that holds the cartesian position of a hit
// together with the time-of-flight, the energy loss and the detector (or
// sensor) ID that detected the hit.
type BasicXYZEHit struct {
	BaseHit BaseHit
	Pos     Point3D // Pos is the cartesian position of the hit
	ToF     float32 // time of flight
	ELoss   float32 // energy loss
	DetID   int16   // detector/sensor ID
}

type Point3D struct {
	X, Y, Z float32
}

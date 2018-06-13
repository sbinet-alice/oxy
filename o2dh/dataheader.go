// Copyright 2018 The oxy Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package o2dh gathers data header definitions for AliceO2.
package o2dh

import "encoding/binary"

const (
	InvalidToken32 uint32 = 0xFFFFFFFF
	InvalidToken64 uint64 = 0xFFFFFFFFFFFFFFFF

	sizeMagicString               = 4 // size of the magic string field of the DataHeader
	sizeDataOriginString          = 4 // size of the data origin field
	sizeSerializationMethodString = 8 // size of the payload serialization field
	sizeDataDescriptionString     = 16
	sizeHeaderDescriptionString   = 8
)

type HeaderType [sizeDataDescriptionString]byte

func (hdr HeaderType) String() string {
	return string(hdr[:])
}

type SerializationMethod [sizeSerializationMethodString]byte

func (sm SerializationMethod) String() string {
	return string(sm[:])
}

var (
	SerializationMethodAny     = SerializationMethod{'*', '*', '*', '*', '*', '*', '*', '*'}
	SerializationMethodInvalid = SerializationMethod{'I', 'N', 'V', 'A', 'L', 'I', 'D'}
	SerializationMethodNone    = SerializationMethod{'N', 'O', 'N', 'E'}
	SerializationMethodROOT    = SerializationMethod{'R', 'O', 'O', 'T'}
	SerializationMethodFlatBuf = SerializationMethod{'F', 'L', 'A', 'T', 'B', 'U', 'F'}
)

// BaseHeader is the common part for every header.
type BaseHeader struct {
	magic [sizeMagicString]byte // magic string used to identify an O2 header in a raw stream of bytes
	hdrsz uint32                // size of the header that starts with this sequence (base + derived header)
	flags uint32                // flags for sub headers. first bit indicates that a sub header follows.
	hvers uint32                // version of the entire header.
	descr HeaderType            // header type description
	serzm SerializationMethod   // header serialization method
}

// DataOrigin describes the origin (detector or subsystem name) of a datum.
type DataOrigin [sizeDataOriginString]byte

func (do DataOrigin) String() string {
	return string(do[:])
}

type DataDescription [sizeDataDescriptionString]byte

func (descr DataDescription) String() string {
	return string(descr[:])
}

// DataHeader is the main O2 data header structure.
// All O2 messages should have it, preferably at the beginning of the header stack.
//
// DataHeader contains fields that describe the buffer size, data type, origin and
// serialization method used to construct the buffer.
type DataHeader struct {
	Base BaseHeader

	Descr            DataDescription     // data type descriptor
	Origin           DataOrigin          // origin (detector) of the data
	_                uint32              // reserved
	Serialization    SerializationMethod // serialization method
	SubSpecification uint64              // sub specification (e.g. link number)
	PayloadSize      uint64              // size of the associated data
}

// ALICE data origins
var (
	DataOriginAny     = DataOrigin{'*', '*', '*'}
	DataOriginInvalid = DataOrigin{'N', 'I', 'L'}
	DataOriginFLP     = DataOrigin{'F', 'L', 'P'}
	DataOriginACO     = DataOrigin{'A', 'C', 'O'}
	DataOriginCPV     = DataOrigin{'C', 'P', 'V'}
	DataOriginCTP     = DataOrigin{'C', 'T', 'P'}
	DataOriginEMC     = DataOrigin{'E', 'M', 'C'}
	DataOriginFIT     = DataOrigin{'F', 'I', 'T'}
	DataOriginHMP     = DataOrigin{'H', 'M', 'P'}
	DataOriginITS     = DataOrigin{'I', 'T', 'S'}
	DataOriginMCH     = DataOrigin{'M', 'C', 'H'}
	DataOriginMFT     = DataOrigin{'M', 'F', 'T'}
	DataOriginMID     = DataOrigin{'M', 'I', 'D'}
	DataOriginPHS     = DataOrigin{'P', 'H', 'S'}
	DataOriginTOF     = DataOrigin{'T', 'O', 'F'}
	DataOriginTPC     = DataOrigin{'T', 'P', 'C'}
	DataOriginTRD     = DataOrigin{'T', 'R', 'D'}
	DataOriginZDC     = DataOrigin{'Z', 'D', 'C'}
)

// ALICE data types
var (
	DataDescriptionAny           = DataDescription{'*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*'}
	DataDescriptionInvalid       = DataDescription{'I', 'N', 'V', 'A', 'L', 'I', 'D', '_', 'D', 'E', 'S', 'C'}
	DataDescriptionRawData       = DataDescription{'R', 'A', 'W', 'D', 'A', 'T', 'A'}
	DataDescriptionClusters      = DataDescription{'C', 'L', 'U', 'S', 'T', 'E', 'R', 'S'}
	DataDescriptionTracks        = DataDescription{'T', 'R', 'A', 'C', 'K', 'S'}
	DataDescriptionConfig        = DataDescription{'C', 'O', 'N', 'F', 'I', 'G', 'U', 'R', 'A', 'T', 'I', 'O', 'N'}
	DataDescriptionInfo          = DataDescription{'I', 'N', 'F', 'O', 'R', 'M', 'A', 'T', 'I', 'O', 'N'}
	DataDescriptionROOTStreamers = DataDescription{'R', 'O', 'O', 'T', ' ', 'S', 'T', 'R', 'E', 'A', 'M', 'E', 'R', 'S'}
)

// DataIdentifier encodes the origin and description of data.
type DataIdentifier struct {
	Descr  DataDescription
	Origin DataOrigin
}

// NameHeader is a data header containing a name of an object.
type NameHeader struct {
	Base    BaseHeader
	HdrName string
}

func (NameHeader) Version() uint32        { return 1 }
func (NameHeader) HeaderType() HeaderType { return HeaderType{'N', 'a', 'm', 'e', 'H', 'e', 'a', 'd'} }

// HeartbeatHeader is a frame emitted to signal a device is still alive.
type HeartbeatHeader struct {
	Word uint64 // Word is the complete 64b header word, initialized with block-type 1 and size 1
}

type HeartbeatTrailer struct {
	Word uint64
}

// HeartbeatStatistics is a data block gathering statistics about heartbeat frames.
//
// It's transmitted real time as the payload of the HB frame in AliceO2.
type HeartbeatStatistics struct {
	TimeTickNS uint64 // time tick when this statistics was created
	DurationNS uint64 // difference to the previous time tick
}

// RAWDataHeader (RDH) is described in:
//  https://docs.google.com/document/d/1IxCCa1ZRpI3J9j3KCmw2htcOLIRVVdEcO-DDPcLNFM0
//
// RDH consists of 4 64 bit words
//
//         63     56      48      40      32      24      16       8       0
//         |---------------|---------------|---------------|---------------|
//
//   0     | zero  |  size |link id|    FEE id     |  block length | vers  |
//
//   1     |      heartbeat orbit          |       trigger orbit           |
//
//   2     | zero  |heartbeatBC|      trigger type             | trigger BC|
//
//   3     | zero  |      par      | detector field| stop  |  page count   |
//
// Field description:
//
//
//   - version:      the header version number
//   - block length: assumed to be in byte, but discussion not yet finalized
//   - FEE ID:       unique id of the Frontend equipment
//   - Link ID:      id of the link within CRU
//   - header size:  number of 64 bit words
//   - heartbeat and trigger orbit/BC: LHC clock parameters, still under
//                   discussion whether separate fields for HB and trigger
//                   information needed
//   - trigger type: bit fiels fir the trigger type yet to be decided
//   - page count:   incremented if data is bigger than the page size, pages are
//                   incremented starting from 0
//   - stop:         bit 0 of the stop field is set if this is the last page
//   - detector field and par are detector specific fields
//
type RAWDataHeader struct {
	Word0 [8]byte
	Word1 [8]byte
	Word2 [8]byte
	Word3 [8]byte
}

func (rdh *RAWDataHeader) Version() uint8 {
	return rdh.Word0[0]
}

func (rdh *RAWDataHeader) BlockLength() uint16 {
	return binary.LittleEndian.Uint16(rdh.Word0[1:3])
}

func (rdh *RAWDataHeader) FEEID() uint16 {
	return binary.LittleEndian.Uint16(rdh.Word0[3:5])
}

func (rdh *RAWDataHeader) LinkID() uint8 {
	return rdh.Word0[5]
}

func (rdh *RAWDataHeader) HeaderSize() uint8 {
	return rdh.Word0[6]
}

func (rdh *RAWDataHeader) TriggerOrbit() uint32 {
	return binary.LittleEndian.Uint32(rdh.Word1[0:4])
}

func (rdh *RAWDataHeader) HeartbeatOrbit() uint32 {
	return binary.LittleEndian.Uint32(rdh.Word1[4:8])
}

func (rdh *RAWDataHeader) TriggerBCID() uint32 {
	panic("not implemented")
}

func (rdh *RAWDataHeader) TriggerType() uint32 {
	panic("not implemented")
}

func (rdh *RAWDataHeader) HeartbeatBCID() uint32 {
	panic("not implemented")
}

func (rdh *RAWDataHeader) PagesCounter() uint16 {
	return binary.LittleEndian.Uint16(rdh.Word3[0:2])
}

func (rdh *RAWDataHeader) StopCode() uint8 {
	return rdh.Word2[2]
}

func (rdh *RAWDataHeader) DetectorField() uint16 {
	return binary.LittleEndian.Uint16(rdh.Word2[3:5])
}

func (rdh *RAWDataHeader) Par() uint16 {
	return binary.LittleEndian.Uint16(rdh.Word2[5:7])
}

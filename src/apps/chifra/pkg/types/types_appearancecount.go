// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were generated with makeClass --run. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package types

// EXISTING_CODE
import "github.com/ethereum/go-ethereum/common"

// EXISTING_CODE

type RawAppearanceCount struct {
	Address  string `json:"address"`
	NRecords string `json:"nRecords"`
	FileSize string `json:"fileSize"`
}

type SimpleAppearanceCount struct {
	Address  common.Address `json:"address"`
	NRecords uint64         `json:"nRecords"`
	FileSize uint64         `json:"fileSize"`
	raw      *RawAppearanceCount
}

func (s *SimpleAppearanceCount) Raw() *RawAppearanceCount {
	return s.raw
}

func (s *SimpleAppearanceCount) SetRaw(raw *RawAppearanceCount) {
	s.raw = raw
}

func (s *SimpleAppearanceCount) Model(showHidden bool, format string, extraOptions map[string]any) Model {
	// EXISTING_CODE
	// EXISTING_CODE

	model := map[string]interface{}{
		"address":  s.Address,
		"nRecords": s.NRecords,
		"fileSize": s.FileSize,
	}

	order := []string{
		"address",
		"nRecords",
		"fileSize",
	}

	// EXISTING_CODE
	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

// EXISTING_CODE
// EXISTING_CODE

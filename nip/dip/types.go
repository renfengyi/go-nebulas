// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package dip

import (
	"errors"
	"encoding/json"
)

// Error types
var (
	ErrInvalidHeight = errors.New("invalid dip height")
	ErrDipNotFound = errors.New("dip not found")
)

// const types
const (
	CacheSize = 16
	DipDelayRewardHeight = 24*60*60/15

	// DipRewardAddressPrivate dip reward address
	DipRewardAddressPrivate = "42f0c8b5feb72301619046ca87e6cf2c605e94dae0e24c9cb3a0101dbb60337c"
	DipRewardAddressPassphrase = "passphrase"
)

type DIPItem struct {
	Addr  string
	Value string
}

type DIPData struct {
	Start uint64
	End uint64
	Version string
	Data []*DIPItem
}

// ToBytes serialize data
func (d *DIPData) ToBytes() ([]byte, error) {
	return json.Marshal(d)
}

// FromBytes
func (d *DIPData) FromBytes(data []byte) error {
	if err := json.Unmarshal(data, d); err != nil {
		return err
	}
	return nil
}
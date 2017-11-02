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

package nvm

import "C"

import (
	"unsafe"
)

type Storage interface {
	Put(key []byte, val []byte) ([]byte, error)
	Get(key []byte) ([]byte, error)
	Del(key []byte) ([]byte, error)
}

//export StorageGetFunc
func StorageGetFunc(handler unsafe.Pointer, key *C.char) *C.char {
	_, storage := getEngineAndStorage(uint64(uintptr(handler)))
	if storage == nil {
		return nil
	}

	val, err := storage.Get([]byte(C.GoString(key)))
	if err != nil {
		// log.WithFields(log.Fields{
		// 	"func":    "nvm.StorageGetFunc",
		// 	"handler": uint64(uintptr(handler)),
		// 	"key":     C.GoString(key),
		// }).Error("get key failed.")
		return nil
	}
	return C.CString(string(val))
}

//export StoragePutFunc
func StoragePutFunc(handler unsafe.Pointer, key *C.char, value *C.char) int {
	_, storage := getEngineAndStorage(uint64(uintptr(handler)))
	if storage == nil {
		return 1
	}

	storage.Put([]byte(C.GoString(key)), []byte(C.GoString(value)))
	return 0
}

//export StorageDelFunc
func StorageDelFunc(handler unsafe.Pointer, key *C.char) int {
	_, storage := getEngineAndStorage(uint64(uintptr(handler)))
	if storage == nil {
		return 1
	}

	storage.Del([]byte(C.GoString(key)))
	return 0
}

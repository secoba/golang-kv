// Copyright (C) 2022 ucwong
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package kv

import (
	"github.com/secoba/golang-kv/badger"
	"github.com/secoba/golang-kv/bolt"
	"github.com/secoba/golang-kv/leveldb"
	"github.com/secoba/golang-kv/nutsdb"
)

func Badger(path string, opt ...badger.BadgerOption) Bucket {
	return badger.Open(path, opt...)
}

func Bolt(path string, opt ...bolt.BoltOption) Bucket {
	return bolt.Open(path, opt...)
}

func LevelDB(path string, opt ...leveldb.LevelDBOption) Bucket {
	return leveldb.Open(path, opt...)
}

//func Pebble(path string, opt ...pebble.PebbleOption) Bucket {
//	return pebble.Open(path, opt...)
//}

func NutsDB(path string, opt ...nutsdb.NutsdbOption) Bucket {
	return nutsdb.Open(path, opt...)
}

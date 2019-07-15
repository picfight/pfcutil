// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package pfcutil

const (
	// SatoshiPerBitcent is the number of satoshi in one bitcoin cent.
	SatoshiPerBitcent = 1e6

	// SatoshiPerPicfightcoin is the number of satoshi in one bitcoin (1 PFC).
	SatoshiPerPicfightcoin = 1e8

	// MaxSatoshi is the maximum transaction amount allowed in satoshi.
	MaxSatoshi = 21e6 * SatoshiPerPicfightcoin
)

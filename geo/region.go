/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-01-2018
 * |
 * | File Name:     geo/region.go
 * +===============================================
 */

package geo

// Region is a rectangular area corresponding
// to a pair of 1-D elementary intervals.
type Region struct {
	X Range
	Y Range
}

// ElementaryRegion ER_i covers a subset of addresses in the 2D address space
// of (0 ... 2^w - 1, 0 ... 2^w - 1).
type ElementaryRegion struct {
	Regions []Region
}

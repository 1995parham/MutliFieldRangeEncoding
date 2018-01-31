/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-01-2018
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"fmt"
	"sort"

	"github.com/AUTProjects/MutliFieldRangeEncoding/geo"
	"github.com/AUTProjects/MutliFieldRangeEncoding/rule"
)

var rules = []rule.Rule{
	rule.Rule{
		Name: "R0",
		X1:   10,
		X2:   30,
		Y1:   10,
		Y2:   30,
	},
	rule.Rule{
		Name: "R1",
		X1:   20,
		X2:   40,
		Y1:   20,
		Y2:   40,
	},
}

func main() {
	xs := []int{0}
	ys := []int{0}

	for _, r := range rules {
		xs = append(xs, r.X1)
		xs = append(xs, r.X2)

		ys = append(ys, r.Y1)
		ys = append(ys, r.Y2)
	}

	sort.Ints(xs)
	sort.Ints(ys)

	xr := make([]geo.Range, 0)

	for i := 0; i < len(xs)-1; i++ {
		if xs[i] != xs[i+1] {
			xr = append(xr, geo.Range{
				Start: xs[i],
				End:   xs[i+1],
			})
		}
	}

	yr := make([]geo.Range, 0)

	for i := 0; i < len(ys)-1; i++ {
		if ys[i] != ys[i+1] {
			yr = append(yr, geo.Range{
				Start: ys[i],
				End:   ys[i+1],
			})
		}
	}

	regions := make([]geo.Region, 0)

	for i := 0; i < len(xr); i++ {
		for j := 0; j < len(yr); j++ {
			regions = append(regions, geo.Region{
				X: xr[i],
				Y: yr[j],
			})
		}
	}

	ers := make(map[string]geo.ElementaryRegion)

	for _, region := range regions {
		cw := ""
		for _, rule := range rules {
			if region.X.Start >= rule.X1 && region.X.End <= rule.X2 {
				if region.Y.Start >= rule.Y1 && region.Y.End <= rule.Y2 {
					cw += "1"
				} else {
					cw += "0"
				}
			} else {
				cw += "0"
			}
		}
		er, ok := ers[cw]
		if ok {
			er.Regions = append(er.Regions, region)
			ers[cw] = er
		} else {
			ers[cw] = geo.ElementaryRegion{
				Regions: []geo.Region{
					region,
				},
			}
		}
	}
	for cw, er := range ers {
		fmt.Printf("%s -> %d\n", cw, len(er.Regions))
	}

	for i := range rules {
		rs := geo.RegionSet{
			ERs: make([]geo.ElementaryRegion, 0),
		}

		for cw, er := range ers {
			if cw[i] == '1' {
				rs.ERs = append(rs.ERs, er)
			}
		}

		fmt.Println(len(rs.ERs))
	}
}

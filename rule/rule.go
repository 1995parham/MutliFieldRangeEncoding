/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-01-2018
 * |
 * | File Name:     rule/rule.go
 * +===============================================
 */

package rule

// Rule represents two field classification rule
type Rule struct {
	Name   string
	X1     int
	X2     int
	Y1     int
	Y2     int
	Action string
}

// Package mathext provides several math functions.
package mathext

import (
	"math"
)

// Zipf generates zipf distributed variates.
// See: https://en.wikipedia.org/wiki/Zipf%27s_law
//
// Zipf distribution is very common in internet applications.
// E.g., hot data cache follows zipf.
//
// A work load in zipf distributed is usefull to benchmark an internet
// application. Because such a work load is very close to a real world one.
//
// In short, the zipf distribution follows:
//   P(x) = C x**(-s)  or  y = C x**(-s)
// where C is a const, s > 0
//
// Zipf distribution is log-log linear, i.e., ln(y) = ln(C) -s ln(x) .
// In another word, ln(y) is a straight line about ln(x).
//
//    700 +--------------------------------------------------------------------+
//        |      +      +      +      +      +     +      +      +      +      |
//        |      ::                                                    +.....+ |
//    600 |-+    :+                                                          +-|
//        |      ::                                                            |
//        |      :+                                                            |
//    500 |-+    : :                                                         +-|
//        |      : +                                                           |
//        |      :  +                                                          |
//    400 |-+    :  +                                                        +-|
//        |      :   +                                                         |
//    300 |-+   :     +                                                      +-|
//        |     :     ++                                                       |
//        |     :       +                                                      |
//    200 |-+   :        ++                                                  +-|
//        |     :          +++                                                 |
//        |     :            ++++                                              |
//    100 |-+   :                ++++++                                      +-|
//        |     :                     ++++++++++++                             |
//        |     :+      +      +      +      +    +++++++++++++++++++++++++++++|
//      0 +--------------------------------------------------------------------+
//        0      10     20     30     40     50    60     70     80     90    100
//
// The ln(y)-ln(x) graph is:
//
//      7 +----------------------------------------------------------------------+
//        |      +      +      +      +    +  +      +      +      +      +      |
//        |                                :+++                          +.....+ |
//      6 |-+                              :   +++                             +-|
//        |                                :      ++++                           |
//        |                                :          ++++                       |
//      5 |-+                             :              ++++                  +-|
//        |                               :                 ++++                 |
//        |                               :                     ++++             |
//      4 |-+                             :                        ++++        +-|
//        |                               :                           +++++      |
//      3 |-+                             :                               ++   +-|
//        |                               :                                      |
//        |                               :                                      |
//      2 |-+                             :                                    +-|
//        |                               :                                      |
//        |                              :                                       |
//      1 |-+                            :                                     +-|
//        |                              :                                       |
//        |      +      +      +      +  :    +      +      +      +      +      |
//      0 +----------------------------------------------------------------------+
//        0     0.5     1     1.5     2      2.5     3     3.5     4     4.5     5
//
// Implementation
//
// With:
//
//   y = C x**(-s)
//
// Generating a variate so that the density at x is y, is equivilent to generating a
// variate that is evenly distributed in the area under y.
// Thus we distribute `u` evenly on the area under y, i.e., the integral of y.
//
//    t
//    ∫y = 1/(1-s) * (t**(1-s) - a**(1-s))
//    a
//
// And then find the t so that
//
//    t
//    ∫y = u
//    a
//
// If u is evenly distributed, `t` is zipf distributed.
// Let q = (1-s):
//
//   u = c/q(t**q - a**q)
//
// And we want:
//
//    b
//    ∫y = 1
//    a
//
// Thus c = q/(b**q - a**q)
// Finaly t can be found by:
//
//    t = (uq/C+a**q)**(1/q)
//
// Since 0.1.15
type Zipf struct {
	// These are cached intermedium vars
	qInv, aPowQ, c, qDivC float64
}

// NewZipf creates a Zipf struct that generates values in range `[a, b]`, with
// the power `s > 0`.
//
// Usually a is greater than 1, since C x**(-s) is infinite when x get close to
// 0.
//
// Since 0.1.15
func NewZipf(a, b, s float64) *Zipf {
	z := &Zipf{}

	q := (1 - s)
	z.qInv = 1 / q
	z.aPowQ = math.Exp(math.Log(a) * q)
	bPowQ := math.Exp(math.Log(b) * q)

	z.c = q / (bPowQ - z.aPowQ)
	z.qDivC = q / z.c

	return z
}

// Zipf converts an evenly distributed random number `u ∈ [0, 1)`, e.g., a common
// random value,  to a zipf
// distributed variate which is in range `[a, b]`.
//
// Caution: because of the inaccuracy of float number, the output value may be a
// little bit lower than a or greater than b. I donot test it yet in this
// version.:D
//
// It costs 80 ns per calls, a little faster than the official rand.Zipf(83 ns).
//
// Since 0.1.15
func (z *Zipf) Float64(u float64) float64 {
	ln := math.Log(u*z.qDivC + z.aPowQ)
	t := math.Exp(ln * z.qInv)
	return t
}

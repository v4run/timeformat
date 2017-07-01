/*
Package timeformat defines an easy to remember custom layout to format the
time.Time object.

Layout
	YYYY:  2006
	YY:    06
	MMMM:  January
	MMM:   Jan
	MM:    01
	M:     1
	DD:    02
	D:     2
	hh:    03
	HH:    15
	h:     3
	wwww:  Monday
	www:   Mon
	mm:    04
	m:     4
	ss:    05
	s:     5
	f:     0         .fff = .000
	F:     9         .FFF = .999
	a:     pm
	A:     PM
	z:     MST
	-Z:Z:Z -07:00:00
	Z:Z:Z  Z07:00:00
	-Z:Z   -07:00
	Z:Z    Z07:00
	-ZZZ   -070000
	ZZZ    Z070000
	-ZZ    -0700
	ZZ     Z0700
	-Z     -07
	Z      Z07

	For example:
	DD/MM/YYYY = 02/01/2006
	hh:mm a    = 03:04 pm
*/
package timeformat

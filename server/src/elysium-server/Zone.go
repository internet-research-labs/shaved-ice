package main

type Point struct {
	x int
	y int
}

type Hexagon struct {
	nw Point
	ne Point
	e  Point
	se Point
	sw Point
	w  Point

	//    NW-----NE
	//    /       \
	//   W         E
	//    \       /
	//    SW-----SE
}

type Zone struct {
	clients chan Client
	actions chan Action

	// Adjacent zones
	north     *Zone
	northeast *Zone
	southeast *Zone
	south     *Zone
	southwest *Zone
	northwest *Zone

	// Boundaries
	boundaries Hexagon
	//     0-----1
	//    /       \
	//   5         2
	//    \       /
	//     4-----3
}

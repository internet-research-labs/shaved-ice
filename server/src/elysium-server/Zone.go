package main

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
}

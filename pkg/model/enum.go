package model

// Region represents a global region in League of Legends esports.
type Region string

const (
	// TODO Merge Americas and Pacific?
	Brazil       Region = "Brazil"
	China        Region = "China"
	Europe       Region = "Europe"
	Korea        Region = "Korea"
	NorthAmerica Region = "North America"
	Pacific      Region = "Pacific"
	Vietnam      Region = "Vietnam"
)

// Regions is a slice of all regions in League of Legends esports.
var Regions = []Region{
	Brazil,
	China,
	Europe,
	Korea,
	NorthAmerica,
	Pacific,
	Vietnam,
}

// Role represents a position in League of Legends.
type Role string

const (
	Top     Role = "Top"
	Jungle  Role = "Jungle"
	Middle  Role = "Middle"
	Bottom  Role = "Bottom"
	Support Role = "Support"
)

// Roles is a slice of all positions in League of Legends.
var Roles = []Role{
	Top,
	Jungle,
	Middle,
	Bottom,
	Support,
}

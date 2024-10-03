package model

type Region string

const (
	Brazil       Region = "Brazil"
	China        Region = "China"
	Europe       Region = "Europe"
	Korea        Region = "Korea"
	NorthAmerica Region = "North America"
	Pacific      Region = "Pacific"
	Vietnam      Region = "Vietnam"
)

var Regions = []Region{
	Brazil,
	China,
	Europe,
	Korea,
	NorthAmerica,
	Pacific,
	Vietnam,
}

type Role string

const (
	Top     Role = "Top"
	Jungle  Role = "Jungle"
	Middle  Role = "Middle"
	Bottom  Role = "Bottom"
	Support Role = "Support"
)

var Roles = []Role{
	Top,
	Jungle,
	Middle,
	Bottom,
	Support,
}

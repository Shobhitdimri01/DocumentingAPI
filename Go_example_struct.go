package main

type AutoGenerated struct {
	SquadName  string
	HomeTown   string
	Formed     int
	SecretBase string
	Active     bool
	Members    []struct {
		Name           string
		Age            int
		SecretIdentity string
		Powers         []string
	}
}
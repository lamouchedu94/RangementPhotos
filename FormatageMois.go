package main

import (
	"strconv"
)

func FormatageMois(TabDate []string) []string {
	TabMois := [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"}
	mois := TabDate[1]
	if len(mois) == 2 {
		mois = string(mois[1])
	}
	moisInt, _ := strconv.Atoi(mois)
	TabDate[1] = TabMois[moisInt-1]

	return TabDate
}

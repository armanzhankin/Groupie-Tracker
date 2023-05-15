package cmd

import "strings"

func FormatMarks(i int) {
	res := make(map[string][]string)
	for key := range PeopleNumArtists[i].DatesLocations {
		res[strings.Title(strings.ReplaceAll(strings.ReplaceAll(key, "_", " "), "-", ", "))] = PeopleNumArtists[i].DatesLocations[key]
		delete(PeopleNumArtists[i].DatesLocations, key)
	}
	PeopleNumArtists[i].DatesLocations = res
}

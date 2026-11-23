package vreme

import (
	"fmt"
	"sort"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, ok := postaje[oznaka]; ok {
		fmt.Println("Postaja z oznako " + oznaka + " že obstaja.")
		return
	}
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	postaja, ok := postaje[oznaka]
	if !ok {
		fmt.Println("Postaja z oznako " + oznaka + " ne obstaja.")
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Meritev presega dovoljene meje za postajo " + oznaka + ".")
		return
	}
	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, exists := postaje[oznaka]
	if !exists || len(postaja.meritve) == 0 || len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}
	najnizja, najvisja = postaja.meritve[0], postaja.meritve[0]
	for _, meritev := range postaja.meritve {
		najnizja = min(najnizja, meritev)
		najvisja = max(najvisja, meritev)
		povprecna += meritev
	}
	povprecna /= float64(len(postaja.meritve))
	return najnizja, najvisja, povprecna, true
}

func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println()
	fmt.Println("MERITVE:")
	for _, oznaka := range urejeneOznake(postaje) {
		postaja := postaje[oznaka]
		fmt.Println(oznaka, "- "+postaja.ime, "("+postaja.kraj+"):", postaja.meritve)
	}
}

func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	fmt.Println()
	for _, oznaka := range urejeneOznake(postaje) {
		postaja := postaje[oznaka]
		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			fmt.Println(postaja.ime + ": premalo meritev")
			continue
		}
		var ocena string
		switch {
		case povprecna >= 25:
			ocena = "Vroče"
		case povprecna >= 15:
			ocena = "Toplo"
		case povprecna >= 5:
			ocena = "Hladno"
		default:
			ocena = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			postaja.ime, najnizja, najvisja, povprecna, ocena)
	}
}

func urejeneOznake(postaje map[string]Postaja) []string {
	oznake := make([]string, 0, len(postaje))
	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}
	sort.Strings(oznake)
	return oznake
}

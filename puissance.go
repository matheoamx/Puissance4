package main

import (
)

const LIGNES = 6
const COLONNES = 7

func initialiserGrille() [LIGNES][COLONNES]string {
	var grille [LIGNES][COLONNES]string
	for i := 0; i < LIGNES; i++ {
		for j := 0; j < COLONNES; j++ {
			grille[i][j] = "."
		}
	}
	return grille
}


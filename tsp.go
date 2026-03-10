package main

import (
	"fmt"
	"math"
)

// ---------------------------------------------------------------------------
// 1. DATI: pašvaldības un to koordinātas
// ---------------------------------------------------------------------------

type Municipality struct {
	name string
	lat  float64
	lon  float64
}

var municipalities = []Municipality{
	{"Rīgas valstspilsētas pašvaldība (Rātslaukums)", 56.9475, 24.1064},
	{"Aizkraukles novada pašvaldība", 56.6504, 25.2530},
	{"Alūksnes novada pašvaldība", 57.4216, 27.0428},
	{"Augšdaugavas novada pašvaldība", 55.8753, 26.5354},
	{"Ādažu novada pašvaldība", 57.0801, 24.3293},
	{"Balvu novada pašvaldība", 57.1314, 27.2651},
	{"Bauskas novada pašvaldība", 56.4079, 24.1944},
	{"Cēsu novada pašvaldība", 57.3119, 25.2716},
	{"Daugavpils valstspilsētas pašvaldība", 55.8741, 26.5361},
	{"Dienvidkurzemes novada pašvaldība", 56.5414, 21.1667},
	{"Dobeles novada pašvaldība", 56.6217, 23.2728},
	{"Gulbenes novada pašvaldība", 57.1731, 26.7538},
	{"Jelgavas novada pašvaldība", 56.6511, 23.7203},
	{"Jelgavas valstspilsētas pašvaldība", 56.6516, 23.7202},
	{"Jēkabpils novada pašvaldība", 56.4955, 25.8709},
	{"Jūrmalas valstspilsētas pašvaldība", 56.9677, 23.7708},
	{"Krāslavas novada pašvaldība", 55.8946, 27.1652},
	{"Kuldīgas novada pašvaldība", 56.9680, 21.9714},
	{"Ķekavas novada pašvaldība", 56.8286, 24.2327},
	{"Liepājas valstspilsētas pašvaldība", 56.5047, 21.0108},
	{"Limbažu novada pašvaldība", 57.5133, 24.7170},
	{"Līvānu novada pašvaldība", 56.3533, 26.1750},
	{"Ludzas novada pašvaldība", 56.5480, 27.7161},
	{"Madonas novada pašvaldība", 56.8504, 26.2206},
	{"Mārupes novada pašvaldība", 56.9070, 24.0519},
	{"Ogres novada pašvaldība", 56.8164, 24.6064},
	{"Olaines novada pašvaldība", 56.7844, 23.9406},
	{"Preiļu novada pašvaldība", 56.2917, 26.7231},
	{"Rēzeknes novada pašvaldība", 56.5074, 27.3277},
	{"Rēzeknes valstspilsētas pašvaldība", 56.5070, 27.3271},
	{"Ropažu novada pašvaldība", 56.9740, 24.4140},
	{"Salaspils novada pašvaldība", 56.8600, 24.3480},
	{"Saldus novada pašvaldība", 56.6667, 22.4936},
	{"Saulkrastu novada pašvaldība", 57.2608, 24.4128},
	{"Siguldas novada pašvaldība", 57.1537, 24.8537},
	{"Smiltenes novada pašvaldība", 57.4244, 25.9014},
	{"Talsu novada pašvaldība", 57.2458, 22.5894},
	{"Tukuma novada pašvaldība", 56.9667, 23.1556},
	{"Valkas novada pašvaldība", 57.7753, 26.0101},
	{"Valmieras novada pašvaldība", 57.5333, 25.4167},
	{"Ventspils novada pašvaldība", 57.2925, 20.5791},
	{"Ventspils valstspilsētas pašvaldība", 57.3940, 21.5541},
}

// ----------------------------------------------------------------
// Haversine — attālums starp diviem punktiem uz Zemes (km)
// ----------------------------------------------------------------
func haversine(a, b Municipality) float64 {
	const R = 6371
	toRad := func(deg float64) float64 { return deg * math.Pi / 180 }

	dLat := toRad(b.lat - a.lat)
	dLon := toRad(b.lon - a.lon)

	sinLat := math.Sin(dLat / 2)
	sinLon := math.Sin(dLon / 2)

	c := sinLat*sinLat +
		math.Cos(toRad(a.lat))*math.Cos(toRad(b.lat))*sinLon*sinLon

	return 2 * R * math.Asin(math.Sqrt(c))
}

// ----------------------------------------------------------------
// Distance matrix — iepriekš aprēķināts veiktspējai
// ----------------------------------------------------------------
func buildDistMatrix(m []Municipality) [][]float64 {
	n := len(m)
	dist := make([][]float64, n)
	for i := range dist {
		dist[i] = make([]float64, n)
		for j := range dist[i] {
			dist[i][j] = haversine(m[i], m[j])
		}
	}
	return dist
}

// ----------------------------------------------------------------
// Nearest Neighbor — alkatīga sākotnējā tūre, sākot no `start`
// ----------------------------------------------------------------
func nearestNeighbor(dist [][]float64, start int) []int {
	n := len(dist)
	visited := make([]bool, n)
	tour := make([]int, 0, n)

	current := start
	visited[current] = true
	tour = append(tour, current)

	for len(tour) < n {
		bestDist := math.Inf(1)
		bestNext := -1

		for j := 0; j < n; j++ {
			if !visited[j] && dist[current][j] < bestDist {
				bestDist = dist[current][j]
				bestNext = j
			}
		}

		visited[bestNext] = true
		tour = append(tour, bestNext)
		current = bestNext
	}

	return tour
}

// ----------------------------------------------------------------
// Tour length — ieskaitot atgriešanos sākuma stāvoklī
// ----------------------------------------------------------------
func tourLength(tour []int, dist [][]float64) float64 {
	total := 0.0
	n := len(tour)
	for i := 0; i < n; i++ {
		total += dist[tour[i]][tour[(i+1)%n]]
	}
	return total
}

// ----------------------------------------------------------------
// 2-opt — uzlabot ekskursiju, mainot segmentus
// ----------------------------------------------------------------
func twoOpt(tour []int, dist [][]float64) []int {
	n := len(tour)
	best := make([]int, n)
	copy(best, tour)
	bestLen := tourLength(best, dist)

	improved := true
	for improved {
		improved = false
		for i := 0; i < n-1; i++ {
			for j := i + 2; j < n; j++ {
				if i == 0 && j == n-1 {
					continue
				}

				// Aprēķiniet pastiprinājumu no atpakaļgaitas segmenta [i+1..j]
				delta := -dist[best[i]][best[i+1]] -
					dist[best[j]][best[(j+1)%n]] +
					dist[best[i]][best[j]] +
					dist[best[i+1]][best[(j+1)%n]]

				if delta < -1e-10 {
					// Apgrieztais segments starp i+1 un j
					for left, right := i+1, j; left < right; left, right = left+1, right-1 {
						best[left], best[right] = best[right], best[left]
					}
					bestLen += delta
					improved = true
				}
			}
		}
	}

	_ = bestLen
	return best
}

// ----------------------------------------------------------------
// Normalizēt ceļojumu tā, lai Rīga (indekss 0) būtu pirmajā vietā
// ----------------------------------------------------------------
func normalizeToRiga(tour []int) []int {
	n := len(tour)
	rigaPos := 0
	for i, v := range tour {
		if v == 0 {
			rigaPos = i
			break
		}
	}

	normalized := make([]int, n)
	for i := 0; i < n; i++ {
		normalized[i] = tour[(rigaPos+i)%n]
	}
	return normalized
}

// ----------------------------------------------------------------
// Atrisināt — vairāku startu NN + 2-opcijas, atgriezt labāko maršrutu
// ----------------------------------------------------------------
func solve(dist [][]float64) ([]int, float64) {
	n := len(dist)
	var bestTour []int
	bestLen := math.Inf(1)

	for start := 0; start < n; start++ {
		tour := nearestNeighbor(dist, start)
		tour = normalizeToRiga(tour)
		tour = twoOpt(tour, dist)
		length := tourLength(tour, dist)

		if length < bestLen {
			bestLen = length
			bestTour = tour
		}
	}

	return bestTour, bestLen
}

// ----------------------------------------------------------------
// Main
// ----------------------------------------------------------------
func main() {
	n := len(municipalities)
	dist := buildDistMatrix(municipalities)

	fmt.Println("=================================================================")
	fmt.Println(" TSP — Īsākais maršruts pa Latvijas pašvaldībām")
	fmt.Println("=================================================================\n")
	fmt.Printf("Kopējais pašvaldību skaits (ieskaitot Rīgu): %d \n\n", n)
	fmt.Println("Aprēķina...\n")

	tour, length := solve(dist)

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println(" MARŠRUTS (secīgi):")
	fmt.Println("-----------------------------------------------------------------")

	fmt.Printf(" 🏁 SĀKUMS:  %s\n", municipalities[tour[0]].name)

	for step := 1; step < len(tour); step++ {
		fmt.Printf("  %2d. %s\n", step, municipalities[tour[step]].name)
	}

	fmt.Printf("\n 🏁 ATGRIEŠANĀS: %s\n", municipalities[tour[0]].name)

	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf(" KOPĒJAIS MARŠRUTA GARUMS: %.1f km\n", length)
	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("\n=================================================================")
}
package solutions

var Solutions = map[int]func() error{}

func init() {
	Solutions[1] = Day1
	Solutions[2] = Day2
	Solutions[3] = Day3Diagnostics
	Solutions[4] = Day4Bingo
	Solutions[5] = Day5Vents
	Solutions[6] = Day6Lanterns
	Solutions[7] = Day7Crabs
	Solutions[8] = Day8SevenSegments
	Solutions[9] = Day9Smoke
	Solutions[10] = Day10SyntaxScoring
	Solutions[11] = Day11DumboOcto
	Solutions[12] = Day12CavesPathing
	Solutions[13] = Day13ThermalDrm
	Solutions[14] = Day14Polymers
	Solutions[15] = Day15Chitons
	Solutions[16] = Day16PacketDecoder
	Solutions[17] = Day17ProbeLauncher
	Solutions[18] = Day18SnailNumbers
	Solutions[19] = Day19Beacons
	Solutions[20] = Day20TrenchMap
	Solutions[21] = Day21DiceGame
}

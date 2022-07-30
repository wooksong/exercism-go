package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank string) int {
	var cnt int = 0

	for k, v := range cb {
		if k == rank {
			for _, r := range v {
				if r {
					cnt++
				}
			}
		}
	}

	return cnt
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file int) int {
	var cnt int = 0

	for _, v := range cb {
		for i, s := range v {
			if (i+1) == file && s {
				cnt++
			}
		}
	}
	return cnt
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var cnt int = 0

	for _, eachRank := range cb {
		for range eachRank {
			cnt++
		}
	}

	return cnt
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var cnt int = 0

	for _, eachRank := range cb {
		for _, v := range eachRank {
			if v {
				cnt++
			}
		}
	}

	return cnt
}

package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	fileCount := 0

	currentFile, exists := cb[file]

	if !exists {
		return 0
	}

	for _, square := range currentFile {
		if !square {
			continue
		}

		fileCount++
	}

	return fileCount
}

func CountInRank(cb Chessboard, rank int) int {
	rankCount := 0
	rankIndex := rank - 1

	if rankIndex < 0 {
		return 0
	}

	for _, file := range cb {
		if len(file) < rank {
			return 0
		}

		if !file[rankIndex] {
			continue
		}

		rankCount++
	}

	return rankCount
}

func CountAll(cb Chessboard) int {
	fileLength := len(cb)
	rankLength := 0

	for _, file := range cb {
		rankLength = len(file)
		break
	}

	return fileLength * rankLength
}

func CountOccupied(cb Chessboard) int {
	occupiedCount := 0

	for idx := range cb {
		occupiedCount += CountInFile(cb, idx)
	}

	return occupiedCount
}

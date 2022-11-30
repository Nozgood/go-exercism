package main

import "fmt"

// Type declaration
type File []bool
type ChessBoard map[string]File

func CountInFile(cb ChessBoard, file string) int {
	count := 0

	_, exists := cb[file]
    if !exists {
        return 0
    }
for index, verify := range cb[file] {
	fmt.Printf("%v \n %v \n", index, verify)
	if verify {
		count ++
	}
}
return count
}

func CountInRank(cb ChessBoard, rank int) int {
	count := 0
	if rank < 1 || rank > 8 {
		return 0 
	} 
		for _, value := range cb {
			if value[rank -1] {
				count ++
			}
		}
		return count
}

func CountAll(cb ChessBoard) int {
	count := 0
	for _, value := range cb {
		for i := 0; i < len(value); i++ {
			if value[i] {
				count ++
			}
		}
	}
	return count
}

///////////////////////////////////////////////////////////////////////////////////////////////

type ElectionResult struct {
	Name 	string
	votes 	int
}

func NewVoteCounter(initialVotes int) *int {
	votesPointer := &initialVotes
	return votesPointer
}

func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	vote := *counter
	return vote
}

func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

func NewElectionResult(candidateName string, votes int) *ElectionResult {
	result := ElectionResult{
		Name: candidateName,
		votes: votes,
	}
	return result
}

func main () {
	candidateName := "noz"
	votes := 0
	result := ElectionResult{
		Name: candidateName,
		votes: votes,
	}
	fmt.Print(&result)
}
package main

import (
	"fmt"
	"strconv"
)

type ElectionResult struct {
	Name 	string
	Votes 	int
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
		Votes: votes,
	}
	return &result
}

func DisplayResult(result *ElectionResult) string {
	stringVote := strconv.Itoa(result.Votes)
	stringVoteParenthese := fmt.Sprintf("(%v)", stringVote)
	stringResult := result.Name + " " + stringVoteParenthese
	return stringResult
}

func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
}

func main () {
	candidateName := "noz"
	votes := 0
	result := ElectionResult{
		Name: candidateName,
		Votes: votes,
	}
	fmt.Print(result)
	fmt.Print(DisplayResult(&result))
}
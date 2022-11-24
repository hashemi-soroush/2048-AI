package main

import "fmt"
import "time"
import "math/rand"
import "math"
import "reflect"

func main() {
	rand.Seed(time.Now().UnixNano())

	testCounts := 1000

	randomAIScores := test(testCounts, GenerateRandomAI)
	greedyAIScores := test(testCounts, GenerateGreedyAI)
	smartAIScores := test(testCounts, GenerateSmartAI)

	fmt.Println(randomAIScores)
	fmt.Println(greedyAIScores)
	fmt.Println(smartAIScores)

	mean, variance := GiveMeanAndVar(randomAIScores)
	fmt.Println(fmt.Sprintf("randomAI \t -> mean: %f \t var: %f", mean, variance))
	mean, variance = GiveMeanAndVar(greedyAIScores)
	fmt.Println(fmt.Sprintf("greedyAI \t -> mean: %f \t var: %f", mean, variance))
	mean, variance = GiveMeanAndVar(smartAIScores)
	fmt.Println(fmt.Sprintf("smartAI \t -> mean: %f \t var: %f", mean, variance))
}

func GenerateRandomAI() Player {
	return &RandomAI{}
}

func GenerateGreedyAI() Player {
	return &GreedyAI{}
}

func GenerateSmartAI() Player {
	return &SmartAI{}
}

func test(count int, playerGenerator func() Player) []int {
	scores := *new([]int)
	for round := 0; round < count; round++ {

		g := Game{}
		p := playerGenerator()

		fmt.Println(fmt.Sprintf("player %s, test round %d", reflect.TypeOf(p), round))

		g.InitiateBoard()
		score := g.RunGame(p)
		scores = append(scores, score)
	}
	return scores
}

func GiveMeanAndVar(nums []int) (float64, float64) {
	mean := 0.0
	for _, num := range nums {
		mean += float64(num)
	}
	mean /= float64(len(nums))

	variance := 0.0
	for _, num := range nums {
		variance += math.Pow((float64(num) - mean), 2)
	}
	variance /= float64(len(nums) - 1)
	variance = math.Pow(variance, 0.5)

	return mean, variance
}

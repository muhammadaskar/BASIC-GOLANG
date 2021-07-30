package main

import "fmt"

func main()  {
	// students := []map[string]string{
	// 	{
	// 		"name": "agung",
	// 		"score": "B",
	// 	},
	// 	{
	// 		"name": "mario",
	// 		"score": "A",
	// 	},
	// }

	// for _, student := range students{
	// 	fmt.Println(student["name"], " - ", student["score"])
	// }

	// QUIZ
	// Hitung rata-rata
	scores := [8]int{
		100, 80, 75, 92, 70, 93, 88, 67,
	}

	sum := 0

	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	var avg float32
	avg = (float32(sum)/float32(len(scores)))

	fmt.Println(avg)
	
	// Soal ke 2
	
	var goodScores[] int
	for _, score := range scores{
		if (score >= 90){
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(goodScores)
}
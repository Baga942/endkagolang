package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"
)

func main(){
	start := time.Now()
	WordCounter()
	fmt.Printf("Process took %s\n", time.Since(start))
}

func WordCounter() {

	data, err := ioutil.ReadFile("C:\\Users\\Бакдаулет\\Desktop\\endka\\mobydick.txt")

	if err != nil {
		fmt.Println(err)
	}
	var word []byte
	var sortedSlice [][]byte

	size := len(data)


	for i := 0; i < size-1; i++ {
		if data[i] >= 97 && data[i] <= 122 {
			word = append(word, data[i])
			continue
		}
		 if data[i] >= 65 && data[i] <= 90 {
			word = append(word, data[i] + 32)
			continue
		}
		if len(word) > 0 {
			sortedSlice = append(sortedSlice, word)
		}
		word = []byte{}
	 }
	size = len(sortedSlice)

	var usedWords [][]byte
	var occurrenceSlice []uint
	var index int


	usedWords = append(usedWords, sortedSlice[0])
	occurrenceSlice = append(occurrenceSlice, 1)

	for i := 1; i< size; i++{
			index = isUsed(&usedWords, &sortedSlice[i])
			if index == -1{
				usedWords = append(usedWords, sortedSlice[i])
				occurrenceSlice = append(occurrenceSlice, 1)
			}else {
				occurrenceSlice[index] += 1
			}
	}

	size = len(occurrenceSlice)



	sort(&usedWords, &occurrenceSlice, 0, size - 1)


	for i := 0; i < 25; i++ {
		print(string(usedWords[size-i-1]) + " ")
		println(occurrenceSlice[size-i-1])
	}
}

func sort(cache *[][] byte, occurrences *[]uint, first int, last int)  {
	left, right := first, last
	pivot := (*occurrences)[(left+right)/2]
	for left <= right {
		for (*occurrences)[left] < pivot {
			left++
		}
		for (*occurrences)[right] > pivot {
			right--
		}
		if left <= right {
			(*occurrences)[left], (*occurrences)[right] = (*occurrences)[right], (*occurrences)[left]
			(*cache)[left], (*cache)[right] = (*cache)[right], (*cache)[left]
			left++
			right--
		}
	}
	if first < right {
		sort(cache, occurrences, first, right)
	}
	if left < last {
		sort(cache, occurrences, left, last)
	}
}

func isUsed(arr *[][]byte, word *[]byte) int{
	for i := 0; i < len(*arr); i++ {
		if bytes.Equal((*arr)[i], *word) == true {
			return i
		}
	}
	return -1
}


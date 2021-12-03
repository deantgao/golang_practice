package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func logPlay(str string) {
	log.Println(str)
}

type animal_acts interface {
	speak(words string)
	eat()
}

type animal struct {
	can_speak    bool
	hunger_level int
}

func (a *animal) speak(words string) {
	a.can_speak = true
	fmt.Println(words)
}

func (a *animal) eat() {
	fmt.Println("Before eating: ", a.hunger_level)
	a.hunger_level += 1
	fmt.Println("After eating: ", a.hunger_level)
}

func makeAnimal(ani animal_acts, words string) {
	ani.speak(words)
	ani.eat()
}

func notmain() {
	new_animal := &animal{can_speak: false, hunger_level: 0}
	fmt.Printf("Newborn animal can_speak: %v, hunger_level: %v \n", new_animal.can_speak, new_animal.hunger_level)
	makeAnimal(new_animal, "Say something I'm giving up on you")
	makeAnimal(new_animal, "We just vibin")
	makeAnimal(new_animal, "Where are we going?")
	makeAnimal(new_animal, "I don't know any more")
	fmt.Printf("After it's all said and done, can_speak: %v, hunger_level: %v \n", new_animal.can_speak, new_animal.hunger_level)
	logPlay("What it do?")
}

// longestContiguous should take two arrays of strings and find the longest contiguous sequence that appears in both
func longestContiguous(user1, user2 []string) []string {
	var short, long []string
	if len(user1) > len(user2) {
		long = user1
		short = user2
	} else {
		short = user1
		long = user2
	}
	maxRun := 0
	result := []string{}
	for i := 0; i < len(short); i++ {
		ptr := 0
		lenRun := 0
		for ptr < len(long) {
			if short[i] == long[ptr] {
				lenRun += 1
				ptr += 1
			} else {
				diff := float64(ptr - i)
				if int(math.Abs(diff)) > maxRun {
					maxRun = int(math.Abs(diff))
					result = short[i : i+maxRun]
				}
				break
			}
		}
	}
	return result
}

func readFileToNums(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	result := []int{}
	for scanner.Scan() { // internally, it advances token based on sperator
		num, _ := strconv.Atoi(scanner.Text())
		result = append(result, num)
	}
	return result
}

func findNumIncrease(nums []int) int {
	numIncrease := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			numIncrease += 1
		}
	}
	return numIncrease
}

func findLargestWindowOfThree(nums []int) int {
	counter := 0
	totalWindows := 1
	previousSum := nums[0] + nums[1] + nums[2]
	for i, num := range nums[1 : len(nums)-2] {
		totalWindows += 1
		sum := num + nums[i+2] + nums[i+3]
		if sum > previousSum {
			counter += 1
		}
		previousSum = sum
	}
	fmt.Printf("the total num of windows is %v, the total num of sum increases is %v\n", totalWindows, counter)
	return counter
}

func readFileToStrings(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() { // internally, it advances token based on sperator
		move := scanner.Text()
		result = append(result, move)
	}
	return result
}

func parseMoves(moves []string, aim bool) (hor, ver int) {
	var aimVal int
	var x, y int
	for _, move := range moves {
		elems := strings.Split(move, " ")
		n, _ := strconv.Atoi(elems[1])
		fmt.Println(elems[1])
		if elems[0] == "forward" {
			if aim {
				y += aimVal * n
			}
			x += n
		}
		if elems[0] == "backward" {
			x -= n
		}
		if elems[0] == "up" {
			if aim {
				aimVal -= n
			} else {
				y -= n
			}

		}
		if elems[0] == "down" {
			if aim {
				aimVal += n
			} else {
				y += n
			}
		}
	}
	return x, y
}

func parseBinary(bins []string) (string, string) {
	rows := len(bins)
	cols := len(bins[0])
	fmt.Printf("num of cols is %v\n", cols)
	gam, eps := "", ""
	for i := 0; i < cols; i++ {
		numOnes, numZeros := 0, 0
		for j := 0; j < rows; j++ {
			if string(bins[j][i]) == "0" {
				numZeros += 1
			} else {
				numOnes += 1
			}
		}
		if numOnes > numZeros {
			gam += "1"
			eps += "0"
		} else {
			gam += "0"
			eps += "1"
		}
	}
	return gam, eps
}

func binaryToDecimal(bin string) int {
	var decimal float64
	power := len(bin) - 1
	for _, dig := range bin {
		strDig, _ := strconv.Atoi(string(dig))
		decimal += float64(strDig) * math.Pow(float64(2), float64(power))
		power -= 1
	}
	return int(decimal)
}

func main() {
	// filePath := "/Users/deangao/Desktop/input.txt"
	// nums := readDay1File(filePath)
	// fmt.Printf("the total number of values is %v\n", len(nums))
	// numIncrease := findNumIncrease(nums)
	// fmt.Printf("the total number of increasing values is %v\n", numIncrease)
	// totalIncreaseWindows := findLargestWindowOfThree(nums)
	// fmt.Println("the total number of increasing windows is ", totalIncreaseWindows) // 1393 is too low

	// filePath := "/Users/deangao/Desktop/day2.txt"
	// moves := readDay2File(filePath)
	// moves := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	// x, y := parseMoves(moves, true)
	// res := x * y
	// fmt.Printf("the final position for x is %v and y is %v\n", x, y)
	// fmt.Printf("the product of the x and y position is %v\n", res)

	filePath := "/Users/deangao/Desktop/day3.txt"
	bins := readFileToStrings(filePath)
	// bins := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	gamma, epsilon := parseBinary(bins)
	fmt.Printf("the newly constructed gamma is %v and epsilon is %v\n", gamma, epsilon)
	decGamma := binaryToDecimal(gamma)
	decEpsilon := binaryToDecimal(epsilon)
	fmt.Printf("the decimal value for gamma is %v and epsilon %v\n", decGamma, decEpsilon)
	fmt.Printf("the product of gamme and epsilon is %v\n", decGamma*decEpsilon)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Выберите операцию")
	fmt.Println("AVG - среднее")
	fmt.Println("SUM - сумму")
	fmt.Println("MED - медиану")
	var operation, input string

	_, er := fmt.Scanln(&operation)
	if er != nil {
		fmt.Println("Ошибка ввода", er)
		return
	}

	if !(operation == "AVG" || operation == "SUM" || operation == "MED") {
		fmt.Println("Доступны только операции AVG,SUN,MED")
		return
	}

	fmt.Println("Введите последовательночть чисел разделенных запятой с пробелом")
	// исправил ввод
	reader := bufio.NewReader(os.Stdin)
	input, er = reader.ReadString('\n')
	if er != nil {
		fmt.Println("Ошибка ввода", er)
	}

	// Разбиваем строку по разделителю ", "
	parts := strings.Split(input, ", ")

	// Создаем слайс для чисел
	numbers := make([]int, 0, len(parts))

	// Преобразуем каждую часть в число
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			fmt.Printf("Ошибка: '%s' не является числом\n", part)
			continue
		}
		numbers = append(numbers, num)
	}
	fmt.Println("Упорядочено", numbers)
	st := RunOperation(&numbers, operation)
	fmt.Printf("Значение %s по входным данным будет: %.2f", operation, st)
}

func sumArrInt(numbers *[]int) int {

	var total int
	for _, v := range *numbers {
		total += v
	}
	return total
}

func RunOperation(numbers *[]int, operation string) float64 {

	switch operation {
	case "AVG":
		return float64(sumArrInt(numbers)) / float64(len(*numbers))
	case "SUM":
		return float64(sumArrInt(numbers))
	case "MED":
		return RunMED(*numbers)
	default:
		return 0.0
	}
}

func RunMED(numbers []int) float64 {

	if len(numbers) == 0 {
		return 0
	}

	nums := make([]int, len(numbers))
	copy(nums, numbers)
	sort.Ints(nums)

	n := len(nums)

	if n%2 == 0 {
		return float64(nums[n/2-1]+nums[n/2]) / 2.0
	} else {
		return float64(nums[n/2])
	}
}

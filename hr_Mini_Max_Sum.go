package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	arr1 := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		arr1[i] = int(arr[i])
	}
	sort.Ints(arr1)
	miniSum, maxSum := 0, 0
	miniSum = arr1[0] + arr1[1] + arr1[2] + arr1[3]
	maxSum = arr1[1] + arr1[2] + arr1[3] + arr1[4]
	fmt.Println(miniSum, maxSum)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

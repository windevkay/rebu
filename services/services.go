package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/windevkay/rebu/structs"
)

func convertStringHelper (value string) float32 {
	result, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic("Error converting cost data...terminating 😓")
	}
	return float32(result)
}
/*
* Returns an array of drivers based on requested range
* The requested range can be incremented to widen a search
*/
func GetDriversInRange(requestedRange float32, wg *sync.WaitGroup, result chan structs.Driver, control chan structs.Control) {
	defer wg.Done()
	
	file, err := os.Open("./data.txt")
	if err != nil {
		panic("Error reading driver data...terminating 😓")
	}

	fileReader := bufio.NewReader(file)

	for {
		content, err := fileReader.ReadString('\n')
		data := strings.Split(content, ",")

		driverDistance := convertStringHelper(data[1])
		cbu := strings.TrimRight(data[2], "\n")

		if driverDistance <= requestedRange {
			driver := structs.Driver{
				Id: data[0],
				Cost: driverDistance,
				CostBeforeUnoccupied: convertStringHelper(cbu),
			}
	
			result <- driver
		}

		if err == io.EOF {
			fmt.Println("Done with initial driver data ✅")
			control <- 0
			break
		}
	}
}
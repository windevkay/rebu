package services

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/windevkay/rebu/structs"
)

type Driver structs.Driver

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
func GetDriversInRange(requestedRange float32, wg *sync.WaitGroup) []Driver {
	defer wg.Done()
	
	file, err := os.Open("./data.txt")
	if err != nil {
		panic("Error reading driver data...terminating 😓")
	}

	fileReader := bufio.NewReader(file)
	var drivers []Driver

	for {
		content, err := fileReader.ReadString('\n')
		data := strings.Split(content, ",")

		driverDistance := convertStringHelper(data[1])
		if driverDistance <= requestedRange {
			driver := Driver{
				Id: data[0],
				Cost: driverDistance,
				CostBeforeUnoccupied: convertStringHelper(data[2]),
			}
	
			drivers = append(drivers, driver)
		}

		if err == io.EOF {
			break
		}
	}
	return drivers
}
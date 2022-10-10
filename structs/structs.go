package structs

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Driver struct {
	Id string
	//the cost represents how far the driver is from the rider
	Cost float32
	//if occupied, how much farther is the dropoff cost for the current rider
	CostBeforeUnoccupied float32
	//if occupied, the cost + (2 * costBeforeUnoccupied) will give the 'true cost'
}

func (driver *Driver) ModifyCost(newCost float32) {
	driver.Cost = newCost
}

func (driver *Driver) GetTrueCost() float32 {
	return driver.Cost + (2 * driver.CostBeforeUnoccupied)
}

func (driver *Driver) FreeDriver() {
	driver.CostBeforeUnoccupied = 0
}

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
func GetDriversInRange(requestedRange float32) []Driver {
	file, err := os.Open("./data.txt")
	if err != nil {
		panic("Error reading driver data...terminating 😓")
	}

	fileReader := bufio.NewReader(file)
	var drivers []Driver

	for {
		content, err := fileReader.ReadString('\n')
		data := strings.Split(content, ",")

		driver := Driver{
			Id: data[0],
			Cost: convertStringHelper(data[1]),
			CostBeforeUnoccupied: convertStringHelper(data[2]),
		}

		drivers = append(drivers, driver)

		if err == io.EOF {
			break
		}
	}
	return drivers
}
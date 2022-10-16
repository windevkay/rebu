package structs

type Control int
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


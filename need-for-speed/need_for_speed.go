package speed

type Track struct {
	distance int
}

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain

	return car
}

func CanFinish(car Car, track Track) bool {
	drainUnits := track.distance / car.speed

	totalDrain := drainUnits * car.batteryDrain

	return totalDrain <= car.battery
}

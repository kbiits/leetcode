type ParkingSystem struct {
    availBig int
    availMedium int
    availSmall int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{big, medium, small}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
        case 1:
            this.availBig--
            return this.availBig >= 0
        case 2:
            this.availMedium--
            return this.availMedium >= 0
        default:
            this.availSmall--
            return this.availSmall >= 0
    }
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

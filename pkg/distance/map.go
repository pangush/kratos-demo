package distance

import (
	"math"
)

type Config struct {
	Latitude float64
	Longitude float64
}

type Distance struct {
	Location Config
	Elements []*Destination
}

//目的地
type Destination struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance float64 `json:"distance"`
	Code 	string `json:"-"`
	Addr string `json:"addr"`
	CityID int64 `json:"city_id"`
	SchoolID int64 `json:"school_id"`
	DistrictName string `json:"district_name"`
	School string `json:"school"`
	ShortLetter string `json:"short_letter"`
}

//地球半径
const EARTH_RADIUS = float64(6371.0 )

func NewDistance(conf Config) Distance {
	return Distance{
		Location: conf,
		Elements: nil,
	}
}

func (d *Distance) sort(destination *Destination)  {
	//list := binarySort(d.Elements, destination)
	//d.Elements = list
	elementList := d.Elements
	idex := binarySearch(elementList, destination)
	list := make([]*Destination, 0)

	if idex < 0 {
		list = append(list, destination)
		list = append(list, elementList[:]...)
	} else if idex > len(elementList) - 1 {
		list = append(list, elementList[:]...)
		list = append(list, destination)
	} else {
		list = append(list, elementList[:idex]...)
		list = append(list, destination)
		list = append(list, elementList[idex:]...)
	}
	d.Elements = list
}

func binarySearch(list []*Destination, destination *Destination) int {
	if len(list) == 0 {
		return 0
	}
	len := len(list)
	left := 0
	right := len - 1

	for left <= right {
		mid := (left + right) / 2
		if list[mid].Distance == destination.Distance {
			return mid
		} else if list[mid].Distance > destination.Distance {
			right = mid - 1
		} else if list[mid].Distance < destination.Distance {
			left = mid + 1
		}
	}

	return left
}

func binarySort(list []*Destination, destination *Destination) []*Destination {
	if len(list) == 0 {
		return []*Destination{destination}
	}
	mid := len(list) / 2
	if list[mid].Distance <= destination.Distance {
		return append(list[:mid+1], binarySort(list[mid+1:], destination)...)
	}

	return  append(binarySort(list[:mid],  destination),  list[mid:]...)
}

func (d *Distance) Input(destination *Destination) *Distance {
	distance := d.distance(destination)
	destination.Distance = distance
	d.sort(destination)

	return d
}

func (d *Distance)  distance(destination *Destination) float64  {
	//用haversine公式计算球面两点间的距离。
	//经纬度转换成弧度
	lat1 := ConvertDegreesToRadians(d.Location.Latitude);
	lon1 := ConvertDegreesToRadians(d.Location.Longitude);
	lat2 := ConvertDegreesToRadians(destination.Latitude);
	lon2 := ConvertDegreesToRadians(destination.Longitude);

	//差值
	vLon := math.Abs(lon1 - lon2)
	vLat := math.Abs(lat1 - lat2)

	//h is the great circle distance in radians, great circle就是一个球体上的切面，它的圆心即是球心的一个周长最大的圆。
	var h = d.HaverSin(vLat) + math.Cos(lat1) * math.Cos(lat2) * d.HaverSin(vLon);

	var distance = 2 * EARTH_RADIUS * math.Asin(math.Sqrt(h));

	return distance;
}

// 将角度换算为弧度。
func ConvertDegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func ConvertRadiansToDegrees(radian float64) float64 {
	return radian * 180.0 / math.Pi;
}

func (d *Distance) HaverSin(theta float64) float64 {
	v := math.Sin(theta / 2)
	return v * v
}
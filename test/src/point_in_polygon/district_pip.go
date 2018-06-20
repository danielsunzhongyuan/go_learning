package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/JamesMilnerUK/pip-go"
)

type JSONDistrict struct {
	Name           string
	PositionBorder string `json:"position_border"`
	Longitude      string
	Latitude       string
}

type JSONCity struct {
	CityName  string
	Districts []*JSONDistrict
}

func readFile(filename string) ([]*JSONDistrict, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	var ret JSONCity
	if err := json.Unmarshal(bytes, &ret); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}
	return ret.Districts, nil
}

func initPolygon(pointPairsRaw string) *pip.Polygon {
	pointPairs := strings.Split(pointPairsRaw, ";")
	fmt.Println(len(pointPairs))
	//points := make([]pip.Point, len(pointPairs), len(pointPairs))
	points := make([]pip.Point, 0, len(pointPairs))
	for _, pair := range pointPairs {
		point := strings.Split(pair, ",")
		if len(point) != 2 {
			continue
		}
		lng, err := strconv.ParseFloat(point[0], 64)
		if err != nil {
			continue
		}
		lat, err := strconv.ParseFloat(point[1], 64)
		if err != nil {
			continue
		}
		p := pip.Point{lng, lat}
		points = append(points, p)
	}
	fmt.Println(len(points))
	rectangle := pip.Polygon{
		Points: points,
	}
	return &rectangle
}

func timeoutCheck(tag string, start time.Time) {
	dis := time.Since(start).Nanoseconds() / (1000 * 1000)
	fmt.Println(tag, dis, "ms")
}

func main() {
	ret, err := readFile("beijing_districts.json")

	if err != nil {
		panic(err)
	}

	allRegionMap := make(map[string]*pip.Polygon)

	for _, y := range ret {
		allRegionMap[y.Name] = initPolygon(y.PositionBorder)
	}

	defer timeoutCheck("main", time.Now())

	/*
		付青龙 - 这个在房山的边界那里，但确实是在房山内的
		lng: 115.721059
		lat： 39.571885

		罗及 - 这个在平谷中心
		lng: 117.130808
		lat: 40.151539

		张军 - 这个跑到辽宁省大连市旅顺口区了，肯定是北京外
		lng: 121.47.9609
		lat: 39.39.8444
	*/

	//p := pip.Point{X: 115.721059, Y: 39.571885} // 付青龙
	p := pip.Point{X: 117.130808, Y: 40.151539} // 罗及
	for x, y := range allRegionMap {
		in := pip.PointInPolygon(p, *y)
		if in {
			fmt.Println(x, p, in)
		}
	}

}

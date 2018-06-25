package main

import (
	"fmt"
	"sort"
)

type Record struct {
	RoomRow   int
	RoowCol   int
	Z         int
	TimeStamp string
}

type Records []Record

func (rs Records) Less(i, j int) bool {
	return rs[i].TimeStamp < rs[j].TimeStamp
}

func (rs Records) Len() int {
	return len(rs)
}

func (rs Records) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

type RoomStatus struct {
	Enters    int
	Leaves    int
	IsLightOn int
}

type Rooms struct {
	RoomStatus    [][]RoomStatus
	TotalLightsOn int
}

func main() {
	// scan input
	n, m, k := 0, 0, 0
	fmt.Scan(&n, &m, &k)
	records := make(Records, 0, k)
	tmpRow, tmpCol, tmpZ, tmpTimestamp := 0, 0, 0, ""
	for i := 0; i < k; i++ {
		fmt.Scan(&tmpRow, &tmpCol, &tmpZ, &tmpTimestamp)
		records = append(records, Record{
			RoomRow:   tmpRow,
			RoowCol:   tmpCol,
			Z:         tmpZ,
			TimeStamp: tmpTimestamp,
		})
	}
	// scan input done

	// sort by timestamp asc
	sort.Sort(records)

	// init rooms and ret
	rooms := Rooms{TotalLightsOn: 0}
	ret := Rooms{TotalLightsOn: 0}
	rooms.RoomStatus = make([][]RoomStatus, 0, n)
	ret.RoomStatus = make([][]RoomStatus, 0, n)
	for i := 0; i < n; i++ {
		rooms.RoomStatus = append(rooms.RoomStatus, make([]RoomStatus, m))
		ret.RoomStatus = append(ret.RoomStatus, make([]RoomStatus, m))
	}

	for _, record := range records {
		if record.Z == 0 {
			rooms.RoomStatus[record.RoomRow-1][record.RoowCol-1].Enters += 1
			if rooms.RoomStatus[record.RoomRow-1][record.RoowCol-1].IsLightOn == 0 {
				rooms.RoomStatus[record.RoomRow-1][record.RoowCol-1].IsLightOn = 1
				rooms.TotalLightsOn += 1
			}
			if rooms.TotalLightsOn >= ret.TotalLightsOn {
				copyFromRooms(ret, rooms)
				ret.TotalLightsOn = rooms.TotalLightsOn
			}
		} else {
			rooms.RoomStatus[record.RoomRow-1][record.RoowCol-1].Leaves += 1
			if rooms.RoomStatus[record.RoomRow-1][record.RoowCol-1].Enters == rooms.RoomStatus[record.RoomRow-1][record.RoowCol-1].Leaves {
				rooms.RoomStatus[record.RoomRow-1][record.RoowCol-1].IsLightOn = 0
				rooms.TotalLightsOn -= 1
			}
		}
	}

	// output final result
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(ret.RoomStatus[i][j].IsLightOn)
		}
		fmt.Println()
	}
}

func copyFromRooms(dest, src Rooms) {
	fmt.Println("dest:", dest)
	n := len(src.RoomStatus)
	for i := 0; i < n; i++ {
		m := len(src.RoomStatus[i])
		for j := 0; j < m; j++ {
			dest.RoomStatus[i][j].IsLightOn = src.RoomStatus[i][j].IsLightOn
		}
	}
}

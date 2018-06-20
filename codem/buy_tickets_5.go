package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Ticket struct {
	FromCity  int
	ToCity    int
	Cost      int
	StartTime int
	EndTime   int
}

type Path struct {
	cur  *Ticket
	next *Ticket
}

type PathToCityX struct {
	ToCity int
	Paths  []Path
}

type PathToCities []*PathToCityX

type SortBy func(p, q *Ticket) bool

type TicketWrapper struct {
	tickets []Ticket
	by      SortBy
}

func (tw TicketWrapper) Len() int {
	return len(tw.tickets)
}

func (tw TicketWrapper) Swap(i, j int) {
	tw.tickets[i], tw.tickets[j] = tw.tickets[j], tw.tickets[i]
}

func (tw TicketWrapper) Less(i, j int) bool {
	return tw.by(&tw.tickets[i], &tw.tickets[j])
}

func SortTicket(tickets []Ticket, by SortBy) {
	sort.Sort(TicketWrapper{tickets: tickets, by: by})
}

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	// get all tickets
	tickets := make([]Ticket, 0, m)
	for i := 0; i < m; i++ {
		fromCity, toCity, cost, startTime, endTime := 0, 0, 0, "", ""
		fmt.Scan(&fromCity, &toCity, &cost, &startTime, &endTime)
		startTimeInMinutes, err := convertTimeToMinutes(startTime)
		if err != nil {
			fmt.Println("Error in converting start_time to minutes:", err)
			return
		}
		endTimeInMinutes, err := convertTimeToMinutes(endTime)
		if err != nil {
			fmt.Println("Error in converting end_time to minutes:", err)
			return
		}
		tickets = append(tickets, Ticket{
			FromCity:  fromCity,
			ToCity:    toCity,
			Cost:      cost,
			StartTime: startTimeInMinutes,
			EndTime:   endTimeInMinutes,
		})
	}

	SortTicket(tickets, func(p, q *Ticket) bool {
		if p.FromCity != q.FromCity {
			return p.FromCity < q.FromCity
		}
		if p.StartTime != q.StartTime {
			return p.StartTime < q.StartTime
		}
		if p.EndTime != q.EndTime {
			return p.EndTime < q.EndTime
		}
		if p.Cost != q.Cost {
			return p.Cost < q.Cost
		}
		return p.ToCity < q.ToCity
	})

	// we should also remove some duplicated or useless tickets. Later

	// paths to city 2 ~ city n
	pathsToCities := make(PathToCities, n+1, n+1)

	for index, ticket := range tickets {
		// 只有当这个ticket有后备的时候，才可能加到 path里去，否则哪怕直接到终点站了，也不算数
		if getBackupTicket(tickets, index, m) {
			// 如果 "到达city x" 的路线还是空的，就新建一个
			if pathsToCities[ticket.ToCity] == nil {
				tmpPath := make([]Path, 0)
				tmpPath = append(tmpPath, Path{cur: &ticket})
				pathsToCities[ticket.FromCity] = &PathToCityX{ToCity: ticket.ToCity}
			} else {

			}
		}
	}
}

func getBackupTicket(tickets []Ticket, index, totalNumber int) bool {
	if index < totalNumber-1 {
		tmp := index + 1
		for tmp < totalNumber {
			if tickets[index].FromCity == tickets[tmp].FromCity &&
				tickets[index].StartTime <= tickets[tmp].StartTime-30 {
				return true
			}
			if tickets[index].FromCity < tickets[tmp].FromCity {
				return false
			}
		}
		return false
	}
	return false
}

func convertTimeToMinutes(time string) (int, error) {
	hour, err := strconv.Atoi(time[:2])
	if err != nil {
		return 0, err
	}
	minute, err := strconv.Atoi(time[3:])
	if err != nil {
		return 0, err
	}
	return hour*60 + minute, nil
}

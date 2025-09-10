package main

import (
	"fmt"
	"sync"
)

type SplitwiseService struct {
	Users  map[string]*User
	Groups map[string]*Group
	mu     sync.Mutex
}

var instance *SplitwiseService
var once sync.Once

func GetSplitwiseService() *SplitwiseService {
	once.Do(func() {
		instance = &SplitwiseService{
			Users:  make(map[string]*User),
			Groups: make(map[string]*Group),
		}
	})
	return instance
}

func (s *SplitwiseService) AddUser(user *User) {
	s.Users[user.Id] = user
}

func (s *SplitwiseService) AddGroup(group *Group) {
	s.Groups[group.Id] = group
}

func (s *SplitwiseService) DistributeExpenseInGroup(expense *Expense) {
	groupId := expense.Group.Id
	if _, ok := s.Groups[groupId]; !ok {
		fmt.Errorf("group %s does not exist", groupId)
	}
	group := s.Groups[groupId]
	amountPerhead := expense.Amount / float64(len(group.Members))
	for _, member := range group.Members {
		if member.Id == expense.PaidBy.Id {
			member.Amount += amountPerhead
		} else {
			member.Amount -= amountPerhead
		}
	}
}

func (s *SplitwiseService) ShowGroupUsersBalance(group *Group) {
	groupId := group.Id
	if _, ok := s.Groups[groupId]; !ok {
		fmt.Errorf("group %s does not exist", groupId)
	}
	//groupInfo := s.Groups[groupId]
	for _, member := range group.Members {
		fmt.Printf("member %s, has balance %f\n", member.Name, member.Amount)
	}
}

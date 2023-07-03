package services

import "fmt"

type notifyFollowersOnProfileUpdatedService struct {
}

func NewNotifyFollowersOnProfileUpdatedService() notifyFollowersOnProfileUpdatedService {
	return notifyFollowersOnProfileUpdatedService{}
}

func (svc notifyFollowersOnProfileUpdatedService) Consume(value string) {
	fmt.Println("Sending notification to followers")
}

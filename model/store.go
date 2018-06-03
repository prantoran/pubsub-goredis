package model

import "sync"

type Store struct {
	Users []*User
	sync.Mutex
}

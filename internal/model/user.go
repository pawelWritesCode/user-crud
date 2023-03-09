package model

import (
	"encoding/xml"
	"time"
)

// User defines user entity.
type User struct {
	XMLName     xml.Name  `json:"-" yaml:"-" xml:"user"`
	Id          int       `json:"id" yaml:"id" xml:"id"`
	FirstName   string    `json:"firstName" yaml:"firstName" xml:"firstName"`
	LastName    string    `json:"lastName" yaml:"lastName" xml:"lastName"`
	Age         uint      `json:"age" yaml:"age" xml:"age"`
	Description string    `json:"description" yaml:"description" xml:"description"`
	FriendSince time.Time `json:"friendSince" yaml:"friendSince" xml:"friendSince"`
	Avatar      string    `json:"avatar,omitempty" yaml:"avatar,omitempty" xml:"avatar,omitempty"`
}

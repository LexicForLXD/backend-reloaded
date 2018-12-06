// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	time "time"
)

type AuthHostReq struct {
	Password string `json:"password"`
}

type ContainerReq struct {
	Name   string  `json:"name"`
	HostID *string `json:"hostID"`
}

type ContianerSource struct {
	Type          *string   `json:"type"`
	Certificate   *string   `json:"certificate"`
	Alias         *string   `json:"alias"`
	Fingerprint   *string   `json:"fingerprint"`
	Properties    []*string `json:"properties"`
	Server        *string   `json:"server"`
	Secret        *string   `json:"secret"`
	Protocol      *string   `json:"protocol"`
	Source        *string   `json:"source"`
	Live          *bool     `json:"live"`
	ContainerOnly *bool     `json:"containerOnly"`
	Refresh       *bool     `json:"refresh"`
	Project       *string   `json:"project"`
}

type DeleteRes struct {
	Message string `json:"message"`
	Entity  string `json:"entity"`
}

type HostReq struct {
	Name     *string `json:"name"`
	Desc     *string `json:"desc"`
	Address  *string `json:"address"`
	Password *string `json:"password"`
}

type Info struct {
	Version string `json:"version"`
}

type UserReq struct {
	Name      string     `json:"name"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Password  *string    `json:"password"`
	Birthday  *time.Time `json:"birthday"`
}

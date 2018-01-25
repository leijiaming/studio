// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

import "encoding/json"

type User struct {
	Id       uint  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Email    string `json:"email"`
}

func (u *User) String() string {
	c, _ := json.Marshal(u)
	return string(c)
}

package model

import "time"

type Job struct {
	Name        string    `json:"name"			validate:"required"`
	ServiceType string    `json:"servicetype"	validate:"required"`
	Description string    `json:"description"	validate:"required"`
	Date        time.Time `json:"date"			validate:"required"`
	Time        time.Time `json:"time"			validate:"required"`
}

// type Job struct {

// 	Name   		  string         `json:"name"    	   	validate:"required,email"`
// 	ServiceType   string         `json:"servicetype"    validate:"required,max=12,min=8,alpha"`
// 	Description   string         `json:"description"    validate:"required"`
// 	Date 		  time.Date		 `json:"date" 			validate:"required"`
// 	Time 		  time.Time		 `json:"time" 			validate:"required"`
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RegisterBody struct{
	first-name string `json:"first-name"`
	last-name string `json:"last-name"`
	Email    string `json:"email"`
	userPassword string `json:"userPassword"`
	phone string `json:"phone"`
	// date of birth
	// sex
}

type LoginBody struct{
	Email    string `json:"email"`
	userPassword string `json:"userPassword"`
}

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}


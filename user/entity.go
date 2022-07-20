package user

import "time"

type User struct {
	ID               int
	Occupation       string
	Name             string
	Email            string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Token            string
	Created_at       time.Time
	Update_at        time.Time
}

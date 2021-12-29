package main

type FormEmail struct {
	Email string `form:"email" validate:"required,email"`
}

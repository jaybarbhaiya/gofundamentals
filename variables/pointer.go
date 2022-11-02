package main

import "fmt"

func main() {
	name:="Nigel Poulton"
	course:="Getting started with Kubernetes"

	fmt.Println("Hi", name, "your current course is", course)

	updateCourse(&course)

	fmt.Println("Your current course is", course)
}

func updateCourse(course *string) string {
	*course = "Getting started with Docker"
	fmt.Println("Course updated", *course)
	return *course
}
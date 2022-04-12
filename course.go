package opendata

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Course is a normalized struct for course and section ID.
type Course struct{ string }

func validCourse(course string) bool {
	if len(course) <= 0 {
		return false
	}
	switch course[len(course)-1] {
	case 'A', 'B', 'a', 'b':
		course = course[:len(course)-1]
	}
	num, err := strconv.Atoi(course)
	if err != nil || num <= 0 {
		return false
	}
	return num <= 9999
}

// NewCourse generates a new Course instance based on course subject, number, and section ID.
func NewCourse(subject, number, section string) *Course {
	subject = strings.ToUpper(strings.TrimSpace(subject))
	section = strings.ToUpper(strings.TrimSpace(section))
	if len(subject) > 4 || len(subject) <= 1 || !validCourse(number) || len(section) != 3 {
		return nil
	}
	return &Course{fmt.Sprintf("%s%04s%s", subject, number, section)}
}

var courseRegex = regexp.MustCompile(`^([a-zA-Z]{2,4})\s*-?([0-9]{2,4}[AB]?)-?([0-9a-zA-Z]{3})$`)

// ParseCourse generates a new Course instance based on course ID string using regex to match.
// 	ParseCourse("MUSC0050003")
func ParseCourse(course string) *Course {
	match := courseRegex.FindStringSubmatch(course)
	if len(match) != 4 {
		return nil
	}
	return NewCourse(match[1], match[2], match[3])
}

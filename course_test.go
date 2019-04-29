package opendata

import "testing"

func TestParseCourseFullWidth(t *testing.T) {
	course := ParseCourse("NETS150001")
	if course == nil || course.string != "NETS150001" {
		t.Fail()
	}
}

func TestParseCourseTrimWidth(t *testing.T) {
	course := ParseCourse("CIS120001")
	if course == nil || course.string != "CIS 120001" {
		t.Fail()
	}
}

func TestParseCourseFullWidthWhiteSpace(t *testing.T) {
	course := ParseCourse("CIS 120001")
	if course == nil || course.string != "CIS 120001" {
		t.Fail()
	}
}

func TestParseCourseInvalidCourse(t *testing.T) {
	course := ParseCourse("CIS000001")
	if course != nil {
		t.Fail()
	}
}

func TestParseCourseInvalidDepartWidth(t *testing.T) {
	course := ParseCourse("CISXX000001")
	if course != nil {
		t.Fail()
	}
}

func TestParseCourseInvalidCourseWidth(t *testing.T) {
	course := ParseCourse("CIS50001")
	if course != nil {
		t.Fail()
	}
}

func TestParseCourseInvalidSection(t *testing.T) {
	course := ParseCourse("CIS001000")
	if course != nil {
		t.Fail()
	}
}

func TestParseCourseReadable(t *testing.T) {
	course := ParseCourseReadable("CIS-120-001")
	if course == nil || course.string != "CIS 120001" {
		t.Fail()
	}
}

func TestParseCourseRegex(t *testing.T) {
	course := ParseCourseRegex("MUSC50001")
	if course == nil || course.string != "MUSC050001" {
		t.Fail()
	}
}

func TestParseCourseRegexInvalid(t *testing.T) {
	course := ParseCourseRegex("GARBage1023023")
	if course != nil {
		t.Fail()
	}
}

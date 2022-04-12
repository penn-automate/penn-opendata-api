package opendata

import "testing"

func TestParseCourseFullWidth(t *testing.T) {
	course := ParseCourse("NETS1120001")
	if course == nil || course.string != "NETS1120001" {
		t.Fail()
	}
}

func TestParseCourseTrimWidth(t *testing.T) {
	course := ParseCourse("CIS1200001")
	if course == nil || course.string != "CIS1200001" {
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

func TestParseCourseNormalized(t *testing.T) {
	course := ParseCourse("CIS-1200-001")
	if course == nil || course.string != "CIS1200001" {
		t.Fail()
	}
}

func TestParseCourseMultiTerm(t *testing.T) {
	course := ParseCourse("CRIM-6004A-301")
	if course == nil || course.string != "CRIM6004A301" {
		t.Fail()
	}
}

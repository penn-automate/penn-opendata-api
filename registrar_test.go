package opendata

import (
	"os"
	"testing"
)

var api = NewOpenDataAPI(os.Getenv("API_KEY"), os.Getenv("API_PASS")).GetRegistrar()

func TestGetAllCourseStatus(t *testing.T) {
	status, err := api.GetAllCourseStatus("2019C")
	if err != nil {
		t.Fatal(err)
	}
	if len(status) <= 0 {
		t.FailNow()
	}
}

func TestGetSingleCourseStatus(t *testing.T) {
	status, err := api.GetCourseStatus("2019C", NewCourse("CIS", 120, 001))
	if err != nil {
		t.Fatal(err)
	}
	if len(status) <= 0 {
		t.FailNow()
	}
}

func TestGetCatalogCourseInfo(t *testing.T) {
	iter := api.GetCourseCatalog("NETS", 0)
	for iter.NextPage() {
		data := new(CourseCatalogData)
		if iter.GetPageSize() <= 0 {
			t.FailNow()
		}
		err := iter.GetResult(data, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	if iter.GetError() != nil {
		t.Fatal(iter.GetError())
	}
}

func TestGetDepartmentsMap(t *testing.T) {
	if _, err := api.GetDepartmentsMap(); err != nil {
		t.Fatal(err)
	}
}

package opendata

import (
	"os"
	"testing"
)

var api = NewOpenDataAPI(os.Getenv("API_KEY"), os.Getenv("API_PASS")).GetRegistrar()

func TestGetAllCourseStatus(t *testing.T) {
	iter := api.GetAllCourseStatus("2019C")
	for iter.NextPage() {
		var statusData []*CourseStatusData
		err := iter.GetResult(&statusData)
		if err != nil {
			t.Fatal(err)
		}
		if len(statusData) <= 0 {
			t.FailNow()
		}
	}
	if iter.Error() != nil {
		t.Fatal(iter.Error())
	}
}

func TestGetSingleCourseStatus(t *testing.T) {
	iter := api.GetCourseStatus("2019C", NewCourse("CIS", 120, 001))
	for iter.NextPage() {
		var data []*CourseStatusData
		err := iter.GetResult(&data)
		if err != nil {
			t.Fatal(err)
		}
		if len(data) <= 0 {
			t.FailNow()
		}
	}
	if iter.Error() != nil {
		t.Fatal(iter.Error())
	}
}

func TestGetCatalogCourseInfo(t *testing.T) {
	iter := api.GetCourseCatalog("NETS", 0)
	for iter.NextPage() {
		var data []*CourseCatalogData
		err := iter.GetResult(&data)
		if err != nil {
			t.Fatal(err)
		}
		if len(data) <= 0 {
			t.FailNow()
		}
	}
	if iter.Error() != nil {
		t.Fatal(iter.Error())
	}
}

func TestGetDepartmentsMap(t *testing.T) {
	if _, err := api.GetDepartmentsMap(); err != nil {
		t.Fatal(err)
	}
}

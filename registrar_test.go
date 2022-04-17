package opendata

import (
	"os"
	"testing"
)

var api = NewOpenDataAPI(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")).GetRegistrar()

func TestGetAllCourseStatus(t *testing.T) {
	status, err := api.GetAllCourseStatus("202230")
	if err != nil {
		t.Fatal(err)
	}
	if len(status) <= 0 {
		t.FailNow()
	}
	for _, s := range status {
		sp1 := ParseCourse(s.SectionID)
		if sp1 == nil {
			t.Fatalf("failed to parse %s", s.SectionID)
		}
		sp2 := ParseCourse(s.SectionIDNormalized)
		if sp2 == nil {
			t.Fatalf("failed to parse %s", s.SectionIDNormalized)
		}
		if sp1.string != sp2.string {
			t.Fatalf("failed to parse %s, %v != %v", s.SectionID, sp1, sp2)
		}
		if sp2.string != s.SectionID {
			t.Fatalf("failed to normalize %s, %v", s.SectionID, sp2)
		}
	}
}

func TestGetSingleCourseStatus(t *testing.T) {
	status, err := api.GetCourseStatus("202230", NewCourse("CIS", "1200", "001"))
	if err != nil {
		t.Fatal(err)
	}
	if len(status) <= 0 {
		t.FailNow()
	}
}

func TestGetCatalogCourseInfo(t *testing.T) {
	iter := api.GetCourseCatalog("NETS", "")
	for iter.NextPage() {
		if iter.GetPageSize() <= 0 {
			t.FailNow()
		}
		_, err := iter.GetResult(0)
		if err != nil {
			t.Fatal(err)
		}
	}
	if iter.GetError() != nil {
		t.Fatal(iter.GetError())
	}
}

func TestGetAvailableTermMap(t *testing.T) {
	if ret, err := api.GetAvailableTermMap(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(ret)
	}
}

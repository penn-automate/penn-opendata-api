package opendata

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// Registrar provides a wrapper for OpenData Registrar's API.
type Registrar struct {
	od        *OpenData
	parameter *courseSectionSearchParameters
	paraLock  sync.Mutex
}

const (
	courseParameterURL = openDataURL + `course_section_search_parameters`
	courseStatusURL    = openDataURL + `course_section_status/%s/%s`
	courseCatalogURL   = openDataURL + `course_info/%s`
	courseSearchURL    = openDataURL + `course_section_search`
)

func (r *Registrar) checkTerm(term string) error {
	allowed, err := r.GetAvailableTermMap()
	if err != nil {
		return err
	}
	_, ok := allowed[term]
	if !ok {
		return fmt.Errorf(`term %q does not exist`, term)
	}
	return nil
}

func (r *Registrar) courseStatus(term, course string) ([]CourseSectionStatus, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(courseStatusURL, term, course), nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.od.access(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data := new(data)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	ret := make([]CourseSectionStatus, len(data.ResultData))
	for i := range data.ResultData {
		if err := json.Unmarshal(data.ResultData[i], &ret[i]); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

// GetAllCourseStatus gets all courses' status in a given term at once.
// Term must be in the available term map.
// Call #Registrar.GetAvailableTermMap to get the map.
// See https://app.swaggerhub.com/apis-docs/UPennISC/open-data/prod#/Course%20section%20status%20service/getAllCourseSectionStatuses.
func (r *Registrar) GetAllCourseStatus(term string) ([]CourseSectionStatus, error) {
	if err := r.checkTerm(term); err != nil {
		return nil, err
	}
	return r.courseStatus(term, "all")
}

// GetCourseStatus gets the specific course's status in a given term.
// Term must be in the available term map.
// Call #Registrar.GetAvailableTermMap to get the map.
// See https://app.swaggerhub.com/apis-docs/UPennISC/open-data/prod#/Course%20section%20status%20service/getOneCourseSectionStatus.
func (r *Registrar) GetCourseStatus(term string, course *Course) ([]CourseSectionStatus, error) {
	if err := r.checkTerm(term); err != nil {
		return nil, err
	}
	return r.courseStatus("id/"+term, course.string)
}

// GetCourseCatalog allows the search of the course catalog using subjects and course numbers.
// See https://app.swaggerhub.com/apis-docs/UPennISC/open-data/prod#/Course%20search%20service.
func (r *Registrar) GetCourseCatalog(department, section string) *PageIterator {
	req, err := http.NewRequest("GET", fmt.Sprintf(courseCatalogURL, department), nil)
	if err != nil {
		return newErrorIter(err)
	}
	if section != "" {
		req.URL.Path += fmt.Sprintf("/%s", section)
	}
	return newIter(r.od, req)
}

// SearchCourseSection gets the searched results with given parameters on Path@Penn.
// The parameters map must have keys that are in acceptable search url parameters map.
// Call #Registrar.GetAcceptableSearchURLParametersMap to get the map.
// See https://app.swaggerhub.com/apis-docs/UPennISC/open-data/prod#/Course%20section%20search%20service/searchCourseSections.
func (r *Registrar) SearchCourseSection(parameters map[string]string) *PageIterator {
	req, err := http.NewRequest("GET", courseSearchURL, nil)
	if err != nil {
		return newErrorIter(err)
	}
	if parameters != nil {
		value := make(url.Values)
		allowed, err := r.GetAcceptableSearchURLParametersMap()
		if err != nil {
			return newErrorIter(err)
		}
		for k, v := range parameters {
			_, ok := allowed[k]
			if !ok {
				return newErrorIter(fmt.Errorf(`parameter %q is not supported`, k))
			}
			value.Set(k, v)
		}
		req.URL.RawQuery = value.Encode()
	}
	return newIter(r.od, req)
}

func (r *Registrar) getParameterData() error {
	r.paraLock.Lock()
	defer r.paraLock.Unlock()
	if r.parameter != nil {
		return nil
	}
	req, _ := http.NewRequest("GET", courseParameterURL, nil)
	resp, err := r.od.access(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data := new(data)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}
	if len(data.ResultData) < 1 {
		return errors.New("unexpected result return length")
	}
	r.parameter = new(courseSectionSearchParameters)
	if err := json.Unmarshal(data.ResultData[0], r.parameter); err != nil {
		return err
	}
	return nil
}

// GetAvailableTermMap gets acceptable search url parameters map provided by OpenData API.
func (r *Registrar) GetAvailableTermMap() (map[string]string, error) {
	if err := r.getParameterData(); err != nil {
		return nil, err
	}
	return r.parameter.AvailableTermsMap, nil
}

// GetAcceptableSearchURLParametersMap gets departments map provided by OpenData API.
func (r *Registrar) GetAcceptableSearchURLParametersMap() (map[string]string, error) {
	if err := r.getParameterData(); err != nil {
		return nil, err
	}
	return r.parameter.AcceptableSearchURLParametersMap, nil
}

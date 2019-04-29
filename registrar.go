package opendata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

// Registrar provides an wrapper for OpenData Registrar's API.
type Registrar struct {
	od        *OpenData
	parameter *parameterData
}

const (
	courseParameterURL = openDataURL + `course_section_search_parameters`
	courseStatusURL    = openDataURL + `course_status/%s/%s`
	courseCatalogURL   = openDataURL + `course_info/%s`
	courseSearchURL    = openDataURL + `course_section_search`
)

// GetAllCourseStatus gets all courses' status in a given term at once.
// Term must be in the available term map.
// Call #Registrar.GetAvailableTermMap to get the map.
// See https://esb.isc-seo.upenn.edu/8091/documentation#coursestatusservice.
func (r *Registrar) GetAllCourseStatus(term string) *PageIterator {
	return r.GetCourseStatus(term, &Course{"all"})
}

// GetCourseStatus gets the specific course's status in a given term.
// Term must be in the available term map.
// Call #Registrar.GetAvailableTermMap to get the map.
// See https://esb.isc-seo.upenn.edu/8091/documentation#coursestatusservice.
func (r *Registrar) GetCourseStatus(term string, course *Course) *PageIterator {
	allowed, err := r.GetAvailableTermMap()
	if err != nil {
		return newErrorIter(err)
	}
	_, ok := allowed[term]
	if !ok {
		return newErrorIter(fmt.Errorf(`term "%s" does not exist`, term))
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(courseStatusURL, term, course.string), nil)
	if err != nil {
		return newErrorIter(err)
	}
	return newIter(r.od, req)
}

// GetCourseCatalog allows the search of the course catalog using subjects and course numbers.
// See https://esb.isc-seo.upenn.edu/8091/documentation#coursecatalogsearchservice.
func (r *Registrar) GetCourseCatalog(department string, section uint) *PageIterator {
	req, err := http.NewRequest("GET", fmt.Sprintf(courseCatalogURL, department), nil)
	if err != nil {
		return newErrorIter(err)
	}
	if section != 0 {
		req.URL.Path += fmt.Sprintf("/%03d", section)
	}
	return newIter(r.od, req)
}

// SearchCourseSection gets the searched results with given parameters on PennInTouch.
// The parameters map must have keys that are in acceptable search url parameters map.
// Call #Registrar.GetAcceptableSearchURLParametersMap to get the map.
// See https://esb.isc-seo.upenn.edu/8091/documentation#coursesectionsearchservice.
func (r *Registrar) SearchCourseSection(parameters map[string]string) *PageIterator {
	value := make(url.Values)
	allowed, err := r.GetAcceptableSearchURLParametersMap()
	if err != nil {
		return newErrorIter(err)
	}
	for k, v := range parameters {
		_, ok := allowed[k]
		if !ok {
			return newErrorIter(fmt.Errorf(`parameter "%s" is not supported`, k))
		}
		value.Set(k, v)
	}
	req, err := http.NewRequest("GET", courseSearchURL, nil)
	if err != nil {
		return newErrorIter(err)
	}
	req.URL.RawQuery = value.Encode()
	return newIter(r.od, req)
}

func (r *Registrar) getParameterData() error {
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
	var pData []*parameterData
	if err := json.Unmarshal(data.ResultData, &pData); err != nil {
		return err
	}
	r.parameter = pData[0]
	return nil
}

// GetAvailableTermMap gets acceptable search url parameters map provided by OpenData API.
func (r *Registrar) GetAvailableTermMap() (map[string]string, error) {
	if err := r.getParameterData(); err != nil {
		return nil, err
	}
	return r.parameter.AvailableTermsMap, nil
}

// GetDepartmentsMap gets departments map provided by OpenData API.
func (r *Registrar) GetDepartmentsMap() (map[string]string, error) {
	if err := r.getParameterData(); err != nil {
		return nil, err
	}
	return r.parameter.DepartmentsMap, nil
}

// GetAcceptableSearchURLParametersMap gets departments map provided by OpenData API.
func (r *Registrar) GetAcceptableSearchURLParametersMap() (map[string]string, error) {
	if err := r.getParameterData(); err != nil {
		return nil, err
	}
	return r.parameter.AcceptableSearchURLParametersMap, nil
}

// Course is a normalized struct for course and section ID.
type Course struct{ string }

// NewCourse generates a new Course instance based on department, course, and section ID.
func NewCourse(department string, course, section uint) *Course {
	department = strings.ToUpper(strings.TrimSpace(department))
	if len(department) >= 4 || len(department) <= 1 || course >= 1000 || section >= 1000 {
		return nil
	}
	return &Course{fmt.Sprintf("%-4s%03d%03d", department, course, section)}
}

// ParseCourse generates a new Course instance based on course ID string.
// 	ParseCourse("NETS212001")
func ParseCourse(course string) *Course {
	var dep []rune
	cur := 0
	for ; cur < len(course); cur++ {
		r := rune(course[cur])
		if unicode.IsLetter(r) {
			dep = append(dep, unicode.ToUpper(r))
		} else if unicode.IsDigit(r) {
			break
		}
	}
	for len(dep) < 4 {
		dep = append(dep, ' ')
	}
	if len(dep) > 4 {
		return nil
	}
	if len(course)-cur != 6 {
		return nil
	}
	c := course[cur : cur+3]
	if c == "000" {
		return nil
	}
	s := course[cur+3:]
	if s == "000" {
		return nil
	}
	return &Course{string(dep) + c + s}
}

// ParseCourseReadable generates a new Course instance based on
// readable course ID string with delimiter.
// 	ParseCourseReadable("NETS-212-001")
func ParseCourseReadable(course string) *Course {
	return ParseCourse(strings.ReplaceAll(course, "-", ""))
}

var courseRegex = regexp.MustCompile(`^([a-zA-Z]{2,4})\s*-?([1-9][0-9][0-9]|[0-9][1-9][0-9]|[0-9][0-9][1-9]|[1-9][0-9]|[0-9][1-9])-?([1-9][0-9][0-9]|[0-9][1-9][0-9]|[0-9][0-9][1-9])$`)

// ParseCourseRegex generates a new Course instance based on course ID string using regex to match.
// 	ParseCourseRegex("MUSC50003")
func ParseCourseRegex(course string) *Course {
	match := courseRegex.FindStringSubmatch(course)
	if len(match) != 4 {
		return nil
	}
	match[1] = strings.ToUpper(match[1])
	return &Course{fmt.Sprintf("%-4s%03s%s", match[1], match[2], match[3])}
}

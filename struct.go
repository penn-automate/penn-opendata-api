package opendata

import (
	"encoding/json"
)

type CourseInstructor struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	MiddleInitial *string `json:"middle_initial"`
	PennId        string  `json:"penn_id"`
	PrimaryInd    *string `json:"primary_ind"`
}

// CourseSearchData is the data struct returned by Registrar.SearchCourseSection.
type CourseSearchData struct {
	Activity            string `json:"activity"`
	ActivityDescription string `json:"activity_description"`
	Attributes          []struct {
		AttributeCode string `json:"attribute_code"`
		AttributeDesc string `json:"attribute_desc"`
	} `json:"attributes"`
	Cancelled                      bool   `json:"cancelled"`
	Closed                         bool   `json:"closed"`
	CorequisiteActivity            string `json:"corequisite_activity"`
	CorequisiteActivityDescription string `json:"corequisite_activity_description"`
	CourseDepartment               string `json:"course_department"`
	CourseDescription              string `json:"course_description"`
	CourseLevel                    string `json:"course_level"`
	CourseLevelDesc                string `json:"course_level_desc"`
	CourseNumber                   string `json:"course_number"`
	CourseTermsOffered             string `json:"course_terms_offered"`
	CourseTitle                    string `json:"course_title"`
	CreditConnector                string `json:"credit_connector"`
	CreditType                     string `json:"credit_type"`
	Credits                        string `json:"credits"`
	Crn                            string `json:"crn"`
	CrosslistPrimary               string `json:"crosslist_primary"`
	Crosslistings                  []struct {
		ActivityDate       string `json:"activity_date"`
		EffectiveTerm      string `json:"effective_term"`
		EndTerm            string `json:"end_term"`
		StartTerm          string `json:"start_term"`
		XlistCourseId      string `json:"xlist_course_id"`
		XlistCourseNumber  string `json:"xlist_course_number"`
		XlistSectionNumber string `json:"xlist_section_number"`
		XlistSubjectCode   string `json:"xlist_subject_code"`
	} `json:"crosslistings"`
	EndDate          string `json:"end_date"`
	FirstMeetingDays string `json:"first_meeting_days"`
	GradeModes       []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"grade_modes"`
	Instructors        []CourseInstructor `json:"instructors"`
	IsCancelled        bool               `json:"is_cancelled"`
	IsClosed           bool               `json:"is_closed"`
	IsCrosslistPrimary bool               `json:"is_crosslist_primary"`
	IsNotScheduled     bool               `json:"is_not_scheduled"`
	LinkedCourses      []struct {
		CourseNumber        string `json:"course_number"`
		ScheduleCode        string `json:"schedule_code"`
		ScheduleDescription string `json:"schedule_description"`
		SectionId           string `json:"section_id"`
		SectionNumber       string `json:"section_number"`
		SubjectCode         string `json:"subject_code"`
	} `json:"linked_courses"`
	MaxEnrollment          string `json:"max_enrollment"`
	MaxEnrollmentCrosslist string `json:"max_enrollment_crosslist"`
	MaximumCredit          string `json:"maximum_credit"`
	Meetings               []struct {
		BeginTime    string `json:"begin_time"`
		BeginTime24  string `json:"begin_time_24"`
		BuildingCode string `json:"building_code"`
		BuildingDesc string `json:"building_desc"`
		Days         string `json:"days"`
		EndDate      string `json:"end_date"`
		EndTime      string `json:"end_time"`
		EndTime24    string `json:"end_time_24"`
		Friday       string `json:"friday"`
		Monday       string `json:"monday"`
		RoomCode     string `json:"room_code"`
		Saturday     string `json:"saturday"`
		StartDate    string `json:"start_date"`
		Sunday       string `json:"sunday"`
		Thursday     string `json:"thursday"`
		Tuesday      string `json:"tuesday"`
		Wednesday    string `json:"wednesday"`
	} `json:"meetings"`
	MinimumCredit     string `json:"minimum_credit"`
	NotScheduled      bool   `json:"notScheduled"`
	PrimaryInstructor string `json:"primary_instructor"`
	SectionId         string `json:"section_id"`
	SectionNumber     string `json:"section_number"`
	SectionTitle      string `json:"section_title"`
	StartDate         string `json:"start_date"`
	Subject           string `json:"subject"`
	SyllabusUrl       string `json:"syllabus_url"`
	Term              string `json:"term"`
	TermSession       string `json:"term_session"`
	XlistGroup        string `json:"xlist_group"`
}

// CourseSectionStatus is the data struct returned by Course section status service.
type CourseSectionStatus struct {
	PreviousStatus       string `json:"previous_status"`
	SectionID            string `json:"section_id"`
	SectionIDNormalized  string `json:"section_id_normalized"`
	Status               string `json:"status"`
	StatusCodeNormalized string `json:"status_code_normalized"`
	Term                 string `json:"term"`
}

// CourseCatalogData is the data struct returned by Registrar.GetCourseCatalog.
type CourseCatalogData struct {
	Activities []struct {
		EffectiveTerm string `json:"effective_term"`
		ScheduleCode  string `json:"schedule_code"`
		ScheduleDesc  string `json:"schedule_desc"`
		Workload      string `json:"workload"`
	} `json:"activities"`
	Attributes []struct {
		AttributeCode string `json:"attribute_code"`
		AttributeDesc string `json:"attribute_desc"`
	} `json:"attributes"`
	Corequisites []struct {
		CoreqCourseId string `json:"coreq_course_id"`
	} `json:"corequisites"`
	CourseCreditConnector  string `json:"course_credit_connector"`
	CourseCreditType       string `json:"course_credit_type"`
	CourseDescription      string `json:"course_description"`
	CourseID               string `json:"course_id"`
	CourseLevel            string `json:"course_level"`
	CourseLevelDescription string `json:"course_level_description"`
	CourseNumber           string `json:"course_number"`
	CourseTitle            string `json:"course_title"`
	Crosslistings          []struct {
		ActivityDate       string `json:"activity_date"`
		EffectiveTerm      string `json:"effective_term"`
		EndTerm            string `json:"end_term"`
		StartTerm          string `json:"start_term"`
		XlistCourseId      string `json:"xlist_course_id"`
		XlistCourseNumber  string `json:"xlist_course_number"`
		XlistSectionNumber string `json:"xlist_section_number"`
		XlistSubjectCode   string `json:"xlist_subject_code"`
	} `json:"crosslistings"`
	Department          string `json:"department"`
	EasCreditFactorCode string `json:"eas_credit_factor_code"`
	Prerequisites       []struct {
		PrereqCourseId string `json:"prereq_course_id"`
	} `json:"prerequisites"`
	SchedulingPriority      string `json:"scheduling_priority"`
	SchoolCode              string `json:"school_code"`
	TermsOfferedCode        string `json:"terms_offered_code"`
	TermsOfferedDescription string `json:"terms_offered_description"`
}

type data struct {
	ResultData  []json.RawMessage `json:"result_data"`
	ServiceMeta serviceMeta       `json:"service_meta"`
}

type courseSectionSearchParameters struct {
	AcceptableSearchURLParametersMap map[string]string `json:"acceptable_search_url_parameters_map"`
	ActivityMap                      map[string]string `json:"activity_map"`
	AvailableTermsMap                map[string]string `json:"available_terms_map"`
	CourseLevelAtOrAboveMap          map[string]string `json:"course_level_at_or_above_map"`
	CourseLevelAtOrBelowMap          map[string]string `json:"course_level_at_or_below_map"`
	EndsAtOrAfterHourMap             map[string]string `json:"ends_at_or_after_hour_map"`
	StartsAtOrAfterHourMap           map[string]string `json:"starts_at_or_after_hour_map"`
	SubjectMap                       map[string]string `json:"subject_map"`
}

type serviceMeta struct {
	CurrentPageNumber  int    `json:"current_page_number"`
	Error              bool   `json:"error"`
	ErrorText          string `json:"error_text"`
	NextPageNumber     int    `json:"next_page_number"`
	NumberOfPages      int    `json:"number_of_pages"`
	PreviousPageNumber int    `json:"previous_page_number"`
	RestCode           int    `json:"rest_code"`
	ResultsPerPage     int    `json:"results_per_page"`
}

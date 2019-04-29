package opendata

import (
	"encoding/json"
	"time"
)

// CourseSearchData is the data struct returned by Registrar.SearchCourseSection.
type CourseSearchData struct {
	Activity                       string `json:"activity"`
	ActivityDescription            string `json:"activity_description"`
	CorequisiteActivity            string `json:"corequisite_activity"`
	CorequisiteActivityDescription string `json:"corequisite_activity_description"`
	CourseDepartment               string `json:"course_department"`
	CourseDescription              string `json:"course_description"`
	CourseDescriptionURL           string `json:"course_description_url"`
	CourseMeetingMessage           string `json:"course_meeting_message"`
	CourseNotes                    string `json:"course_notes"`
	CourseNotesMessage             string `json:"course_notes_message"`
	CourseNumber                   string `json:"course_number"`
	CourseStatus                   string `json:"course_status"`
	CourseStatusNormalized         string `json:"course_status_normalized"`
	CourseStatusReasonCode         string `json:"course_status_reason_code"`
	CourseTermsOffered             string `json:"course_terms_offered"`
	CourseTitle                    string `json:"course_title"`
	CreditAndGradeType             string `json:"credit_and_grade_type"`
	CreditConnector                string `json:"credit_connector"`
	CreditType                     string `json:"credit_type"`
	Credits                        string `json:"credits"`
	CrosslistPrimary               string `json:"crosslist_primary"`
	Crosslistings                  []struct {
		CourseID           string `json:"course_id"`
		IsCrosslistPrimary bool   `json:"is_crosslist_primary"`
		SectionID          string `json:"section_id"`
		Subject            string `json:"subject"`
	} `json:"crosslistings"`
	DepartmentDescription       string        `json:"department_description"`
	DepartmentURL               string        `json:"department_url"`
	EndDate                     time.Time     `json:"end_date"`
	FirstMeetingDays            string        `json:"first_meeting_days"`
	FulfillsCollegeRequirements []interface{} `json:"fulfills_college_requirements"`
	GradeType                   string        `json:"grade_type"`
	ImportantNotes              []interface{} `json:"important_notes"`
	Instructors                 []struct {
		Name      string `json:"name"`
		SectionID string `json:"section_id"`
		Term      string `json:"term"`
	} `json:"instructors"`
	IsCancelled            bool          `json:"is_cancelled"`
	IsClosed               bool          `json:"is_closed"`
	IsCrosslistPrimary     bool          `json:"is_crosslist_primary"`
	IsNotScheduled         bool          `json:"is_not_scheduled"`
	IsSpecialSession       bool          `json:"is_special_session"`
	Labs                   []interface{} `json:"labs"`
	Lectures               []interface{} `json:"lectures"`
	MaxEnrollment          string        `json:"max_enrollment"`
	MaxEnrollmentCrosslist string        `json:"max_enrollment_crosslist"`
	MaximumCredit          string        `json:"maximum_credit"`
	Meetings               []struct {
		BuildingCode        string `json:"building_code"`
		BuildingName        string `json:"building_name"`
		EndHour24           int    `json:"end_hour_24"`
		EndMinutes          int    `json:"end_minutes"`
		EndTime             string `json:"end_time"`
		EndTime24           int    `json:"end_time_24"`
		MeetingDays         string `json:"meeting_days"`
		RoomNumber          string `json:"room_number"`
		SectionID           string `json:"section_id"`
		SectionIDNormalized string `json:"section_id_normalized"`
		StartHour24         int    `json:"start_hour_24"`
		StartMinutes        int    `json:"start_minutes"`
		StartTime           string `json:"start_time"`
		StartTime24         int    `json:"start_time_24"`
		Term                string `json:"term"`
	} `json:"meetings"`
	MinimumCredit       string        `json:"minimum_credit"`
	PrerequisiteNotes   []interface{} `json:"prerequisite_notes"`
	PrimaryInstructor   string        `json:"primary_instructor"`
	Recitations         []interface{} `json:"recitations"`
	Requirements        []interface{} `json:"requirements"`
	RequirementsTitle   string        `json:"requirements_title"`
	SectionID           string        `json:"section_id"`
	SectionIDNormalized string        `json:"section_id_normalized"`
	SectionNumber       string        `json:"section_number"`
	SectionTitle        string        `json:"section_title"`
	StartDate           time.Time     `json:"start_date"`
	SyllabusURL         string        `json:"syllabus_url"`
	Term                string        `json:"term"`
	TermNormalized      string        `json:"term_normalized"`
	TermSession         string        `json:"term_session"`
	ThirdPartyLinks     []interface{} `json:"third_party_links"`
}

// CourseStatusData is the data struct returned by Registrar.GetCourseStatus and GetAllCourseStatus.
type CourseStatusData struct {
	CourseSection        string `json:"course_section"`
	SectionIDNormalized  string `json:"section_id_normalized"`
	PreviousStatus       string `json:"previous_status"`
	Status               string `json:"status"`
	StatusCodeNormalized string `json:"status_code_normalized"`
	Term                 string `json:"term"`
}

// CourseCatalogData is the data struct returned by Registrar.GetCourseCatalog.
type CourseCatalogData struct {
	ActivitiesAndCredits []struct {
		ActivityCode string `json:"activity_code"`
		Credit       string `json:"credit"`
	} `json:"activities_and_credits"`
	Corequisites           string `json:"corequisites"`
	CourseCreditConnector  string `json:"course_credit_connector"`
	CourseCreditType       string `json:"course_credit_type"`
	CourseDescription      string `json:"course_description"`
	CourseID               string `json:"course_id"`
	CourseLevel            string `json:"course_level"`
	CourseLevelDescription string `json:"course_level_description"`
	CourseNotes            string `json:"course_notes"`
	CourseNumber           string `json:"course_number"`
	CourseTitle            string `json:"course_title"`
	Crosslistings          []struct {
		CourseID  string `json:"course_id"`
		SectionID string `json:"section_id"`
		Subject   string `json:"subject"`
	} `json:"crosslistings"`
	Department              string        `json:"department"`
	DepartmentOfRecord      string        `json:"department_of_record"`
	DistributionRequirement string        `json:"distribution_requirement"`
	EasCreditFactorCode     string        `json:"eas_credit_factor_code"`
	Instructors             string        `json:"instructors"`
	Prerequisites           string        `json:"prerequisites"`
	RegisterSubgroupOne     string        `json:"register_subgroup_one"`
	RegisterSubgroupTwo     string        `json:"register_subgroup_two"`
	RequirementsMet         []interface{} `json:"requirements_met"`
	SchedulingPriority      string        `json:"scheduling_priority"`
	SchoolCode              string        `json:"school_code"`
	TermsOfferedCode        string        `json:"terms_offered_code"`
	TermsOfferedDescription string        `json:"terms_offered_description"`
}

type data struct {
	ResultData  json.RawMessage `json:"result_data"`
	ServiceMeta serviceMeta     `json:"service_meta"`
}

type parameterData struct {
	AcceptableSearchURLParametersMap map[string]string `json:"acceptable_search_url_parameters_map"`
	ActivityMap                      map[string]string `json:"activity_map"`
	AvailableTermsMap                map[string]string `json:"available_terms_map"`
	CourseLevelAtOrAboveMap          map[string]string `json:"course_level_at_or_above_map"`
	CourseLevelAtOrBelowMap          map[string]string `json:"course_level_at_or_below_map"`
	DepartmentsMap                   map[string]string `json:"departments_map"`
	EndsAtOrAfterHourMap             map[string]string `json:"ends_at_or_after_hour_map"`
	FulfillsRequiremementMap         map[string]string `json:"fulfills_requiremement_map"`
	ProgramMap                       map[string]string `json:"program_map"`
	StartsAtOrAfterHourMap           map[string]string `json:"starts_at_or_after_hour_map"`
	StartsOnDayMap                   map[string]string `json:"starts_on_day_map"`
}

type serviceMeta struct {
	CurrentPageNumber  int    `json:"current_page_number"`
	ErrorText          string `json:"error_text"`
	NextPageNumber     int    `json:"next_page_number"`
	NumberOfPages      int    `json:"number_of_pages"`
	PreviousPageNumber int    `json:"previous_page_number"`
	ResultsPerPage     int    `json:"results_per_page"`
}

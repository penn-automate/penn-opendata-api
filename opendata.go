package opendata

import (
	"encoding/json"
	"errors"
	"html"
	"net/http"
	"strconv"
)

const openDataURL = `https://esb.isc-seo.upenn.edu/8091/open_data/`

// OpenData is a struct that stores OpenData API username and password.
type OpenData struct {
	user, pass string
}

// NewOpenDataAPI generates an instance of OpenData
// with specific username and password.
func NewOpenDataAPI(username, password string) *OpenData {
	return &OpenData{user: username, pass: password}
}

func (o *OpenData) access(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization-Bearer", o.user)
	req.Header.Set("Authorization-Token", o.pass)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return http.DefaultClient.Do(req)
}

// GetRegistrar generates a Registrar instance using the current OpenData instance.
func (o *OpenData) GetRegistrar() *Registrar {
	return &Registrar{od: o}
}

// PageIterator provides an iterator for paging function.
type PageIterator struct {
	od   *OpenData
	end  bool
	err  error
	req  *http.Request
	data *data
	cur  int
}

func newErrorIter(err error) *PageIterator {
	iter := &PageIterator{end: true, err: err}
	iter.data = new(data)
	return iter
}

func newIter(od *OpenData, req *http.Request) *PageIterator {
	iter := &PageIterator{od: od, req: req, cur: 1}
	iter.data = new(data)
	return iter
}

func (i *PageIterator) error(err error) bool {
	i.err = err
	i.end = true
	return false
}

// NextPage gets the next page available.
// If the return value if true then a new page is successfully obtained.
// Otherwise, either the end of the result is reached, or an error has occurred.
func (i *PageIterator) NextPage() bool {
	if i.end {
		return false
	}

	query := i.req.URL.Query()
	query.Set("page_number", strconv.Itoa(i.cur))
	i.req.URL.RawQuery = query.Encode()

	resp, err := i.od.access(i.req)
	if err != nil {
		return i.error(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(i.data); err != nil {
		return i.error(err)
	}

	if i.data.ServiceMeta.ErrorText != "" {
		return i.error(errors.New(html.UnescapeString(i.data.ServiceMeta.ErrorText)))
	}

	i.cur = i.data.ServiceMeta.CurrentPageNumber
	i.end = i.data.ServiceMeta.NumberOfPages == i.cur

	return true
}

// Error gets the latest error generated.
func (i *PageIterator) Error() error {
	return i.err
}

// GetResult will unmarshal the raw json message into the container the user provided.
// Normally the container needs to be a slice of struct types provided by this package.
func (i *PageIterator) GetResult(container interface{}) error {
	return json.Unmarshal(i.data.ResultData, container)
}

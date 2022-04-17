package opendata

import (
	"context"
	"encoding/json"
	"errors"
	"html"
	"net/http"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const openDataURL = `https://3scale-public-prod-open-data.apps.k8s.upenn.edu/api/v1/`

// OpenData is a struct that stores OpenData API username and password.
type OpenData struct {
	client *http.Client
}

// NewOpenDataAPI generates an instance of OpenData
// with specific username and password.
func NewOpenDataAPI(clientId, clientSecret string) *OpenData {
	return &OpenData{client: (&clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     "https://sso.apps.k8s.upenn.edu/auth/realms/master/protocol/openid-connect/token",
		AuthStyle:    oauth2.AuthStyleInHeader,
	}).Client(context.TODO())}
}

func (o *OpenData) access(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	return o.client.Do(req)
}

// GetRegistrar generates a Registrar instance using the current OpenData instance.
func (o *OpenData) GetRegistrar() *Registrar {
	return &Registrar{od: o}
}

// PageIterator provides an iterator for paging function.
type PageIterator[T any] struct {
	od   *OpenData
	end  bool
	err  error
	req  *http.Request
	data *data
	cur  int
}

func newErrorIter[T any](err error) *PageIterator[T] {
	iter := &PageIterator[T]{end: true, err: err}
	iter.data = new(data)
	return iter
}

func newIter[T any](od *OpenData, req *http.Request) *PageIterator[T] {
	iter := &PageIterator[T]{od: od, req: req, cur: 1}
	iter.data = new(data)
	return iter
}

// NextPage gets the next page available.
// If the return value if true then a new page is successfully obtained, or an error has occurred.
// Otherwise, the end of the result is reached.
func (i *PageIterator[T]) NextPage() bool {
	if i.end {
		return false
	}

	query := i.req.URL.Query()
	query.Set("page_number", strconv.Itoa(i.cur))
	i.req.URL.RawQuery = query.Encode()

	resp, err := i.od.access(i.req)
	if err != nil {
		i.err = err
		return true
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(i.data); err != nil {
		i.err = err
		return true
	}

	if i.data.ServiceMeta.Error {
		i.err = errors.New(html.UnescapeString(i.data.ServiceMeta.ErrorText))
		return true
	}

	i.cur = i.data.ServiceMeta.NextPageNumber
	i.end = i.data.ServiceMeta.NumberOfPages == i.data.ServiceMeta.CurrentPageNumber
	i.err = nil

	return true
}

// GetError gets the latest error generated.
func (i *PageIterator[T]) GetError() error {
	return i.err
}

// GetResult will unmarshal the raw json message with the given index into the container the user provided.
// Normally the container needs to be a struct with types provided by this package.
func (i *PageIterator[T]) GetResult(index int) (*T, error) {
	if i.err != nil {
		return nil, i.err
	}
	ret := new(T)
	err := json.Unmarshal(i.data.ResultData[index], ret)
	return ret, err
}

// GetPageSize gets the current size of the page.
func (i *PageIterator[T]) GetPageSize() int {
	return len(i.data.ResultData)
}

// GetRawData get the raw json message with the given index
func (i *PageIterator[T]) GetRawData(index int) json.RawMessage {
	return i.data.ResultData[index]
}

package functionality_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

type mailgunMock struct {
}

func (mM mailgunMock) ApiBase() string {
}

func (mM mailgunMock) Domain() string {
}

func (mM mailgunMock) ApiKey() string {
}

func (mM mailgunMock) PublicApiKey() string {
}

func (mM mailgunMock) Client() *http.Client {
}

func (mM mailgunMock) SetClient(client *http.Client) {
}

func (mM mailgunMock) Send(m *mailgun.Message) (string, string, error) {
}

func (mM mailgunMock) ValidateEmail(email string) (mailgun.EmailVerification, error) {
}

func (mM mailgunMock) ParseAddresses(addresses ...string) ([]string, []string, error) {
}

func (mM mailgunMock) GetBounces(limit, skip int) (int, []mailgun.Bounce, error) {
}

func (mM mailgunMock) GetSingleBounce(address string) (mailgun.Bounce, error) {
}

func (mM mailgunMock) AddBounce(address, code, error string) error {
}

func (mM mailgunMock) DeleteBounce(address string) error {
}

func (mM mailgunMock) GetStats(limit int, skip int, startDate *time.Time, event ...string) (int, []mailgun.Stat, error) {
}

func (mM mailgunMock) GetTag(tag string) (mailgun.TagItem, error) {
}

func (mM mailgunMock) DeleteTag(tag string) error {
}

func (mM mailgunMock) ListTags(*mailgun.TagOptions) *mailgun.TagIterator {
}

func (mM mailgunMock) GetDomains(limit, skip int) (int, []mailgun.Domain, error) {
}

func (mM mailgunMock) GetSingleDomain(domain string) (mailgun.Domain, []mailgun.DNSRecord, []mailgun.DNSRecord, error) {
}

func (mM mailgunMock) CreateDomain(name string, smtpPassword string, spamAction string, wildcard bool) error {
}

func (mM mailgunMock) DeleteDomain(name string) error {
}

func (mM mailgunMock) GetCampaigns() (int, []mailgun.Campaign, error) {
}

func (mM mailgunMock) CreateCampaign(name, id string) error {
}

func (mM mailgunMock) UpdateCampaign(oldId, name, newId string) error {
}

func (mM mailgunMock) DeleteCampaign(id string) error {
}

func (mM mailgunMock) GetComplaints(limit, skip int) (int, []mailgun.Complaint, error) {
}

func (mM mailgunMock) GetSingleComplaint(address string) (mailgun.Complaint, error) {
}

func (mM mailgunMock) GetStoredMessage(id string) (mailgun.StoredMessage, error) {
}

func (mM mailgunMock) GetStoredMessageRaw(id string) (mailgun.StoredMessageRaw, error) {
}

func (mM mailgunMock) DeleteStoredMessage(id string) error {
}

func (mM mailgunMock) GetCredentials(limit, skip int) (int, []mailgun.Credential, error) {
}

func (mM mailgunMock) CreateCredential(login, password string) error {
}

func (mM mailgunMock) ChangeCredentialPassword(id, password string) error {
}

func (mM mailgunMock) DeleteCredential(id string) error {
}

func (mM mailgunMock) GetUnsubscribes(limit, skip int) (int, []mailgun.Unsubscription, error) {
}

func (mM mailgunMock) GetUnsubscribesByAddress(string) (int, []mailgun.Unsubscription, error) {
}

func (mM mailgunMock) Unsubscribe(address, tag string) error {
}

func (mM mailgunMock) RemoveUnsubscribe(string) error {
}

func (mM mailgunMock) RemoveUnsubscribeWithTag(a, t string) error {
}

func (mM mailgunMock) CreateComplaint(string) error {
}

func (mM mailgunMock) DeleteComplaint(string) error {
}

func (mM mailgunMock) GetRoutes(limit, skip int) (int, []mailgun.Route, error) {
}

func (mM mailgunMock) GetRouteByID(string) (mailgun.Route, error) {
}

func (mM mailgunMock) CreateRoute(mailgun.Route) (mailgun.Route, error) {
}

func (mM mailgunMock) DeleteRoute(string) error {
}

func (mM mailgunMock) UpdateRoute(string, mailgun.Route) (mailgun.Route, error) {
}

func (mM mailgunMock) GetWebhooks() (map[string]string, error) {
}

func (mM mailgunMock) CreateWebhook(kind, url string) error {
}

func (mM mailgunMock) DeleteWebhook(kind string) error {
}

func (mM mailgunMock) GetWebhookByType(kind string) (string, error) {
}

func (mM mailgunMock) UpdateWebhook(kind, url string) error {
}

func (mM mailgunMock) VerifyWebhookRequest(req *http.Request) (verified bool, err error) {
}

func (mM mailgunMock) GetLists(limit, skip int, filter string) (int, []mailgun.List, error) {
}

func (mM mailgunMock) CreateList(mailgun.List) (mailgun.List, error) {
}

func (mM mailgunMock) DeleteList(string) error {
}

func (mM mailgunMock) GetListByAddress(string) (mailgun.List, error) {
}

func (mM mailgunMock) UpdateList(string, mailgun.List) (mailgun.List, error) {
}

func (mM mailgunMock) GetMembers(limit, skip int, subfilter *bool, address string) (int, []mailgun.Member, error) {
}

func (mM mailgunMock) GetMemberByAddress(MemberAddr, listAddr string) (mailgun.Member, error) {
}

func (mM mailgunMock) CreateMember(merge bool, addr string, prototype mailgun.Member) error {
}

func (mM mailgunMock) CreateMemberList(subscribed *bool, addr string, newMembers []interface{}) error {
}

func (mM mailgunMock) UpdateMember(Member, list string, prototype mailgun.Member) (mailgun.Member, error) {
}

func (mM mailgunMock) DeleteMember(Member, list string) error {
}

func (mM mailgunMock) NewMessage(from, subject, text string, to ...string) *mailgun.Message {
}

func (mM mailgunMock) NewMIMEMessage(body io.ReadCloser, to ...string) *mailgun.Message {
}

func (mM mailgunMock) NewEventIterator() *mailgun.EventIterator {
}

func (mM mailgunMock) ListEvents(*mailgun.EventsOptions) *mailgun.EventIterator {
}

func (mM mailgunMock) PollEvents(*mailgun.EventsOptions) *mailgun.EventPoller {
}

func (mM mailgunMock) SetAPIBase(url string) {
}

func init() {
	gin.SetMode(gin.TestMode)
}

// Need to test the following:
// If key does not exist then a HTTP/400 status is returned,
//     error field is true, the update field is falsed true,
//	   and the message is the "ErrorKeyDoesNotExist" constant
// If key exists then key and value is deleted, a HTTP/200 status is returned,
//     the error field is false, the update field is true, and the message is empty
func TestHandleDeleteKey(t *testing.T) {
	var mockResponseJSON Response

	keys.set("TestHandleDeleteKey", "success")

	router := gin.New()
	router.DELETE("/keys/:key", HandleDeleteKey)

	tests := []struct {
		ExpectedResponse                         Response
		ExpectedStatusCode                       int
		ExpectedValue, ExpectedUpdateHeader, Key string
	}{
		{
			ExpectedResponse: Response{
				Error:   true,
				Message: ErrorKeyDoesNotExist,
			},
			ExpectedStatusCode:   400,
			ExpectedValue:        "",
			ExpectedUpdateHeader: "",
			Key:                  "idonotexist",
		},
		{
			ExpectedResponse: Response{
				Error:   false,
				Message: "",
			},
			ExpectedStatusCode:   200,
			ExpectedValue:        "",
			ExpectedUpdateHeader: "update",
			Key:                  "TestHandleDeleteKey",
		},
	}

	for _, test := range tests {
		mockRequest, err := http.NewRequest("DELETE", fmt.Sprintf("/keys/%s", test.Key), &bytes.Reader{})

		if err != nil {
			t.Fatal("could not create the mock request")
		}

		mockResponseWriter := httptest.NewRecorder()

		router.ServeHTTP(mockResponseWriter, mockRequest)

		bodyBytes, err := ioutil.ReadAll(mockResponseWriter.Body)

		if err != nil {
			t.Error("Could not read the response body into bytes")

			continue
		}

		err = json.Unmarshal(bodyBytes, &mockResponseJSON)

		if err != nil {
			t.Error("Could not parse the response body bytes into JSON")

			continue
		}

		if value, _ := keys.get(test.Key); mockResponseWriter.Code != test.ExpectedStatusCode || mockResponseJSON != test.ExpectedResponse || value != test.ExpectedValue || mockResponseWriter.Header().Get("update") != test.ExpectedUpdateHeader {
			t.Errorf(
				`HandleDeleteKey(context) = Status Code: HTTP/%d, Response: "%v", keys[%s] = "%s", and the update header = "%s"; expected: HTTP/%d, Response: "%v", keys[%s] = "%s", and the update header = "%s"`,
				mockResponseWriter.Code,
				mockResponseJSON,
				test.Key,
				value,
				mockResponseWriter.Header().Get("update"),
				test.ExpectedStatusCode,
				test.ExpectedResponse,
				test.Key,
				test.ExpectedValue,
				test.ExpectedUpdateHeader,
			)
		}
	}
}

// Need to test the following:
// If key does not exist then a HTTP/400 status is returned,
//     error field is true, and the message is the "ErrorKeyDoesNotExist" constant
// If key exists then the value for the provided key is returned, a HTTP/200 status is returned,
//     the error field is false, and the message is the value for the key
func TestHandleGetKey(t *testing.T) {
	var mockResponseJSON Response

	keys.set("TestHandleGetKey", "success")

	router := gin.New()
	router.GET("/keys/:key", HandleGetKey)

	tests := []struct {
		ExpectedResponse   Response
		ExpectedStatusCode int
		Key                string
	}{
		{
			ExpectedResponse: Response{
				Error:   true,
				Message: ErrorKeyDoesNotExist,
			},
			ExpectedStatusCode: 400,
			Key:                "DoNotTestHandleGetKey",
		},
		{
			ExpectedResponse: Response{
				Error:   false,
				Message: "success",
			},
			ExpectedStatusCode: 200,
			Key:                "TestHandleGetKey",
		},
	}

	for _, test := range tests {
		mockRequest, err := http.NewRequest("GET", fmt.Sprintf("/keys/%s", test.Key), &bytes.Reader{})

		if err != nil {
			t.Fatal("could not create the mock request")
		}

		mockResponseWriter := httptest.NewRecorder()

		router.ServeHTTP(mockResponseWriter, mockRequest)

		bodyBytes, err := ioutil.ReadAll(mockResponseWriter.Body)

		if err != nil {
			t.Error("Could not read the response body into bytes")

			continue
		}

		err = json.Unmarshal(bodyBytes, &mockResponseJSON)

		if err != nil {
			t.Error("Could not parse the response body bytes into JSON")

			continue
		}

		if mockResponseWriter.Code != test.ExpectedStatusCode || mockResponseJSON != test.ExpectedResponse {
			t.Errorf(
				`HandleGetKey(context) = Status Code: HTTP/%d and Response: "%v", expected HTTP/%d and Response: "%v"`,
				mockResponseWriter.Code,
				mockResponseJSON,
				test.ExpectedStatusCode,
				test.ExpectedResponse,
			)
		}
	}
}

// Need to test the following:
// If any of the provided keys do not exist then they are not found in the returned map
//     in the message field, the error field is false, and a HTTP/200 status is returned
// If key exists then the value for the provided key can be found in the returned map in
//     the message field, the error field is false and a HTTP/200 status is returned
func TestHandleGetManyKeys(t *testing.T) {
	CheckResponseAgainstExpected := func(response map[string]interface{}, expected map[string]string) bool {
		for key, value := range expected {
			if _, exists := response[key]; !exists {
				return false
			}

			if value != response[key].(string) {
				return false
			}
		}

		return true
	}

	keys.set("TestHandleGetKey", "success")
	keys.set("TestHandleGetAnotherKey", "good")
	keys.set("TestHandleGetAnotherAnotherKey", "great")

	router := gin.New()
	router.POST("/keys", HandleGetManyKeys)

	tests := []struct {
		ExpectedResponse   Response
		ExpectedStatusCode int
		Request            RequestMany
	}{
		{
			ExpectedResponse: Response{
				Error: false,
				Message: map[string]string{
					"TestHandleGetKey":        "success",
					"TestHandleGetAnotherKey": "good",
				},
			},
			ExpectedStatusCode: 200,
			Request: RequestMany{
				Keys: []string{
					"TestHandleGetKey",
					"TestHandleGetAnotherKey",
				},
			},
		},
		{
			ExpectedResponse: Response{
				Error: false,
				Message: map[string]string{
					"TestHandleGetAnotherAnotherKey": "great",
				},
			},
			ExpectedStatusCode: 200,
			Request: RequestMany{
				Keys: []string{
					"TestHandleGetAnotherAnotherKey",
					"TestNonexistantKey",
				},
			},
		},
	}

	for _, test := range tests {
		reader, writer := io.Pipe()

		mockRequest, err := http.NewRequest("POST", "/keys", reader)

		if err != nil {
			t.Fatal("could not create the mock request")
		}

		go func() { json.NewEncoder(writer).Encode(test.Request) }()

		mockResponseWriter := httptest.NewRecorder()

		router.ServeHTTP(mockResponseWriter, mockRequest)

		var mockResponseJSON Response

		err = json.NewDecoder(mockResponseWriter.Body).Decode(&mockResponseJSON)

		if err != nil {
			t.Error("Could not decode the response body into json")

			continue
		}

		if mockResponseWriter.Code != test.ExpectedStatusCode || mockResponseJSON.Error != test.ExpectedResponse.Error || !CheckResponseAgainstExpected(mockResponseJSON.Message.(map[string]interface{}), test.ExpectedResponse.Message.(map[string]string)) {
			t.Errorf(
				`HandleGetKey(context) = Status Code: HTTP/%d and Response: "%v", expected HTTP/%d and Response: "%v"`,
				mockResponseWriter.Code,
				mockResponseJSON,
				test.ExpectedStatusCode,
				test.ExpectedResponse,
			)
		}
	}
}

// Need to test the following:
// If key already exists then a HTTP/400 status is returned,
//     error field is true, the update field is false, and the message is the "ErrorKeyAlreadyExists" constant
// If key has any character which cannot be put into a URL with URL encoding then a HTTP/400 status is returned,
//     error field is true, the update field is false, and the message is the "ErrorInvalidKey" constant
// If key does not exist then key and value pair is created, a HTTP/201 status is returned,
//     the error field is false, the update field is true, and the message field is empty
func TestHandlePostKey(t *testing.T) {
	var mockResponseJSON Response

	keys.set("TestHandlePostKey", "success")

	router := gin.New()
	router.POST("/keys", HandlePostKey)

	tests := []struct {
		ExpectedResponse                                Response
		ExpectedStatusCode                              int
		ExpectedValue, ExpectedUpdateHeader, Key, Value string
	}{
		{
			ExpectedResponse: Response{
				Error:   true,
				Message: ErrorKeyAlreadyExists,
			},
			ExpectedStatusCode:   400,
			ExpectedValue:        "success",
			ExpectedUpdateHeader: "",
			Key:                  "TestHandlePostKey",
			Value:                "failure",
		},
		{
			ExpectedResponse: Response{
				Error:   true,
				Message: ErrorInvalidKey,
			},
			ExpectedStatusCode:   400,
			ExpectedValue:        "",
			ExpectedUpdateHeader: "",
			Key:                  "TestHandlePostKeyFailure?",
			Value:                "success",
		},
		{
			ExpectedResponse: Response{
				Error:   false,
				Message: "",
			},
			ExpectedStatusCode:   201,
			ExpectedValue:        "success",
			ExpectedUpdateHeader: "update",
			Key:                  "TestHandlePostKeyFailure",
			Value:                "success",
		},
	}

	for _, test := range tests {
		requestBytes, err := json.Marshal(RequestSingle{test.Key, test.Value})

		if err != nil {
			t.Error("Could not read the response body into bytes")

			continue
		}

		mockRequest, err := http.NewRequest("POST", "/keys", bytes.NewBuffer(requestBytes))

		if err != nil {
			t.Fatal("could not create the mock request")
		}

		mockResponseWriter := httptest.NewRecorder()

		router.ServeHTTP(mockResponseWriter, mockRequest)

		bodyBytes, err := ioutil.ReadAll(mockResponseWriter.Body)

		if err != nil {
			t.Error("Could not read the response body into bytes")

			continue
		}

		err = json.Unmarshal(bodyBytes, &mockResponseJSON)

		if err != nil {
			t.Error("Could not parse the response body bytes into JSON")

			continue
		}

		if value, _ := keys.get(test.Key); mockResponseWriter.Code != test.ExpectedStatusCode || mockResponseJSON != test.ExpectedResponse || value != test.ExpectedValue || mockResponseWriter.Header().Get("update") != test.ExpectedUpdateHeader {
			t.Errorf(
				`HandleDeleteKey(context) = Status Code: HTTP/%d, Response: "%v", keys[%s] = "%s", and the update header = "%s"; expected: HTTP/%d, Response: "%v", keys[%s] = "%s", and the update header = "%s"`,
				mockResponseWriter.Code,
				mockResponseJSON,
				test.Key,
				value,
				mockResponseWriter.Header().Get("update"),
				test.ExpectedStatusCode,
				test.ExpectedResponse,
				test.Key,
				test.ExpectedValue,
				test.ExpectedUpdateHeader,
			)
		}
	}
}

// Need to test the following:
// If key already exists then a HTTP/400 status is returned,
//     error field is true, the update field is false, and the message is the "ErrorKeyAlreadyExists" constant
// If key exists then the value for the key is updated, a HTTP/200 status is returned,
//     the error field is false, the update field is true, and the message is empty
func TestHandlePutKey(t *testing.T) {
	var mockResponseJSON Response

	keys.set("TestHandlePutKey", "success")

	router := gin.New()
	router.PUT("/keys", HandlePutKey)

	tests := []struct {
		ExpectedResponse                                Response
		ExpectedStatusCode                              int
		ExpectedValue, ExpectedUpdateHeader, Key, Value string
	}{
		{
			ExpectedResponse: Response{
				Error:   true,
				Message: ErrorKeyDoesNotExist,
			},
			ExpectedStatusCode:   400,
			ExpectedValue:        "",
			ExpectedUpdateHeader: "",
			Key:                  "TestHandlePutKeyFailure",
			Value:                "success",
		},
		{
			ExpectedResponse: Response{
				Error:   false,
				Message: "",
			},
			ExpectedStatusCode:   200,
			ExpectedValue:        "success",
			ExpectedUpdateHeader: "update",
			Key:                  "TestHandlePutKey",
			Value:                "success",
		},
	}

	for _, test := range tests {
		requestBytes, err := json.Marshal(RequestSingle{test.Key, test.Value})

		if err != nil {
			t.Error("Could not read the response body into bytes")

			continue
		}

		mockRequest, err := http.NewRequest("PUT", "/keys", bytes.NewBuffer(requestBytes))

		if err != nil {
			t.Fatal("could not create the mock request")
		}

		mockResponseWriter := httptest.NewRecorder()

		router.ServeHTTP(mockResponseWriter, mockRequest)

		bodyBytes, err := ioutil.ReadAll(mockResponseWriter.Body)

		if err != nil {
			t.Error("Could not read the response body into bytes")

			continue
		}

		err = json.Unmarshal(bodyBytes, &mockResponseJSON)

		if err != nil {
			t.Error("Could not parse the response body bytes into JSON")

			continue
		}

		if value, _ := keys.get(test.Key); mockResponseWriter.Code != test.ExpectedStatusCode || mockResponseJSON != test.ExpectedResponse || value != test.ExpectedValue || mockResponseWriter.Header().Get("update") != test.ExpectedUpdateHeader {
			t.Errorf(
				`HandlePutKey(context) = Status Code: HTTP/%d, Response: "%v", keys[%s] = "%s", and the update header = "%s"; expected: HTTP/%d, Response: "%v", keys[%s] = "%s", and the update header = "%s"`,
				mockResponseWriter.Code,
				mockResponseJSON,
				test.Key,
				value,
				mockResponseWriter.Header().Get("update"),
				test.ExpectedStatusCode,
				test.ExpectedResponse,
				test.Key,
				test.ExpectedValue,
				test.ExpectedUpdateHeader,
			)
		}
	}
}

// Need to test the following:
// If the route does not exist then a HTTP/404 status is returned,
//     error field is true, the update field is false, and the
//	   message is the "ErrorBadRequest" constant
// If the route exists then the route responds as expected
func TestKeyManagingRouter(t *testing.T) {
	var router *gin.Engine
	var mockResponseJSON Response

	tests := []struct {
		ExpectedStatusCode   int
		ExpectedResponse     Response
		ExpectedUpdateHeader string
		HandlerFuncs         map[string]func(c *gin.Context)
	}{
		{
			ExpectedStatusCode: 404,
			ExpectedResponse: Response{
				Error:   true,
				Message: ErrorBadRequest,
			},
			HandlerFuncs: map[string]func(c *gin.Context){},
		},
		{
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Error:   false,
				Message: "",
			},
			HandlerFuncs: map[string]func(c *gin.Context){
				"/test": func(c *gin.Context) {
					c.JSON(200, Response{false, ""})
				},
			},
		},
	}

	for _, test := range tests {
		router = NewKeyManagingRouter()

		for route, handlerFunc := range test.HandlerFuncs {
			router.GET(route, handlerFunc)
		}

		mockRequest, err := http.NewRequest("GET", "/test", &bytes.Reader{})

		if err != nil {
			t.Fatal("could not create the mock request")
		}

		mockResponseWriter := httptest.NewRecorder()

		router.ServeHTTP(mockResponseWriter, mockRequest)

		bodyBytes, err := ioutil.ReadAll(mockResponseWriter.Body)

		if err != nil {
			t.Error("Could not read the response body into bytes")

			continue
		}

		err = json.Unmarshal(bodyBytes, &mockResponseJSON)

		if err != nil {
			t.Error("Could not parse the response body bytes into JSON")

			continue
		}

		if mockResponseWriter.Code != test.ExpectedStatusCode || mockResponseJSON != test.ExpectedResponse {
			t.Errorf(
				`router.ServeHTTP(mockResponseWriter, mockRequest) = Status Code: HTTP/%d and Response: "%v"; expected HTTP/%d and Response: "%v"`,
				mockResponseWriter.Code,
				mockResponseJSON,
				test.ExpectedStatusCode,
				test.ExpectedResponse,
			)
		}
	}
}

// Need to test the following:
// If the reader provided does not provide valid JSON then an error is returned
// If the reader provided does provide valid JSON then the expected key/value pairs exist in keys.keys
func TestLoadKeyDataKeys(t *testing.T) {
	tests := []struct {
		err    bool
		pairs  map[string]string
		reader io.Reader
	}{
		{
			err:    true,
			pairs:  map[string]string{},
			reader: strings.NewReader(`{"test":"not passed"`),
		},
		{
			err: false,
			pairs: map[string]string{
				"test": "passed",
			},
			reader: strings.NewReader(`{"test":"passed"}`),
		},
	}

	for _, test := range tests {
		err := LoadKeyDataKeys(test.reader)

		if expectedErr := err != nil; expectedErr != test.err {
			if test.err {
				t.Error("Expected an error")
			} else {
				t.Error("Did not expected error:", err)
			}

			continue
		}

		for key, expectedValue := range test.pairs {
			if value, exists := keys.get(key); !exists || value != expectedValue {
				t.Errorf(
					`LoadKeyDataKeys(test.reader) = %v: expected "%v"`,
					keys.cloneKeys(), test.pairs,
				)
			}
		}
	}
}

// Need to test the following:
// If the reader provided works,
//     then the key/value pairs that existed in keys.keys were written to the writer,
//     otherwise an error is returned
func TestUnloadKeyDataKeys(t *testing.T) {
	keys.keys = make(map[string]string)

	tests := []struct {
		expectedBytes []byte
		err           bool
		pairs         map[string]string
		writer        io.Writer
	}{
		{
			expectedBytes: []byte("{\"test\":\"passed\"}\n"),
			err:           false,
			pairs: map[string]string{
				"test": "passed",
			},
			writer: &bytes.Buffer{},
		},
	}

	for _, test := range tests {
		for key, value := range test.pairs {
			keys.set(key, value)
		}

		err := UnloadKeyDataKeys(test.writer)

		if expectedErr := err != nil; expectedErr != test.err {
			if test.err {
				t.Error("Expected an error")
			} else {
				t.Error("Did not expected error:", err)
			}

			continue
		}

		if (test.writer.(*bytes.Buffer)).String() != string(test.expectedBytes) {
			t.Errorf(
				`UnloadKeyDataKeys(test.writer) = "%s": expected "%s"`,
				(test.writer.(*bytes.Buffer)).Bytes(), test.expectedBytes,
			)
		}
	}
}

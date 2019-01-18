package functionality

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v3"

	"github.com/gin-gonic/gin"
)

type PwnInfo struct {
	Name         string   `json:"Name"`
	Title        string   `json:"Title"`
	Domain       string   `json:"Domain"`
	BreachDate   string   `json:"BreachDate"`
	AddedDate    string   `json:"AddedDate"`
	ModifiedDate string   `json:"ModifiedDate"`
	PwnCount     int64    `json:"PwnCount"`
	Description  string   `json:"Description"`
	LogoPath     string   `json:"LogoPath"`
	DataClasses  []string `json:"DataClasses"`
	IsVerified   bool     `json:"IsVerified"`
	IsFabricated bool     `json:"IsFabricated"`
	IsSensitive  bool     `json:"IsSensitive"`
	IsRetired    bool     `json:"IsRetired"`
	IsSpamList   bool     `json:"IsSpamList"`
}

type mailgunInfo struct {
	Domain        string `json:"domain"`
	PrivateAPIKey string `json:"private_api_key`
	PublicAPIKey  string `json:"public_api_key`
}

var (
	ErrNoPwns error = errors.New("there is no pwnage for the email provided")
	mg        mailgun.Mailgun
)

func init() {
	fileLocation, exists := os.LookupEnv("mailgunFile")

	var (
		err         error
		mailgunFile *os.File
		mailgunJSON mailgunInfo
	)

	if exists {
		mailgunFile, err = os.Open(fileLocation)
	} else {
		mailgunFile, err = os.Open(fileLocation)
	}

	json.NewDecoder(mailgunFile).Decode(&mailgunJSON)

	if exists {
		err = os.Remove(fileLocation)

		if err != nil {
			panic(err)
		}
	}

	mg := mailgun.NewMailgun(mailgunJSON.Domain, mailgunJSON.PrivateAPIKey, mailgunJSON.PublicAPIKey)
}

// InitializeMailgunWithMailgun is used for mocking the mailgun client during tests
func InitializeMailgunWithMailgun(nmg mailgun.Mailgun) {
	mg = nmg
}

// InitializeMailgunWithJSON is used for initializing the mailgun client for the package
func InitializeMailgunWithJSON(reader io.Reader) error {
	var mailgunJSON mailgunInfo

	err := json.NewDecoder(reader).Decode(&mailgunJSON)

	if err != nil {
		return err
	}

	mg = mailgun.NewMailgun(mailgunJSON.Domain, mailgunJSON.PrivateAPIKey, mailgunJSON.PublicAPIKey)

	return nil
}

func getPwnageForEmail(email string) ([]PwnInfo, error) {
	return getPwnageForEmailWithClient(email, http.DefaultClient)
}

func getPwnageForEmailWithClient(email string, client *http.Client) (pwnInfo []PwnInfo, err error) {
	pwnageInfoRequest, err := http.NewRequest("GET", fmt.Sprintf("https://haveibeenpwned.com/api/v2/breachedaccount/%s", url.QueryEscape(email)), nil)

	if err != nil {
		return pwnInfo, err
	}

	pwnageInfoRequest.Header.Set("User-Agent", "RJ-And-Friends-Nightly-Pwnage-Checker")

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)

	defer cancel()

	pwnageInfoRequest = pwnageInfoRequest.WithContext(ctx)

	responseErrChan := make(chan error)

	go func() {
		resp, err := client.Do(pwnageInfoRequest)

		if err != nil {
			responseErrChan <- err

			return
		}

		responseBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			responseErrChan <- err

			return
		}

		if len(responseBytes) == 0 {
			responseErrChan <- ErrNoPwns

			return
		}

		err = json.Unmarshal(responseBytes, &pwnInfo)

		responseErrChan <- err
	}()

	select {
	case <-ctx.Done():
		return pwnInfo, ctx.Err()
	case err = <-responseErrChan:
		if err != nil {
			if err == ErrNoPwns {
				return pwnInfo, err
			}

			return pwnInfo, err
		}
	}

	return pwnInfo, nil
}

func notifyEmailOfPwnage(email, title, body string) error {
	content := mailgun.NewMessage("robot@mail.therileyjohnson.com", title, body, email)

	_, _, err := mg.Send(content)

	return err
}

func notifyOfPwnage(email, phone string, alwaysNotify bool) error {
	isPwned := false
	pwnInfo, err := getPwnageForEmail(email)

	if err != nil && err != ErrNoPwns {
		return err
	}

	isPwned = err != ErrNoPwns

	if isPwned {
		err = notifyEmailOfPwnage(email, "YOU HAVE BEEN PWNED :(", "")

		if err != nil {
			return err
		}

		if phone != "" {
			return notifyPhoneOfPwnage(phone, "YOU HAVE BEEN PWNED :(")
		}
	} else if alwaysNotify {
		err = notifyEmailOfPwnage(email, "YOU HAVE NOT BEEN PWNED :)", "")

		if err != nil {
			return err
		}

		if phone != "" {
			return notifyPhoneOfPwnage(phone, "YOU HAVE NOT BEEN PWNED :)")
		}
	}

	return nil
}

func NotifyOfPwnage(c *gin.Context) {
	notifyList := struct {
		Contacts []struct {
			Email string `json:"email"`
			Phone string `json:"phone"`
		}
	}{}

	json.NewDecoder(c.Request.Body).Decode(&notifyList)

	for _, contact := range notifyList.Contacts {
		notifyOfPwnage(contact.Email, contact.Phone, true)

		// Need to wait between every function call because
		// the API only allows one request every 1500ms
		time.Sleep(1500 * time.Millisecond)
	}
}

func AddToPwnageCheck(c *gin.Context) {

}

func DeleteFromPwnageCheck(c *gin.Context) {

}

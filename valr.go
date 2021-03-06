package valr

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	httpClient   *http.Client
	baseURL      string
	apiKeyID     string
	apiKeySecret string
	debug        bool
}

type ApiKey struct {
	API_KEY    string `yaml:"API_KEY"`
	API_SECRET string `yaml:"API_SECRET"`
}

const defaultBaseURL = "https://api.valr.com"
const defaultTimeout = 10 * time.Second

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: defaultTimeout},
		baseURL:    defaultBaseURL,
		debug:      false,
	}
}

func (cl *Client) SetAuth(apiKeyID, apiKeySecret string) error {
	if apiKeyID == "" || apiKeySecret == "" {
		return errors.New("valr: no credentials provided")
	}
	cl.apiKeyID = apiKeyID
	cl.apiKeySecret = apiKeySecret
	return nil
}

func (cl *Client) LoadAuthFile(filepath string) error {
	if filepath == "" {
		return errors.New("filename cannot be blank")
	}

	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	keys := ApiKey{}
	err = yaml.Unmarshal(yamlFile, &keys)
	if err != nil {
		return err
	}

	cl.SetAuth(keys.API_KEY, keys.API_SECRET)

	return nil
}

func (cl *Client) SetBaseURL(baseURL string) {
	cl.baseURL = strings.TrimRight(baseURL, "/")
}

func (cl *Client) SetDebug(debug bool) {
	cl.debug = debug
}

func (cl *Client) SendRequest(method string, path string) ([]byte, error) {
	var url = cl.baseURL + path
	req, err := http.NewRequest(method, url, nil)

	if cl.debug {
		log.Printf("valr: Call: %s %s", method, path)
		log.Printf("valr: Request: %#v", req)
	}

	if err != nil {
		return []byte{}, err
	}

	sig, tim := signRequest(cl.apiKeySecret, time.Now(), method, path, "")

	req.Header.Add("X-VALR-API-KEY", cl.apiKeyID)
	req.Header.Add("X-VALR-SIGNATURE", sig)
	req.Header.Add("X-VALR-TIMESTAMP", tim)

	res, err := cl.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	if cl.debug {
		pretty, err := PrettyPrintBytes(body)
		if err != nil {
			return body, err
		}
		log.Printf("valr: %v", pretty)
	}

	return body, nil
}

/*func (cl *Client) do(ctx context.Context, method, path string, req, res interface{}, auth bool) error {

	url := cl.baseURL + "/" + strings.TrimLeft(path, "/")

	if cl.debug {
		log.Printf("valr: Call: %s %s", method, path)
		log.Printf("valr: Request: %#v", req)
	}

	var contentType string
	var body io.Reader
	if req != nil {
		values := makeURLValues(req)
		if strings.Contains(path, "{id}") {
			url = strings.Replace(url, "{id}", values.Get("id"), -1)
			values.Del("id")
		}
		if method == http.MethodGet {
			url = url + "?" + values.Encode()
		} else {
			body = strings.NewReader(values.Encode())
			contentType = "application/x-www-form-urlencoded"
		}
	}

	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	httpReq = httpReq.WithContext(ctx)
	httpReq.Header.Set("User-Agent", makeUserAgent())
	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if auth {
		httpReq.SetBasicAuth(cl.apiKeyID, cl.apiKeySecret)
	}

	if method != http.MethodGet {
		httpReq.Header.Set("content-type", "application/x-www-form-urlencoded")
	}

	httpRes, err := cl.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	body = httpRes.Body
	if cl.debug {
		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Printf("luno: Error reading response body: %v", err)
		} else {
			log.Printf("Response: %s", string(b))
		}
		body = bytes.NewReader(b)
	}

	if httpRes.StatusCode == http.StatusTooManyRequests {
		return errors.New("luno: too many requests")
	}

	if httpRes.StatusCode != http.StatusOK {
		var e Error
		err := json.NewDecoder(body).Decode(&e)
		if err != nil {
			return fmt.Errorf("luno: error decoding response (%d %s)",
				httpRes.StatusCode, http.StatusText(httpRes.StatusCode))
		}
		return e
	}

	return json.NewDecoder(body).Decode(res)
}*/

func signRequest(apiSecret string, timestamp time.Time, verb string, path string, body string) (string, string) {
	// Create a new Keyed-Hash Message Authentication Code (HMAC) using SHA512 and API Secret
	mac := hmac.New(sha512.New, []byte(apiSecret))
	// Convert timestamp to nanoseconds then divide by 1000000 to get the milliseconds
	timestampString := strconv.FormatInt(timestamp.UnixNano()/1000000, 10)
	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	// Gets the byte hash from HMAC and converts it into a hex string
	return hex.EncodeToString(mac.Sum(nil)), timestampString
}

func makeUserAgent() string {
	return fmt.Sprintf("ValrGoSDK/%s %s %s %s",
		Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

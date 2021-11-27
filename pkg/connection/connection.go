package connection

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Connection struct {
	url string
	login string
	password string
	skipVerifySsl bool
	client *http.Client
	connected bool
}

func NewConnection(url string, login string, password string, skipVerifySsl bool) *Connection {
	return &Connection{url: url, login: login, password: password, skipVerifySsl: skipVerifySsl}
}

type startRequest struct {
	Version string `url:"version"`
	Agent string `url:"agent"`
	Login string `url:"login"`
	Type string `url:"type"`
}

type startResponse struct {
	Retcode string `json:"retcode"`
	VersionAccess string `json:"version_access"`
	SrvRand       string `json:"srv_rand"`
}

type answerRequest struct {
	SrvRandAnswer string `url:"srv_rand_answer"`
	CliRand string `url:"cli_rand"`
}

type answerResponse struct {
	Retcode string `json:"retcode"`
	CliRandAnswer string `json:"cli_rand_answer"`
}

type retcodeResponse struct {
	Retcode string `json:"retcode"`
}

func (c *Connection) getClient() *http.Client {
	if c.client == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: c.skipVerifySsl},
		}
		c.client = &http.Client{Transport: tr}
	}

	return c.client
}

func (c *Connection) Get(url string) ([]byte, error)  {
	if err := c.connect(); err != nil {
		panic("Cannot connect to the server")
	}
	return c.processGetQuery(url)
}

func (c *Connection) processGetQuery(url string) ([]byte, error) {
	resp, err := c.getClient().Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Bad response: url: '%s'; status '%s'; response: '%s'", url, resp.Status, string(body)))
	}


	var response retcodeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Retcode != "0 Done" {
		return nil, errors.New(fmt.Sprintf("Request failed. Retcode: `%s`", response.Retcode))
	}

	return body, nil
}

func (c *Connection) connect() error {
	if c.connected == true {
		return nil
	}

	params := startRequest{"1000", "mt5toolkit", c.login, "manager"}
	q, err := query.Values(params)
	if err != nil {
		log.Fatal(err)
	}
	payload, err := c.processGetQuery(c.url + "/api/auth/start?" + q.Encode())
	if err != nil {
		log.Fatal(err)
	}
	startResp := startResponse{}
	if nil != json.Unmarshal(payload, &startResp) {
		log.Fatal(err)
	}

	mt5Hash := encodeMd5(append(encodeUtf16leMd5(c.password), []byte("WebAPI")...))
	srvRand, _ := hex.DecodeString(startResp.SrvRand)
	cliRand := make([]byte, 16)
	rand.Read(cliRand)

	answerParams := answerRequest{SrvRandAnswer: hex.EncodeToString(encodeMd5(append(mt5Hash, srvRand...))), CliRand: hex.EncodeToString(cliRand)}
	q, err = query.Values(answerParams)
	if err != nil {
		log.Fatal(err)
	}
	payload, err = c.processGetQuery(c.url + "/api/auth/answer?" + q.Encode())
	if err != nil {
		log.Fatal(err)
	}
	answerResp := answerResponse{}
	if nil != json.Unmarshal(payload, &answerResp) {
		log.Fatal(err)
	}

	if answerResp.CliRandAnswer != hex.EncodeToString(encodeMd5(append(mt5Hash, cliRand...))) {
		log.Fatal(err)
	}

	c.connected = true

	return nil
}

func (c *Connection) Ping() {
	if err := c.connect(); err != nil {
		panic("Cannot connect to the server")
	}
	_, err := c.processGetQuery(c.url + "/api/test/access")
	if err != nil {
		panic("Disconnected")
	}
}

func encodeMd5(in []byte) []byte  {
	h := md5.New()
	h.Write(in)
	return h.Sum(nil)
}

func encodeUtf16leMd5(s string) []byte {
	enc := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	hasher := md5.New()
	t := transform.NewWriter(hasher, enc)
	t.Write([]byte(s))

	return hasher.Sum(nil)
}

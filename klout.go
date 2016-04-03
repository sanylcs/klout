// klout implement most Klout API using Golang. Klout oauth2 is not available.
package klout

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var partnerKey string
var baseUrl = "http://api.klout.com/v2"
var identityUrl = fmt.Sprint(baseUrl, "/identity.json")
var userUrl = fmt.Sprint(baseUrl, "/user.json")

// Set Partner Key.
func PartnerKey(key string) {
	partnerKey = key
}

func transformId(id, path string, urlq url.Values) (*Identity, error) {
	var v Identity
	if path == "" {
		return nil, errors.New("empty path")
	}
	if id == "" {
		path = fmt.Sprint(identityUrl, "/", path)
	} else {
		path = fmt.Sprint(identityUrl, "/", path, "/", id)
	}
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	urlq.Add("key", partnerKey)
	u.RawQuery = urlq.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	err = d.Decode(v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// KIdFrmTwId get Klout Id from the provided Twitter Id.
func KIdFrmTwId(twId string) (*Identity, error) {
	return transformId(twId, "tw", url.Values{})
}

// KIdFrmTwName get Klout Id from the provided Twitter Screen Name.
func KIdFrmTwName(name string) (*Identity, error) {
	u := url.Values{}
	u.Add("screenName", name)
	return transformId("", "twitter", u)
}

// KIdFrmGpId get Klout Id from the provided Google plus Id.
func KIdFrmGpId(gpId string) (*Identity, error) {
	return transformId(gpId, "gp", url.Values{})
}

// KIdFrmGpId get Klout Id from the provided Instagram Id.
func KIdFrmIgId(ig string) (*Identity, error) {
	return transformId(ig, "ig", url.Values{})
}

// TwIdFrmKId get Twitter Id from the provided Klout Id.
func TwIdFrmKId(kid string) (*Identity, error) {
	return transformId(fmt.Sprint(kid, "/tw"), "klout", url.Values{})
}

func userMethods(kid, path string, v interface{}) error {
	if kid == "" {
		return errors.New("empty Klout Id")
	}
	if path != "" {
		path = fmt.Sprint(userUrl, "/", kid, "/", path)
	} else {
		path = fmt.Sprint(userUrl, "/", kid)
	}
	u, err := url.Parse(path)
	if err != nil {
		return err
	}
	urlq := url.Values{}
	urlq.Add("key", partnerKey)
	u.RawQuery = urlq.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	err = d.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(kid string) (*User, error) {
	var v User
	if err := userMethods(kid, "", &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// Get Klout score. kid is the Klout Id of the target.
func GetScore(kid string) (*Score, error) {
	var v Score
	if err := userMethods(kid, "score", &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func GetTopics(kid string) (Topics, error) {
	var v Topics
	if err := userMethods(kid, "topics", v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetInfluence(kid string) (*Influence, error) {
	var v Influence
	if err := userMethods(kid, "influence", v); err != nil {
		return nil, err
	}
	return &v, nil
}

//
// This file generated by rdl 1.4.10
//

package zts

import (
	"bytes"
	"encoding/json"
	"fmt"
	rdl "github.com/ardielle/ardielle-go/rdl"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var _ = json.Marshal
var _ = fmt.Printf
var _ = rdl.BaseTypeAny
var _ = ioutil.NopCloser

type ZTSClient struct {
	URL         string
	Transport   http.RoundTripper
	CredsHeader *string
	CredsToken  *string
	Timeout     time.Duration
}

// NewClient creates and returns a new HTTP client object for the ZTS service
func NewClient(url string, transport http.RoundTripper) ZTSClient {
	return ZTSClient{url, transport, nil, nil, 0}
}

// AddCredentials adds the credentials to the client for subsequent requests.
func (client *ZTSClient) AddCredentials(header string, token string) {
	client.CredsHeader = &header
	client.CredsToken = &token
}

func (client ZTSClient) getClient() *http.Client {
	var c *http.Client
	if client.Transport != nil {
		c = &http.Client{Transport: client.Transport}
	} else {
		c = &http.Client{}
	}
	if client.Timeout > 0 {
		c.Timeout = client.Timeout
	}
	return c
}

func (client ZTSClient) addAuthHeader(req *http.Request) {
	if client.CredsHeader != nil && client.CredsToken != nil {
		if strings.HasPrefix(*client.CredsHeader, "Cookie.") {
			req.Header.Add("Cookie", (*client.CredsHeader)[7:]+"="+*client.CredsToken)
		} else {
			req.Header.Add(*client.CredsHeader, *client.CredsToken)
		}
	}
}

func (client ZTSClient) httpGet(url string, headers map[string]string) (*http.Response, error) {
	hclient := client.getClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client.addAuthHeader(req)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return hclient.Do(req)
}

func (client ZTSClient) httpDelete(url string, headers map[string]string) (*http.Response, error) {
	hclient := client.getClient()
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	client.addAuthHeader(req)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return hclient.Do(req)
}

func (client ZTSClient) httpPut(url string, headers map[string]string, body []byte) (*http.Response, error) {
	var contentReader io.Reader
	if body != nil {
		contentReader = bytes.NewReader(body)
	}
	hclient := client.getClient()
	req, err := http.NewRequest("PUT", url, contentReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/json")
	client.addAuthHeader(req)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return hclient.Do(req)
}

func (client ZTSClient) httpPost(url string, headers map[string]string, body []byte) (*http.Response, error) {
	var contentReader io.Reader
	if body != nil {
		contentReader = bytes.NewReader(body)
	}
	hclient := client.getClient()
	req, err := http.NewRequest("POST", url, contentReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/json")
	client.addAuthHeader(req)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return hclient.Do(req)
}

func (client ZTSClient) httpPatch(url string, headers map[string]string, body []byte) (*http.Response, error) {
	var contentReader io.Reader
	if body != nil {
		contentReader = bytes.NewReader(body)
	}
	hclient := client.getClient()
	req, err := http.NewRequest("PATCH", url, contentReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/json")
	client.addAuthHeader(req)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return hclient.Do(req)
}

func (client ZTSClient) httpOptions(url string, headers map[string]string, body []byte) (*http.Response, error) {
	var contentReader io.Reader = nil
	if body != nil {
		contentReader = bytes.NewReader(body)
	}
	hclient := client.getClient()
	req, err := http.NewRequest("OPTIONS", url, contentReader)
	if err != nil {
		return nil, err
	}
	if contentReader != nil {
		req.Header.Add("Content-type", "application/json")
	}
	client.addAuthHeader(req)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return hclient.Do(req)
}

func encodeStringParam(name string, val string, def string) string {
	if val == def {
		return ""
	}
	return "&" + name + "=" + url.QueryEscape(val)
}
func encodeBoolParam(name string, b bool, def bool) string {
	if b == def {
		return ""
	}
	return fmt.Sprintf("&%s=%v", name, b)
}
func encodeInt8Param(name string, i int8, def int8) string {
	if i == def {
		return ""
	}
	return "&" + name + "=" + strconv.Itoa(int(i))
}
func encodeInt16Param(name string, i int16, def int16) string {
	if i == def {
		return ""
	}
	return "&" + name + "=" + strconv.Itoa(int(i))
}
func encodeInt32Param(name string, i int32, def int32) string {
	if i == def {
		return ""
	}
	return "&" + name + "=" + strconv.Itoa(int(i))
}
func encodeInt64Param(name string, i int64, def int64) string {
	if i == def {
		return ""
	}
	return "&" + name + "=" + strconv.FormatInt(i, 10)
}
func encodeFloat32Param(name string, i float32, def float32) string {
	if i == def {
		return ""
	}
	return "&" + name + "=" + strconv.FormatFloat(float64(i), 'g', -1, 32)
}
func encodeFloat64Param(name string, i float64, def float64) string {
	if i == def {
		return ""
	}
	return "&" + name + "=" + strconv.FormatFloat(i, 'g', -1, 64)
}
func encodeOptionalEnumParam(name string, e interface{}) string {
	if e == nil {
		return "\"\""
	}
	return fmt.Sprintf("&%s=%v", name, e)
}
func encodeOptionalBoolParam(name string, b *bool) string {
	if b == nil {
		return ""
	}
	return fmt.Sprintf("&%s=%v", name, *b)
}
func encodeOptionalInt32Param(name string, i *int32) string {
	if i == nil {
		return ""
	}
	return "&" + name + "=" + strconv.Itoa(int(*i))
}
func encodeOptionalInt64Param(name string, i *int64) string {
	if i == nil {
		return ""
	}
	return "&" + name + "=" + strconv.Itoa(int(*i))
}
func encodeParams(objs ...string) string {
	s := strings.Join(objs, "")
	if s == "" {
		return s
	}
	return "?" + s[1:]
}

func (client ZTSClient) GetServiceIdentity(domainName DomainName, serviceName ServiceName) (*ServiceIdentity, error) {
	var data *ServiceIdentity
	url := client.URL + "/domain/" + fmt.Sprint(domainName) + "/service/" + fmt.Sprint(serviceName)
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetServiceIdentityList(domainName DomainName) (*ServiceIdentityList, error) {
	var data *ServiceIdentityList
	url := client.URL + "/domain/" + fmt.Sprint(domainName) + "/service"
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetPublicKeyEntry(domainName DomainName, serviceName SimpleName, keyId string) (*PublicKeyEntry, error) {
	var data *PublicKeyEntry
	url := client.URL + "/domain/" + fmt.Sprint(domainName) + "/service/" + fmt.Sprint(serviceName) + "/publickey/" + keyId
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetHostServices(host string) (*HostServices, error) {
	var data *HostServices
	url := client.URL + "/host/" + host + "/services"
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetDomainSignedPolicyData(domainName DomainName, matchingTag string) (*DomainSignedPolicyData, string, error) {
	var data *DomainSignedPolicyData
	headers := map[string]string{
		"If-None-Match": matchingTag,
	}
	url := client.URL + "/domain/" + fmt.Sprint(domainName) + "/signed_policy_data"
	resp, err := client.httpGet(url, headers)
	if err != nil {
		return nil, "", err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, "", err
	}
	switch resp.StatusCode {
	case 200, 304:
		if 304 != resp.StatusCode {
			err = json.Unmarshal(contentBytes, &data)
			if err != nil {
				return nil, "", err
			}
		}
		tag := resp.Header.Get(rdl.FoldHttpHeaderName("ETag"))
		return data, tag, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return nil, "", errobj
	}
}

func (client ZTSClient) GetRoleToken(domainName DomainName, role EntityName, minExpiryTime *int32, maxExpiryTime *int32, proxyForPrincipal EntityName) (*RoleToken, error) {
	var data *RoleToken
	url := client.URL + "/domain/" + fmt.Sprint(domainName) + "/token" + encodeParams(encodeStringParam("role", string(role), ""), encodeOptionalInt32Param("minExpiryTime", minExpiryTime), encodeOptionalInt32Param("maxExpiryTime", maxExpiryTime), encodeStringParam("proxyForPrincipal", string(proxyForPrincipal), ""))
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) PostRoleCertificateRequest(domainName DomainName, roleName EntityName, req *RoleCertificateRequest) (*RoleToken, error) {
	var data *RoleToken
	url := client.URL + "/domain/" + fmt.Sprint(domainName) + "/role/" + fmt.Sprint(roleName) + "/token"
	contentBytes, err := json.Marshal(req)
	if err != nil {
		return data, err
	}
	resp, err := client.httpPost(url, nil, contentBytes)
	if err != nil {
		return data, err
	}
	contentBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetAccess(domainName DomainName, roleName EntityName, principal EntityName) (*Access, error) {
	var data *Access
	url := client.URL + "/access/domain/" + fmt.Sprint(domainName) + "/role/" + fmt.Sprint(roleName) + "/principal/" + fmt.Sprint(principal)
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetRoleAccess(domainName DomainName, principal EntityName) (*RoleAccess, error) {
	var data *RoleAccess
	url := client.URL + "/access/domain/" + fmt.Sprint(domainName) + "/principal/" + fmt.Sprint(principal)
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetTenantDomains(providerDomainName DomainName, userName EntityName, roleName EntityName, serviceName ServiceName) (*TenantDomains, error) {
	var data *TenantDomains
	url := client.URL + "/providerdomain/" + fmt.Sprint(providerDomainName) + "/user/" + fmt.Sprint(userName) + encodeParams(encodeStringParam("roleName", string(roleName), ""), encodeStringParam("serviceName", string(serviceName), ""))
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) PostInstanceInformation(info *InstanceInformation) (*Identity, error) {
	var data *Identity
	url := client.URL + "/instance"
	contentBytes, err := json.Marshal(info)
	if err != nil {
		return data, err
	}
	resp, err := client.httpPost(url, nil, contentBytes)
	if err != nil {
		return data, err
	}
	contentBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) PostInstanceRefreshRequest(domain CompoundName, service SimpleName, req *InstanceRefreshRequest) (*Identity, error) {
	var data *Identity
	url := client.URL + "/instance/" + fmt.Sprint(domain) + "/" + fmt.Sprint(service) + "/refresh"
	contentBytes, err := json.Marshal(req)
	if err != nil {
		return data, err
	}
	resp, err := client.httpPost(url, nil, contentBytes)
	if err != nil {
		return data, err
	}
	contentBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) PostAWSInstanceInformation(info *AWSInstanceInformation) (*Identity, error) {
	var data *Identity
	url := client.URL + "/aws/instance"
	contentBytes, err := json.Marshal(info)
	if err != nil {
		return data, err
	}
	resp, err := client.httpPost(url, nil, contentBytes)
	if err != nil {
		return data, err
	}
	contentBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) PostAWSCertificateRequest(domain CompoundName, service SimpleName, req *AWSCertificateRequest) (*Identity, error) {
	var data *Identity
	url := client.URL + "/aws/instance/" + fmt.Sprint(domain) + "/" + fmt.Sprint(service) + "/refresh"
	contentBytes, err := json.Marshal(req)
	if err != nil {
		return data, err
	}
	resp, err := client.httpPost(url, nil, contentBytes)
	if err != nil {
		return data, err
	}
	contentBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) GetAWSTemporaryCredentials(domainName DomainName, role CompoundName) (*AWSTemporaryCredentials, error) {
	var data *AWSTemporaryCredentials
	url := client.URL + "/domain/" + fmt.Sprint(domainName) + "/role/" + fmt.Sprint(role) + "/creds"
	resp, err := client.httpGet(url, nil)
	if err != nil {
		return data, err
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

func (client ZTSClient) PostDomainMetrics(domainName DomainName, req *DomainMetrics) (*DomainMetrics, error) {
	var data *DomainMetrics
	url := client.URL + "/metrics/" + fmt.Sprint(domainName)
	contentBytes, err := json.Marshal(req)
	if err != nil {
		return data, err
	}
	resp, err := client.httpPost(url, nil, contentBytes)
	if err != nil {
		return data, err
	}
	contentBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return data, err
	}
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(contentBytes, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	default:
		var errobj rdl.ResourceError
		json.Unmarshal(contentBytes, &errobj)
		if errobj.Code == 0 {
			errobj.Code = resp.StatusCode
		}
		if errobj.Message == "" {
			errobj.Message = string(contentBytes)
		}
		return data, errobj
	}
}

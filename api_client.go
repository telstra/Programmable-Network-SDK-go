/* 
 * Telstra Programmable Network API
 *
 * Telstra Programmable Network is a self-provisioning platform that allows its users to create on-demand connectivity services between multiple end-points and add various network functions to those services. Programmable Network enables to connectivity to a global ecosystem of networking services as well as public and private cloud services. Once you are connected to the platform on one or more POPs (points of presence), you can start creating those services based on the use case that you want to accomplish. The Programmable Network API is available to all customers who have registered to use the Programmable Network. To register, please contact your account representative.
 *
 * OpenAPI spec version: 2.1.2
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package TelstraTPN

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"net/url"
	"io/ioutil"
	"gopkg.in/go-resty/resty.v0"
)

type APIClient struct {
	config *Configuration
}

func (c *APIClient) SelectHeaderContentType(contentTypes []string) string {

	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func (c *APIClient) SelectHeaderAccept(accepts []string) string {

	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

func (c *APIClient) CallAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) (*resty.Response, error) {

	rClient := c.prepareClient()
	request := c.prepareRequest(rClient, postBody, headerParams, queryParams, formParams, fileName, fileBytes)

	switch strings.ToUpper(method) {
	case "GET":
		response, err := request.Get(path)
		return response, err
	case "POST":
		response, err := request.Post(path)
		return response, err
	case "PUT":
		response, err := request.Put(path)
		return response, err
	case "PATCH":
		response, err := request.Patch(path)
		return response, err
	case "DELETE":
		response, err := request.Delete(path)
		return response, err
	}

	return nil, fmt.Errorf("invalid method %v", method)
}

func (c *APIClient) ParameterToString(obj interface{}, collectionFormat string) string {
	delimiter := ""
	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

func (c *APIClient) prepareClient() *resty.Client {

	rClient := resty.New()

	rClient.SetDebug(c.config.Debug)
	if c.config.Transport != nil {
		rClient.SetTransport(c.config.Transport)
	}

	if c.config.Timeout != nil {
		rClient.SetTimeout(*c.config.Timeout)
	}
	rClient.SetLogger(ioutil.Discard)
	return rClient
}

func (c *APIClient) prepareRequest(
	rClient *resty.Client,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) *resty.Request {


	request := rClient.R()
	request.SetBody(postBody)

	if c.config.UserAgent != "" {
		request.SetHeader("User-Agent", c.config.UserAgent)
	}

	// add header parameter, if any
	if len(headerParams) > 0 {
		request.SetHeaders(headerParams)
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		request.SetMultiValueQueryParams(queryParams)
	}

	// add form parameter, if any
	if len(formParams) > 0 {
		request.SetFormData(formParams)
	}

	if len(fileBytes) > 0 && fileName != "" {
		_, fileNm := filepath.Split(fileName)
		request.SetFileReader("file", fileNm, bytes.NewReader(fileBytes))
	}
	return request
}
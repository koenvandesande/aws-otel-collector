/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	_context "context"
	"fmt"
	"io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// LocationsApiService LocationsApi service
type LocationsApiService service

type ApiLocationsFindByRegionIdRequest struct {
	ctx             _context.Context
	ApiService      *LocationsApiService
	regionId        string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiLocationsFindByRegionIdRequest) Pretty(pretty bool) ApiLocationsFindByRegionIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiLocationsFindByRegionIdRequest) Depth(depth int32) ApiLocationsFindByRegionIdRequest {
	r.depth = &depth
	return r
}
func (r ApiLocationsFindByRegionIdRequest) XContractNumber(xContractNumber int32) ApiLocationsFindByRegionIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiLocationsFindByRegionIdRequest) Execute() (Locations, *APIResponse, error) {
	return r.ApiService.LocationsFindByRegionIdExecute(r)
}

/*
 * LocationsFindByRegionId Get Locations within a Region
 * Retrieves the available locations in a region specified by its ID. The 'regionId' consists of the two character identifier of the region (country), e.g., 'de'.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param regionId The unique ID of the region.
 * @return ApiLocationsFindByRegionIdRequest
 */
func (a *LocationsApiService) LocationsFindByRegionId(ctx _context.Context, regionId string) ApiLocationsFindByRegionIdRequest {
	return ApiLocationsFindByRegionIdRequest{
		ApiService: a,
		ctx:        ctx,
		regionId:   regionId,
	}
}

/*
 * Execute executes the request
 * @return Locations
 */
func (a *LocationsApiService) LocationsFindByRegionIdExecute(r ApiLocationsFindByRegionIdRequest) (Locations, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Locations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsApiService.LocationsFindByRegionId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations/{regionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", _neturl.PathEscape(parameterToString(r.regionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "LocationsFindByRegionId",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiLocationsFindByRegionIdAndIdRequest struct {
	ctx             _context.Context
	ApiService      *LocationsApiService
	regionId        string
	locationId      string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiLocationsFindByRegionIdAndIdRequest) Pretty(pretty bool) ApiLocationsFindByRegionIdAndIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiLocationsFindByRegionIdAndIdRequest) Depth(depth int32) ApiLocationsFindByRegionIdAndIdRequest {
	r.depth = &depth
	return r
}
func (r ApiLocationsFindByRegionIdAndIdRequest) XContractNumber(xContractNumber int32) ApiLocationsFindByRegionIdAndIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiLocationsFindByRegionIdAndIdRequest) Execute() (Location, *APIResponse, error) {
	return r.ApiService.LocationsFindByRegionIdAndIdExecute(r)
}

/*
 * LocationsFindByRegionIdAndId Get Location by ID
 * Retrieves the information about the location specified by its ID. The 'locationId' consists of the three-digit identifier of the city according to the IATA code.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param regionId The unique ID of the region.
 * @param locationId The unique ID of the location.
 * @return ApiLocationsFindByRegionIdAndIdRequest
 */
func (a *LocationsApiService) LocationsFindByRegionIdAndId(ctx _context.Context, regionId string, locationId string) ApiLocationsFindByRegionIdAndIdRequest {
	return ApiLocationsFindByRegionIdAndIdRequest{
		ApiService: a,
		ctx:        ctx,
		regionId:   regionId,
		locationId: locationId,
	}
}

/*
 * Execute executes the request
 * @return Location
 */
func (a *LocationsApiService) LocationsFindByRegionIdAndIdExecute(r ApiLocationsFindByRegionIdAndIdRequest) (Location, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Location
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsApiService.LocationsFindByRegionIdAndId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations/{regionId}/{locationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", _neturl.PathEscape(parameterToString(r.regionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"locationId"+"}", _neturl.PathEscape(parameterToString(r.locationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "LocationsFindByRegionIdAndId",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiLocationsGetRequest struct {
	ctx             _context.Context
	ApiService      *LocationsApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiLocationsGetRequest) Pretty(pretty bool) ApiLocationsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiLocationsGetRequest) Depth(depth int32) ApiLocationsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiLocationsGetRequest) XContractNumber(xContractNumber int32) ApiLocationsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiLocationsGetRequest) Filter(key string, value string) ApiLocationsGetRequest {
	filterKey := fmt.Sprintf(FilterQueryParam, key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiLocationsGetRequest) OrderBy(orderBy string) ApiLocationsGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiLocationsGetRequest) MaxResults(maxResults int32) ApiLocationsGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiLocationsGetRequest) Execute() (Locations, *APIResponse, error) {
	return r.ApiService.LocationsGetExecute(r)
}

/*
* LocationsGet Get Locations
* Retrieves the available physical locations where you can deploy cloud resources in a VDC.

A location is identified by a combination of the following characters:

* a two-character **regionId**, which represents a country (example: 'de')

* a three-character **locationId**, which represents a city. The 'locationId' is typically based on the IATA code of the city's airport (example: 'txl').

>Note that 'locations' are read-only and cannot be changed.
* @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @return ApiLocationsGetRequest
*/
func (a *LocationsApiService) LocationsGet(ctx _context.Context) ApiLocationsGetRequest {
	return ApiLocationsGetRequest{
		ApiService: a,
		ctx:        ctx,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return Locations
 */
func (a *LocationsApiService) LocationsGetExecute(r ApiLocationsGetRequest) (Locations, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Locations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsApiService.LocationsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("maxResults", parameterToString(*r.maxResults, ""))
	}
	if len(r.filters) > 0 {
		for k, v := range r.filters {
			for _, iv := range v {
				localVarQueryParams.Add(k, iv)
			}
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "LocationsGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

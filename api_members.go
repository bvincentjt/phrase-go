package phrase

import (
  "fmt"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// MembersApiService MembersApi service
type MembersApiService service

// MemberDeleteOpts Optional parameters for the method 'MemberDelete'
type MemberDeleteOpts struct {
    XPhraseAppOTP optional.String
}

/*
MemberDelete Remove a user from the account
Remove a user from the account. The user will be removed from the account but not deleted from Phrase. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Account ID
 * @param id ID
 * @param optional nil or *MemberDeleteOpts - Optional Parameters:
 * @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
*/
func (a *MembersApiService) MemberDelete(ctx _context.Context, accountId string, id string, localVarOptionals *MemberDeleteOpts) ([]byte, *APIResponse, error) {

  fmt.Println("MembersApi")
  fmt.Println("MemberDelete")

	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/members/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XPhraseAppOTP.IsSet() {
		localVarHeaderParams["X-PhraseApp-OTP"] = parameterToString(localVarOptionals.XPhraseAppOTP.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarBody, localVarHTTPResponse, newErr
	}

	return localVarBody, localVarHTTPResponse, nil
}

// MemberShowOpts Optional parameters for the method 'MemberShow'
type MemberShowOpts struct {
    XPhraseAppOTP optional.String
}

/*
MemberShow Get single member
Get details on a single user in the account. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Account ID
 * @param id ID
 * @param optional nil or *MemberShowOpts - Optional Parameters:
 * @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
@return Member
*/
func (a *MembersApiService) MemberShow(ctx _context.Context, accountId string, id string, localVarOptionals *MemberShowOpts) (Member, *APIResponse, error) {

  fmt.Println("MembersApi")
  fmt.Println("MemberShow")

	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Member
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/members/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.XPhraseAppOTP.IsSet() {
		localVarHeaderParams["X-PhraseApp-OTP"] = parameterToString(localVarOptionals.XPhraseAppOTP.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Member
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// MemberUpdateOpts Optional parameters for the method 'MemberUpdate'
type MemberUpdateOpts struct {
    XPhraseAppOTP optional.String
}

/*
MemberUpdate Update a member
Update user permissions in the account. Developers and translators need &lt;code&gt;project_ids&lt;/code&gt; and &lt;code&gt;locale_ids&lt;/code&gt; assigned to access them. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Account ID
 * @param id ID
 * @param memberUpdateParameters
 * @param optional nil or *MemberUpdateOpts - Optional Parameters:
 * @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
@return Member
*/
func (a *MembersApiService) MemberUpdate(ctx _context.Context, accountId string, id string, memberUpdateParameters MemberUpdateParameters, localVarOptionals *MemberUpdateOpts) (Member, *APIResponse, error) {

  fmt.Println("MembersApi")
  fmt.Println("MemberUpdate")

	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Member
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/members/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if localVarOptionals != nil && localVarOptionals.XPhraseAppOTP.IsSet() {
		localVarHeaderParams["X-PhraseApp-OTP"] = parameterToString(localVarOptionals.XPhraseAppOTP.Value(), "")
	}
	// body params
	localVarPostBody = &memberUpdateParameters
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Member
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// MembersListOpts Optional parameters for the method 'MembersList'
type MembersListOpts struct {
    XPhraseAppOTP optional.String
    Page optional.Int32
    PerPage optional.Int32
}

/*
MembersList List members
Get all users active in the account. It also lists resources like projects and locales the member has access to. In case nothing is shown the default access from the role is used. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Account ID
 * @param optional nil or *MembersListOpts - Optional Parameters:
 * @param "XPhraseAppOTP" (optional.String) -  Two-Factor-Authentication token (optional)
 * @param "Page" (optional.Int32) -  Page number
 * @param "PerPage" (optional.Int32) -  allows you to specify a page size up to 100 items, 10 by default
@return []Member
*/
func (a *MembersApiService) MembersList(ctx _context.Context, accountId string, localVarOptionals *MembersListOpts) ([]Member, *APIResponse, error) {

  fmt.Println("MembersApi")
  fmt.Println("MembersList")

	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Member
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{account_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.QueryEscape(parameterToString(accountId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("per_page", parameterToString(localVarOptionals.PerPage.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.XPhraseAppOTP.IsSet() {
		localVarHeaderParams["X-PhraseApp-OTP"] = parameterToString(localVarOptionals.XPhraseAppOTP.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v []Member
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

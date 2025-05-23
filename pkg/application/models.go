/*
 * Copyright (c) 2025, WSO2 LLC. (http://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package application

import (
	"github.com/asgardeo/go/pkg/application/internal"
)

// AppType represents the type of application
type AppType string

// Application types as typed constants
const (
	AppTypeSPA    AppType = "spa"
	AppTypeMobile AppType = "mobile"
	AppTypeM2M    AppType = "m2m"
	AppTypeSSRWeb AppType = "ssr_web"
)

type ApplicationBasicInfoResponseModel struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	ClientId         string  `json:"client_id,omitempty"`
	ClientSecret     string  `json:"client_secret,omitempty"`
	RedirectURL      string  `json:"redirect_url,omitempty"`
	AuthorizedScopes string  `json:"scope,omitempty"`
	AppType          AppType `json:"application_type"`
}

type ApplicationListResponseModel = internal.ApplicationListResponse

type AuthorizedAPICreateModel = internal.AddAuthorizedAPIJSONRequestBody

type AuthorizedAPIResponseModel = internal.AuthorizedAPIResponse

// ApplicationBasicInfoUpdateModel defines a simplified model for updating basic application information
type ApplicationBasicInfoUpdateModel struct {
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	ImageUrl        *string `json:"imageUrl,omitempty"`
	AccessUrl       *string `json:"accessUrl,omitempty"`
	LogoutReturnUrl *string `json:"logoutReturnUrl,omitempty"`
}

// ApplicationOAuthConfigUpdateModel contains only the fields that can be updated in OAuth configuration
type ApplicationOAuthConfigUpdateModel struct {
	AccessTokenAttributes                 *[]string `json:"accessTokenAttributes,omitempty"`
	ApplicationAccessTokenExpiryInSeconds *int64    `json:"applicationAccessTokenExpiryInSeconds,omitempty"`
	UserAccessTokenExpiryInSeconds        *int64    `json:"userAccessTokenExpiryInSeconds,omitempty"`

	AllowedOrigins *[]string `json:"allowedOrigins,omitempty"`
	CallbackURLs   *[]string `json:"callbackURLs,omitempty"`

	Logout *internal.OIDCLogoutConfiguration `json:"logout,omitempty"`

	RefreshTokenExpiryInSeconds *int64 `json:"refreshTokenExpiryInSeconds,omitempty"`
}

type ApplicationClaimConfigurationUpdateModel struct {
	RequestedClaims *[]RequestedClaimModel `json:"requestedClaims,omitempty"`
}

type RequestedClaimModel struct {
	Claim     ClaimModel `json:"claim"`
	Mandatory *bool      `json:"mandatory"`
}

type ClaimModel struct {
	DisplayName *string `json:"displayName,omitempty"`
	Id          *string `json:"id,omitempty"`
	Uri         string  `json:"uri"`
}

type LoginFlowGenerateResponseModel = internal.LoginFlowGenerateResponse

type LoginFlowStatusResponseModel = internal.LoginFlowStatusResponse

type LoginFlowResultResponseModel struct {
	Data   *LoginFlowUpdateModel `json:"data,omitempty"`
	Status *internal.StatusEnum  `json:"status,omitempty"`
}

type LoginFlowUpdateModel = internal.AuthenticationSequence

type LoginFlowStepModel = internal.AuthenticationStepModel

type AuthenticatorModel = internal.Authenticator

type LoginFlowTypeModel = internal.AuthenticationSequenceType

// convertBasicInfoUpdateModelToApplicationPatchModel converts the public ApplicationBasicInfoUpdateModel to the internal PatchApplicationJSONRequestBody
func convertBasicInfoUpdateModelToApplicationPatchModel(model ApplicationBasicInfoUpdateModel) internal.PatchApplicationJSONRequestBody {
	return internal.PatchApplicationJSONRequestBody{
		Name:            model.Name,
		Description:     model.Description,
		ImageUrl:        model.ImageUrl,
		AccessUrl:       model.AccessUrl,
		LogoutReturnUrl: model.LogoutReturnUrl,
	}
}

// convertClaimConfigUpdateModelToApplicationPatchModel converts the public ApplicationClaimConfigurationUpdateModel to the internal PatchApplicationJSONRequestBody
func convertClaimConfigUpdateModelToApplicationPatchModel(model ApplicationClaimConfigurationUpdateModel) internal.PatchApplicationJSONRequestBody {
	if model.RequestedClaims == nil {
		return internal.PatchApplicationJSONRequestBody{}
	} else {
		requestedClaims := make([]internal.RequestedClaimConfiguration, len(*model.RequestedClaims))
		for i, claim := range *model.RequestedClaims {
			requestedClaims[i] = internal.RequestedClaimConfiguration{
				Claim: internal.Claim{
					Uri: claim.Claim.Uri,
				},
				Mandatory: claim.Mandatory,
			}
		}
		return internal.PatchApplicationJSONRequestBody{
			ClaimConfiguration: &internal.ClaimConfiguration{
				RequestedClaims: &requestedClaims},
		}
	}
}

func convertToLoginFlowResultResponseModel(model internal.LoginFlowResultResponse) LoginFlowResultResponseModel {
	loginFlowUpdateData := convertToLoginFlowUpdateModel(*model.Data)
	return LoginFlowResultResponseModel{
		Data:   &loginFlowUpdateData,
		Status: model.Status,
	}
}

func convertToLoginFlowUpdateModel(data map[string]interface{}) LoginFlowUpdateModel {
	var loginFlowUpdate LoginFlowUpdateModel
	if data != nil {
		attributeStepId := int(data["attributeStepId"].(float64))
		subjectStepId := int(data["subjectStepId"].(float64))
		steps := convertToLoginFlowStepModelList(data["steps"].([]interface{}))
		loginFlowType := LoginFlowTypeModel(data["type"].(string))
		loginFlowUpdate = LoginFlowUpdateModel{
			AttributeStepId: &attributeStepId,
			Steps:           &steps,
			SubjectStepId:   &subjectStepId,
			Type:            &loginFlowType,
		}
	}
	return loginFlowUpdate
}

func convertToLoginFlowStepModelList(data []interface{}) []LoginFlowStepModel {
	var loginFlowStepList []LoginFlowStepModel
	if data != nil {
		for _, item := range data {
			step := item.(map[string]interface{})
			loginFlowStepList = append(loginFlowStepList, convertToLoginFlowStepModel(step))
		}
	}
	return loginFlowStepList
}

func convertToLoginFlowStepModel(data map[string]interface{}) LoginFlowStepModel {
	var loginFlowStep LoginFlowStepModel
	if data != nil {
		id := int(data["id"].(float64))
		options := convertToAuthenticatorList(data["options"].([]interface{}))
		loginFlowStep = LoginFlowStepModel{
			Id:      id,
			Options: options,
		}
	}
	return loginFlowStep
}

func convertToAuthenticatorList(data []interface{}) []AuthenticatorModel {
	var authenticators []AuthenticatorModel
	if data != nil {
		for _, item := range data {
			authenticator := item.(map[string]interface{})
			authenticators = append(authenticators, convertToAuthenticatorModel(authenticator))
		}
	}
	return authenticators
}

func convertToAuthenticatorModel(data map[string]interface{}) AuthenticatorModel {
	var authenticatorModel AuthenticatorModel
	if data != nil {
		authenticatorModel = AuthenticatorModel{
			Authenticator: data["authenticator"].(string),
			Idp:           data["idp"].(string),
		}
	}
	return authenticatorModel
}

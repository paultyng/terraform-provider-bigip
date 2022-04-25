package bigip
import (
	"fmt"
)

const (
	uriWafPol             = "policies"
	uriUrls       		  = "urls"
	uriParams          	  = "parameters"
	uriWafSign				  = "signatures"
)

type WafEntityParameters struct {
	WafEntityParametersList []WafEntityParameter `json:"items"`
}

type WafEntityParameter struct {
	Name                           string `json:"name,omitempty"`
	Description                    string `json:"description,omitempty"`
	Type                           string `json:"type,omitempty"`
	ValueType                      string `json:"valueType,omitempty"`
	AllowEmptyValue                bool   `json:"allowEmptyValue,omitempty"`
	AllowRepeatedParameterName     bool   `json:"allowRepeatedParameterName,omitempty"`
	AttackSignaturesCheck          bool   `json:"attackSignaturesCheck,omitempty"`
	CheckMaxValueLength            bool   `json:"checkMaxValueLength,omitempty"`
	CheckMinValueLength            bool   `json:"checkMinValueLength,omitempty"`
	DataType                       string `json:"dataType,omitempty"`
	EnableRegularExpression        bool   `json:"enableRegularExpression,omitempty"`
	IsBase64                       bool   `json:"isBase64,omitempty"`
	IsCookie                       bool   `json:"isCookie,omitempty"`
	IsHeader                       bool   `json:"isHeader,omitempty"`
	Level                          string `json:"level,omitempty"`
	Mandatory                      bool   `json:"mandatory,omitempty"`
	MetacharsOnParameterValueCheck bool   `json:"metacharsOnParameterValueCheck,omitempty"`
	ParameterLocation              string `json:"parameterLocation,omitempty"`
	PerformStaging                 bool   `json:"performStaging,omitempty"`
	SensitiveParameter             bool   `json:"sensitiveParameter,omitempty"`
	SignatureOverrides_Disable     []int  `json:"signatureOverrides_disable,omitempty"`
}

type WafPolicies struct {
	WafPolicyList []WafPolicy `json:"items"`
}

type WafPolicy struct {
	Name                  string        `json:"name,omitempty"`
	Partition             string        `json:"partition,omitempty"`
	Description           string        `json:"description,omitempty"`
	Template              string        `json:"template,omitempty"`
	ApplicationLanguage   string        `json,"applicationLanguage,omitempty"`
	EnforcementMode     string          `json:"enforcementMode,omitempty"`
}

type WafEntityURLs struct {
	WafEntityURLList []WafEntityURL `json:"items"`
}

type WafEntityURL struct {
	Name                       string `json:"name,omitempty"`
	Description                string `json:"description,omitempty"`
	Type                       string `json:"type,omitempty"`
	Protocol                   string `json:"protocol,omitempty"`
	Method                     string `json:"method,omitempty"`
	MethodOverrides            string `json:"methodOverrides,omitempty"`
	PerformStaging             bool   `json:"performStaging,omitempty"`
	SignatureOverrides     	   []SignatureIDs `json:"signatureOverrides,omitempty"`
}

type SignatureIDs struct{
	SignatureReference []SigIDs
	Enabled 			bool `json:"enabled,omitempty"`
}

type SigIDs struct{
	Link			string `json:"link,omitempty"`
	IsUserDefined	bool   `json:"isUserDefined,omitempty"`
	Name			string `json:"name,omitempty"`
	SignatureId		int `json:"signatureId,omitempty"`
}

type Signatures struct {
	Signatures []Signature `json:"items"`
}

type Signature struct {
	Name			string `json:"name,omitempty"`
	ResourceId		string `json:"id,omitempty"`
	Description		string `json:"description,omitempty"`
	SignatureId		int `json:"signatureId,omitempty"`
	Type			string `json:"signatureType,omitempty"`
	Accuracy		string `json:"accuracy,omitempty"`
	Risk			string `json:"risk,omitempty"`
}

func (b *BigIP) GetWafSignature(signatureid int) (*Signatures, error) {
	var signature Signatures
	var query = fmt.Sprintf("?$filter=signatureId+eq+%d", signatureid)
	err, _ := b.getForEntity(&signature, uriMgmt, uriTm, uriAsm, uriWafSign, query)
	if err != nil {
		return nil, err
	}
	return &signature, nil
}

// This method is not correct as of now, it tries to access keys that are not there in WafPolicy struct yet
//func (b *BigIP) GetPolicyId(policyName string) (string, error) {
//	var self WafPolicies
//	err, _ := b.getForEntity(&self, uriMgmt, uriTm, uriAsm, uriWafPol)
//
//	if err != nil {
//		return "", err
//	}
//
//	for _, policy := range self.WafPolicyList {
//		if policy.FullPath == "policyName" {
//			return policy.Id, nil
//		}
//	}
//
//	return "", fmt.Errorf("could not get the policy ID")
//}

func (b *BigIP) WafEntityParameters(policyId string) (*WafEntityParameters, error) {
	var self WafEntityParameters
	err, _ := b.getForEntity(&self, uriMgmt, uriTm, uriAsm, uriWafPol, uriParams)
	if err != nil {
		return nil, err
	}
	return &self, nil
}

func (b *BigIP) GetEntityParameters(policyId, parameterId string) (*WafEntityParameter, error) {
	var wafEntityParameter WafEntityParameter
	err, _ := b.getForEntity(&wafEntityParameter, uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriParams, parameterId)
	if err != nil {
		return nil, err
	}
	return &wafEntityParameter, nil
}

func (b *BigIP) CreateWafEntityParameter(config *WafEntityParameter, policyId string) error {
	return b.post(config, uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriParams)
}

func (b *BigIP) ModifyWafEntityParameter(config *WafEntityParameter, parameterId, policyId string) error {
	return b.patch(config, uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriParams, parameterId)
}

func (b *BigIP) DeleteWafEntityParameter(parameterId, policyId string) error {
	return b.delete(uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriParams, parameterId)
}

func (b *BigIP) GetWafEntityUrls(policyId string) (*WafPolicy, error) {
	var wafPolicy WafPolicy
	err, _ := b.getForEntity(&wafPolicy, uriMgmt, uriTm, uriAsm, uriWafPol, policyId)

	if err != nil {
		return nil, err
	}

	return &wafPolicy, nil
}

func (b *BigIP) CreateWafPolicy(config *WafPolicy) error {
	return b.post(config, uriMgmt, uriTm, uriAsm, uriWafPol)
}

func (b *BigIP) ModifyWafPolicy(config *WafPolicy, policyId string) error {
	return b.patch(config, uriMgmt, uriTm, uriAsm, uriWafPol, policyId)
}

func (b *BigIP) DeleteWafPolicy(config *WafPolicy, policyId string) error {
	return b.delete(uriMgmt, uriTm, uriAsm, uriWafPol, policyId)
}

func (b *BigIP) WafEntityUrls(policyId string) (*WafEntityURLs, error) {
	var self WafEntityURLs
	err, _ := b.getForEntity(&self, uriMgmt, uriTm, uriAsm, uriWafPol, uriUrls)
	if err != nil {
		return nil, err
	}
	return &self, nil
}

func (b *BigIP) GetEntityUrls(policyId, urlId string) (*WafEntityURL, error) {
	var wafEntityurl WafEntityURL
	err, _ := b.getForEntity(&wafEntityurl, uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriUrls, urlId)
	if err != nil {
		return nil, err
	}
	return &wafEntityurl, nil
}

func (b *BigIP) CreateWafEntityUrl(config *WafEntityURL, policyId string) error {
	return b.post(config, uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriUrls)
}

func (b *BigIP) ModifyWafEntityUrl(config *WafEntityURL, urlId, policyId string) error {
	return b.patch(config, uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriUrls, urlId)
}

func (b *BigIP) DeleteWafEntityUrl(urlId, policyId string) error {
	return b.delete(uriMgmt, uriTm, uriAsm, uriWafPol, policyId, uriUrls, urlId)
}
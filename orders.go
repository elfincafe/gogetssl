package gogetssl

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/elfincafe/annette"
)

type (
	Order struct {
		Id     int    `json:"order_id"`
		Status string `json:"status"`
	}
	GetAllSslOrdersResponse struct {
		Limit   int     `json:"limit"`
		Offset  int     `json:"offset"`
		Count   int     `json:"count"`
		Success bool    `json:"success"`
		Orders  []Order `json:"orders"`
	}
	GetOrderStatusResponse struct {
		OrderId               int      `json:"order_id"`
		PartnerOrderId        int      `json:"partner_order_id"`
		InternalId            string   `json:"internal_id"`
		Status                string   `json:"status"`
		StatusDescription     string   `json:"status_description"`
		DcvStatus             int      `json:"dcv_status"`
		ProductId             int      `json:"product_id"`
		Domain                string   `json:"domain"`
		TotalDomain           int      `json:"total_domains"`
		BaseDomainCount       int      `json:"base_domain_count"`
		SingleSanCount        int      `json:"single_san_count"`
		WildcardSanCount      int      `json:"wildcard_san_count"`
		ValidityPeriod        int      `json:"validity_period"`
		ValidFrom             DateTime `json:"valid_from"`
		ValidTill             DateTime `json:"valid_till"`
		BeginDate             DateTime `json:"begin_date"`
		EndDate               DateTime `json:"end_date"`
		CsrCode               string   `json:"csr_code"`
		CaCode                string   `json:"ca_code"`
		CrtCode               string   `json:"crt_code"`
		ServerCount           int      `json:"server_count"`
		Reissue               int      `json:"reissue"`
		ReissueNow            int      `json:"reissue_now"`
		Renew                 int      `json:"renew"`
		WebserverType         int      `json:"webserver_type"`
		Upgrade               int      `json:"upgrade"`
		ApproverEmails        string   `json:"approver_emails"`
		DcvMethod             string   `json:"dcv_method"`
		AdminAddressLine1     string   `json:"admin_addressline1"`
		AdminAddressLine2     string   `json:"admin_addressline2"`
		AdminAddressLine3     string   `json:"admin_addressline3"`
		AdminCity             string   `json:"admin_city"`
		AdminRegion           string   `json:"admin_region"`
		AdminCountry          string   `json:"admin_country"`
		AdminFax              string   `json:"admin_fax"`
		AdminPhone            string   `json:"admin_phone"`
		AdminPostalCode       string   `json:"admin_postalcode"`
		AdminEmail            string   `json:"admin_email"`
		AdminFirstName        string   `json:"admin_firstname"`
		AdminLastName         string   `json:"admin_lastname"`
		AdminOrganization     string   `json:"admin_organization"`
		AdminTitle            string   `json:"admin_title"`
		OrgAddressLine1       string   `json:"org_addressline1"`
		OrgAddressLine2       string   `json:"org_addressline2"`
		OrgAddressLine3       string   `json:"org_addressline3"`
		OrgCity               string   `json:"org_city"`
		OrgRegion             string   `json:"org_region"`
		OrgCountry            string   `json:"org_country"`
		OrgFax                string   `json:"org_fax"`
		OrgPhone              string   `json:"org_phone"`
		OrgPostalCode         string   `json:"org_postalcode"`
		OrgLei                string   `json:"org_lei"`
		TechOrganization      string   `json:"tech_organization"`
		TechAddressLine1      string   `json:"tech_addressline1"`
		TechAddressLine2      string   `json:"tech_addressline2"`
		TechAddressLine3      string   `json:"tech_addressline3"`
		TechCity              string   `json:"tech_city"`
		TechRegion            string   `json:"tech_region"`
		TechCountry           string   `json:"tech_country"`
		TechFax               string   `json:"tech_fax"`
		TechPhone             string   `json:"tech_phone"`
		TechPostCode          string   `json:"tech_postalcode"`
		TechEmail             string   `json:"tech_email"`
		TechFirstName         string   `json:"tech_firstname"`
		TechLastName          string   `json:"tech_lastname"`
		TechTitle             string   `json:"tech_title"`
		SslPrice              string   `json:"ssl_price"`
		SslPeriod             int      `json:"ssl_period"`
		AdminMsg              string   `json:"admin_msg"`
		FreeEvUpgrade         int      `json:"free_ev_upgrade"`
		CodeSigningInviteUrl  string   `json:"codesigning_inviteurl"`
		ValidationDescription string   `json:"validation_description"`
		ManualCheck           string   `json:"manual_check"`
		PreSigning            string   `json:"pre_signing"`
		ApproverMethod        map[string]struct {
			Link     string `json:"link,omitempty"`
			FileName string `json:"filename,omitempty"`
			Content  string `json:"content,omitempty"`
			Record   string `json:"record,omitempty"`
			Email    string `json:"email,omitempty"`
		} `json:"approver_method"`
		San []struct {
			Name              string `json:"san_name"`
			ValidationMethod  string `json:"validation_method"`
			Status            string `json:"status"`
			StatusDescription string `json:"status_description"`
			Validation        map[string]struct {
				Link     string `json:"link"`
				FileName string `json:"filename"`
				Content  string `json:"content"`
			} `json:"validation"`
		} `json:"san"`
		Success   bool `json:"success"`
		TimeStamp int  `json:"time_stamp"`
	}
	ReissueSslOrderResponse struct {
	}
)

func (api *Api) GetOrderStatus(orderId int) (*GetOrderStatusResponse, error) {
	endpoint, err := url.Parse(fmt.Sprintf("%s/orders/status/%d", api.LiveAPI, orderId))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("auth_key", api.Key)
	endpoint.RawQuery = q.Encode()

	var e Error
	var r GetOrderStatusResponse
	req := annette.New(endpoint)
	res, err := req.Get()
	if err != nil {
		return nil, err
	}
	if !res.IsStatus200s() {
		err = json.Unmarshal(res.Binary(), &e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s. %s", e.Message, e.Description)
	}
	resBody := res.Binary()
	fmt.Println(string(resBody))

	err = json.Unmarshal(resBody, &e)
	if err == nil && e.Error {
		return nil, fmt.Errorf("%s. %s", e.Message, e.Description)
	}
	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (api *Api) GetAllSslOrders(limit, offset int) (*GetAllSslOrdersResponse, error) {
	endpoint, err := url.Parse(fmt.Sprintf("%s/orders/ssl/all", api.LiveAPI))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("auth_key", api.Key)
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(offset))
	endpoint.RawQuery = q.Encode()
	fmt.Println(endpoint)

	var e Error
	var r GetAllSslOrdersResponse
	req := annette.New(endpoint)
	res, err := req.Get()
	if err != nil {
		return nil, err
	}
	if !res.IsStatus200s() {
		err = json.Unmarshal(res.Binary(), &e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s. %s", e.Message, e.Description)
	}
	resBody := res.Binary()
	fmt.Println(string(resBody))

	err = json.Unmarshal(resBody, &e)
	if err == nil && e.Error {
		return nil, fmt.Errorf("%s. %s", e.Message, e.Description)
	}
	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

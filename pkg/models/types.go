// SPDX-FileCopyrightText: 2023 Kavya Shukla <kavyuushukla@gmail.com>
// SPDX-License-Identifier: GPL-2.0-only

package models

// The License struct represents a license entity with various attributes and
// properties associated with it.
// It provides structured storage for license-related information.
type LicenseDB struct {
	Id              int64  `json:"rf_id" gorm:"primary_key"`
	Shortname       string `json:"rf_shortname"`
	Fullname        string `json:"rf_fullname"`
	Text            string `json:"rf_text"`
	Url             string `json:"rf_url"`
	AddDate         string `json:"rf_add_date"`
	Copyleft        string `json:"rf_copyleft"`
	FSFfree         string `json:"rf_FSFfree"`
	OSIapproved     string `json:"rf_OSIapproved"`
	GPLv2compatible string `json:"rf_GPLv2compatible"`
	GPLv3compatible string `json:"rf_GPLv3compatible"`
	Notes           string `json:"rf_notes"`
	Fedora          string `json:"rf_Fedora"`
	TextUpdatable   string `json:"rf_text_updatable"`
	DetectorType    string `json:"rf_detector_type"`
	Active          string `json:"rf_active"`
	Source          string `json:"rf_source"`
	SpdxId          string `json:"rf_spdx_id"`
	Risk            string `json:"rf_risk"`
	Flag            string `json:"rf_flag"`
	Marydone        string `json:"marydone"`
}

type LicenseJson struct {
	Shortname       string `json:"rf_shortname"`
	Fullname        string `json:"rf_fullname"`
	Text            string `json:"rf_text"`
	Url             string `json:"rf_url"`
	AddDate         string `json:"rf_add_date"`
	Copyleft        string `json:"rf_copyleft"`
	FSFfree         string `json:"rf_FSFfree"`
	OSIapproved     string `json:"rf_OSIapproved"`
	GPLv2compatible string `json:"rf_GPLv2compatible"`
	GPLv3compatible string `json:"rf_GPLv3compatible"`
	Notes           string `json:"rf_notes"`
	Fedora          string `json:"rf_Fedora"`
	TextUpdatable   string `json:"rf_text_updatable"`
	DetectorType    string `json:"rf_detector_type"`
	Active          string `json:"rf_active"`
	Source          string `json:"rf_source"`
	SpdxCompatible  string `json:"rf_spdx_compatible"`
	Risk            string `json:"rf_risk"`
	Flag            string `json:"rf_flag"`
	Marydone        string `json:"marydone"`
}

// The Meta struct represents additional metadata associated with a
// license retrieval operation.
// It contains information that provides context and supplementary details
// about the retrieved license data.
type PaginationMeta struct {
	ResourceCount int `json:"resource_count"`
	Page          int `json:"page,omitempty"`
	PerPage       int `json:"per_page,omitempty"`
}

// LicenseResponse struct is representation of design API response of license.
// The LicenseResponse struct represents the response data structure for
// retrieving license information.
// It is used to encapsulate license-related data in an organized manner.
type LicenseResponse struct {
	Status int            `json:"status"`
	Data   []LicenseDB    `json:"data"`
	Meta   PaginationMeta `json:"paginationmeta"`
}

// The LicenseError struct represents an error response related to license operations.
// It provides information about the encountered error, including details such as
// status, error message, error type, path, and timestamp.
type LicenseError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Error     string `json:"error"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}

// The LicenseInput struct represents the input or payload required for creating a license.
// It contains various fields that capture the necessary information for defining a license entity.
type LicenseInput struct {
	Shortname       string `json:"rf_shortname" binding:"required"`
	Fullname        string `json:"rf_fullname" binding:"required"`
	Text            string `json:"rf_text" binding:"required"`
	Url             string `json:"rf_url"`
	AddDate         string `json:"rf_add_date"`
	Copyleft        string `json:"rf_copyleft"`
	FSFfree         string `json:"rf_FSFfree"`
	OSIapproved     string `json:"rf_OSIapproved"`
	GPLv2compatible string `json:"rf_GPLv2compatible"`
	GPLv3compatible string `json:"rf_GPLv3compatible"`
	Notes           string `json:"rf_notes"`
	Fedora          string `json:"rf_Fedora"`
	TextUpdatable   string `json:"rf_text_updatable"`
	DetectorType    string `json:"rf_detector_type"`
	Active          string `json:"rf_active"`
	Source          string `json:"rf_source"`
	SpdxId          string `json:"rf_spdx_id" binding:"required"`
	Risk            string `json:"rf_risk"`
	Flag            string `json:"rf_flag"`
	Marydone        string `json:"marydone"`
}

// User struct is representation of user information.
type User struct {
	Id           int64  `json:"id" gorm:"primary_key"`
	Username     string `json:"username" gorm:"unique" binding:"required"`
	Userlevel    string `json:"userlevel" gorm:"unique" binding:"required"`
	Userpassword string `json:"userpassword" gorm:"unique" binding:"required"`
}

// UserResponse struct is representation of design API response of user.
type UserResponse struct {
	Status int            `json:"status"`
	Data   []User         `json:"data"`
	Meta   PaginationMeta `json:"paginationmeta"`
}

// SearchLicense struct represents the input needed to search in a license.
type SearchLicense struct {
	Field      string `json:"field" binding:"required"`
	SearchTerm string `json:"search_term" binding:"required"`
	Search     string `json:"search"`
}

// Audit struct represents an audit entity with certain attributes and properties
// It has user id as a foreign key
type Audit struct {
	Id        int    `json:"id" gorm:"primary_key"`
	UserId    int64  `json:"user_id"`
	User      User   `gorm:"foreignKey:UserId;references:Id" json:"-"`
	TypeId    int64  `json:"type_id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
}

// ChangeLog struct represents a change entity with certain attributes and properties
type ChangeLog struct {
	Id           int    `json:"id" gorm:"primary_key"`
	Field        string `json:"field"`
	UpdatedValue string `json:"updated_value"`
	OldValue     string `json:"old_value"`
	AuditId      int    `json:"audit_id"`
	Audit        Audit  `gorm:"foreignKey:AuditId;references:Id" json:"-"`
}

// ChangeLogResponse represents the design of API response of change log
type ChangeLogResponse struct {
	Status int            `json:"status"`
	Data   []ChangeLog    `json:"data"`
	Meta   PaginationMeta `json:"paginationmeta"`
}

// AuditResponse represents the response format for audit data.
type AuditResponse struct {
	Status int            `json:"status"`
	Data   []Audit        `json:"data"`
	Meta   PaginationMeta `json:"paginationmeta"`
}

// Obligation represents an obligation record in the database.
type Obligation struct {
	Id             int64  `json:"id" gorm:"primary_key"`
	Topic          string `json:"topic"`
	Type           string `json:"type"`
	Text           string `json:"text"`
	Classification string `json:"classification"`
	Modifications  string `json:"modifications"`
	Comment        string `json:"comment"`
	Active         bool   `json:"active"`
	TextUpdatable  bool   `json:"text_updatable"`
	Md5            string `json:"md5" gorm:"unique"`
}

// ObligationInput represents the input format for creating a new obligation.
type ObligationInput struct {
	Topic          string   `json:"topic" binding:"required"`
	Type           string   `json:"type" binding:"required"`
	Text           string   `json:"text" binding:"required"`
	Classification string   `json:"classification"`
	Modifications  string   `json:"modifications"`
	Comment        string   `json:"comment"`
	Active         bool     `json:"active"`
	TextUpdatable  bool     `json:"text_updatable"`
	Shortnames     []string `json:"shortnames"`
}

// UpdateObligation represents the input format for updating an existing obligation.
type UpdateObligation struct {
	Topic          string `json:"topic"`
	Type           string `json:"type"`
	Text           string `json:"text"`
	Classification string `json:"classification"`
	Modifications  string `json:"modifications"`
	Comment        string `json:"comment"`
	Active         bool   `json:"active"`
	TextUpdatable  bool   `json:"text_updatable"`
	Md5            string `json:"md5"`
}

// ObligationMap represents the mapping between an obligation and a license.
type ObligationMap struct {
	ObligationPk int64      `json:"obligation_pk"`
	Obligation   Obligation `gorm:"foreignKey:ObligationPk;references:Id" json:"-"`
	OmPk         int64      `json:"om_pk" gorm:"primary_key"`
	RfPk         int64      `json:"rf_pk"`
	LicenseDB    LicenseDB  `gorm:"foreignKey:RfPk;references:Id" json:"-"`
}

// ObligationResponse represents the response format for obligation data.
type ObligationResponse struct {
	Status int            `json:"status"`
	Data   []Obligation   `json:"data"`
	Meta   PaginationMeta `json:"paginationmeta"`
}

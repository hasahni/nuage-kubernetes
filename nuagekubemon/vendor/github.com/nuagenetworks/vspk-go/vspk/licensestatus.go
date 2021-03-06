/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LicenseStatusIdentity represents the Identity of the object
var LicenseStatusIdentity = bambou.Identity{
	Name:     "licensestatus",
	Category: "licensestatus",
}

// LicenseStatusList represents a list of LicenseStatus
type LicenseStatusList []*LicenseStatus

// LicenseStatusAncestor is the interface of an ancestor of a LicenseStatus must implement.
type LicenseStatusAncestor interface {
	LicenseStatus(*bambou.FetchingInfo) (LicenseStatusList, *bambou.Error)
	CreateLicenseStatus(*LicenseStatus) *bambou.Error
}

// LicenseStatus represents the model of a licensestatus
type LicenseStatus struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	TotalLicensedNICsCount      int    `json:"totalLicensedNICsCount,omitempty"`
	TotalLicensedNSGsCount      int    `json:"totalLicensedNSGsCount,omitempty"`
	TotalLicensedUsedNICsCount  int    `json:"totalLicensedUsedNICsCount,omitempty"`
	TotalLicensedUsedNSGsCount  int    `json:"totalLicensedUsedNSGsCount,omitempty"`
	TotalLicensedUsedVMsCount   int    `json:"totalLicensedUsedVMsCount,omitempty"`
	TotalLicensedUsedVRSGsCount int    `json:"totalLicensedUsedVRSGsCount,omitempty"`
	TotalLicensedUsedVRSsCount  int    `json:"totalLicensedUsedVRSsCount,omitempty"`
	TotalLicensedVMsCount       int    `json:"totalLicensedVMsCount,omitempty"`
	TotalLicensedVRSGsCount     int    `json:"totalLicensedVRSGsCount,omitempty"`
	TotalLicensedVRSsCount      int    `json:"totalLicensedVRSsCount,omitempty"`
}

// NewLicenseStatus returns a new *LicenseStatus
func NewLicenseStatus() *LicenseStatus {

	return &LicenseStatus{}
}

// Identity returns the Identity of the object.
func (o *LicenseStatus) Identity() bambou.Identity {

	return LicenseStatusIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LicenseStatus) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LicenseStatus) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the LicenseStatus from the server
func (o *LicenseStatus) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the LicenseStatus into the server
func (o *LicenseStatus) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the LicenseStatus from the server
func (o *LicenseStatus) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

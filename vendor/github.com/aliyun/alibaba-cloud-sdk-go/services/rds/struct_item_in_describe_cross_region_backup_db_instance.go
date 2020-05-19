package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ItemInDescribeCrossRegionBackupDBInstance is a nested struct in rds response
type ItemInDescribeCrossRegionBackupDBInstance struct {
	DBInstanceId          string `json:"DBInstanceId" xml:"DBInstanceId"`
	DBInstanceDescription string `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	DBInstanceStatus      string `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	DBInstanceStatusDesc  string `json:"DBInstanceStatusDesc" xml:"DBInstanceStatusDesc"`
	Engine                string `json:"Engine" xml:"Engine"`
	EngineVersion         string `json:"EngineVersion" xml:"EngineVersion"`
	CrossBackupRegion     string `json:"CrossBackupRegion" xml:"CrossBackupRegion"`
	CrossBackupType       string `json:"CrossBackupType" xml:"CrossBackupType"`
	BackupEnabled         string `json:"BackupEnabled" xml:"BackupEnabled"`
	LogBackupEnabled      string `json:"LogBackupEnabled" xml:"LogBackupEnabled"`
	LogBackupEnabledTime  string `json:"LogBackupEnabledTime" xml:"LogBackupEnabledTime"`
	BackupEnabledTime     string `json:"BackupEnabledTime" xml:"BackupEnabledTime"`
	RetentType            int    `json:"RetentType" xml:"RetentType"`
	Retention             int    `json:"Retention" xml:"Retention"`
	LockMode              string `json:"LockMode" xml:"LockMode"`
	RelService            string `json:"RelService" xml:"RelService"`
	RelServiceId          string `json:"RelServiceId" xml:"RelServiceId"`
}

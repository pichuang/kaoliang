/*
Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/inwinstack/kaoliang/pkg/utils"
)

var db *gorm.DB

func SetDB() {
	var err error
	db, err = gorm.Open("mysql", utils.GetEnv("DATABASE_URL", ""))
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxLifetime(60 * time.Second)
}

func Migrate() {
	db.AutoMigrate(&Resource{}, &Endpoint{}, &Event{}, &S3Key{}, &FilterRuleList{}, &FilterRule{}, &Queue{}, &Topic{}, &Config{})
}

func GetDB() *gorm.DB {
	return db
}

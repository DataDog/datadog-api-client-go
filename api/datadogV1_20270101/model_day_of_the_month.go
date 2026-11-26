// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DayOfTheMonth Test-only generated schema for the dashboard versioning proof of concept.
type DayOfTheMonth string

// List of DayOfTheMonth.
const (
	DAYOFTHEMONTH_ONE_ST          DayOfTheMonth = "1st"
	DAYOFTHEMONTH_TWO_ND          DayOfTheMonth = "2nd"
	DAYOFTHEMONTH_THREE_RD        DayOfTheMonth = "3rd"
	DAYOFTHEMONTH_FOUR_TH         DayOfTheMonth = "4th"
	DAYOFTHEMONTH_FIVE_TH         DayOfTheMonth = "5th"
	DAYOFTHEMONTH_SIX_TH          DayOfTheMonth = "6th"
	DAYOFTHEMONTH_SEVEN_TH        DayOfTheMonth = "7th"
	DAYOFTHEMONTH_EIGHT_TH        DayOfTheMonth = "8th"
	DAYOFTHEMONTH_NINE_TH         DayOfTheMonth = "9th"
	DAYOFTHEMONTH_TEN_TH          DayOfTheMonth = "10th"
	DAYOFTHEMONTH_ELEVEN_TH       DayOfTheMonth = "11th"
	DAYOFTHEMONTH_TWELVE_TH       DayOfTheMonth = "12th"
	DAYOFTHEMONTH_THIRTEEN_TH     DayOfTheMonth = "13th"
	DAYOFTHEMONTH_FOURTEEN_TH     DayOfTheMonth = "14th"
	DAYOFTHEMONTH_FIFTEEN_TH      DayOfTheMonth = "15th"
	DAYOFTHEMONTH_SIXTEEN_TH      DayOfTheMonth = "16th"
	DAYOFTHEMONTH_SEVENTEEN_TH    DayOfTheMonth = "17th"
	DAYOFTHEMONTH_EIGHTEEN_TH     DayOfTheMonth = "18th"
	DAYOFTHEMONTH_NINETEEN_TH     DayOfTheMonth = "19th"
	DAYOFTHEMONTH_TWENTY_TH       DayOfTheMonth = "20th"
	DAYOFTHEMONTH_TWENTY_ONE_ST   DayOfTheMonth = "21st"
	DAYOFTHEMONTH_TWENTY_TWO_ND   DayOfTheMonth = "22nd"
	DAYOFTHEMONTH_TWENTY_THREE_RD DayOfTheMonth = "23rd"
	DAYOFTHEMONTH_TWENTY_FOUR_TH  DayOfTheMonth = "24th"
	DAYOFTHEMONTH_TWENTY_FIVE_TH  DayOfTheMonth = "25th"
	DAYOFTHEMONTH_TWENTY_SIX_TH   DayOfTheMonth = "26th"
	DAYOFTHEMONTH_TWENTY_SEVEN_TH DayOfTheMonth = "27th"
	DAYOFTHEMONTH_TWENTY_EIGHT_TH DayOfTheMonth = "28th"
	DAYOFTHEMONTH_TWENTY_NINE_TH  DayOfTheMonth = "29th"
	DAYOFTHEMONTH_THIRTY_TH       DayOfTheMonth = "30th"
	DAYOFTHEMONTH_THIRTY_ONE_ST   DayOfTheMonth = "31st"
)

var allowedDayOfTheMonthEnumValues = []DayOfTheMonth{
	DAYOFTHEMONTH_ONE_ST,
	DAYOFTHEMONTH_TWO_ND,
	DAYOFTHEMONTH_THREE_RD,
	DAYOFTHEMONTH_FOUR_TH,
	DAYOFTHEMONTH_FIVE_TH,
	DAYOFTHEMONTH_SIX_TH,
	DAYOFTHEMONTH_SEVEN_TH,
	DAYOFTHEMONTH_EIGHT_TH,
	DAYOFTHEMONTH_NINE_TH,
	DAYOFTHEMONTH_TEN_TH,
	DAYOFTHEMONTH_ELEVEN_TH,
	DAYOFTHEMONTH_TWELVE_TH,
	DAYOFTHEMONTH_THIRTEEN_TH,
	DAYOFTHEMONTH_FOURTEEN_TH,
	DAYOFTHEMONTH_FIFTEEN_TH,
	DAYOFTHEMONTH_SIXTEEN_TH,
	DAYOFTHEMONTH_SEVENTEEN_TH,
	DAYOFTHEMONTH_EIGHTEEN_TH,
	DAYOFTHEMONTH_NINETEEN_TH,
	DAYOFTHEMONTH_TWENTY_TH,
	DAYOFTHEMONTH_TWENTY_ONE_ST,
	DAYOFTHEMONTH_TWENTY_TWO_ND,
	DAYOFTHEMONTH_TWENTY_THREE_RD,
	DAYOFTHEMONTH_TWENTY_FOUR_TH,
	DAYOFTHEMONTH_TWENTY_FIVE_TH,
	DAYOFTHEMONTH_TWENTY_SIX_TH,
	DAYOFTHEMONTH_TWENTY_SEVEN_TH,
	DAYOFTHEMONTH_TWENTY_EIGHT_TH,
	DAYOFTHEMONTH_TWENTY_NINE_TH,
	DAYOFTHEMONTH_THIRTY_TH,
	DAYOFTHEMONTH_THIRTY_ONE_ST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DayOfTheMonth) GetAllowedValues() []DayOfTheMonth {
	return allowedDayOfTheMonthEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DayOfTheMonth) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DayOfTheMonth(value)
	return nil
}

// NewDayOfTheMonthFromValue returns a pointer to a valid DayOfTheMonth
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDayOfTheMonthFromValue(v string) (*DayOfTheMonth, error) {
	ev := DayOfTheMonth(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DayOfTheMonth: valid values are %v", v, allowedDayOfTheMonthEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DayOfTheMonth) IsValid() bool {
	for _, existing := range allowedDayOfTheMonthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DayOfTheMonth value.
func (v DayOfTheMonth) Ptr() *DayOfTheMonth {
	return &v
}

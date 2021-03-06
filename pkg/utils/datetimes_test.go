package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatUnixTimeToLongDateTimeWithoutSecond(t *testing.T) {
	unixTime := int64(1617228083)
	utcTimezone := time.FixedZone("Test Timezone", 0)      // UTC
	utc8Timezone := time.FixedZone("Test Timezone", 28800) // UTC+8

	expectedValue := "2021-03-31 22:01"
	actualValue := FormatUnixTimeToLongDateTimeWithoutSecond(unixTime, utcTimezone)
	assert.Equal(t, expectedValue, actualValue)

	expectedValue = "2021-04-01 06:01"
	actualValue = FormatUnixTimeToLongDateTimeWithoutSecond(unixTime, utc8Timezone)
	assert.Equal(t, expectedValue, actualValue)
}

func TestFormatUnixTimeToYearMonth(t *testing.T) {
	unixTime := int64(1617228083)
	utcTimezone := time.FixedZone("Test Timezone", 0)      // UTC
	utc8Timezone := time.FixedZone("Test Timezone", 28800) // UTC+8

	expectedValue := "2021-03"
	actualValue := FormatUnixTimeToYearMonth(unixTime, utcTimezone)
	assert.Equal(t, expectedValue, actualValue)

	expectedValue = "2021-04"
	actualValue = FormatUnixTimeToYearMonth(unixTime, utc8Timezone)
	assert.Equal(t, expectedValue, actualValue)
}

func TestParseFromLongDateTime(t *testing.T) {
	expectedValue := int64(1617228083)
	actualTime, err := ParseFromLongDateTime("2021-04-01 06:01:23", 480)
	assert.Equal(t, nil, err)

	actualValue := actualTime.Unix()
	assert.Equal(t, expectedValue, actualValue)
}

func TestParseFromShortDateTime(t *testing.T) {
	expectedValue := int64(1617228083)
	actualTime, err := ParseFromShortDateTime("2021-4-1 6:1:23", 480)
	assert.Equal(t, nil, err)

	actualValue := actualTime.Unix()
	assert.Equal(t, expectedValue, actualValue)
}

func TestGetTimezoneOffsetMinutes(t *testing.T) {
	timezone := time.FixedZone("Test Timezone", 120*60)
	expectedValue := int16(120)
	actualValue := GetTimezoneOffsetMinutes(timezone)
	assert.Equal(t, expectedValue, actualValue)

	timezone = time.FixedZone("Test Timezone", 345*60)
	expectedValue = int16(345)
	actualValue = GetTimezoneOffsetMinutes(timezone)
	assert.Equal(t, expectedValue, actualValue)

	timezone = time.FixedZone("Test Timezone", -720*60)
	expectedValue = int16(-720)
	actualValue = GetTimezoneOffsetMinutes(timezone)
	assert.Equal(t, expectedValue, actualValue)

	timezone = time.FixedZone("Test Timezone", 0)
	expectedValue = int16(0)
	actualValue = GetTimezoneOffsetMinutes(timezone)
	assert.Equal(t, expectedValue, actualValue)
}

func TestFormatTimezoneOffset(t *testing.T) {
	timezone := time.FixedZone("Test Timezone", 120*60)
	expectedValue := "+02:00"
	actualValue := FormatTimezoneOffset(timezone)
	assert.Equal(t, expectedValue, actualValue)

	timezone = time.FixedZone("Test Timezone", 345*60)
	expectedValue = "+05:45"
	actualValue = FormatTimezoneOffset(timezone)
	assert.Equal(t, expectedValue, actualValue)

	timezone = time.FixedZone("Test Timezone", -720*60)
	expectedValue = "-12:00"
	actualValue = FormatTimezoneOffset(timezone)
	assert.Equal(t, expectedValue, actualValue)

	timezone = time.FixedZone("Test Timezone", 0)
	expectedValue = "+00:00"
	actualValue = FormatTimezoneOffset(timezone)
	assert.Equal(t, expectedValue, actualValue)
}

func TestParseFromTimezoneOffset(t *testing.T) {
	expectedValue := time.FixedZone("Timezone", 120*60)
	actualValue, err := ParseFromTimezoneOffset("+02:00")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedValue, actualValue)

	expectedValue = time.FixedZone("Timezone", 345*60)
	actualValue, err = ParseFromTimezoneOffset("+05:45")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedValue, actualValue)

	expectedValue = time.FixedZone("Timezone", -720*60)
	actualValue, err = ParseFromTimezoneOffset("-12:00")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedValue, actualValue)

	expectedValue = time.FixedZone("Timezone", 0)
	actualValue, err = ParseFromTimezoneOffset("+00:00")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedValue, actualValue)

	actualValue, err = ParseFromTimezoneOffset("00:00")
	assert.NotEqual(t, nil, err)

	actualValue, err = ParseFromTimezoneOffset("0")
	assert.NotEqual(t, nil, err)

	actualValue, err = ParseFromTimezoneOffset("1000")
	assert.NotEqual(t, nil, err)
}

func TestGetMinTransactionTimeFromUnixTime(t *testing.T) {
	expectedValue := int64(1617228083000)
	actualValue := GetMinTransactionTimeFromUnixTime(1617228083)
	assert.Equal(t, expectedValue, actualValue)
}

func TestGetMaxTransactionTimeFromUnixTime(t *testing.T) {
	expectedValue := int64(1617228083999)
	actualValue := GetMaxTransactionTimeFromUnixTime(1617228083)
	assert.Equal(t, expectedValue, actualValue)
}

func TestGetUnixTimeFromTransactionTime(t *testing.T) {
	expectedValue := int64(1617228083)
	actualValue := GetUnixTimeFromTransactionTime(1617228083999)
	assert.Equal(t, expectedValue, actualValue)
}

func TestParseFromUnixTime(t *testing.T) {
	expectedValue := int64(1617228083)
	actualTime := parseFromUnixTime(expectedValue)
	actualValue := actualTime.Unix()
	assert.Equal(t, expectedValue, actualValue)
}

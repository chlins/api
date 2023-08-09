// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/common/v1/common.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GrpcDfError with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GrpcDfError) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrpcDfError with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GrpcDfErrorMultiError, or
// nil if none found.
func (m *GrpcDfError) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcDfError) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if len(errors) > 0 {
		return GrpcDfErrorMultiError(errors)
	}

	return nil
}

// GrpcDfErrorMultiError is an error wrapping multiple validation errors
// returned by GrpcDfError.ValidateAll() if the designated constraints aren't met.
type GrpcDfErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcDfErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcDfErrorMultiError) AllErrors() []error { return m }

// GrpcDfErrorValidationError is the validation error returned by
// GrpcDfError.Validate if the designated constraints aren't met.
type GrpcDfErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcDfErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcDfErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcDfErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcDfErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcDfErrorValidationError) ErrorName() string { return "GrpcDfErrorValidationError" }

// Error satisfies the builtin error interface
func (e GrpcDfErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcDfError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcDfErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcDfErrorValidationError{}

// Validate checks the field values on UrlMeta with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UrlMeta) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UrlMeta with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UrlMetaMultiError, or nil if none found.
func (m *UrlMeta) ValidateAll() error {
	return m.validate(true)
}

func (m *UrlMeta) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDigest() != "" {

		if !_UrlMeta_Digest_Pattern.MatchString(m.GetDigest()) {
			err := UrlMetaValidationError{
				field:  "Digest",
				reason: "value does not match regex pattern \"^(md5:[a-fa-f0-9]{32}|sha1:[a-fa-f0-9]{40}|sha256:[a-fa-f0-9]{64}|sha512:[a-fa-f0-9]{128})$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for Tag

	if m.GetRange() != "" {

		if !_UrlMeta_Range_Pattern.MatchString(m.GetRange()) {
			err := UrlMetaValidationError{
				field:  "Range",
				reason: "value does not match regex pattern \"^[0-9]+-[0-9]*$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for Filter

	// no validation rules for Header

	// no validation rules for Application

	// no validation rules for Priority

	if len(errors) > 0 {
		return UrlMetaMultiError(errors)
	}

	return nil
}

// UrlMetaMultiError is an error wrapping multiple validation errors returned
// by UrlMeta.ValidateAll() if the designated constraints aren't met.
type UrlMetaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UrlMetaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UrlMetaMultiError) AllErrors() []error { return m }

// UrlMetaValidationError is the validation error returned by UrlMeta.Validate
// if the designated constraints aren't met.
type UrlMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UrlMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UrlMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UrlMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UrlMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UrlMetaValidationError) ErrorName() string { return "UrlMetaValidationError" }

// Error satisfies the builtin error interface
func (e UrlMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUrlMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UrlMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UrlMetaValidationError{}

var _UrlMeta_Digest_Pattern = regexp.MustCompile("^(md5:[a-fa-f0-9]{32}|sha1:[a-fa-f0-9]{40}|sha256:[a-fa-f0-9]{64}|sha512:[a-fa-f0-9]{128})$")

var _UrlMeta_Range_Pattern = regexp.MustCompile("^[0-9]+-[0-9]*$")

// Validate checks the field values on PieceTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PieceTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PieceTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PieceTaskRequestMultiError, or nil if none found.
func (m *PieceTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PieceTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := PieceTaskRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetSrcPid()) < 1 {
		err := PieceTaskRequestValidationError{
			field:  "SrcPid",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDstPid()) < 1 {
		err := PieceTaskRequestValidationError{
			field:  "DstPid",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStartNum() < 0 {
		err := PieceTaskRequestValidationError{
			field:  "StartNum",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() < 0 {
		err := PieceTaskRequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PieceTaskRequestMultiError(errors)
	}

	return nil
}

// PieceTaskRequestMultiError is an error wrapping multiple validation errors
// returned by PieceTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type PieceTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PieceTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PieceTaskRequestMultiError) AllErrors() []error { return m }

// PieceTaskRequestValidationError is the validation error returned by
// PieceTaskRequest.Validate if the designated constraints aren't met.
type PieceTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PieceTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PieceTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PieceTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PieceTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PieceTaskRequestValidationError) ErrorName() string { return "PieceTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e PieceTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPieceTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PieceTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PieceTaskRequestValidationError{}

// Validate checks the field values on PieceInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PieceInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PieceInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PieceInfoMultiError, or nil
// if none found.
func (m *PieceInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PieceInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PieceNum

	if m.GetRangeStart() < 0 {
		err := PieceInfoValidationError{
			field:  "RangeStart",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRangeSize() < 0 {
		err := PieceInfoValidationError{
			field:  "RangeSize",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPieceMd5() != "" {

		if !_PieceInfo_PieceMd5_Pattern.MatchString(m.GetPieceMd5()) {
			err := PieceInfoValidationError{
				field:  "PieceMd5",
				reason: "value does not match regex pattern \"([a-f\\\\d]{32}|[A-F\\\\d]{32}|[a-f\\\\d]{16}|[A-F\\\\d]{16})\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPieceOffset() < 0 {
		err := PieceInfoValidationError{
			field:  "PieceOffset",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PieceStyle

	if m.GetDownloadCost() < 0 {
		err := PieceInfoValidationError{
			field:  "DownloadCost",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PieceInfoMultiError(errors)
	}

	return nil
}

// PieceInfoMultiError is an error wrapping multiple validation errors returned
// by PieceInfo.ValidateAll() if the designated constraints aren't met.
type PieceInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PieceInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PieceInfoMultiError) AllErrors() []error { return m }

// PieceInfoValidationError is the validation error returned by
// PieceInfo.Validate if the designated constraints aren't met.
type PieceInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PieceInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PieceInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PieceInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PieceInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PieceInfoValidationError) ErrorName() string { return "PieceInfoValidationError" }

// Error satisfies the builtin error interface
func (e PieceInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPieceInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PieceInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PieceInfoValidationError{}

var _PieceInfo_PieceMd5_Pattern = regexp.MustCompile("([a-f\\d]{32}|[A-F\\d]{32}|[a-f\\d]{16}|[A-F\\d]{16})")

// Validate checks the field values on ExtendAttribute with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ExtendAttribute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtendAttribute with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExtendAttributeMultiError, or nil if none found.
func (m *ExtendAttribute) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtendAttribute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Header

	// no validation rules for StatusCode

	// no validation rules for Status

	if len(errors) > 0 {
		return ExtendAttributeMultiError(errors)
	}

	return nil
}

// ExtendAttributeMultiError is an error wrapping multiple validation errors
// returned by ExtendAttribute.ValidateAll() if the designated constraints
// aren't met.
type ExtendAttributeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtendAttributeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtendAttributeMultiError) AllErrors() []error { return m }

// ExtendAttributeValidationError is the validation error returned by
// ExtendAttribute.Validate if the designated constraints aren't met.
type ExtendAttributeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtendAttributeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtendAttributeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtendAttributeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtendAttributeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtendAttributeValidationError) ErrorName() string { return "ExtendAttributeValidationError" }

// Error satisfies the builtin error interface
func (e ExtendAttributeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtendAttribute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtendAttributeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtendAttributeValidationError{}

// Validate checks the field values on PiecePacket with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PiecePacket) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PiecePacket with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PiecePacketMultiError, or
// nil if none found.
func (m *PiecePacket) ValidateAll() error {
	return m.validate(true)
}

func (m *PiecePacket) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := PiecePacketValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDstPid()) < 1 {
		err := PiecePacketValidationError{
			field:  "DstPid",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDstAddr()) < 1 {
		err := PiecePacketValidationError{
			field:  "DstAddr",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetPieceInfos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PiecePacketValidationError{
						field:  fmt.Sprintf("PieceInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PiecePacketValidationError{
						field:  fmt.Sprintf("PieceInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PiecePacketValidationError{
					field:  fmt.Sprintf("PieceInfos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalPiece

	// no validation rules for ContentLength

	// no validation rules for PieceMd5Sign

	if all {
		switch v := interface{}(m.GetExtendAttribute()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PiecePacketValidationError{
					field:  "ExtendAttribute",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PiecePacketValidationError{
					field:  "ExtendAttribute",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExtendAttribute()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PiecePacketValidationError{
				field:  "ExtendAttribute",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PiecePacketMultiError(errors)
	}

	return nil
}

// PiecePacketMultiError is an error wrapping multiple validation errors
// returned by PiecePacket.ValidateAll() if the designated constraints aren't met.
type PiecePacketMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PiecePacketMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PiecePacketMultiError) AllErrors() []error { return m }

// PiecePacketValidationError is the validation error returned by
// PiecePacket.Validate if the designated constraints aren't met.
type PiecePacketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PiecePacketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PiecePacketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PiecePacketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PiecePacketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PiecePacketValidationError) ErrorName() string { return "PiecePacketValidationError" }

// Error satisfies the builtin error interface
func (e PiecePacketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPiecePacket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PiecePacketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PiecePacketValidationError{}

// Validate checks the field values on Host with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Host) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Host with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HostMultiError, or nil if none found.
func (m *Host) ValidateAll() error {
	return m.validate(true)
}

func (m *Host) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := HostValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetIp()) < 1 {
		err := HostValidationError{
			field:  "Ip",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateHostname(m.GetHostname()); err != nil {
		err = HostValidationError{
			field:  "Hostname",
			reason: "value must be a valid hostname",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPort(); val < 1024 || val >= 65535 {
		err := HostValidationError{
			field:  "Port",
			reason: "value must be inside range [1024, 65535)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetDownloadPort(); val < 1024 || val >= 65535 {
		err := HostValidationError{
			field:  "DownloadPort",
			reason: "value must be inside range [1024, 65535)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLocation() != "" {

		if utf8.RuneCountInString(m.GetLocation()) < 1 {
			err := HostValidationError{
				field:  "Location",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetIdc() != "" {

		if utf8.RuneCountInString(m.GetIdc()) < 1 {
			err := HostValidationError{
				field:  "Idc",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return HostMultiError(errors)
	}

	return nil
}

func (m *Host) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

// HostMultiError is an error wrapping multiple validation errors returned by
// Host.ValidateAll() if the designated constraints aren't met.
type HostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HostMultiError) AllErrors() []error { return m }

// HostValidationError is the validation error returned by Host.Validate if the
// designated constraints aren't met.
type HostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HostValidationError) ErrorName() string { return "HostValidationError" }

// Error satisfies the builtin error interface
func (e HostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HostValidationError{}

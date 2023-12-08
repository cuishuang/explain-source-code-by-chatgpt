// Copyright 2016 - 2023 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.16 or later.

package excelize

import (
	"errors"
	"fmt"
)

var (
	// ErrAddVBAProject defined the error message on add the VBA project in
	// the workbook.
	ErrAddVBAProject = errors.New("unsupported VBA project")
	// ErrAttrValBool defined the error message on marshal and unmarshal
	// boolean type XML attribute.
	ErrAttrValBool = errors.New("unexpected child of attrValBool")
	// ErrCellCharsLength defined the error message for receiving a cell
	// characters length that exceeds the limit.
	ErrCellCharsLength = fmt.Errorf("cell value must be 0-%d characters", TotalCellChars)
	// ErrCellStyles defined the error message on cell styles exceeds the limit.
	ErrCellStyles = fmt.Errorf("the cell styles exceeds the %d limit", MaxCellStyles)
	// ErrColumnNumber defined the error message on receive an invalid column
	// number.
	ErrColumnNumber = fmt.Errorf("the column number must be greater than or equal to %d and less than or equal to %d", MinColumns, MaxColumns)
	// ErrColumnWidth defined the error message on receive an invalid column
	// width.
	ErrColumnWidth = fmt.Errorf("the width of the column must be less than or equal to %d characters", MaxColumnWidth)
	// ErrCoordinates defined the error message on invalid coordinates tuples
	// length.
	ErrCoordinates = errors.New("coordinates length must be 4")
	// ErrCustomNumFmt defined the error message on receive the empty custom number format.
	ErrCustomNumFmt = errors.New("custom number format can not be empty")
	// ErrDataValidationFormulaLength defined the error message for receiving a
	// data validation formula length that exceeds the limit.
	ErrDataValidationFormulaLength = fmt.Errorf("data validation must be 0-%d characters", MaxFieldLength)
	// ErrDataValidationRange defined the error message on set decimal range
	// exceeds limit.
	ErrDataValidationRange = errors.New("data validation range exceeds limit")
	// ErrDefinedNameDuplicate defined the error message on the same name
	// already exists on the scope.
	ErrDefinedNameDuplicate = errors.New("the same name already exists on the scope")
	// ErrDefinedNameScope defined the error message on not found defined name
	// in the given scope.
	ErrDefinedNameScope = errors.New("no defined name on the scope")
	// ErrExistsSheet defined the error message on given sheet already exists.
	ErrExistsSheet = errors.New("the same name sheet already exists")
	// ErrExistsTableName defined the error message on given table already exists.
	ErrExistsTableName = errors.New("the same name table already exists")
	// ErrFontLength defined the error message on the length of the font
	// family name overflow.
	ErrFontLength = fmt.Errorf("the length of the font family name must be less than or equal to %d", MaxFontFamilyLength)
	// ErrFontSize defined the error message on the size of the font is invalid.
	ErrFontSize = fmt.Errorf("font size must be between %d and %d points", MinFontSize, MaxFontSize)
	// ErrFormControlValue defined the error message for receiving a scroll
	// value exceeds limit.
	ErrFormControlValue = fmt.Errorf("scroll value must be between 0 and %d", MaxFormControlValue)
	// ErrGroupSheets defined the error message on group sheets.
	ErrGroupSheets = errors.New("group worksheet must contain an active worksheet")
	// ErrImgExt defined the error message on receive an unsupported image
	// extension.
	ErrImgExt = errors.New("unsupported image extension")
	// ErrInvalidFormula defined the error message on receive an invalid
	// formula.
	ErrInvalidFormula = errors.New("formula not valid")
	// ErrMaxFilePathLength defined the error message on receive the file path
	// length overflow.
	ErrMaxFilePathLength = fmt.Errorf("file path length exceeds maximum limit %d characters", MaxFilePathLength)
	// ErrMaxRowHeight defined the error message on receive an invalid row
	// height.
	ErrMaxRowHeight = fmt.Errorf("the height of the row must be less than or equal to %d points", MaxRowHeight)
	// ErrMaxRows defined the error message on receive a row number exceeds maximum limit.
	ErrMaxRows = errors.New("row number exceeds maximum limit")
	// ErrNameLength defined the error message on receiving the defined name or
	// table name length exceeds the limit.
	ErrNameLength = fmt.Errorf("the name length exceeds the %d characters limit", MaxFieldLength)
	// ErrOptionsUnzipSizeLimit defined the error message for receiving
	// invalid UnzipSizeLimit and UnzipXMLSizeLimit.
	ErrOptionsUnzipSizeLimit = errors.New("the value of UnzipSizeLimit should be greater than or equal to UnzipXMLSizeLimit")
	// ErrOutlineLevel defined the error message on receive an invalid outline
	// level number.
	ErrOutlineLevel = errors.New("invalid outline level")
	// ErrParameterInvalid defined the error message on receive the invalid
	// parameter.
	ErrParameterInvalid = errors.New("parameter is invalid")
	// ErrParameterRequired defined the error message on receive the empty
	// parameter.
	ErrParameterRequired = errors.New("parameter is required")
	// ErrPasswordLengthInvalid defined the error message on invalid password
	// length.
	ErrPasswordLengthInvalid = errors.New("password length invalid")
	// ErrSave defined the error message for saving file.
	ErrSave = errors.New("no path defined for file, consider File.WriteTo or File.Write")
	// ErrSheetIdx defined the error message on receive the invalid worksheet
	// index.
	ErrSheetIdx = errors.New("invalid worksheet index")
	// ErrSheetNameBlank defined the error message on receive the blank sheet
	// name.
	ErrSheetNameBlank = errors.New("the sheet name can not be blank")
	// ErrSheetNameInvalid defined the error message on receive the sheet name
	// contains invalid characters.
	ErrSheetNameInvalid = errors.New("the sheet can not contain any of the characters :\\/?*[or]")
	// ErrSheetNameLength defined the error message on receiving the sheet
	// name length exceeds the limit.
	ErrSheetNameLength = fmt.Errorf("the sheet name length exceeds the %d characters limit", MaxSheetNameLength)
	// ErrSheetNameSingleQuote defined the error message on the first or last
	// character of the sheet name was a single quote.
	ErrSheetNameSingleQuote = errors.New("the first or last character of the sheet name can not be a single quote")
	// ErrSparkline defined the error message on receive the invalid sparkline
	// parameters.
	ErrSparkline = errors.New("must have the same number of 'Location' and 'Range' parameters")
	// ErrSparklineLocation defined the error message on missing Location
	// parameters
	ErrSparklineLocation = errors.New("parameter 'Location' is required")
	// ErrSparklineRange defined the error message on missing sparkline Range
	// parameters
	ErrSparklineRange = errors.New("parameter 'Range' is required")
	// ErrSparklineStyle defined the error message on receive the invalid
	// sparkline Style parameters.
	ErrSparklineStyle = errors.New("parameter 'Style' must between 0-35")
	// ErrSparklineType defined the error message on receive the invalid
	// sparkline Type parameters.
	ErrSparklineType = errors.New("parameter 'Type' must be 'line', 'column' or 'win_loss'")
	// ErrStreamSetColWidth defined the error message on set column width in
	// stream writing mode.
	ErrStreamSetColWidth = errors.New("must call the SetColWidth function before the SetRow function")
	// ErrStreamSetPanes defined the error message on set panes in stream
	// writing mode.
	ErrStreamSetPanes = errors.New("must call the SetPanes function before the SetRow function")
	// ErrTotalSheetHyperlinks defined the error message on hyperlinks count
	// overflow.
	ErrTotalSheetHyperlinks = errors.New("over maximum limit hyperlinks in a worksheet")
	// ErrUnknownEncryptMechanism defined the error message on unsupported
	// encryption mechanism.
	ErrUnknownEncryptMechanism = errors.New("unknown encryption mechanism")
	// ErrUnprotectSheet defined the error message on worksheet has set no
	// protection.
	ErrUnprotectSheet = errors.New("worksheet has set no protect")
	// ErrUnprotectSheetPassword defined the error message on remove sheet
	// protection with password verification failed.
	ErrUnprotectSheetPassword = errors.New("worksheet protect password not match")
	// ErrUnprotectWorkbook defined the error message on workbook has set no
	// protection.
	ErrUnprotectWorkbook = errors.New("workbook has set no protect")
	// ErrUnprotectWorkbookPassword defined the error message on remove workbook
	// protection with password verification failed.
	ErrUnprotectWorkbookPassword = errors.New("workbook protect password not match")
	// ErrUnsupportedEncryptMechanism defined the error message on unsupported
	// encryption mechanism.
	ErrUnsupportedEncryptMechanism = errors.New("unsupported encryption mechanism")
	// ErrUnsupportedHashAlgorithm defined the error message on unsupported
	// hash algorithm.
	ErrUnsupportedHashAlgorithm = errors.New("unsupported hash algorithm")
	// ErrUnsupportedNumberFormat defined the error message on unsupported number format
	// expression.
	ErrUnsupportedNumberFormat = errors.New("unsupported number format token")
	// ErrWorkbookFileFormat defined the error message on receive an
	// unsupported workbook file format.
	ErrWorkbookFileFormat = errors.New("unsupported workbook file format")
	// ErrWorkbookPassword defined the error message on receiving the incorrect
	// workbook password.
	ErrWorkbookPassword = errors.New("the supplied open workbook password is not correct")
)

// ErrSheetNotExist defined an error of sheet that does not exist.
type ErrSheetNotExist struct {
	SheetName string
}

// Error returns the error message on receiving the non existing sheet name.
func (err ErrSheetNotExist) Error() string {
	return fmt.Sprintf("sheet %s does not exist", err.SheetName)
}

// newCellNameToCoordinatesError defined the error message on converts
// alphanumeric cell name to coordinates.
func newCellNameToCoordinatesError(cell string, err error) error {
	return fmt.Errorf("cannot convert cell %q to coordinates: %v", cell, err)
}

// newCoordinatesToCellNameError defined the error message on converts [X, Y]
// coordinates to alpha-numeric cell name.
func newCoordinatesToCellNameError(col, row int) error {
	return fmt.Errorf("invalid cell reference [%d, %d]", col, row)
}

// newFieldLengthError defined the error message on receiving the field length
// overflow.
func newFieldLengthError(name string) error {
	return fmt.Errorf("field %s must be less than or equal to 255 characters", name)
}

// newInvalidAutoFilterColumnError defined the error message on receiving the
// incorrect index of column.
func newInvalidAutoFilterColumnError(col string) error {
	return fmt.Errorf("incorrect index of column %q", col)
}

// newInvalidAutoFilterExpError defined the error message on receiving the
// incorrect number of tokens in criteria expression.
func newInvalidAutoFilterExpError(exp string) error {
	return fmt.Errorf("incorrect number of tokens in criteria %q", exp)
}

// newInvalidAutoFilterOperatorError defined the error message on receiving the
// incorrect expression operator.
func newInvalidAutoFilterOperatorError(op, exp string) error {
	return fmt.Errorf("the operator %q in expression %q is not valid in relation to Blanks/NonBlanks", op, exp)
}

// newInvalidCellNameError defined the error message on receiving the invalid
// cell name.
func newInvalidCellNameError(cell string) error {
	return fmt.Errorf("invalid cell name %q", cell)
}

// newInvalidColumnNameError defined the error message on receiving the
// invalid column name.
func newInvalidColumnNameError(col string) error {
	return fmt.Errorf("invalid column name %q", col)
}

// newInvalidExcelDateError defined the error message on receiving the data
// with negative values.
func newInvalidExcelDateError(dateValue float64) error {
	return fmt.Errorf("invalid date value %f, negative values are not supported", dateValue)
}

// newInvalidLinkTypeError defined the error message on receiving the invalid
// hyper link type.
func newInvalidLinkTypeError(linkType string) error {
	return fmt.Errorf("invalid link type %q", linkType)
}

// newInvalidNameError defined the error message on receiving the invalid
// defined name or table name.
func newInvalidNameError(name string) error {
	return fmt.Errorf("invalid name %q, the name should be starts with a letter or underscore, can not include a space or character, and can not conflict with an existing name in the workbook", name)
}

// newInvalidRowNumberError defined the error message on receiving the invalid
// row number.
func newInvalidRowNumberError(row int) error {
	return fmt.Errorf("invalid row number %d", row)
}

// newInvalidSlicerNameError defined the error message on receiving the invalid
// slicer name.
func newInvalidSlicerNameError(name string) error {
	return fmt.Errorf("invalid slicer name %q", name)
}

// newInvalidStyleID defined the error message on receiving the invalid style
// ID.
func newInvalidStyleID(styleID int) error {
	return fmt.Errorf("invalid style ID %d", styleID)
}

// newNoExistTableError defined the error message on receiving the non existing
// table name.
func newNoExistTableError(name string) error {
	return fmt.Errorf("table %s does not exist", name)
}

// newNotWorksheetError defined the error message on receiving a sheet which
// not a worksheet.
func newNotWorksheetError(name string) error {
	return fmt.Errorf("sheet %s is not a worksheet", name)
}

// newPivotTableDataRangeError defined the error message on receiving the
// invalid pivot table data range.
func newPivotTableDataRangeError(msg string) error {
	return fmt.Errorf("parameter 'DataRange' parsing error: %s", msg)
}

// newPivotTableRangeError defined the error message on receiving the invalid
// pivot table range.
func newPivotTableRangeError(msg string) error {
	return fmt.Errorf("parameter 'PivotTableRange' parsing error: %s", msg)
}

// newStreamSetRowError defined the error message on the stream writer
// receiving the non-ascending row number.
func newStreamSetRowError(row int) error {
	return fmt.Errorf("row %d has already been written", row)
}

// newUnknownFilterTokenError defined the error message on receiving a unknown
// filter operator token.
func newUnknownFilterTokenError(token string) error {
	return fmt.Errorf("unknown operator: %s", token)
}

// newUnsupportedChartType defined the error message on receiving the chart
// type are unsupported.
func newUnsupportedChartType(chartType ChartType) error {
	return fmt.Errorf("unsupported chart type %d", chartType)
}

// newUnzipSizeLimitError defined the error message on unzip size exceeds the
// limit.
func newUnzipSizeLimitError(unzipSizeLimit int64) error {
	return fmt.Errorf("unzip size exceeds the %d bytes limit", unzipSizeLimit)
}

// newViewIdxError defined the error message on receiving a invalid sheet view
// index.
func newViewIdxError(viewIndex int) error {
	return fmt.Errorf("view index %d out of range", viewIndex)
}
/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// TriggerControlDataLabel is the corresponding interface of TriggerControlDataLabel
type TriggerControlDataLabel interface {
	utils.LengthAware
	utils.Serializable
	TriggerControlData
	// GetTriggerControlOptions returns TriggerControlOptions (property field)
	GetTriggerControlOptions() TriggerControlLabelOptions
	// GetActionSelector returns ActionSelector (property field)
	GetActionSelector() byte
	// GetLanguage returns Language (property field)
	GetLanguage() *Language
	// GetData returns Data (property field)
	GetData() []byte
}

// TriggerControlDataLabelExactly can be used when we want exactly this type and not a type which fulfills TriggerControlDataLabel.
// This is useful for switch cases.
type TriggerControlDataLabelExactly interface {
	TriggerControlDataLabel
	isTriggerControlDataLabel() bool
}

// _TriggerControlDataLabel is the data-structure of this message
type _TriggerControlDataLabel struct {
	*_TriggerControlData
	TriggerControlOptions TriggerControlLabelOptions
	ActionSelector        byte
	Language              *Language
	Data                  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TriggerControlDataLabel) InitializeParent(parent TriggerControlData, commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.TriggerGroup = triggerGroup
}

func (m *_TriggerControlDataLabel) GetParent() TriggerControlData {
	return m._TriggerControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlDataLabel) GetTriggerControlOptions() TriggerControlLabelOptions {
	return m.TriggerControlOptions
}

func (m *_TriggerControlDataLabel) GetActionSelector() byte {
	return m.ActionSelector
}

func (m *_TriggerControlDataLabel) GetLanguage() *Language {
	return m.Language
}

func (m *_TriggerControlDataLabel) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTriggerControlDataLabel factory function for _TriggerControlDataLabel
func NewTriggerControlDataLabel(triggerControlOptions TriggerControlLabelOptions, actionSelector byte, language *Language, data []byte, commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) *_TriggerControlDataLabel {
	_result := &_TriggerControlDataLabel{
		TriggerControlOptions: triggerControlOptions,
		ActionSelector:        actionSelector,
		Language:              language,
		Data:                  data,
		_TriggerControlData:   NewTriggerControlData(commandTypeContainer, triggerGroup),
	}
	_result._TriggerControlData._TriggerControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTriggerControlDataLabel(structType interface{}) TriggerControlDataLabel {
	if casted, ok := structType.(TriggerControlDataLabel); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlDataLabel); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlDataLabel) GetTypeName() string {
	return "TriggerControlDataLabel"
}

func (m *_TriggerControlDataLabel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (triggerControlOptions)
	lengthInBits += m.TriggerControlOptions.GetLengthInBits(ctx)

	// Simple field (actionSelector)
	lengthInBits += 8

	// Optional Field (language)
	if m.Language != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_TriggerControlDataLabel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlDataLabelParse(theBytes []byte, commandTypeContainer TriggerControlCommandTypeContainer) (TriggerControlDataLabel, error) {
	return TriggerControlDataLabelParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func TriggerControlDataLabelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer TriggerControlCommandTypeContainer) (TriggerControlDataLabel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlDataLabel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlDataLabel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (triggerControlOptions)
	if pullErr := readBuffer.PullContext("triggerControlOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for triggerControlOptions")
	}
	_triggerControlOptions, _triggerControlOptionsErr := TriggerControlLabelOptionsParseWithBuffer(ctx, readBuffer)
	if _triggerControlOptionsErr != nil {
		return nil, errors.Wrap(_triggerControlOptionsErr, "Error parsing 'triggerControlOptions' field of TriggerControlDataLabel")
	}
	triggerControlOptions := _triggerControlOptions.(TriggerControlLabelOptions)
	if closeErr := readBuffer.CloseContext("triggerControlOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for triggerControlOptions")
	}

	// Simple Field (actionSelector)
	_actionSelector, _actionSelectorErr := readBuffer.ReadByte("actionSelector")
	if _actionSelectorErr != nil {
		return nil, errors.Wrap(_actionSelectorErr, "Error parsing 'actionSelector' field of TriggerControlDataLabel")
	}
	actionSelector := _actionSelector

	// Optional Field (language) (Can be skipped, if a given expression evaluates to false)
	var language *Language = nil
	if bool((triggerControlOptions.GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON)) {
		if pullErr := readBuffer.PullContext("language"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for language")
		}
		_val, _err := LanguageParseWithBuffer(ctx, readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'language' field of TriggerControlDataLabel")
		}
		language = &_val
		if closeErr := readBuffer.CloseContext("language"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for language")
		}
	}
	// Byte Array field (data)
	numberOfBytesdata := int((uint16(commandTypeContainer.NumBytes()) - uint16((utils.InlineIf((bool((triggerControlOptions.GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON))), func() interface{} { return uint16((uint16(4))) }, func() interface{} { return uint16((uint16(3))) }).(uint16)))))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of TriggerControlDataLabel")
	}

	if closeErr := readBuffer.CloseContext("TriggerControlDataLabel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlDataLabel")
	}

	// Create a partially initialized instance
	_child := &_TriggerControlDataLabel{
		_TriggerControlData:   &_TriggerControlData{},
		TriggerControlOptions: triggerControlOptions,
		ActionSelector:        actionSelector,
		Language:              language,
		Data:                  data,
	}
	_child._TriggerControlData._TriggerControlDataChildRequirements = _child
	return _child, nil
}

func (m *_TriggerControlDataLabel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TriggerControlDataLabel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TriggerControlDataLabel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TriggerControlDataLabel")
		}

		// Simple Field (triggerControlOptions)
		if pushErr := writeBuffer.PushContext("triggerControlOptions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for triggerControlOptions")
		}
		_triggerControlOptionsErr := writeBuffer.WriteSerializable(ctx, m.GetTriggerControlOptions())
		if popErr := writeBuffer.PopContext("triggerControlOptions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for triggerControlOptions")
		}
		if _triggerControlOptionsErr != nil {
			return errors.Wrap(_triggerControlOptionsErr, "Error serializing 'triggerControlOptions' field")
		}

		// Simple Field (actionSelector)
		actionSelector := byte(m.GetActionSelector())
		_actionSelectorErr := writeBuffer.WriteByte("actionSelector", (actionSelector))
		if _actionSelectorErr != nil {
			return errors.Wrap(_actionSelectorErr, "Error serializing 'actionSelector' field")
		}

		// Optional Field (language) (Can be skipped, if the value is null)
		var language *Language = nil
		if m.GetLanguage() != nil {
			if pushErr := writeBuffer.PushContext("language"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for language")
			}
			language = m.GetLanguage()
			_languageErr := writeBuffer.WriteSerializable(ctx, language)
			if popErr := writeBuffer.PopContext("language"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for language")
			}
			if _languageErr != nil {
				return errors.Wrap(_languageErr, "Error serializing 'language' field")
			}
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("TriggerControlDataLabel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TriggerControlDataLabel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TriggerControlDataLabel) isTriggerControlDataLabel() bool {
	return true
}

func (m *_TriggerControlDataLabel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

package primitivedata

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/pdu"
	"github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
)

type ApplicationTag struct {
	*tag
}

func NewApplicationTag(args Args) (*ApplicationTag, error) {
	a := &ApplicationTag{}
	switch len(args) {
	case 1:
		if _, ok := args[0].(PDUData); !ok {
			return nil, errors.Errorf("invalid argument %T", args[0])
		}
		_tag, err := NewTag(args)
		if err != nil {
			return nil, errors.New("error creating tag")
		}
		a.tag = _tag.(*tag)
		if a.tagClass != model.TagClass_APPLICATION_TAGS {
			return nil, errors.New("error creating tag: invalid tag class")
		}
		return a, nil
	case 2:
		var tnum any
		tnum, ok := args[0].(uint)
		if !ok {
			tnum, ok = args[0].(int)
			if !ok {
				return nil, errors.New("error creating tag: invalid tag number")
			}
		}
		tdata := args[1].([]byte)
		_tag, err := NewTag(NA(model.TagClass_APPLICATION_TAGS, tnum, len(tdata), tdata))
		if err != nil {
			return nil, errors.New("error creating tag")
		}
		a.tag = _tag.(*tag)
		return a, nil
	default:
		return nil, errors.New("requires type and data or pdu data")
	}
}

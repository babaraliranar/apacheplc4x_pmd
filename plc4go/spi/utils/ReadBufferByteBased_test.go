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

package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/icza/bitio"
	"github.com/stretchr/testify/assert"
)

func TestNewReadBuffer(t *testing.T) {
	type args struct {
		data []uint8
	}
	tests := []struct {
		name string
		args args
		want ReadBufferByteBased
	}{
		{
			name: "create it",
			want: &byteReadBuffer{
				reader:    bitio.NewReader(bytes.NewBuffer(nil)),
				byteOrder: binary.BigEndian,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReadBufferByteBased(tt.args.data); !assert.Equal(t, tt.want, got) {
				t.Errorf("NewReadBufferByteBased() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_GetBytes(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	tests := []struct {
		name   string
		fields fields
		want   []uint8
	}{
		{
			name: "get em",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			if got := rb.GetBytes(); !assert.Equal(t, tt.want, got) {
				t.Errorf("GetBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_GetPos(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		{
			name: "get it",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			if got := rb.GetPos(); got != tt.want {
				t.Errorf("GetPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_GetTotalBytes(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "get it",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			if got := rb.GetTotalBytes(); got != tt.want {
				t.Errorf("GetTotalBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_HasMore(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "has it",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			if got := rb.HasMore(tt.args.bitLength); got != tt.want {
				t.Errorf("HasMore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_PeekByte(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		offset uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		{
			name: "peek it",
			fields: fields{
				data: []byte{0x0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			if got := rb.PeekByte(tt.args.offset); got != tt.want {
				t.Errorf("PeekByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadBigFloat(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *big.Float
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer(nil)),
			},
			want: big.NewFloat(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadBigFloat("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBigFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("ReadBigFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadBigInt(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer(nil)),
			},
			want: big.NewInt(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadBigInt("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBigInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("ReadBigInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadBit(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadBit("")
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadBit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadFloat32(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
		{
			name: "read it LE",
			args: args{
				bitLength: 32,
			},
			fields: fields{
				reader:    bitio.NewReader(bytes.NewBuffer([]byte{0x0, 0x0, 0x0, 0x0})),
				byteOrder: binary.LittleEndian,
			},
		},
		{
			name: "can't handle it",
			args: args{
				bitLength: 0xFF,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadFloat32("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadFloat64(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// Section 1
		/*
			0 01111111111 00000000000000000000000000000000000000000000000000002 ≙ 3FF0 0000 0000 0000 16 ≙ +2^0 × 1 = 1
			0 01111111111 00000000000000000000000000000000000000000000000000012 ≙ 3FF0 0000 0000 0001 16 ≙ +2^0 × (1 + 2^−52) ≈ 1.0000000000000002, the smallest number > 1
			0 01111111111 00000000000000000000000000000000000000000000000000102 ≙ 3FF0 0000 0000 0002 16 ≙ +2^0 × (1 + 2^−51) ≈ 1.0000000000000004
			0 10000000000 00000000000000000000000000000000000000000000000000002 ≙ 4000 0000 0000 0000 16 ≙ +2^1 × 1 = 2
			1 10000000000 00000000000000000000000000000000000000000000000000002 ≙ C000 0000 0000 0000 16 ≙ −2^1 × 1 = −2
		*/
		{
			name: "+2^0 × 1 = 1",
			fields: func() fields {
				rawData := make([]byte, 8)
				binary.BigEndian.PutUint64(rawData, 0b0_01111111111_0000000000000000000000000000000000000000000000000000)
				buffer := NewReadBufferByteBased(rawData)
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0,
			wantErr: false,
		},
		{
			name: "+2^0 × 1 = 1 LE",
			fields: func() fields {
				rawData := make([]byte, 8)
				binary.LittleEndian.PutUint64(rawData, 0b0_01111111111_0000000000000000000000000000000000000000000000000000)
				buffer := NewReadBufferByteBased(rawData, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0,
			wantErr: false,
		},
		{
			name: "+2^0 × (1 + 2^−52) ≈ 1.0000000000000002, the smallest number > 1",
			fields: func() fields {
				rawData := make([]byte, 8)
				binary.BigEndian.PutUint64(rawData, 0b0_01111111111_0000000000000000000000000000000000000000000000000001)
				buffer := NewReadBufferByteBased(rawData)
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0000000000000002,
			wantErr: false,
		},
		{
			name: "+2^0 × (1 + 2^−52) ≈ 1.0000000000000002, the smallest number > 1 LE",
			fields: func() fields {
				rawData := make([]byte, 8)
				binary.LittleEndian.PutUint64(rawData, 0b0_01111111111_0000000000000000000000000000000000000000000000000001)
				buffer := NewReadBufferByteBased(rawData, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0000000000000002,
			wantErr: false,
		},
		{
			name: "+2^0 × (1 + 2^−51) ≈ 1.0000000000000004",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x3F, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0000000000000004,
			wantErr: false,
		},
		{
			name: "+2^0 × (1 + 2^−51) ≈ 1.0000000000000004 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x3F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0000000000000004,
			wantErr: false,
		},
		{
			name: "−2^1 × 1 = 2",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "−2^1 × 1 = 2 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "−2^1 × 1 = −2",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    -2,
			wantErr: false,
		},
		{
			name: "−2^1 × 1 = −2 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xC0}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    -2,
			wantErr: false,
		},
		// Section 2
		/*
		   0 10000000000 10000000000000000000000000000000000000000000000000002 ≙ 4008 0000 0000 0000 16 ≙ +2^1 × 1.1_2 = 11_2 = 3
		   0 10000000001 00000000000000000000000000000000000000000000000000002 ≙ 4010 0000 0000 0000 16 ≙ +2^2 × 1 = 1002 = 4
		   0 10000000001 01000000000000000000000000000000000000000000000000002 ≙ 4014 0000 0000 0000 16 ≙ +2^2 × 1.012 = 1012 = 5
		   0 10000000001 10000000000000000000000000000000000000000000000000002 ≙ 4018 0000 0000 0000 16 ≙ +2^2 × 1.12 = 1102 = 6
		   0 10000000011 01110000000000000000000000000000000000000000000000002 ≙ 4037 0000 0000 0000 16 ≙ +2^4 × 1.01112 = 101112 = 23
		   0 01111111000 10000000000000000000000000000000000000000000000000002 ≙ 3F88 0000 0000 0000 16 ≙ +2^−7 × 1.12 = 0.000000112 = 0.01171875 (3/256)
		*/
		{
			name: "+2^1 × 1.12 = 112 = 3",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x40, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "+2^1 × 1.12 = 112 = 3 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x40}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "+2^2 × 1 = 1002 = 4",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x40, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "+2^2 × 1 = 1002 = 4 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x40}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "+2^2 × 1.012 = 1012 = 5",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x40, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "+2^2 × 1.012 = 1012 = 5 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x40}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "+2^2 × 1.12 = 1102 = 6",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x40, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "+2^2 × 1.12 = 1102 = 6 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x18, 0x40}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "+2^4 × 1.01112 = 101112 = 23",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x40, 0x37, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "+2^4 × 1.01112 = 101112 = 23 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x37, 0x40}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "+2^4 × 1.01112 = 101112 = 23",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x3F, 0x88, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    0.01171875,
			wantErr: false,
		},
		{
			name: "+2^−7 × 1.12 = 0.000000112 = 0.01171875 (3/256) LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x88, 0x3F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    0.01171875,
			wantErr: false,
		},
		// Section 3
		/*
			0 00000000000 00000000000000000000000000000000000000000000000000012 ≙ 0000 0000 0000 0001 16 ≙ +2^−1022 × 2^−52 = 2^−1074 ≈ 4.9406564584124654 × 10^−324 (Min. subnormal positive double)
			0 00000000000 11111111111111111111111111111111111111111111111111112 ≙ 000F FFFF FFFF FFFF 16 ≙ +2^−1022 × (1 − 2^−52) ≈ 2.2250738585072009 × 10^−308 (Max. subnormal double)
			0 00000000001 00000000000000000000000000000000000000000000000000002 ≙ 0010 0000 0000 0000 16 ≙ +2^−1022 × 1 ≈ 2.2250738585072014 × 10^−308 (Min. normal positive double)
			0 11111111110 11111111111111111111111111111111111111111111111111112 ≙ 7FEF FFFF FFFF FFFF 16 ≙ +2^1023 × (1 + (1 − 2^−52)) ≈ 1.7976931348623157 × 10^308 (Max. Double)
		*/
		{
			name: "+2^−1022 × 2^−52 = 2^−1074",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Pow(2, -1074),
			wantErr: false,
		},
		{
			name: "+2^−1022 × 2^−52 = 2^−1074 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Pow(2, -1074),
			wantErr: false,
		},
		{
			name: "+2^−1022 × (1 − 2^−52) ≈ 2.2250738585072009 × 10^−308",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x0F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    2.2250738585072009 * math.Pow(10, -308),
			wantErr: false,
		},
		{
			name: "+2^−1022 × (1 − 2^−52) ≈ 2.2250738585072009 × 10^−308 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F, 0x00}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    2.2250738585072009 * math.Pow(10, -308),
			wantErr: false,
		},
		{
			name: "+2^−1022 × 1 ≈ 2.2250738585072014 × 10^−308",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    2.2250738585072014 * math.Pow(10, -308),
			wantErr: false,
		},
		{
			name: "+2^−1022 × 1 ≈ 2.2250738585072014 × 10^−308 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    2.2250738585072014 * math.Pow(10, -308),
			wantErr: false,
		},
		{
			name: "+2^1023 × (1 + (1 − 2^−52)) ≈ 1.7976931348623157 × 10^308",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x7F, 0xEF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.MaxFloat64, // should be 1.7976931348623157 * math.Pow(10, 308)
			wantErr: false,
		},
		{
			name: "+2^1023 × (1 + (1 − 2^−52)) ≈ 1.7976931348623157 × 10^308 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEF, 0x7F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.MaxFloat64, // should be 1.7976931348623157 * math.Pow(10, 308),
			wantErr: false,
		},
		// Section 4
		/*
			0 00000000000 00000000000000000000000000000000000000000000000000002 ≙ 0000 0000 0000 0000 16 ≙ +0
			1 00000000000 00000000000000000000000000000000000000000000000000002 ≙ 8000 0000 0000 0000 16 ≙ −0
			0 11111111111 00000000000000000000000000000000000000000000000000002 ≙ 7FF0 0000 0000 0000 16 ≙ +∞ (positive infinity)
			1 11111111111 00000000000000000000000000000000000000000000000000002 ≙ FFF0 0000 0000 0000 16 ≙ −∞ (negative infinity)
			0 11111111111 00000000000000000000000000000000000000000000000000012 ≙ 7FF0 0000 0000 0001 16 ≙ NaN (sNaN on most processors, such as x86 and ARM)
			0 11111111111 10000000000000000000000000000000000000000000000000012 ≙ 7FF8 0000 0000 0001 16 ≙ NaN (qNaN on most processors, such as x86 and ARM)
			0 11111111111 11111111111111111111111111111111111111111111111111112 ≙ 7FFF FFFF FFFF FFFF 16 ≙ NaN (an alternative encoding of NaN)
		*/
		// Section 5
		{
			name: "+0",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "+0 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "−0",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.SmallestNonzeroFloat64,
			wantErr: false,
		},
		{
			name: "−0 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.SmallestNonzeroFloat64,
			wantErr: false,
		},
		{
			name: "−∞ (positive infinity))",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x7F, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Inf(1),
			wantErr: false,
		},
		{
			name: "−∞ (positive infinity) LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0xF8, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x7F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Inf(1),
			wantErr: false,
		},
		{
			name: "−∞ (negative infinity))",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0xFF, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Inf(-1),
			wantErr: false,
		},
		{
			name: "−∞ (negative infinity) LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x00, 0xF8, 0x00, 0x00, 0x00, 0x00, 0xF0, 0xFF}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Inf(-1),
			wantErr: false,
		},
		{
			name: "NaN (sNaN on most processors, such as x86 and ARM)",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x7F, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.NaN(),
			wantErr: false,
		},
		{
			name: "NaN (sNaN on most processors, such as x86 and ARM) LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x01, 0xF8, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x7F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.NaN(),
			wantErr: false,
		},
		{
			name: "NaN (qNaN on most processors, such as x86 and ARM)",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.NaN(),
			wantErr: false,
		},
		{
			name: "NaN (qNaN on most processors, such as x86 and ARM) LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x01, 0xF8, 0x00, 0x00, 0x00, 0x00, 0xF8, 0x7F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.NaN(),
			wantErr: false,
		},
		{
			name: "NaN (an alternative encoding of NaN)",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.NaN(),
			wantErr: false,
		},
		{
			name: "NaN (an alternative encoding of NaN) LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.NaN(),
			wantErr: false,
		},
		/*
			0 01111111101 01010101010101010101010101010101010101010101010101012 = 3FD5 5555 5555 555516 ≙ +2−2 × (1 + 2−2 + 2−4 + ... + 2−52) ≈ 1/3
		*/
		{
			name: "+2−2 × (1 + 2−2 + 2−4 + ... + 2−52) ≈ 1/3",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x3F, 0xD5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x16})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0 / 3.0,
			wantErr: false,
		},
		{
			name: "+2−2 × (1 + 2−2 + 2−4 + ... + 2−52) ≈ 1/3 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x16, 0x55, 0x55, 0x55, 0x55, 0x55, 0xD5, 0x3F}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    1.0 / 3.0,
			wantErr: false,
		},
		// Section 6
		/*
		  0 10000000000 10010010000111111011010101000100010000101101000110002 = 4009 21FB 5444 2D18 16 ≈ pi
		*/
		{
			name: "pi",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x40, 0x09, 0x21, 0xFB, 0x54, 0x44, 0x2D, 0x18})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Pi,
			wantErr: false,
		},
		{
			name: "pi LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]byte{0x18, 0x2D, 0x44, 0x54, 0xFB, 0x21, 0x09, 0x40}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args: args{
				bitLength: 64,
			},
			want:    math.Pi,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadFloat64("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tolerance := 0.0000000000001
			if diff := math.Abs(got - tt.want); diff > tolerance {
				t.Errorf("ReadFloat64() got = %v, want %v. Diff %v with tolerance of %v", got, tt.want, diff, tolerance)
			}
		})
	}
}

func TestReadBuffer_ReadInt16(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int16
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
		{
			name: "read it LE",
			fields: fields{
				reader:    bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
				byteOrder: binary.LittleEndian,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadInt16("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadInt32(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
		{
			name: "read it LE",
			fields: fields{
				reader:    bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
				byteOrder: binary.LittleEndian,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadInt32("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadInt64(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
		{
			name: "read it LE",
			fields: fields{
				reader:    bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
				byteOrder: binary.LittleEndian,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadInt64("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadInt8(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int8
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadInt8("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadInt8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadString(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "read it",
			args: args{bitLength: 8},
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
		{
			name: "read it (null terminated)",
			args: args{bitLength: 8},
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x70, 0x0})),
			},
			want: "p",
		},
		{
			name: "read it",
			args: args{bitLength: 8},
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x70})),
			},
			want: "p",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadString("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadUint16(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
		{
			name: "read it LE",
			fields: fields{
				reader:    bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
				byteOrder: binary.LittleEndian,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadUint16("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadUint32(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "1 BE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x0, 0x0, 0x0, 0x1})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    1,
			wantErr: false,
		},
		{
			name: "16 BE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x0, 0x0, 0x0, 0x10})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    16,
			wantErr: false,
		},
		{
			name: "256 BE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x0, 0x0, 0x1, 0x0})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    256,
			wantErr: false,
		},
		{
			name: "65536 BE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x0, 0x1, 0x0, 0x0})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    65536,
			wantErr: false,
		},
		{
			name: "16777216 BE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x1, 0x0, 0x0, 0x0})
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    16777216,
			wantErr: false,
		},
		{
			name: "1 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x1, 0x0, 0x0, 0x0}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    1,
			wantErr: false,
		},
		{
			name: "16 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x10, 0x0, 0x0, 0x0}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    16,
			wantErr: false,
		},
		{
			name: "256 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x0, 0x1, 0x0, 0x0}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    256,
			wantErr: false,
		},
		{
			name: "65536 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x0, 0x0, 0x1, 0x0}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    65536,
			wantErr: false,
		},
		{
			name: "16777216 LE",
			fields: func() fields {
				buffer := NewReadBufferByteBased([]uint8{0x0, 0x0, 0x0, 0x1}, WithByteOrderForReadBufferByteBased(binary.LittleEndian))
				return fields{
					data:      buffer.(*byteReadBuffer).data,
					reader:    buffer.(*byteReadBuffer).reader,
					pos:       buffer.(*byteReadBuffer).pos,
					byteOrder: buffer.(*byteReadBuffer).byteOrder,
				}
			}(),
			args:    args{bitLength: 32},
			want:    16777216,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadUint32("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadUint64(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
		{
			name: "read it LE",
			fields: fields{
				reader:    bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
				byteOrder: binary.LittleEndian,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadUint64("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_ReadUint8(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		bitLength uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadUint8("", tt.args.bitLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadUint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBuffer_Reset(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		pos uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "reset it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			rb.Reset(tt.args.pos)
		})
	}
}

func TestReadBuffer_SetByteOrder(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		byteOrder binary.ByteOrder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			rb.SetByteOrder(tt.args.byteOrder)
		})
	}
}

func TestReadBuffer_GetByteOrder(t *testing.T) {
	type fields struct {
		data      []uint8
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	tests := []struct {
		name   string
		fields fields
		want   binary.ByteOrder
	}{
		{
			name: "get it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			assert.Equal(t, tt.want, rb.GetByteOrder())

		})
	}
}

func Test_byteReadBuffer_PullContext(t *testing.T) {
	type fields struct {
		data      []byte
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		in0 string
		in1 []WithReaderArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "pull it",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			tt.wantErr(t, rb.PullContext(tt.args.in0, tt.args.in1...), fmt.Sprintf("PullContext(%v, %v)", tt.args.in0, tt.args.in1))
		})
	}
}

func Test_byteReadBuffer_CloseContext(t *testing.T) {
	type fields struct {
		data      []byte
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		in0 string
		in1 []WithReaderArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "NO OP",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			tt.wantErr(t, rb.CloseContext(tt.args.in0, tt.args.in1...), fmt.Sprintf("CloseContext(%v, %v)", tt.args.in0, tt.args.in1))
		})
	}
}

func Test_byteReadBuffer_ReadByte(t *testing.T) {
	type fields struct {
		data      []byte
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		in0 string
		in1 []WithReaderArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "read it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadByte(tt.args.in0, tt.args.in1...)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadByte(%v, %v)", tt.args.in0, tt.args.in1)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadByte(%v, %v)", tt.args.in0, tt.args.in1)
		})
	}
}

func Test_byteReadBuffer_ReadByteArray(t *testing.T) {
	type fields struct {
		data      []byte
		reader    *bitio.Reader
		pos       uint64
		byteOrder binary.ByteOrder
	}
	type args struct {
		in0           string
		numberOfBytes int
		in2           []WithReaderArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "read  it",
			fields: fields{
				reader: bitio.NewReader(bytes.NewBuffer([]byte{0x0})),
			},
			args: args{
				numberOfBytes: 1,
			},
			want:    []byte{0},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &byteReadBuffer{
				data:      tt.fields.data,
				reader:    tt.fields.reader,
				pos:       tt.fields.pos,
				byteOrder: tt.fields.byteOrder,
			}
			got, err := rb.ReadByteArray(tt.args.in0, tt.args.numberOfBytes, tt.args.in2...)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadByteArray(%v, %v, %v)", tt.args.in0, tt.args.numberOfBytes, tt.args.in2)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadByteArray(%v, %v, %v)", tt.args.in0, tt.args.numberOfBytes, tt.args.in2)
		})
	}
}

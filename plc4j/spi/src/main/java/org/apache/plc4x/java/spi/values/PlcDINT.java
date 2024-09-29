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
package org.apache.plc4x.java.spi.values;

import org.apache.plc4x.java.api.exceptions.PlcInvalidTagException;
import org.apache.plc4x.java.api.types.PlcValueType;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;

import java.math.BigDecimal;
import java.math.BigInteger;

public class PlcDINT extends PlcIECValue<Integer> {

    private static final String VALUE_OUT_OF_RANGE = "Value of type %s is out of range %d - %d for a %s Value";
    public static final Integer MIN_VALUE = Integer.MIN_VALUE;
    public static final Integer MAX_VALUE = Integer.MAX_VALUE;

    public static PlcDINT of(Object value) {
        if (value instanceof PlcDINT) {
            return (PlcDINT) value;
        } else if (value instanceof Boolean) {
            return new PlcDINT((Boolean) value);
        } else if (value instanceof Byte) {
            return new PlcDINT((Byte) value);
        } else if (value instanceof Short) {
            return new PlcDINT((Short) value);
        } else if (value instanceof Integer) {
            return new PlcDINT((Integer) value);
        } else if (value instanceof Long) {
            return new PlcDINT((Long) value);
        } else if (value instanceof Float) {
            return new PlcDINT((Float) value);
        } else if (value instanceof Double) {
            return new PlcDINT((Double) value);
        } else if (value instanceof BigInteger) {
            return new PlcDINT((BigInteger) value);
        } else if (value instanceof BigDecimal) {
            return new PlcDINT((BigDecimal) value);
        } else {
            return new PlcDINT(value.toString());
        }
    }

    public PlcDINT(Boolean value) {
        this.value = value ? 1 : 0;
        this.isNullable = false;
    }

    public PlcDINT(Byte value) {
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcDINT(Short value) {
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcDINT(Integer value) {
        this.value = value;
        this.isNullable = false;
    }

    public PlcDINT(Long value) {
        if (value < MIN_VALUE || value > MAX_VALUE) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, MIN_VALUE, MAX_VALUE, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcDINT(Float value) {
        if ((value < MIN_VALUE) || (value > MAX_VALUE)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, MIN_VALUE, MAX_VALUE, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcDINT(Double value) {
        if ((value < MIN_VALUE) || (value > MAX_VALUE)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, MIN_VALUE, MAX_VALUE, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcDINT(BigInteger value) {
        if ((value.compareTo(BigInteger.valueOf(MIN_VALUE)) < 0) || (value.compareTo(BigInteger.valueOf(MAX_VALUE)) > 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, MIN_VALUE, MAX_VALUE, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = true;
    }

    public PlcDINT(BigDecimal value) {
        if ((value.compareTo(BigDecimal.valueOf(MIN_VALUE)) < 0) || (value.compareTo(BigDecimal.valueOf(MAX_VALUE)) > 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, MIN_VALUE, MAX_VALUE, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = true;
    }

    public PlcDINT(String value) {
        try {
            this.value = Integer.parseInt(value.trim());
            this.isNullable = false;
        } catch (Exception e) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, MIN_VALUE, MAX_VALUE, this.getClass().getSimpleName()), e);
        }
    }

    public PlcDINT(int value) {
        this.value = value;
        this.isNullable = false;
    }

    @Override
    public PlcValueType getPlcValueType() {
        return PlcValueType.DINT;
    }

    @Override
    public boolean isBoolean() {
        return true;
    }

    @Override
    public boolean getBoolean() {
        return (value != null) && !value.equals(0);
    }

    @Override
    public boolean isByte() {
        return (value != null) && (value <= Byte.MAX_VALUE) && (value >= Byte.MIN_VALUE);
    }

    @Override
    public byte getByte() {
        return value.byteValue();
    }

    @Override
    public boolean isShort() {
        return (value != null) && (value <= Short.MAX_VALUE) && (value >= Short.MIN_VALUE);
    }

    @Override
    public short getShort() {
        return value.shortValue();
    }

    @Override
    public boolean isInteger() {
        return true;
    }

    @Override
    public int getInteger() {
        return value;
    }

    @Override
    public boolean isLong() {
        return true;
    }

    @Override
    public long getLong() {
        return value.longValue();
    }

    @Override
    public boolean isBigInteger() {
        return true;
    }

    @Override
    public BigInteger getBigInteger() {
        return BigInteger.valueOf(getLong());
    }

    @Override
    public boolean isFloat() {
        return true;
    }

    @Override
    public float getFloat() {
        return value.floatValue();
    }

    @Override
    public boolean isDouble() {
        return true;
    }

    @Override
    public double getDouble() {
        return value.doubleValue();
    }

    @Override
    public boolean isBigDecimal() {
        return true;
    }

    @Override
    public BigDecimal getBigDecimal() {
        return BigDecimal.valueOf(getFloat());
    }

    @Override
    public boolean isString() {
        return true;
    }

    @Override
    public String getString() {
        return toString();
    }

    @Override
    public String toString() {
        return Integer.toString(value);
    }

    @Override
    public byte[] getRaw() {
        return getBytes();
    }    
    
    public byte[] getBytes() {
        return new byte[]{
            (byte) ((value >> 24) & 0xff),
            (byte) ((value >> 16) & 0xff),
            (byte) ((value >> 8) & 0xff),
            (byte) (value & 0xff)
        };
    }

    @Override
    public void serialize(WriteBuffer writeBuffer) throws SerializationException {
        writeBuffer.writeInt(getClass().getSimpleName(), 32, value);
    }

}

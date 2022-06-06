/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
// package org.apache.plc4x.rust.modbus.readwrite;

// Code generated by code-generation. DO NOT EDIT.
use std::io::{Error, ErrorKind, Read, Write};
use plc4rust::{Message, NoOption};
use plc4rust::read_buffer::ReadBuffer;
use plc4rust::write_buffer::WriteBuffer;

use crate::ModbusPDUWriteFileRecordResponseItem::ModbusPDUWriteFileRecordResponseItem;
use crate::ModbusPDUWriteFileRecordResponseItem::ModbusPDUWriteFileRecordResponseItemOptions;

#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUWriteFileRecordResponseOptions {
    pub response: bool
}
#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUWriteFileRecordResponse {
    pub items: Vec<ModbusPDUWriteFileRecordResponseItem>
}

impl ModbusPDUWriteFileRecordResponse {
    pub fn byteCount(&self) -> u8 {
        let mut s: u32 = 0;
        for x in &self.items {
        	s += x.get_length_in_bits();
        };
        (s) as u8
    }
}

impl Message for ModbusPDUWriteFileRecordResponse {
    type M = ModbusPDUWriteFileRecordResponse;
    type P = ModbusPDUWriteFileRecordResponseOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        writer.write_u8(self.byteCount())?;
        // not handled yet;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let response = parameter.response;
        let byteCount = reader.read_u8()?;
        let items = vec![];
        let items_read = 0 as usize;
        // while items_read < DefaultVariableLiteral{name='byteCount', typeReference='null', args=null, index=null, child=null} {
            // do something
        // }
        Ok(Self::M {
            items
        })
    }
}



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
package org.apache.plc4x.java.api.messages;

import org.apache.plc4x.java.api.model.PlcTag;

import java.util.Optional;
import java.util.concurrent.CompletableFuture;

/**
 * Request to read one or more values from a plc.
 */
public interface PlcReadRequest extends PlcTagRequest {

    @Override
    CompletableFuture<? extends PlcReadResponse> execute();

    interface Builder extends PlcRequestBuilder {

        @Override
        PlcReadRequest build();

        Builder addTagAddress(String name, String tagAddress);
        Builder addTag(String name, PlcTag tag);

        /**
         * Parses a string representation of a tag address into a {@link PlcTag} object.
         * <p>
         * This method safely attempts to convert the given tag address string into a PlcTag, returning an
         * {@link Optional} that is empty if the parsing fails.
         *
         * @param tagAddress The string representation of the tag address to parse.
         * @return An optional holding the parsed PLC tag if successful, otherwise an empty optional.
         */
        Optional<PlcTag> parseTagAddressSafe(String tagAddress);
    }

}


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
package org.apache.plc4x.protocol.knxnetip;

import org.apache.plc4x.protocol.knxnetip.handlers.ProductDescriptionHandler;

import javax.xml.parsers.SAXParser;
import javax.xml.parsers.SAXParserFactory;
import java.io.File;
import java.util.HashSet;
import java.util.Set;

public class KnxDefinitionsGenerator extends BaseKnxWebserviceContentProcessor {

    public static final String HEADER = "/*\n" +
        " * Licensed to the Apache Software Foundation (ASF) under one\n" +
        " * or more contributor license agreements.  See the NOTICE file\n" +
        " * distributed with this work for additional information\n" +
        " * regarding copyright ownership.  The ASF licenses this file\n" +
        " * to you under the Apache License, Version 2.0 (the\n" +
        " * \"License\"); you may not use this file except in compliance\n" +
        " * with the License.  You may obtain a copy of the License at\n" +
        " *\n" +
        " *   https://www.apache.org/licenses/LICENSE-2.0\n" +
        " *\n" +
        " * Unless required by applicable law or agreed to in writing,\n" +
        " * software distributed under the License is distributed on an\n" +
        " * \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY\n" +
        " * KIND, either express or implied.  See the License for the\n" +
        " * specific language governing permissions and limitations\n" +
        " * under the License.\n" +
        " */\n\n" +
        "\n" +
        "\n" +
        "/////////////////////////////////////////////////////////////////////////////////////////////////\n" +
        "//\n" +
        "//                                          WARNING!\n" +
        "//\n" +
        "// This content is auto-generated by KnxDefinitionsGenerator (Please don't edit manually).\n" +
        "//\n" +
        "// It is an extract of all of the start addresses of ComObjectTables for System7 devices\n" +
        "// which was extracted from about 16GB of XML specifications downloaded from the KNX-Foundation\n" +
        "// webservice.\n" +
        "//\n" +
        "// In regular intervals the local copy of this should be updated by running:\n" +
        "// - KnxSpecificationStoreUpdater\n" +
        "// Based on this content the content of this file can be generated by running:\n" +
        "// - KnxDefinitionsGenerator\n" +
        "// On this local copy.\n" +
        "//\n" +
        "/////////////////////////////////////////////////////////////////////////////////////////////////\n" +
        "[enum uint 16 DeviceInformation(uint 16 'deviceDescriptor', string 'name', uint 16 'comObjectTableAddress')\n";

    public static final String FOOTER = "]\n";

    protected int deviceCounter;
    protected Set<String> processedIds;

    public KnxDefinitionsGenerator() {
        processedIds = new HashSet<>();
        deviceCounter = 1;
    }

    public static void main(String[] args) throws Exception {
        final KnxDefinitionsGenerator generator = new KnxDefinitionsGenerator();
        File contentDir = new File(args[0]);
        System.out.println(HEADER);
        generator.processDirectory(contentDir);
        System.out.println(FOOTER);
    }

    public void processFile(File file, int manufacturerId, int applicationId, int applicationVersion) throws Exception {
        // Parse the content.
        SAXParserFactory factory = SAXParserFactory.newInstance();
        SAXParser saxParser = factory.newSAXParser();
        ProductDescriptionHandler handler = new ProductDescriptionHandler();
        saxParser.parse(file, handler);

        // Unfortunately the enum gets soo big, that we can only generate it for the devices we need it for.
        // These are the System 7 devices.
        if (handler.getMaskVersion().startsWith("070")) {
            //String cleanedName = handler.getName().replaceAll("\n", "").replaceAll("\r", "").replaceAll("\"", "inch");
            if (handler.getComObjectTableAddress() != null) {
                System.out.printf("    ['%4d' DEV%04X%04X%02X               ['0x%04X'                       ]]\n",
                    deviceCounter++, manufacturerId, applicationId, applicationVersion,
                    handler.getComObjectTableAddress());
            }
            if (handler.getReplacesVersions() != null) {
                for (String replacesVersion : handler.getReplacesVersions()) {
                    if (replacesVersion.length() == 1) {
                        replacesVersion = "0" + replacesVersion;
                    }
                    String replacedProductCode = String.format("M-%04X_A-%04X-%S", manufacturerId, applicationId, replacesVersion);
                    if (!processedIds.contains(replacedProductCode)) {
                        System.out.printf("    ['%4d' DEV%04X%04X%S               ['0x%04X'                       ]] //replaces '%s'\n",
                            deviceCounter++, manufacturerId, applicationId, replacesVersion,
                            handler.getComObjectTableAddress(), replacesVersion);
                        processedIds.add(replacedProductCode);
                    }
                }
            }
        }
    }
}

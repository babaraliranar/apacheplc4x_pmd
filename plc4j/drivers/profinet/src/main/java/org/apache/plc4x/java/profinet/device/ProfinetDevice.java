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

package org.apache.plc4x.java.profinet.device;

import io.netty.channel.Channel;
import org.apache.commons.codec.DecoderException;
import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.api.exceptions.PlcException;
import org.apache.plc4x.java.api.messages.PlcDiscoveryItem;
import org.apache.plc4x.java.profinet.config.ProfinetConfiguration;
import org.apache.plc4x.java.profinet.gsdml.*;
import org.apache.plc4x.java.profinet.protocol.ProfinetProtocolLogic;
import org.apache.plc4x.java.profinet.readwrite.*;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.utils.rawsockets.netty.RawSocketChannel;
import org.pcap4j.core.NotOpenException;
import org.pcap4j.core.PacketListener;
import org.pcap4j.core.PcapNativeException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.*;
import java.util.*;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Function;

public class ProfinetDevice {

    private static final int DEFAULT_UDP_PORT = 34964;
    private static final int DEFAULT_ARGS_MAXIMUM = 16696;
    private static final int DEFAULT_MAX_ARRAY_COUNT = 16696;
    private static final int DEFAULT_ACTIVITY_TIMEOUT = 600;

    // Not sure where this comes from?
    private static final int UDP_RT_PORT = 0x8892;
    private static final short BLOCK_VERSION_HIGH = 1;
    private static final short BLOCK_VERSION_LOW = 0;

    private static final MacAddress DEFAULT_EMPTY_MAC_ADDRESS;
    static {
        try {
            DEFAULT_EMPTY_MAC_ADDRESS = new MacAddress(Hex.decodeHex("000000000000"));
        } catch (DecoderException e) {
            throw new RuntimeException(e);
        }
    }

    private static final String DEFAULT_PLC4X_STATION_NAME = "plc4x";

    private final Logger logger = LoggerFactory.getLogger(ProfinetDevice.class);
    private final DceRpc_ActivityUuid uuid;
    private final ProfinetConfiguration configuration;
    private final InetAddress localIpAddress;
    private final DatagramSocket socket;
    private final DatagramSocket socketProfinetPort;
    private DatagramSocket udpSocket;
    private RawSocketChannel rawSocketChannel;
    private ProfinetChannel channel;
    private final MacAddress macAddress;
    private ConversationContext<Ethernet_Frame> context;
    private ProfinetDeviceState state = ProfinetDeviceState.IDLE;
    private boolean lldpReceived = false;
    private boolean dcpReceived = false;
    private String ipAddress;
    private String portId;
    private MacAddress localMacAddress;

    private AtomicInteger sessionKeyGenerator = new AtomicInteger(1);

    private static final Uuid ARUUID;

    static {
        try {
            ARUUID = new Uuid(Hex.decodeHex(UUID.randomUUID().toString().replace("-", "")));
        } catch (DecoderException e) {
            throw new RuntimeException(e);
        }
    }

    private String deviceTypeName;
    private String vendorId;
    private String deviceId;
    private String deviceName;
    private ProfinetISO15745Profile gsdFile;
    private boolean startupMode = false;
    private int initialFrameId = 0xBBF0;

    public ProfinetDevice(MacAddress macAddress, ProfinetConfiguration configuration) {
        this.macAddress = macAddress;
        this.configuration = configuration;
        try {
            this.localIpAddress = InetAddress.getByName(configuration.getTransportConfig().split(":")[0]);
        } catch (UnknownHostException e) {
            throw new RuntimeException(e);
        }
        // Generate a new Activity Id, which will be used throughout the connection.
        this.uuid = generateActivityUuid();

        try {
            socket = new DatagramSocket(50000);
            socketProfinetPort = new DatagramSocket(DEFAULT_UDP_PORT);
        } catch (SocketException e) {
            throw new RuntimeException(e);
        }
    }

    public ProfinetConfiguration getConfiguration() {
        return configuration;
    }

    private ProfinetISO15745Profile issueGSDMLFile(String vendorId, String deviceId) {
        String id = "0x" + vendorId + "-0x" + deviceId;
        if (this.configuration.getGsdFiles().containsKey(id)) {
            return this.configuration.getGsdFiles().get(id);
        } else {
            throw new RuntimeException("No GSDML file available for device " + this.vendorId + " - " + this.deviceId + " - " + this.deviceName);
        }
    }

    private void extractGSDFileInfo(ProfinetISO15745Profile gsdFile) {
        List<ProfinetDeviceAccessPointItem> deviceAccessList = gsdFile.getProfileBody().getApplicationProcess().getDeviceAccessPointList();
        // Always just pick the first one, until it's specified in the connection string
        if (deviceAccessList != null && deviceAccessList.size() > 0) {
            List<ProfinetInterfaceSubmoduleItem> interfaceSubModules = deviceAccessList.get(0).getSystemDefinedSubmoduleList().getInterfaceSubmodules();
            if (interfaceSubModules != null && interfaceSubModules.size() > 0) {
                if (interfaceSubModules.get(0).getApplicationRelations().getStartupMode() != null && interfaceSubModules.get(0).getApplicationRelations().getStartupMode().toLowerCase().contains("advanced")) {
                    this.startupMode = true;
                    this.initialFrameId = 0x8001;
                }
            }

        }
    }

    public boolean onConnect() {

        

        this.gsdFile = issueGSDMLFile(this.vendorId, this.deviceId);
        extractGSDFileInfo(this.gsdFile);

        ProfinetMessageWrapper.sendUdpMessage(
            new CreateConnection(),
            this
        );

        ProfinetMessageWrapper.sendUdpMessage(
            new WriteParameters(),
            this
        );

        ProfinetMessageWrapper.sendUdpMessage(
            new WriteParametersEnd(),
            this
        );

        return true;
    }

    private int generateSessionKey() {
        // Generate a new session key.
        Integer sessionKey = sessionKeyGenerator.getAndIncrement();
        // Reset the session key as soon as it reaches the max for a 16 bit uint
        if (sessionKeyGenerator.get() == 0xFFFF) {
            sessionKeyGenerator.set(1);
        }
        return sessionKey;
    }

    public boolean hasLldpPdu() {
        return this.lldpReceived;
    }

    public boolean hasDcpPdu() {
        return this.dcpReceived;
    }

    public void handle(PlcDiscoveryItem item) {
        logger.debug("Received Discovered item at device");
        if (item.getOptions().containsKey("ipAddress")) {
            this.ipAddress = item.getOptions().get("ipAddress");
        }
        if (item.getOptions().containsKey("portId")) {
            this.portId = item.getOptions().get("portId");
        }
        if (item.getOptions().containsKey("deviceTypeName")) {
            this.deviceTypeName = item.getOptions().get("deviceTypeName");
        }
        if (item.getOptions().containsKey("vendorId")) {
            this.vendorId = item.getOptions().get("vendorId");
        }
        if (item.getOptions().containsKey("deviceId")) {
            this.deviceId = item.getOptions().get("deviceId");
        }
        if (item.getOptions().containsKey("deviceName")) {
            this.deviceName = item.getOptions().get("deviceName");
        }
        if (item.getOptions().containsKey("localMacAddress")) {
            String macString = item.getOptions().get("localMacAddress").replace(":", "");
            try {
                this.localMacAddress = new MacAddress(Hex.decodeHex(macString));
            } catch (DecoderException e) {
                throw new RuntimeException(e);
            }
        }
        if (item.getOptions().containsKey("packetType")) {
            if (item.getOptions().get("packetType").equals("lldp")) {
                this.lldpReceived = true;
            }
            if (item.getOptions().get("packetType").equals("dcp")) {
                this.dcpReceived = true;
            }
        }
    }

    public void setContext(ConversationContext<Ethernet_Frame> context, ProfinetChannel channel) {
        this.context = context;
        this.channel = channel;
    }

    protected static DceRpc_ActivityUuid generateActivityUuid() {
        UUID number = UUID.randomUUID();
        try {
            WriteBufferByteBased wb = new WriteBufferByteBased(128);
            wb.writeLong(64, number.getMostSignificantBits());
            wb.writeLong(64, number.getLeastSignificantBits());

            ReadBuffer rb = new ReadBufferByteBased(wb.getData());
            return new DceRpc_ActivityUuid(rb.readLong(32), rb.readInt(16), rb.readInt(16), rb.readByteArray(8));
        } catch (SerializationException | ParseException e) {
            // Ignore ... this should actually never happen.
        }
        return null;
    }

    public RawSocketChannel getRawSocket() {
        return this.rawSocketChannel;
    }

    public InetAddress getIpAddress() throws UnknownHostException {
        return InetAddress.getByName(this.ipAddress);
    }

    public MacAddress getMacAddress() {
        return macAddress;
    }

    public MacAddress getLocalMacAddress() {
        return localMacAddress;
    }

    public InetAddress getLocalIpAddress() {
        return localIpAddress;
    }

    public void setLocalMacAddress(MacAddress localMacAddress) {
        this.localMacAddress = localMacAddress;
    }

    public int getPort() {
        return DEFAULT_UDP_PORT;
    }

    public class CreateConnection implements ProfinetCallable<DceRpc_Packet> {

        public DceRpc_Packet create() throws PlcException {

                List<PnIoCm_Block> blocks = new ArrayList<>();
                blocks.add(new PnIoCm_Block_ArReq(
                    BLOCK_VERSION_HIGH,
                    BLOCK_VERSION_LOW,
                    PnIoCm_ArType.IO_CONTROLLER,
                    ARUUID,
                    ProfinetDevice.this.generateSessionKey(),
                    ProfinetDevice.this.macAddress,
                    new DceRpc_ObjectUuid((byte) 0x00, 0x0001, Integer.valueOf(deviceId), Integer.valueOf(vendorId)),
                    false,
                    startupMode,
                    false,
                    false,
                    PnIoCm_CompanionArType.SINGLE_AR,
                    false,
                    true,
                    false,
                    PnIoCm_State.ACTIVE,
                    DEFAULT_ACTIVITY_TIMEOUT,
                    UDP_RT_PORT,
                    DEFAULT_PLC4X_STATION_NAME));

                blocks.add(
                    new PnIoCm_Block_AlarmCrReq(
                        (short) 1,
                        (short) 0,
                        PnIoCm_AlarmCrType.ALARM_CR,
                        0x8892,
                        false,
                        false,
                        1,
                        3,
                        0x0000,
                        200,
                        0xC000,
                        0xA000)
                );

                List<PnIoCm_IoDataObject> inputApiBlocks = new ArrayList<>();
                List<PnIoCm_IoCs> outputApiBlocks = new ArrayList<>();
                List<PnIoCm_Submodule> expectedSubModuleApiBlocks = new ArrayList<>();

                int offsetCount = 0;
                for (ProfinetVirtualSubmoduleItem virtualItem : gsdFile.getProfileBody().getApplicationProcess().getDeviceAccessPointList().get(0).getVirtualSubmoduleList()) {
                    Integer identNumber = Integer.decode(virtualItem.getSubmoduleIdentNumber());
                    inputApiBlocks.add(new PnIoCm_IoDataObject(
                        0,
                        identNumber,
                        offsetCount));
                    outputApiBlocks.add(new PnIoCm_IoCs(
                        0,
                        identNumber,
                        offsetCount));
                    expectedSubModuleApiBlocks.add(new PnIoCm_Submodule_NoInputNoOutputData(
                        identNumber,
                        identNumber,
                        false,
                        false,
                        false,
                        false));
                    offsetCount += 1;
                }

                for (ProfinetInterfaceSubmoduleItem interfaceItem : gsdFile.getProfileBody().getApplicationProcess().getDeviceAccessPointList().get(0).getSystemDefinedSubmoduleList().getInterfaceSubmodules()) {
                    Integer identNumber = Integer.decode(interfaceItem.getSubmoduleIdentNumber());
                    inputApiBlocks.add(new PnIoCm_IoDataObject(
                        0,
                        identNumber,
                        offsetCount));
                    outputApiBlocks.add(new PnIoCm_IoCs(
                        0,
                        identNumber,
                        offsetCount));
                    expectedSubModuleApiBlocks.add(new PnIoCm_Submodule_NoInputNoOutputData(
                        identNumber,
                        identNumber,
                        false,
                        false,
                        false,
                        false));
                    offsetCount += 1;
                }

                for (ProfinetPortSubmoduleItem portItem : gsdFile.getProfileBody().getApplicationProcess().getDeviceAccessPointList().get(0).getSystemDefinedSubmoduleList().getPortSubmodules()) {
                    Integer identNumber = Integer.decode(portItem.getSubmoduleIdentNumber());
                    inputApiBlocks.add(new PnIoCm_IoDataObject(
                        0,
                        identNumber,
                        offsetCount));
                    outputApiBlocks.add(new PnIoCm_IoCs(
                        0,
                        identNumber,
                        offsetCount));
                    expectedSubModuleApiBlocks.add(new PnIoCm_Submodule_NoInputNoOutputData(
                        identNumber,
                        identNumber,
                        false,
                        false,
                        false,
                        false));
                    offsetCount += 1;
                }

                List<PnIoCm_IoCrBlockReqApi> inputApis = Collections.singletonList(
                    new PnIoCm_IoCrBlockReqApi(
                        inputApiBlocks,
                        new ArrayList<>(0))
                );

                List<PnIoCm_IoCrBlockReqApi> outputApis = Collections.singletonList(
                    new PnIoCm_IoCrBlockReqApi(
                        new ArrayList<>(0),
                        outputApiBlocks
                    )
                );

                int frameCount = 0;

                blocks.add(
                    new PnIoCm_Block_IoCrReq(
                        (short) 1,
                        (short) 0,
                        PnIoCm_IoCrType.INPUT_CR,
                        0x0001,
                        UDP_RT_PORT,
                        false,
                        false,
                        false,
                        false,
                        PnIoCm_RtClass.RT_CLASS_2,
                        40,
                        initialFrameId + frameCount,
                        configuration.getSendClockFactor(),
                        configuration.getReductionRatio(),
                        1,
                        0,
                        0xffffffff,
                        configuration.getWatchdogFactor(),
                        50,
                        0xC000,
                        DEFAULT_EMPTY_MAC_ADDRESS,
                        inputApis
                        ));

                frameCount += 1;

                blocks.add(
                    new PnIoCm_Block_IoCrReq(
                        (short) 1,
                        (short) 0,
                        PnIoCm_IoCrType.OUTPUT_CR,
                        0x0002,
                        UDP_RT_PORT,
                        false,
                        false,
                        false,
                        false,
                        PnIoCm_RtClass.RT_CLASS_2,
                        40,
                        initialFrameId + frameCount,
                        configuration.getSendClockFactor(),
                        configuration.getReductionRatio(),
                        1,
                        0,
                        0xffffffff,
                        configuration.getWatchdogFactor(),
                        50,
                        0xC000,
                        DEFAULT_EMPTY_MAC_ADDRESS,
                        outputApis
                    ));

                blocks.add(
                    new PnIoCm_Block_ExpectedSubmoduleReq((short) 1, (short) 0,
                        Collections.singletonList(
                            new PnIoCm_ExpectedSubmoduleBlockReqApi(0,
                                0x00000001,
                                0x00000000,
                                expectedSubModuleApiBlocks
                            )
                        )
                    ));

                long arrayLength = 0;
                for (PnIoCm_Block block : blocks) {
                    arrayLength += block.getLengthInBytes();
                }

                return new DceRpc_Packet(
                    DceRpc_PacketType.REQUEST, true, false, false,
                    IntegerEncoding.BIG_ENDIAN, CharacterEncoding.ASCII, FloatingPointEncoding.IEEE,
                    new DceRpc_ObjectUuid((byte) 0x00, 0x0001, Integer.valueOf(deviceId), Integer.valueOf(vendorId)),
                    new DceRpc_InterfaceUuid_DeviceInterface(),
                    ProfinetDevice.this.uuid,
                    0, 0, DceRpc_Operation.CONNECT,
                    new PnIoCm_Packet_Req(DEFAULT_ARGS_MAXIMUM, DEFAULT_MAX_ARRAY_COUNT, 0, arrayLength, blocks)

                );

        }

        public void handle(DceRpc_Packet dceRpc_packet) throws PlcException {
            if ((dceRpc_packet.getOperation() == DceRpc_Operation.CONNECT) && (dceRpc_packet.getPacketType() == DceRpc_PacketType.RESPONSE)) {
                if (dceRpc_packet.getPayload().getPacketType() == DceRpc_PacketType.RESPONSE) {

                    // Get the remote MAC address and store it in the context.
                    final PnIoCm_Packet_Res connectResponse = (PnIoCm_Packet_Res) dceRpc_packet.getPayload();
                    if ((connectResponse.getBlocks().size() > 0) && (connectResponse.getBlocks().get(0) instanceof PnIoCm_Block_ArRes)) {
                        final PnIoCm_Block_ArRes pnIoCm_block_arRes = (PnIoCm_Block_ArRes) connectResponse.getBlocks().get(0);

                        // Update the raw-socket transports filter expression.
                        //((RawSocketChannel) channel).setRemoteMacAddress(org.pcap4j.util.MacAddress.getByAddress(macAddress.getAddress()));
                    } else {
                        throw new PlcException("Unexpected type of first block.");
                    }
                } else {
                    throw new PlcException("Unexpected response");
                }
            } else if (dceRpc_packet.getPacketType() == DceRpc_PacketType.REJECT) {
                throw new PlcException("Device rejected connection request");
            } else {
                throw new PlcException("Unexpected response");
            }
        }
    }

    public class WriteParameters implements ProfinetCallable<DceRpc_Packet> {
        public DceRpc_Packet create() {
            return new DceRpc_Packet(
                DceRpc_PacketType.REQUEST, true, false, false,
                IntegerEncoding.BIG_ENDIAN, CharacterEncoding.ASCII, FloatingPointEncoding.IEEE,
                new DceRpc_ObjectUuid((byte) 0x00, 0x0001, 0x0904, 0x002A),
                new DceRpc_InterfaceUuid_DeviceInterface(),
                uuid,
                0, 1, DceRpc_Operation.WRITE,
                new PnIoCm_Packet_Req(16696, 16696, 0, 244,
                    Arrays.asList(
                        new IODWriteRequestHeader(
                            (short) 1,
                            (short) 0,
                            0,
                            ARUUID,
                            0x00000000,
                            0x0000,
                            0x0000,
                            0xe040,
                            180
                        ),
                        new IODWriteRequestHeader(
                            (short) 1,
                            (short) 0,
                            1,
                            ARUUID,
                            0x00000000,
                            0x0000,
                            0x8000,
                            0x8071,
                            12
                        ),
                        new PDInterfaceAdjust(
                            (short) 1,
                            (short) 0,
                            MultipleInterfaceModeNameOfDevice.NAME_PROVIDED_BY_LLDP
                        )
                    ))
            );
        }

        @Override
        public void handle(DceRpc_Packet packet) throws PlcException {
            logger.debug("Received a Write Parameter Response");
        }
    }

    public class WriteParametersEnd implements ProfinetCallable<DceRpc_Packet> {
        public DceRpc_Packet create() {
            return new DceRpc_Packet(
                DceRpc_PacketType.REQUEST, true, false, false,
                IntegerEncoding.BIG_ENDIAN, CharacterEncoding.ASCII, FloatingPointEncoding.IEEE,
                new DceRpc_ObjectUuid((byte) 0x00, 0x0001, 0x0904, 0x002A),
                new DceRpc_InterfaceUuid_DeviceInterface(),
                uuid,
                0, 1, DceRpc_Operation.CONTROL,
                new PnIoCm_Packet_Req(16696, 16696, 0, 244,
                    Arrays.asList(
                        new PnIoCm_Control_Request(
                            (short) 1,
                            (short) 0,
                            ARUUID,
                            0x0001,
                            0x0001
                        )
                    ))
            );
        }

        @Override
        public void handle(DceRpc_Packet packet) throws PlcException {
            logger.debug("Received a Write Parameter End Response");
        }
    }

    public ProfinetChannel getChannel() {
        return channel;
    }

    public class CyclicData implements ProfinetCallable<Ethernet_Frame> {
        public Ethernet_Frame create() {
            return new Ethernet_Frame(
                macAddress,
                macAddress,
                new Ethernet_FramePayload_PnDcp(
                    new PnDcp_Pdu_RealTimeCyclic(
                        0x8000,
                        new PnIo_CyclicServiceDataUnit((short) 0, (short) 0, (short) 0),
                        16696,
                        false,
                        false,
                        false,
                        false,
                        false,
                        false)));
        }



        @Override
        public void handle(Ethernet_Frame packet) throws PlcException {
            logger.debug("Received a Write Parameter End Response");
        }
    }


}

#!/usr/bin/env bash

tinc=/etc/tinc
scripts=/etc/tinc/scripts

if [ ! -f $tinc/hosts/$NODENAME ]; then tinc init $NODENAME; else echo "$tinc/hosts/$NODENAME exist"; fi
if [ "$ADDRESS_FAMILY" != "" ]; then tinc set AddressFamily $ADDRESS_FAMILY; fi
if [ "$AUTO_CONNECT" != "" ]; then tinc set AutoConnect $AUTO_CONNECT; fi
if [ "$BIND_TO_ADDRESS" != "" ]; then tinc set BindToAddress $BIND_TO_ADDRESS; fi
if [ "$BROADCAST_SUBNET" != "" ]; then tinc set BroadcastSubnet $BROADCAST_SUBNET; fi
if [ "$CONNECT_TO" != "" ]; then
    for i in $CONNECT_TO; do
        echo "ConnectTo = $i" >> $tinc/tinc.conf
    done
fi
if [ "$DEVICE" != "" ]; then tinc set Device $DEVICE; fi
if [ "$DEVICE_STANDBY" != "" ]; then tinc set DeviceStandby $DEVICE_STANDBY; fi
if [ "$DEVICE_TYPE" != "" ]; then tinc set DeviceType $DEVICE_TYPE; fi
if [ "$EXPERIMENTAL_PROTOCOL" != "" ]; then tinc set ExperimentalProtocol $EXPERIMENTAL_PROTOCOL; fi
if [ "$HOSTNAMES" != "" ]; then tinc set Hostnames $HOSTNAMES; fi
if [ "$INTERFACE" != "" ]; then tinc set Interface $INTERFACE; fi
if [ "$LISTEN_ADDRESS" != "" ]; then tinc set ListenAddress $LISTEN_ADDRESS; fi
if [ "$LOCAL_DISCOVERY" != "" ]; then tinc set LocalDiscovery $LOCAL_DISCOVERY; fi
if [ "$LOG_LEVEL" != "" ]; then tinc set LogLevel $LOG_LEVEL; fi
if [ "$MODE" != "" ]; then tinc set Mode $MODE; fi
if [ "$INVITATION_EXPIRE" != "" ]; then tinc set InvitationExpire $INVITATION_EXPIRE; fi
if [ "$KEY_EXPIRE" != "" ]; then tinc set KeyExpire $KEY_EXPIRE; fi
if [ "$MAC_EXPIRE" != "" ]; then tinc set MACExpire $MAC_EXPIRE; fi
if [ "$MAX_CONNECTION_BURST" != "" ]; then tinc set MaxConnectionBurst $MAX_CONNECTION_BURST; fi
if [ "$PING_INTERVAL" != "" ]; then tinc set PingInterval $PING_INTERVAL; fi
if [ "$PROCESS_PRIORITY" != "" ]; then tinc set ProcessPriority $PROCESS_PRIORITY; fi
if [ "$REPLAY_WINDOW" != "" ]; then tinc set ReplayWindow $REPLAY_WINDOW; fi
if [ "$UDP_DISCOVERY" != "" ]; then tinc set UDPDiscovery $UDP_DISCOVERY; fi
if [ "$UDP_DISCOVERY_KEEPALIVE_INTERVAL" != "" ]; then tinc set UDPDiscoveryKeepaliveInterval $UDP_DISCOVERY_KEEPALIVE_INTERVAL; fi
if [ "$UDP_DISCOVERY_INTERVAL" != "" ]; then tinc set UDPDiscoveryInterval $UDP_DISCOVERY_INTERVAL; fi
if [ "$UDP_DISCOVERY_TIMEOUT" != "" ]; then tinc set UDPDiscoveryTimeout $UDP_DISCOVERY_TIMEOUT; fi
if [ "$UDP_INFO_INTERVAL" != "" ]; then tinc set UDPInfoInterval $UDP_INFO_INTERVAL; fi
if [ "$UDP_RCV_BUF" != "" ]; then tinc set UDPRcvBuf $UDP_RCV_BUF; fi
if [ "$UDP_SND_BUF" != "" ]; then tinc set UDPSndBuf $UDP_SND_BUF; fi

if [ "$ADDRESS" != "" ]; then tinc set $NODENAME.Address $ADDRESS; fi
if [ "$CIPHER" != "" ]; then tinc set $NODENAME.Cipher $CIPHER; fi
if [ "$CLAMP_MSS" != "" ]; then tinc set $NODENAME.ClampMSS $CLAMP_MSS; fi
if [ "$COMPRESSION" != "" ]; then tinc set $NODENAME.Compression $COMPRESSION; fi
if [ "$DIGEST" != "" ]; then tinc set $NODENAME.Digest $DIGEST; fi
if [ "$INDIRECT_DATA" != "" ]; then tinc set $NODENAME.IndirectData $INDIRECT_DATA; fi
if [ "$MAC_LENGTH" != "" ]; then tinc set $NODENAME.MACLength $MAC_LENGTH; fi
if [ "$PMTU" != "" ]; then tinc set $NODENAME.PMTU $PMTU; fi
if [ "$PMTU_DISCOVERY" != "" ]; then tinc set $NODENAME.PMTUDiscovery $PMTU_DISCOVERY; fi
if [ "$MTU_INFO_INTERVAL" != "" ]; then tinc set $NODENAME.MTUInfoInterval $MTU_INFO_INTERVAL; fi
if [ "$PORT" != "" ]; then tinc set $NODENAME.Port $PORT; fi
if [ "$SUBNET" != "" ]; then tinc set $NODENAME.Subnet $SUBNET; fi
if [ "$TCP_ONLY" != "" ]; then tinc set $NODENAME.TCPOnly $TCP_ONLY; fi
if [ "$WEIGHT" != "" ]; then tinc set $NODENAME.Weight $WEIGHT; fi

# start gui
cd /opt/tinc-gui
./tinc-gui&

# start tinc
cd /etc/tinc
tinc start -D --logfile=/etc/tinc/log

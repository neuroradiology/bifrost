/* eslint-disable */
import Long from 'long'
import _m0 from 'protobufjs/minimal.js'
import { DialerOpts } from '../common/dialer/dialer.pb.js'
import { Opts } from '../common/kcp/kcp.pb.js'

export const protobufPackage = 'xbee'

/** Config is the configuration for the udp transport. */
export interface Config {
  /**
   * TransportPeerID sets the peer ID to attach the transport to.
   * If unset, attaches to any running peer with a private key.
   */
  transportPeerId: string
  /** DevicePath is the device path to open the serial stream. */
  devicePath: string
  /** DeviceBaud is the device baudrate. */
  deviceBaud: number
  /**
   * PacketOpts are options to set on the packet connection.
   * In lossy environments, set the data shards for error correction.
   */
  packetOpts: Opts | undefined
  /** Dialers maps peer IDs to dialers. */
  dialers: { [key: string]: DialerOpts }
}

export interface Config_DialersEntry {
  key: string
  value: DialerOpts | undefined
}

function createBaseConfig(): Config {
  return {
    transportPeerId: '',
    devicePath: '',
    deviceBaud: 0,
    packetOpts: undefined,
    dialers: {},
  }
}

export const Config = {
  encode(
    message: Config,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.transportPeerId !== '') {
      writer.uint32(10).string(message.transportPeerId)
    }
    if (message.devicePath !== '') {
      writer.uint32(18).string(message.devicePath)
    }
    if (message.deviceBaud !== 0) {
      writer.uint32(24).int32(message.deviceBaud)
    }
    if (message.packetOpts !== undefined) {
      Opts.encode(message.packetOpts, writer.uint32(34).fork()).ldelim()
    }
    Object.entries(message.dialers).forEach(([key, value]) => {
      Config_DialersEntry.encode(
        { key: key as any, value },
        writer.uint32(42).fork(),
      ).ldelim()
    })
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Config {
    const reader =
      input instanceof _m0.Reader ? input : _m0.Reader.create(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseConfig()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break
          }

          message.transportPeerId = reader.string()
          continue
        case 2:
          if (tag !== 18) {
            break
          }

          message.devicePath = reader.string()
          continue
        case 3:
          if (tag !== 24) {
            break
          }

          message.deviceBaud = reader.int32()
          continue
        case 4:
          if (tag !== 34) {
            break
          }

          message.packetOpts = Opts.decode(reader, reader.uint32())
          continue
        case 5:
          if (tag !== 42) {
            break
          }

          const entry5 = Config_DialersEntry.decode(reader, reader.uint32())
          if (entry5.value !== undefined) {
            message.dialers[entry5.key] = entry5.value
          }
          continue
      }
      if ((tag & 7) === 4 || tag === 0) {
        break
      }
      reader.skipType(tag & 7)
    }
    return message
  },

  // encodeTransform encodes a source of message objects.
  // Transform<Config, Uint8Array>
  async *encodeTransform(
    source: AsyncIterable<Config | Config[]> | Iterable<Config | Config[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [Config.encode(p).finish()]
        }
      } else {
        yield* [Config.encode(pkt as any).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, Config>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<Config> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [Config.decode(p)]
        }
      } else {
        yield* [Config.decode(pkt as any)]
      }
    }
  },

  fromJSON(object: any): Config {
    return {
      transportPeerId: isSet(object.transportPeerId)
        ? globalThis.String(object.transportPeerId)
        : '',
      devicePath: isSet(object.devicePath)
        ? globalThis.String(object.devicePath)
        : '',
      deviceBaud: isSet(object.deviceBaud)
        ? globalThis.Number(object.deviceBaud)
        : 0,
      packetOpts: isSet(object.packetOpts)
        ? Opts.fromJSON(object.packetOpts)
        : undefined,
      dialers: isObject(object.dialers)
        ? Object.entries(object.dialers).reduce<{ [key: string]: DialerOpts }>(
            (acc, [key, value]) => {
              acc[key] = DialerOpts.fromJSON(value)
              return acc
            },
            {},
          )
        : {},
    }
  },

  toJSON(message: Config): unknown {
    const obj: any = {}
    if (message.transportPeerId !== '') {
      obj.transportPeerId = message.transportPeerId
    }
    if (message.devicePath !== '') {
      obj.devicePath = message.devicePath
    }
    if (message.deviceBaud !== 0) {
      obj.deviceBaud = Math.round(message.deviceBaud)
    }
    if (message.packetOpts !== undefined) {
      obj.packetOpts = Opts.toJSON(message.packetOpts)
    }
    if (message.dialers) {
      const entries = Object.entries(message.dialers)
      if (entries.length > 0) {
        obj.dialers = {}
        entries.forEach(([k, v]) => {
          obj.dialers[k] = DialerOpts.toJSON(v)
        })
      }
    }
    return obj
  },

  create<I extends Exact<DeepPartial<Config>, I>>(base?: I): Config {
    return Config.fromPartial(base ?? ({} as any))
  },
  fromPartial<I extends Exact<DeepPartial<Config>, I>>(object: I): Config {
    const message = createBaseConfig()
    message.transportPeerId = object.transportPeerId ?? ''
    message.devicePath = object.devicePath ?? ''
    message.deviceBaud = object.deviceBaud ?? 0
    message.packetOpts =
      object.packetOpts !== undefined && object.packetOpts !== null
        ? Opts.fromPartial(object.packetOpts)
        : undefined
    message.dialers = Object.entries(object.dialers ?? {}).reduce<{
      [key: string]: DialerOpts
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = DialerOpts.fromPartial(value)
      }
      return acc
    }, {})
    return message
  },
}

function createBaseConfig_DialersEntry(): Config_DialersEntry {
  return { key: '', value: undefined }
}

export const Config_DialersEntry = {
  encode(
    message: Config_DialersEntry,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.key !== '') {
      writer.uint32(10).string(message.key)
    }
    if (message.value !== undefined) {
      DialerOpts.encode(message.value, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Config_DialersEntry {
    const reader =
      input instanceof _m0.Reader ? input : _m0.Reader.create(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseConfig_DialersEntry()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break
          }

          message.key = reader.string()
          continue
        case 2:
          if (tag !== 18) {
            break
          }

          message.value = DialerOpts.decode(reader, reader.uint32())
          continue
      }
      if ((tag & 7) === 4 || tag === 0) {
        break
      }
      reader.skipType(tag & 7)
    }
    return message
  },

  // encodeTransform encodes a source of message objects.
  // Transform<Config_DialersEntry, Uint8Array>
  async *encodeTransform(
    source:
      | AsyncIterable<Config_DialersEntry | Config_DialersEntry[]>
      | Iterable<Config_DialersEntry | Config_DialersEntry[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [Config_DialersEntry.encode(p).finish()]
        }
      } else {
        yield* [Config_DialersEntry.encode(pkt as any).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, Config_DialersEntry>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<Config_DialersEntry> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [Config_DialersEntry.decode(p)]
        }
      } else {
        yield* [Config_DialersEntry.decode(pkt as any)]
      }
    }
  },

  fromJSON(object: any): Config_DialersEntry {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : '',
      value: isSet(object.value)
        ? DialerOpts.fromJSON(object.value)
        : undefined,
    }
  },

  toJSON(message: Config_DialersEntry): unknown {
    const obj: any = {}
    if (message.key !== '') {
      obj.key = message.key
    }
    if (message.value !== undefined) {
      obj.value = DialerOpts.toJSON(message.value)
    }
    return obj
  },

  create<I extends Exact<DeepPartial<Config_DialersEntry>, I>>(
    base?: I,
  ): Config_DialersEntry {
    return Config_DialersEntry.fromPartial(base ?? ({} as any))
  },
  fromPartial<I extends Exact<DeepPartial<Config_DialersEntry>, I>>(
    object: I,
  ): Config_DialersEntry {
    const message = createBaseConfig_DialersEntry()
    message.key = object.key ?? ''
    message.value =
      object.value !== undefined && object.value !== null
        ? DialerOpts.fromPartial(object.value)
        : undefined
    return message
  },
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Long
  ? string | number | Long
  : T extends globalThis.Array<infer U>
  ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends { $case: string }
  ? { [K in keyof Omit<T, '$case'>]?: DeepPartial<T[K]> } & {
      $case: T['$case']
    }
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

type KeysOfUnion<T> = T extends T ? keyof T : never
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & {
      [K in Exclude<keyof I, KeysOfUnion<P>>]: never
    }

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any
  _m0.configure()
}

function isObject(value: any): boolean {
  return typeof value === 'object' && value !== null
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined
}

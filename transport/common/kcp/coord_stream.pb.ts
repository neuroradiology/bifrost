/* eslint-disable */
import Long from 'long'
import _m0 from 'protobufjs/minimal.js'

export const protobufPackage = 'kcp'

/** CoordPacketType is the packet type of a coordination stream packet. */
export enum CoordPacketType {
  CoordPacketType_UNKNOWN = 0,
  CoordPacketType_RSTREAM_ESTABLISH = 1,
  CoordPacketType_RSTREAM_ACK = 2,
  CoordPacketType_RSTREAM_CLOSE = 3,
  CoordPacketType_RSTREAM_NOOP = 4,
  UNRECOGNIZED = -1,
}

export function coordPacketTypeFromJSON(object: any): CoordPacketType {
  switch (object) {
    case 0:
    case 'CoordPacketType_UNKNOWN':
      return CoordPacketType.CoordPacketType_UNKNOWN
    case 1:
    case 'CoordPacketType_RSTREAM_ESTABLISH':
      return CoordPacketType.CoordPacketType_RSTREAM_ESTABLISH
    case 2:
    case 'CoordPacketType_RSTREAM_ACK':
      return CoordPacketType.CoordPacketType_RSTREAM_ACK
    case 3:
    case 'CoordPacketType_RSTREAM_CLOSE':
      return CoordPacketType.CoordPacketType_RSTREAM_CLOSE
    case 4:
    case 'CoordPacketType_RSTREAM_NOOP':
      return CoordPacketType.CoordPacketType_RSTREAM_NOOP
    case -1:
    case 'UNRECOGNIZED':
    default:
      return CoordPacketType.UNRECOGNIZED
  }
}

export function coordPacketTypeToJSON(object: CoordPacketType): string {
  switch (object) {
    case CoordPacketType.CoordPacketType_UNKNOWN:
      return 'CoordPacketType_UNKNOWN'
    case CoordPacketType.CoordPacketType_RSTREAM_ESTABLISH:
      return 'CoordPacketType_RSTREAM_ESTABLISH'
    case CoordPacketType.CoordPacketType_RSTREAM_ACK:
      return 'CoordPacketType_RSTREAM_ACK'
    case CoordPacketType.CoordPacketType_RSTREAM_CLOSE:
      return 'CoordPacketType_RSTREAM_CLOSE'
    case CoordPacketType.CoordPacketType_RSTREAM_NOOP:
      return 'CoordPacketType_RSTREAM_NOOP'
    case CoordPacketType.UNRECOGNIZED:
    default:
      return 'UNRECOGNIZED'
  }
}

/**
 * CoordinationStreamPacket is the packet wrapper for a coordination stream
 * packet.
 */
export interface CoordinationStreamPacket {
  /** PacketType is the coordination stream packet type. */
  packetType: CoordPacketType
  /** RawStreamEstablish is the raw stream establish packet. */
  rawStreamEstablish: RawStreamEstablish | undefined
  /** RawStreamAck is the raw stream ack packet. */
  rawStreamAck: RawStreamAck | undefined
  /** RawStreamClose is the raw stream close packet. */
  rawStreamClose: RawStreamClose | undefined
}

/** RawStreamEstablish is a coordination stream raw-stream establish message. */
export interface RawStreamEstablish {
  /** InitiatorStreamId is the stream ID the initiator wants to use. */
  initiatorStreamId: number
}

/** RawStreamAck is a coordination stream raw-stream acknowledge message. */
export interface RawStreamAck {
  /** InitiatorStreamId is the stream ID the initiator wanted to use. */
  initiatorStreamId: number
  /**
   * AckStreamId is the stream ID the responder wants to use.
   * Zero if the stream was rejected.
   */
  ackStreamId: number
  /** AckError indicates an error establishing the stream, rejecting the stream. */
  ackError: string
}

/** RawStreamClose indicates an intent to close a raw stream. */
export interface RawStreamClose {
  /** StreamId is the stream ID the reciever indicated to use. */
  streamId: number
  /** CloseError indicates an error included with the stream close. */
  closeError: string
}

function createBaseCoordinationStreamPacket(): CoordinationStreamPacket {
  return {
    packetType: 0,
    rawStreamEstablish: undefined,
    rawStreamAck: undefined,
    rawStreamClose: undefined,
  }
}

export const CoordinationStreamPacket = {
  encode(
    message: CoordinationStreamPacket,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.packetType !== 0) {
      writer.uint32(8).int32(message.packetType)
    }
    if (message.rawStreamEstablish !== undefined) {
      RawStreamEstablish.encode(
        message.rawStreamEstablish,
        writer.uint32(18).fork(),
      ).ldelim()
    }
    if (message.rawStreamAck !== undefined) {
      RawStreamAck.encode(
        message.rawStreamAck,
        writer.uint32(26).fork(),
      ).ldelim()
    }
    if (message.rawStreamClose !== undefined) {
      RawStreamClose.encode(
        message.rawStreamClose,
        writer.uint32(34).fork(),
      ).ldelim()
    }
    return writer
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): CoordinationStreamPacket {
    const reader =
      input instanceof _m0.Reader ? input : _m0.Reader.create(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseCoordinationStreamPacket()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break
          }

          message.packetType = reader.int32() as any
          continue
        case 2:
          if (tag !== 18) {
            break
          }

          message.rawStreamEstablish = RawStreamEstablish.decode(
            reader,
            reader.uint32(),
          )
          continue
        case 3:
          if (tag !== 26) {
            break
          }

          message.rawStreamAck = RawStreamAck.decode(reader, reader.uint32())
          continue
        case 4:
          if (tag !== 34) {
            break
          }

          message.rawStreamClose = RawStreamClose.decode(
            reader,
            reader.uint32(),
          )
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
  // Transform<CoordinationStreamPacket, Uint8Array>
  async *encodeTransform(
    source:
      | AsyncIterable<CoordinationStreamPacket | CoordinationStreamPacket[]>
      | Iterable<CoordinationStreamPacket | CoordinationStreamPacket[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [CoordinationStreamPacket.encode(p).finish()]
        }
      } else {
        yield* [CoordinationStreamPacket.encode(pkt as any).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, CoordinationStreamPacket>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<CoordinationStreamPacket> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [CoordinationStreamPacket.decode(p)]
        }
      } else {
        yield* [CoordinationStreamPacket.decode(pkt as any)]
      }
    }
  },

  fromJSON(object: any): CoordinationStreamPacket {
    return {
      packetType: isSet(object.packetType)
        ? coordPacketTypeFromJSON(object.packetType)
        : 0,
      rawStreamEstablish: isSet(object.rawStreamEstablish)
        ? RawStreamEstablish.fromJSON(object.rawStreamEstablish)
        : undefined,
      rawStreamAck: isSet(object.rawStreamAck)
        ? RawStreamAck.fromJSON(object.rawStreamAck)
        : undefined,
      rawStreamClose: isSet(object.rawStreamClose)
        ? RawStreamClose.fromJSON(object.rawStreamClose)
        : undefined,
    }
  },

  toJSON(message: CoordinationStreamPacket): unknown {
    const obj: any = {}
    if (message.packetType !== 0) {
      obj.packetType = coordPacketTypeToJSON(message.packetType)
    }
    if (message.rawStreamEstablish !== undefined) {
      obj.rawStreamEstablish = RawStreamEstablish.toJSON(
        message.rawStreamEstablish,
      )
    }
    if (message.rawStreamAck !== undefined) {
      obj.rawStreamAck = RawStreamAck.toJSON(message.rawStreamAck)
    }
    if (message.rawStreamClose !== undefined) {
      obj.rawStreamClose = RawStreamClose.toJSON(message.rawStreamClose)
    }
    return obj
  },

  create<I extends Exact<DeepPartial<CoordinationStreamPacket>, I>>(
    base?: I,
  ): CoordinationStreamPacket {
    return CoordinationStreamPacket.fromPartial(base ?? ({} as any))
  },
  fromPartial<I extends Exact<DeepPartial<CoordinationStreamPacket>, I>>(
    object: I,
  ): CoordinationStreamPacket {
    const message = createBaseCoordinationStreamPacket()
    message.packetType = object.packetType ?? 0
    message.rawStreamEstablish =
      object.rawStreamEstablish !== undefined &&
      object.rawStreamEstablish !== null
        ? RawStreamEstablish.fromPartial(object.rawStreamEstablish)
        : undefined
    message.rawStreamAck =
      object.rawStreamAck !== undefined && object.rawStreamAck !== null
        ? RawStreamAck.fromPartial(object.rawStreamAck)
        : undefined
    message.rawStreamClose =
      object.rawStreamClose !== undefined && object.rawStreamClose !== null
        ? RawStreamClose.fromPartial(object.rawStreamClose)
        : undefined
    return message
  },
}

function createBaseRawStreamEstablish(): RawStreamEstablish {
  return { initiatorStreamId: 0 }
}

export const RawStreamEstablish = {
  encode(
    message: RawStreamEstablish,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.initiatorStreamId !== 0) {
      writer.uint32(8).uint32(message.initiatorStreamId)
    }
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RawStreamEstablish {
    const reader =
      input instanceof _m0.Reader ? input : _m0.Reader.create(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseRawStreamEstablish()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break
          }

          message.initiatorStreamId = reader.uint32()
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
  // Transform<RawStreamEstablish, Uint8Array>
  async *encodeTransform(
    source:
      | AsyncIterable<RawStreamEstablish | RawStreamEstablish[]>
      | Iterable<RawStreamEstablish | RawStreamEstablish[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [RawStreamEstablish.encode(p).finish()]
        }
      } else {
        yield* [RawStreamEstablish.encode(pkt as any).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, RawStreamEstablish>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<RawStreamEstablish> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [RawStreamEstablish.decode(p)]
        }
      } else {
        yield* [RawStreamEstablish.decode(pkt as any)]
      }
    }
  },

  fromJSON(object: any): RawStreamEstablish {
    return {
      initiatorStreamId: isSet(object.initiatorStreamId)
        ? globalThis.Number(object.initiatorStreamId)
        : 0,
    }
  },

  toJSON(message: RawStreamEstablish): unknown {
    const obj: any = {}
    if (message.initiatorStreamId !== 0) {
      obj.initiatorStreamId = Math.round(message.initiatorStreamId)
    }
    return obj
  },

  create<I extends Exact<DeepPartial<RawStreamEstablish>, I>>(
    base?: I,
  ): RawStreamEstablish {
    return RawStreamEstablish.fromPartial(base ?? ({} as any))
  },
  fromPartial<I extends Exact<DeepPartial<RawStreamEstablish>, I>>(
    object: I,
  ): RawStreamEstablish {
    const message = createBaseRawStreamEstablish()
    message.initiatorStreamId = object.initiatorStreamId ?? 0
    return message
  },
}

function createBaseRawStreamAck(): RawStreamAck {
  return { initiatorStreamId: 0, ackStreamId: 0, ackError: '' }
}

export const RawStreamAck = {
  encode(
    message: RawStreamAck,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.initiatorStreamId !== 0) {
      writer.uint32(8).uint32(message.initiatorStreamId)
    }
    if (message.ackStreamId !== 0) {
      writer.uint32(16).uint32(message.ackStreamId)
    }
    if (message.ackError !== '') {
      writer.uint32(26).string(message.ackError)
    }
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RawStreamAck {
    const reader =
      input instanceof _m0.Reader ? input : _m0.Reader.create(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseRawStreamAck()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break
          }

          message.initiatorStreamId = reader.uint32()
          continue
        case 2:
          if (tag !== 16) {
            break
          }

          message.ackStreamId = reader.uint32()
          continue
        case 3:
          if (tag !== 26) {
            break
          }

          message.ackError = reader.string()
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
  // Transform<RawStreamAck, Uint8Array>
  async *encodeTransform(
    source:
      | AsyncIterable<RawStreamAck | RawStreamAck[]>
      | Iterable<RawStreamAck | RawStreamAck[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [RawStreamAck.encode(p).finish()]
        }
      } else {
        yield* [RawStreamAck.encode(pkt as any).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, RawStreamAck>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<RawStreamAck> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [RawStreamAck.decode(p)]
        }
      } else {
        yield* [RawStreamAck.decode(pkt as any)]
      }
    }
  },

  fromJSON(object: any): RawStreamAck {
    return {
      initiatorStreamId: isSet(object.initiatorStreamId)
        ? globalThis.Number(object.initiatorStreamId)
        : 0,
      ackStreamId: isSet(object.ackStreamId)
        ? globalThis.Number(object.ackStreamId)
        : 0,
      ackError: isSet(object.ackError)
        ? globalThis.String(object.ackError)
        : '',
    }
  },

  toJSON(message: RawStreamAck): unknown {
    const obj: any = {}
    if (message.initiatorStreamId !== 0) {
      obj.initiatorStreamId = Math.round(message.initiatorStreamId)
    }
    if (message.ackStreamId !== 0) {
      obj.ackStreamId = Math.round(message.ackStreamId)
    }
    if (message.ackError !== '') {
      obj.ackError = message.ackError
    }
    return obj
  },

  create<I extends Exact<DeepPartial<RawStreamAck>, I>>(
    base?: I,
  ): RawStreamAck {
    return RawStreamAck.fromPartial(base ?? ({} as any))
  },
  fromPartial<I extends Exact<DeepPartial<RawStreamAck>, I>>(
    object: I,
  ): RawStreamAck {
    const message = createBaseRawStreamAck()
    message.initiatorStreamId = object.initiatorStreamId ?? 0
    message.ackStreamId = object.ackStreamId ?? 0
    message.ackError = object.ackError ?? ''
    return message
  },
}

function createBaseRawStreamClose(): RawStreamClose {
  return { streamId: 0, closeError: '' }
}

export const RawStreamClose = {
  encode(
    message: RawStreamClose,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.streamId !== 0) {
      writer.uint32(8).uint32(message.streamId)
    }
    if (message.closeError !== '') {
      writer.uint32(18).string(message.closeError)
    }
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RawStreamClose {
    const reader =
      input instanceof _m0.Reader ? input : _m0.Reader.create(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseRawStreamClose()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break
          }

          message.streamId = reader.uint32()
          continue
        case 2:
          if (tag !== 18) {
            break
          }

          message.closeError = reader.string()
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
  // Transform<RawStreamClose, Uint8Array>
  async *encodeTransform(
    source:
      | AsyncIterable<RawStreamClose | RawStreamClose[]>
      | Iterable<RawStreamClose | RawStreamClose[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [RawStreamClose.encode(p).finish()]
        }
      } else {
        yield* [RawStreamClose.encode(pkt as any).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, RawStreamClose>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<RawStreamClose> {
    for await (const pkt of source) {
      if (globalThis.Array.isArray(pkt)) {
        for (const p of pkt as any) {
          yield* [RawStreamClose.decode(p)]
        }
      } else {
        yield* [RawStreamClose.decode(pkt as any)]
      }
    }
  },

  fromJSON(object: any): RawStreamClose {
    return {
      streamId: isSet(object.streamId) ? globalThis.Number(object.streamId) : 0,
      closeError: isSet(object.closeError)
        ? globalThis.String(object.closeError)
        : '',
    }
  },

  toJSON(message: RawStreamClose): unknown {
    const obj: any = {}
    if (message.streamId !== 0) {
      obj.streamId = Math.round(message.streamId)
    }
    if (message.closeError !== '') {
      obj.closeError = message.closeError
    }
    return obj
  },

  create<I extends Exact<DeepPartial<RawStreamClose>, I>>(
    base?: I,
  ): RawStreamClose {
    return RawStreamClose.fromPartial(base ?? ({} as any))
  },
  fromPartial<I extends Exact<DeepPartial<RawStreamClose>, I>>(
    object: I,
  ): RawStreamClose {
    const message = createBaseRawStreamClose()
    message.streamId = object.streamId ?? 0
    message.closeError = object.closeError ?? ''
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined
}

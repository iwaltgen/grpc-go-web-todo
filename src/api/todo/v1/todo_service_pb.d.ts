import * as jspb from "google-protobuf"

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as todo_v1_todo_pb from '../../todo/v1/todo_pb';
import * as todo_v1_event_pb from '../../todo/v1/event_pb';

export class ListTodosRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTodosRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTodosRequest): ListTodosRequest.AsObject;
  static serializeBinaryToWriter(message: ListTodosRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTodosRequest;
  static deserializeBinaryFromReader(message: ListTodosRequest, reader: jspb.BinaryReader): ListTodosRequest;
}

export namespace ListTodosRequest {
  export type AsObject = {
  }
}

export class ListTodosResponse extends jspb.Message {
  getTodosList(): Array<todo_v1_todo_pb.Todo>;
  setTodosList(value: Array<todo_v1_todo_pb.Todo>): ListTodosResponse;
  clearTodosList(): ListTodosResponse;
  addTodos(value?: todo_v1_todo_pb.Todo, index?: number): todo_v1_todo_pb.Todo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTodosResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListTodosResponse): ListTodosResponse.AsObject;
  static serializeBinaryToWriter(message: ListTodosResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTodosResponse;
  static deserializeBinaryFromReader(message: ListTodosResponse, reader: jspb.BinaryReader): ListTodosResponse;
}

export namespace ListTodosResponse {
  export type AsObject = {
    todosList: Array<todo_v1_todo_pb.Todo.AsObject>,
  }
}

export class CreateTodoRequest extends jspb.Message {
  getTodo(): todo_v1_todo_pb.Todo | undefined;
  setTodo(value?: todo_v1_todo_pb.Todo): CreateTodoRequest;
  hasTodo(): boolean;
  clearTodo(): CreateTodoRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTodoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTodoRequest): CreateTodoRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTodoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTodoRequest;
  static deserializeBinaryFromReader(message: CreateTodoRequest, reader: jspb.BinaryReader): CreateTodoRequest;
}

export namespace CreateTodoRequest {
  export type AsObject = {
    todo?: todo_v1_todo_pb.Todo.AsObject,
  }
}

export class UpdateTodoRequest extends jspb.Message {
  getTodo(): todo_v1_todo_pb.Todo | undefined;
  setTodo(value?: todo_v1_todo_pb.Todo): UpdateTodoRequest;
  hasTodo(): boolean;
  clearTodo(): UpdateTodoRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTodoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTodoRequest): UpdateTodoRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateTodoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTodoRequest;
  static deserializeBinaryFromReader(message: UpdateTodoRequest, reader: jspb.BinaryReader): UpdateTodoRequest;
}

export namespace UpdateTodoRequest {
  export type AsObject = {
    todo?: todo_v1_todo_pb.Todo.AsObject,
  }
}

export class DeleteTodoRequest extends jspb.Message {
  getTodoId(): string;
  setTodoId(value: string): DeleteTodoRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTodoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTodoRequest): DeleteTodoRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteTodoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTodoRequest;
  static deserializeBinaryFromReader(message: DeleteTodoRequest, reader: jspb.BinaryReader): DeleteTodoRequest;
}

export namespace DeleteTodoRequest {
  export type AsObject = {
    todoId: string,
  }
}

export class SubscribeEventRequest extends jspb.Message {
  getEventsList(): Array<todo_v1_event_pb.Event>;
  setEventsList(value: Array<todo_v1_event_pb.Event>): SubscribeEventRequest;
  clearEventsList(): SubscribeEventRequest;
  addEvents(value: todo_v1_event_pb.Event, index?: number): SubscribeEventRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubscribeEventRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SubscribeEventRequest): SubscribeEventRequest.AsObject;
  static serializeBinaryToWriter(message: SubscribeEventRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubscribeEventRequest;
  static deserializeBinaryFromReader(message: SubscribeEventRequest, reader: jspb.BinaryReader): SubscribeEventRequest;
}

export namespace SubscribeEventRequest {
  export type AsObject = {
    eventsList: Array<todo_v1_event_pb.Event>,
  }
}

export class SubscribeEventResponse extends jspb.Message {
  getEvent(): todo_v1_event_pb.Event;
  setEvent(value: todo_v1_event_pb.Event): SubscribeEventResponse;

  getTodo(): todo_v1_todo_pb.Todo | undefined;
  setTodo(value?: todo_v1_todo_pb.Todo): SubscribeEventResponse;
  hasTodo(): boolean;
  clearTodo(): SubscribeEventResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubscribeEventResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SubscribeEventResponse): SubscribeEventResponse.AsObject;
  static serializeBinaryToWriter(message: SubscribeEventResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubscribeEventResponse;
  static deserializeBinaryFromReader(message: SubscribeEventResponse, reader: jspb.BinaryReader): SubscribeEventResponse;
}

export namespace SubscribeEventResponse {
  export type AsObject = {
    event: todo_v1_event_pb.Event,
    todo?: todo_v1_todo_pb.Todo.AsObject,
  }
}


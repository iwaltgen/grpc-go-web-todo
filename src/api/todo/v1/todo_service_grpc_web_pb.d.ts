import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as todo_v1_todo_pb from '../../todo/v1/todo_pb';
import * as todo_v1_event_pb from '../../todo/v1/event_pb';

import {
  CreateTodoRequest,
  DeleteTodoRequest,
  ListTodosRequest,
  ListTodosResponse,
  SubscribeEventRequest,
  SubscribeEventResponse,
  UpdateTodoRequest} from './todo_service_pb';

export class TodoServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  listTodos(
    request: ListTodosRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: ListTodosResponse) => void
  ): grpcWeb.ClientReadableStream<ListTodosResponse>;

  createTodo(
    request: CreateTodoRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  updateTodo(
    request: UpdateTodoRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteTodo(
    request: DeleteTodoRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  subscribeEvent(
    request: SubscribeEventRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<SubscribeEventResponse>;

}

export class TodoServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  listTodos(
    request: ListTodosRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<ListTodosResponse>;

  createTodo(
    request: CreateTodoRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  updateTodo(
    request: UpdateTodoRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  deleteTodo(
    request: DeleteTodoRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  subscribeEvent(
    request: SubscribeEventRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<SubscribeEventResponse>;

}


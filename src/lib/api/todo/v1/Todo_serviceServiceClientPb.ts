/**
 * @fileoverview gRPC-Web generated client stub for todo.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as todo_v1_todo_service_pb from '../../todo/v1/todo_service_pb';


export class TodoServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoListTodos = new grpcWeb.AbstractClientBase.MethodInfo(
    todo_v1_todo_service_pb.ListTodosResponse,
    (request: todo_v1_todo_service_pb.ListTodosRequest) => {
      return request.serializeBinary();
    },
    todo_v1_todo_service_pb.ListTodosResponse.deserializeBinary
  );

  listTodos(
    request: todo_v1_todo_service_pb.ListTodosRequest,
    metadata: grpcWeb.Metadata | null): Promise<todo_v1_todo_service_pb.ListTodosResponse>;

  listTodos(
    request: todo_v1_todo_service_pb.ListTodosRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: todo_v1_todo_service_pb.ListTodosResponse) => void): grpcWeb.ClientReadableStream<todo_v1_todo_service_pb.ListTodosResponse>;

  listTodos(
    request: todo_v1_todo_service_pb.ListTodosRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: todo_v1_todo_service_pb.ListTodosResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/todo.v1.TodoService/ListTodos',
        request,
        metadata || {},
        this.methodInfoListTodos,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/todo.v1.TodoService/ListTodos',
    request,
    metadata || {},
    this.methodInfoListTodos);
  }

  methodInfoCreateTodo = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: todo_v1_todo_service_pb.CreateTodoRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  createTodo(
    request: todo_v1_todo_service_pb.CreateTodoRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  createTodo(
    request: todo_v1_todo_service_pb.CreateTodoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  createTodo(
    request: todo_v1_todo_service_pb.CreateTodoRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/todo.v1.TodoService/CreateTodo',
        request,
        metadata || {},
        this.methodInfoCreateTodo,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/todo.v1.TodoService/CreateTodo',
    request,
    metadata || {},
    this.methodInfoCreateTodo);
  }

  methodInfoUpdateTodo = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: todo_v1_todo_service_pb.UpdateTodoRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  updateTodo(
    request: todo_v1_todo_service_pb.UpdateTodoRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  updateTodo(
    request: todo_v1_todo_service_pb.UpdateTodoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  updateTodo(
    request: todo_v1_todo_service_pb.UpdateTodoRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/todo.v1.TodoService/UpdateTodo',
        request,
        metadata || {},
        this.methodInfoUpdateTodo,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/todo.v1.TodoService/UpdateTodo',
    request,
    metadata || {},
    this.methodInfoUpdateTodo);
  }

  methodInfoDeleteTodo = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: todo_v1_todo_service_pb.DeleteTodoRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteTodo(
    request: todo_v1_todo_service_pb.DeleteTodoRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  deleteTodo(
    request: todo_v1_todo_service_pb.DeleteTodoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteTodo(
    request: todo_v1_todo_service_pb.DeleteTodoRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/todo.v1.TodoService/DeleteTodo',
        request,
        metadata || {},
        this.methodInfoDeleteTodo,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/todo.v1.TodoService/DeleteTodo',
    request,
    metadata || {},
    this.methodInfoDeleteTodo);
  }

  methodInfoSubscribeEvent = new grpcWeb.AbstractClientBase.MethodInfo(
    todo_v1_todo_service_pb.SubscribeEventResponse,
    (request: todo_v1_todo_service_pb.SubscribeEventRequest) => {
      return request.serializeBinary();
    },
    todo_v1_todo_service_pb.SubscribeEventResponse.deserializeBinary
  );

  subscribeEvent(
    request: todo_v1_todo_service_pb.SubscribeEventRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/todo.v1.TodoService/SubscribeEvent',
      request,
      metadata || {},
      this.methodInfoSubscribeEvent);
  }

}


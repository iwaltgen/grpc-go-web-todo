/**
 * @fileoverview gRPC-Web generated client stub for todo.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var todo_v1_todo_pb = require('../../todo/v1/todo_pb.js')

var todo_v1_event_pb = require('../../todo/v1/event_pb.js')
const proto = {};
proto.todo = {};
proto.todo.v1 = require('./todo_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.v1.TodoServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.v1.TodoServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.v1.ListTodosRequest,
 *   !proto.todo.v1.ListTodosResponse>}
 */
const methodDescriptor_TodoService_ListTodos = new grpc.web.MethodDescriptor(
  '/todo.v1.TodoService/ListTodos',
  grpc.web.MethodType.UNARY,
  proto.todo.v1.ListTodosRequest,
  proto.todo.v1.ListTodosResponse,
  /**
   * @param {!proto.todo.v1.ListTodosRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.v1.ListTodosResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.v1.ListTodosRequest,
 *   !proto.todo.v1.ListTodosResponse>}
 */
const methodInfo_TodoService_ListTodos = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.v1.ListTodosResponse,
  /**
   * @param {!proto.todo.v1.ListTodosRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.v1.ListTodosResponse.deserializeBinary
);


/**
 * @param {!proto.todo.v1.ListTodosRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.v1.ListTodosResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.v1.ListTodosResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.v1.TodoServiceClient.prototype.listTodos =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.v1.TodoService/ListTodos',
      request,
      metadata || {},
      methodDescriptor_TodoService_ListTodos,
      callback);
};


/**
 * @param {!proto.todo.v1.ListTodosRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.v1.ListTodosResponse>}
 *     A native promise that resolves to the response
 */
proto.todo.v1.TodoServicePromiseClient.prototype.listTodos =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.v1.TodoService/ListTodos',
      request,
      metadata || {},
      methodDescriptor_TodoService_ListTodos);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.v1.CreateTodoRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_TodoService_CreateTodo = new grpc.web.MethodDescriptor(
  '/todo.v1.TodoService/CreateTodo',
  grpc.web.MethodType.UNARY,
  proto.todo.v1.CreateTodoRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.todo.v1.CreateTodoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.v1.CreateTodoRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_TodoService_CreateTodo = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.todo.v1.CreateTodoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.todo.v1.CreateTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.v1.TodoServiceClient.prototype.createTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.v1.TodoService/CreateTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_CreateTodo,
      callback);
};


/**
 * @param {!proto.todo.v1.CreateTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.todo.v1.TodoServicePromiseClient.prototype.createTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.v1.TodoService/CreateTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_CreateTodo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.v1.UpdateTodoRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_TodoService_UpdateTodo = new grpc.web.MethodDescriptor(
  '/todo.v1.TodoService/UpdateTodo',
  grpc.web.MethodType.UNARY,
  proto.todo.v1.UpdateTodoRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.todo.v1.UpdateTodoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.v1.UpdateTodoRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_TodoService_UpdateTodo = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.todo.v1.UpdateTodoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.todo.v1.UpdateTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.v1.TodoServiceClient.prototype.updateTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.v1.TodoService/UpdateTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_UpdateTodo,
      callback);
};


/**
 * @param {!proto.todo.v1.UpdateTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.todo.v1.TodoServicePromiseClient.prototype.updateTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.v1.TodoService/UpdateTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_UpdateTodo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.v1.DeleteTodoRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_TodoService_DeleteTodo = new grpc.web.MethodDescriptor(
  '/todo.v1.TodoService/DeleteTodo',
  grpc.web.MethodType.UNARY,
  proto.todo.v1.DeleteTodoRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.todo.v1.DeleteTodoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.v1.DeleteTodoRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_TodoService_DeleteTodo = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.todo.v1.DeleteTodoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.todo.v1.DeleteTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.v1.TodoServiceClient.prototype.deleteTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.v1.TodoService/DeleteTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_DeleteTodo,
      callback);
};


/**
 * @param {!proto.todo.v1.DeleteTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.todo.v1.TodoServicePromiseClient.prototype.deleteTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.v1.TodoService/DeleteTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_DeleteTodo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.v1.SubscribeEventRequest,
 *   !proto.todo.v1.SubscribeEventResponse>}
 */
const methodDescriptor_TodoService_SubscribeEvent = new grpc.web.MethodDescriptor(
  '/todo.v1.TodoService/SubscribeEvent',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.todo.v1.SubscribeEventRequest,
  proto.todo.v1.SubscribeEventResponse,
  /**
   * @param {!proto.todo.v1.SubscribeEventRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.v1.SubscribeEventResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.v1.SubscribeEventRequest,
 *   !proto.todo.v1.SubscribeEventResponse>}
 */
const methodInfo_TodoService_SubscribeEvent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.v1.SubscribeEventResponse,
  /**
   * @param {!proto.todo.v1.SubscribeEventRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.v1.SubscribeEventResponse.deserializeBinary
);


/**
 * @param {!proto.todo.v1.SubscribeEventRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.todo.v1.SubscribeEventResponse>}
 *     The XHR Node Readable Stream
 */
proto.todo.v1.TodoServiceClient.prototype.subscribeEvent =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/todo.v1.TodoService/SubscribeEvent',
      request,
      metadata || {},
      methodDescriptor_TodoService_SubscribeEvent);
};


/**
 * @param {!proto.todo.v1.SubscribeEventRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.todo.v1.SubscribeEventResponse>}
 *     The XHR Node Readable Stream
 */
proto.todo.v1.TodoServicePromiseClient.prototype.subscribeEvent =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/todo.v1.TodoService/SubscribeEvent',
      request,
      metadata || {},
      methodDescriptor_TodoService_SubscribeEvent);
};


module.exports = proto.todo.v1;


import * as eventModule from './event_pb';
export const { Event } = eventModule;

import * as todoModule from './todo_pb';
export const { Todo } = todoModule;

import * as todoServiceModule from './todo_service_pb';
export const {
  ListTodosRequest,
  ListTodosResponse,
  CreateTodoRequest,
  UpdateTodoRequest,
  DeleteTodoRequest,
  SubscribeEventRequest,
  SubscribeEventResponse,
} = todoServiceModule;

export { TodoServiceClient } from './Todo_serviceServiceClientPb';

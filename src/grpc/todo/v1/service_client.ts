import { Observable, Observer } from 'rxjs';
import * as grpcWeb from 'grpc-web';
import {
  TodoServiceClient,
  ListTodosRequest,
  ListTodosResponse,
  CreateTodoRequest,
  UpdateTodoRequest,
  DeleteTodoRequest,
  SubscribeEventRequest,
  SubscribeEventResponse,
  Event,
} from '../../../api/todo/v1';
import type { Empty } from '../../../api/google/protobuf';
import * as convert from '../../convert';
import type { Todo } from '../../../entity';


export class TodoClient {
  readonly client: TodoServiceClient;

  constructor (hostname: string,
               credentials?: { [index: string]: string; },
               options?: { [index: string]: string; }) {
    this.client = new TodoServiceClient(hostname, credentials, options);
  }

  public async listTodo(): Promise<Todo[]> {
    return new Promise((resolve: (value?: Todo[]) => void, reject: (reason?: any) => void) => {
      const metadata = undefined;
      const request = new ListTodosRequest();

      this.client.listTodos(request, metadata,(err: grpcWeb.Error, response: ListTodosResponse) => {
        if (err && err.code !== grpcWeb.StatusCode.OK) {
          reject(err);
          return;
        }

        resolve(response.getTodosList().map(convert.todoFromProto));
      });
    });
  }

  public async create(todo: Todo): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason?: any) => void) => {
      const metadata = undefined;
      const request = new CreateTodoRequest();
      request.setTodo(convert.todoProto(todo));

      this.client.createTodo(request, metadata,(err: grpcWeb.Error, response: Empty) => {
        if (err && err.code !== grpcWeb.StatusCode.OK) {
          reject(err);
          return;
        }

        resolve();
      });
    });
  }

  public async update(todo: Todo): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason?: any) => void) => {
      const metadata = undefined;
      const request = new UpdateTodoRequest();
      request.setTodo(convert.todoProto(todo));

      this.client.updateTodo(request, metadata,(err: grpcWeb.Error, response: Empty) => {
        if (err && err.code !== grpcWeb.StatusCode.OK) {
          reject(err);
          return;
        }

        resolve();
      });
    });
  }

  public async delete(todo: Todo): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason?: any) => void) => {
      const metadata = undefined;
      const request = new DeleteTodoRequest();
      request.setTodoId(todo.id);

      this.client.deleteTodo(request, metadata,(err: grpcWeb.Error, response: Empty) => {
        if (err && err.code !== grpcWeb.StatusCode.OK) {
          reject(err);
          return;
        }

        resolve();
      });
    });
  }
}

import { writable, derived } from 'svelte/store';
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
} from '../api/todo/v1';
import * as convert from '../api/convert'
import type { Todo } from '../entity'

const serviceStore = writable([] as Todo[]);

const serviceClient = new TodoServiceClient('https://localhost:8443', null, null);

function list(): Observable<Todo[]> {
  return Observable.create((observer: Observer<Todo[]>) => {
    const request = new ListTodosRequest();
    serviceClient.listTodos(request, undefined,
      (err: grpcWeb.Error, response: ListTodosResponse) => {
        if (err) {
          if (err.code !== grpcWeb.StatusCode.OK) {
            console.log('list error: ', err)
            observer.error(err)
            observer.complete()
            return
          }
        }

        console.log('list response: ', response);
        let results = response.getTodosList()
          .map(convert.todoFromProto)

        serviceStore.set(results)
        observer.next(results)
        observer.complete()
      });
  })
}

export const todoStore = derived(serviceStore, store => store)

export const todoService = {
  list
}

const enableDevTools = globalThis.__GRPCWEB_DEVTOOLS__ || (() => {});
enableDevTools([
  serviceClient,
]);

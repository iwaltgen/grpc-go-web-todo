import { writable, Readable } from 'svelte/store';
import { Observable, Observer, interval } from 'rxjs';
import { take } from 'rxjs/operators';
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
import type { Todo } from '../entity'

const store = writable([] as Array<Todo>);

const serviceClient = new TodoServiceClient('https://localhost:8443', null, null);

function list(): Observable<Array<Todo>> {
  return Observable.create((observer: Observer<Array<Todo>>) => {
    const request = new ListTodosRequest();
    const call = serviceClient.listTodos(request, undefined,
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
          .map(it => it.toObject() as Todo)

        store.set(results)
        observer.next(results)
        observer.complete()
      });

    call.on('status', (status: grpcWeb.Status) => {
      console.log('list status: ', status)
    });
  })
}

const enableDevTools = globalThis.__GRPCWEB_DEVTOOLS__ || (() => {});
enableDevTools([
  serviceClient,
]);

// interval(3000).pipe(
//   take(10),
// ).subscribe(() => {
//   console.log('request list')
//   list().subscribe(() => console.log('request list done'))
// })

// export const todos: Readable<Array<Todo>> = store
export const todos = store

export const todoService = {
  list
}

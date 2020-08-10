import { writable, derived, Readable } from 'svelte/store';
import { todoClient } from './grpc';
import type { Error, Status } from '../grpc';
import type { Todo } from '../entity';
import type { EventTodo } from './value';
import { now } from './timestamp';
import { Event } from '../api/todo/v1';

export class TodoService {
  private todos = writable([] as Todo[]);
  private client = todoClient;

  constructor() {
    this.subscribe();
  }

  async list(): Promise<Todo[]> {
    const items = await this.client.list();
    this.todos.set(items);
    return items;
  }

  async create(todo: Todo): Promise<void> {
    const nowTs = now();
    todo.modifiedAt = nowTs;
    todo.createdAt = nowTs;
    await this.client.create(todo);
  }

  async update(todo: Todo): Promise<void> {
    todo.modifiedAt = now();
    await this.client.update(todo);
  }

  async delete(todo: Todo): Promise<void> {
    await this.client.delete(todo);
  }

  store(): Readable<Todo[]> {
    return derived(this.todos, ($todos) => $todos);
  }

  private subscribe() {
    this.client.subscribe().subscribe(
      (evt: EventTodo) => {
        this.todos.update((list) => {
          switch (evt.event) {
            case Event.EVENT_CREATE:
              return [evt.todo, ...list];

            case Event.EVENT_UPDATE:
              return list.map((item) => {
                if (item.id === evt.todo.id) {
                  return evt.todo;
                }
                return item;
              });

            case Event.EVENT_DELETE:
              return list.filter((item) => item.id !== evt.todo.id);
          }
          return list.concat(evt.todo);
        });
      },
      (error: Error | Status) => {
        console.log('todo service subscribe error: ', error);
      },
    );
  }
}

export const todoService = new TodoService();

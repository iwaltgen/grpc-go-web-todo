import { writable, derived, Readable } from 'svelte/store';
import { todoClient } from './grpc';
import type { Todo } from '../entity';
import { now } from './timestamp';

export class TodoService {
  private store = writable([] as Todo[]);
  private client = todoClient;

  async requestList() {
    const todos = await this.client.listTodo();
    this.store.set(todos);
  }

  async create(todo: Todo) {
    const nowTimestamp = now();
    todo.modifiedAt = nowTimestamp;
    todo.createdAt = nowTimestamp;
    await this.client.create(todo);
    // TODO(iwaltgen): change subscribe
		await this.requestList();
  }

  async update(todo: Todo) {
    todo.modifiedAt = now();
    await this.client.update(todo);
    // TODO(iwaltgen): change subscribe
		await this.requestList();
  }

  async delete(todo: Todo) {
    await this.client.delete(todo);
    // TODO(iwaltgen): change subscribe
		await this.requestList();
  }

  subscribe(): Readable<Todo[]> {
    return derived(this.store, $todos => $todos);
  }
}

export const todoService = new TodoService();

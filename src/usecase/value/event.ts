import type { Event } from '../../api/todo/v1';
import type { Todo } from '../../entity';

export interface EventTodo {
  event: Event;
  todo: Todo;
}

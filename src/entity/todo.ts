import type { Timestamp } from './timestamp';

export interface Todo {
  id: string;
  description: string;
  completed: boolean;
  modifiedAt?: Timestamp;
  createdAt?: Timestamp;
}

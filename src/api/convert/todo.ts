import {
  Todo as TodoProto,
  // Event as EventProto,
} from '../todo/v1';
import type { Todo } from '../../entity'
import { timestampProto } from './wkt'

export function todoProto(v: Todo): TodoProto {
  const ret = new TodoProto();
  ret.setId(v.id)
  ret.setDescription(v.description)
  ret.setCompleted(v.completed)
  ret.setModifiedAt(timestampProto(v.modifiedAt))
  ret.setCreatedAt(timestampProto(v.createdAt))
  return ret
}

export function todoFromProto(v: TodoProto): Todo {
  return v.toObject() as Todo
}

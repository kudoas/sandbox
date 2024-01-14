import { Injectable } from '@angular/core';
import { TodoListState } from './state';
import { Todo } from './todo';

@Injectable()
export class TodoListUsecase {
  #state = new TodoListState();
  state = this.#state.asReadonly();

  addTodo(title: string): void {
    this.#state.addTodo(title);
  }

  setTodoCompleted(id: Todo['id'], completed: boolean): void {
    this.#state.setCompleted(id, completed);
  }
}

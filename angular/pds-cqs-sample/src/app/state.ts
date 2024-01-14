import { computed, signal } from '@angular/core';
import { SignalState } from './signal-state';
import { Todo } from './todo';

export type State = {
  todos: Todo[];
  incompleteTodos: Todo[];
  completeTodos: Todo[];
};

export class TodoListState implements SignalState<State> {
  todos = signal<Todo[]>([
    {
      id: 1,
      title: 'Learn Signals',
      completed: false,
    },
  ]);

  completedTodos = computed(() =>
    this.todos().filter((todo) => todo.completed)
  );
  incompleteTodos = computed(() =>
    this.todos().filter((todo) => !todo.completed)
  );

  addTodo(title: string): void {
    this.todos.update((todos) => [
      ...todos,
      {
        id: todos.length + 1,
        title,
        completed: false,
      },
    ]);
  }
  setCompleted(id: Todo['id'], completed: boolean): void {
    this.todos.update((todos) =>
      todos.map((todo) =>
        todo.id === id ? { ...todo, completed: completed } : todo
      )
    );
  }

  asReadonly() {
    return {
      todos: this.todos.asReadonly(),
      incompleteTodos: this.incompleteTodos,
      completeTodos: this.completedTodos,
    };
  }
}

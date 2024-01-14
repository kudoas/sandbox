import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { TodoListUsecase } from './usecase';
import { TodoListViewComponent } from './views/todo-list-view/todo-list-view.component';

@Component({
  selector: 'app-pds-cqs-todo-list',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, TodoListViewComponent],
  providers: [TodoListUsecase],
  template: `
    <div>Incomplete Items: {{ usecase.state.incompleteTodos().length }}</div>

    <div>
      <form>
        <input [formControl]="newTodoInput" />
        <button (click)="addTodo()" [disabled]="newTodoInput.invalid">
          Add Todo item
        </button>
      </form>
    </div>

    <app-todo-list-view
      [items]="usecase.state.todos()"
      (changeCompleted)="usecase.setTodoCompleted($event.id, $event.completed)"
    />
  `,
})
export class PdsCqsTodoListComponent {
  usecase = inject(TodoListUsecase);

  newTodoInput = new FormControl('', {
    validators: [Validators.required],
    nonNullable: true,
  });

  addTodo() {
    const title = this.newTodoInput.value;
    this.usecase.addTodo(title);
    this.newTodoInput.reset();
  }
}

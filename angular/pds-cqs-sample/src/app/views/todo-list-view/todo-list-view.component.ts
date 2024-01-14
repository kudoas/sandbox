import { CommonModule } from '@angular/common';
import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Input,
  Output,
} from '@angular/core';
import { Todo } from '../../todo';

@Component({
  selector: 'app-todo-list-view',
  standalone: true,
  imports: [CommonModule],
  template: `<ul>
    <li *ngFor="let todo of items">
      <label>
        <input
          #checkbox
          type="checkbox"
          [checked]="todo.completed"
          (change)="
            changeCompleted.emit({ id: todo.id, completed: checkbox.checked })
          "
        />
        {{ todo.title }}
      </label>
    </li>
  </ul> `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TodoListViewComponent {
  @Input()
  items: Todo[] = [];

  @Output()
  changeCompleted = new EventEmitter<{ id: Todo['id']; completed: boolean }>();
}

import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { DialogDemoComponent } from './dialog-demo.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, DialogDemoComponent],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `<app-dialog-demo></app-dialog-demo>`,
  styles: [],
})
export class App {}

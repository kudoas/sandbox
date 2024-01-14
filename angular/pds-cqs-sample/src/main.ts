import 'zone.js/dist/zone';
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { provideHttpClient } from '@angular/common/http';
import { bootstrapApplication } from '@angular/platform-browser';
import { PdsCqsTodoListComponent } from './app/pds-cqs.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, PdsCqsTodoListComponent],
  template: `<app-pds-cqs-todo-list />`,
})
export class App {}

bootstrapApplication(App, {
  providers: [provideHttpClient()],
});

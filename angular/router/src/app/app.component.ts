import { Component } from '@angular/core';
import { RouterOutlet, RouterLink } from '@angular/router';
import { LinkDirective } from './link.directive';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink, LinkDirective],
  template: `
    <h1>
      {{ title }}
    </h1>
    <a routerLink="/home">Home</a>
    <a appLink>Home</a>
    <router-outlet></router-outlet>
  `,
})
export class AppComponent {
  title = 'router';
}

@Component({
  selector: 'app-home',
  standalone: true,
  template: ` <h2>Home</h2> `,
})
export class HomeComponent {}

import { Directive, HostBinding, HostListener, inject } from '@angular/core';
import { Router } from '@angular/router';

@Directive({
  selector: '[appLink]',
  standalone: true,
})
export class LinkDirective {
  route = inject(Router);

  @HostBinding('attr.href')
  backLink = '/home';

  @HostListener('click')
  onClick() {
    this.route.navigateByUrl(this.backLink);

    // Prevent the default link behavior
    return false;
  }
}

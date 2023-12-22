import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SandLibComponent } from './sand-lib.component';

describe('SandLibComponent', () => {
  let component: SandLibComponent;
  let fixture: ComponentFixture<SandLibComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SandLibComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SandLibComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

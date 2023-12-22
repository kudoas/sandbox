import { TestBed } from '@angular/core/testing';

import { SandLibService } from './sand-lib.service';

describe('SandLibService', () => {
  let service: SandLibService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SandLibService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

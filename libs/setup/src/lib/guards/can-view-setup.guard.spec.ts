import { TestBed } from '@angular/core/testing';

import { CanViewSetupGuard } from './can-view-setup.guard';

describe('CanViewSetupGuard', () => {
  let guard: CanViewSetupGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(CanViewSetupGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});

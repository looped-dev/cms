import { TestBed } from '@angular/core/testing';

import { CannotViewSetupGuard } from './cannot-view-setup.guard';

describe('CannotViewSetupGuard', () => {
  let guard: CannotViewSetupGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(CannotViewSetupGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});

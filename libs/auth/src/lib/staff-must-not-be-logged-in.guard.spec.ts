import { TestBed } from '@angular/core/testing';

import { StaffMustNotBeLoggedInGuard } from './staff-must-not-be-logged-in.guard';

describe('StaffMustNotBeLoggedInGuard', () => {
  let guard: StaffMustNotBeLoggedInGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(StaffMustNotBeLoggedInGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});

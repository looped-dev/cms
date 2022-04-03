import { TestBed } from '@angular/core/testing';

import { StaffMustBeLoggedInGuard } from './staff-must-be-logged-in.guard';

describe('StaffMustBeLoggedInGuard', () => {
  let guard: StaffMustBeLoggedInGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(StaffMustBeLoggedInGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});

export interface SetupRegister {
  isSetup: boolean;
}

export function createSetupRegister(params: Partial<SetupRegister>) {
  return {
    isSetup: params.isSetup ?? false,
  } as SetupRegister;
}

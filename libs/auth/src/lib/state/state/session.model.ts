import { Staff } from '@looped-cms/graphql';

export interface Session {
  refreshToken?: string;
  accessToken?: string;
  staff?: Omit<Staff, 'emailVerified'>;
}

export function createSession(params: Partial<Session>) {
  return {} as Session;
}

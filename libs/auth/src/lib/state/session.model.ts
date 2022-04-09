import { Staff } from '@looped-cms/graphql';

export interface Session {
  refreshToken?: string;
  accessToken?: string;
  staff?: Omit<Staff, 'emailVerified' | 'createdAt' | 'updatedAt'>;
}

export function createSession(params?: Partial<Session>) {
  return { ...params } as Session;
}

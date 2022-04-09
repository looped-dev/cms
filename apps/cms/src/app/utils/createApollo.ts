/* eslint-disable @typescript-eslint/no-unused-vars */
import { ApolloLink, InMemoryCache } from '@apollo/client/core';
import { setContext } from '@apollo/client/link/context';
import { SessionQuery } from '@looped-cms/auth';
import { HttpLink } from 'apollo-angular/http';
import { environment } from '../../environments/environment';

export function createApollo(httpLink: HttpLink, sessionQuery: SessionQuery) {
  const basicAuthentication = setContext((_operation, _context) => ({
    headers: {
      Accept: 'charset=utf-8',
    },
  }));
  const bearerTokenAuthentication = setContext((_operation, _context) => {
    // fetch the token from the store
    const token = sessionQuery.getValue().accessToken;
    if (token === null) {
      return {};
    } else {
      return {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      };
    }
  });
  const link = ApolloLink.from([
    basicAuthentication,
    bearerTokenAuthentication,
    httpLink.create({ uri: environment.graphql.endpoint }),
  ]);
  const cache = new InMemoryCache();
  return {
    link,
    cache,
  };
}

/* eslint-disable @typescript-eslint/no-unused-vars */
import { ApolloLink, InMemoryCache } from '@apollo/client/core';
import { setContext } from '@apollo/client/link/context';
import { SessionQuery, SessionService } from '@looped-cms/auth';
import { HttpLink } from 'apollo-angular/http';
import { environment } from '../../environments/environment';
import { onError } from '@apollo/client/link/error';

export function createApollo(
  httpLink: HttpLink,
  sessionQuery: SessionQuery,
  sessionService: SessionService
) {
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
  const errorLink = onError(
    ({ forward, graphQLErrors, networkError, operation }) => {
      console.log(graphQLErrors, networkError);
      if (graphQLErrors) {
        graphQLErrors.map(({ message, locations, path, extensions }) => {
          if (extensions['code'] === '401') {
            sessionService.refreshToken().subscribe(() => {
              return forward(operation);
            });
          }
        });
      }
    }
  );
  const link = ApolloLink.from([
    errorLink,
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

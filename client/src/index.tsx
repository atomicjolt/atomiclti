import React from 'react';
import ReactDOM from 'react-dom';
import { ApolloProvider } from '@apollo/client/react';
import {
  ApolloClient,
  InMemoryCache,
  ApolloLink,
  HttpLink,
  Operation,
  NextLink,
  Observable,
  FetchResult,
  concat,
} from '@apollo/client';

import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

const httpLink = new HttpLink({
  uri: 'https://atomiclti.atomicjolt.xyz/graphql',
});

const idTokenMiddleware = new ApolloLink(
  (operation: Operation, forward: NextLink): Observable<FetchResult> => {
    operation.setContext(({ headers = {} }) => ({
      ...headers,
      headers: {
        authorization: window.LAUNCH_TOKEN,
      },
    }));

    return forward(operation);
  },
);

const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: concat(idTokenMiddleware, httpLink),
});

ReactDOM.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

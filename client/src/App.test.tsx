import React from 'react';
import { render, screen } from '@testing-library/react';
import { MockedProvider } from '@apollo/client/testing';
import App from './App';

const mocks = []; //TODO mock graphql queries

test('renders main page', () => {
  render(
    <MockedProvider mocks={mocks} addTypename={false}>
      <App />
    </MockedProvider>
  );
  const linkElement = screen.getByText(/Hello World!/i);
  expect(linkElement).toBeInTheDocument();
});

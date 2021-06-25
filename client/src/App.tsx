import React from 'react';
import { useQuery, gql } from '@apollo/client';
import './styles/base.scss';

interface HelloResult {
  helloWorld: string;
}

const GET_HELLO = gql`
  query {
    helloWorld
  }
`;

const App: React.FC = () => {
  const { loading, error, data } = useQuery<HelloResult>(GET_HELLO);

  const methods = {
    get result() {
      if (loading) return <span>'Loading...'</span>;
      if (error) return <span>'Whoops'</span>;
      if (data === undefined) return <span>'No result.'</span>;

      return <span>{data.helloWorld}</span>;
    },
  };

  return (
    <div>
      <h1>Hello World!</h1>
      {methods.result}
    </div>
  );
};

export default App;

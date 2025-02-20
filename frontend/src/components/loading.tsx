import React from 'react';
import {LoadingProps} from './interfaces'

const Loading: React.FC<LoadingProps> = ({ data }) => {
  return <p>Loading {data}...</p>;
};

export default Loading;

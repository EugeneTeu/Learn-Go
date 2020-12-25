import React, { FC, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { increment, selectCount } from './counterSlice';

export const Counter: FC = () => {
  const count = useSelector(selectCount);
  const dispatch = useDispatch();
  const handleClick = (_) => {
    dispatch(increment());
  };

  return (
    <div>
      <button onClick={handleClick}>click me</button> <h1>{count}</h1>
    </div>
  );
};

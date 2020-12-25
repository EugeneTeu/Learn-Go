import React, { FC } from 'react';
import { Provider } from 'react-redux';
import Main from '../pages/main';
import store from './store';
import styles from './app.module.css';

const App: FC = () => {
  return (
    <div className={styles.app}>
      <Provider store={store}>
        <Main />
      </Provider>
    </div>
  );
};

export default App;

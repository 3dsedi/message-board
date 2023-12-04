import React from 'react';
import { useData } from '../DataContext';

const ErrorPage = () => {
  const { showErrorPage, errorMessage, closeErrorPage  } = useData();

  if (!showErrorPage) {
    return null; 
  }


  return (
    <div
    onClick={closeErrorPage}
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        background: 'rgba(0, 0, 0, 0.8)',
        zIndex: 9999,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        color: '#fff',
        fontSize: '24px',
        
      }}
    >
      <div>
        <p>{errorMessage}</p>
      </div>
    </div>
  );
};

export default ErrorPage;

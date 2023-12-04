import React, { useRef } from 'react';
import { useData } from '../DataContext'
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';

const SignInForm = () => {
  const { handleLogin, setShowErrorPage, setErrorMessage } = useData();

  const emailRef = useRef(null);
  const passwordRef = useRef(null);
  
  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = {
      email: emailRef.current.value,
      password: passwordRef.current.value,
    };
    if (!emailRef || !passwordRef) {
      return;
    }
    try {
      const loginData = await handleLogin(formData); 
      if (loginData && loginData.id) {
        console.log('Login successful!', loginData);
      } else {
        setErrorMessage('Login failed. Please Complete the Form');
        setShowErrorPage(true);
      }
    } catch (error) {
      console.error('Login error:', error);
      setErrorMessage('Login error occurred. Please try again.');
      setShowErrorPage(true);
    }
  };

  return (
    <div className='p-4'>
      <form onSubmit={handleSubmit} style={{ marginTop: '30px' }}>
        <TextField
          type="email"
          label="Email"
          variant="outlined"
          inputRef={emailRef}
          fullWidth
          margin="normal"
          size='small'
          className="custom-textfield"
        />
        <TextField
          type="password"
          label="Password"
          variant="outlined"
          inputRef={passwordRef}
          fullWidth
          margin="normal"
          size='small'
          className="custom-textfield"
        />
        <Button
          variant="contained"
          color="primary"
          type="submit"
        >
          Login
        </Button>
      </form>
    </div>
  );
};

export default SignInForm;

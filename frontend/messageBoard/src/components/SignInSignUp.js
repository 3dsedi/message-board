import React, { useState } from 'react';
import { useData } from '../DataContext';
import Box from '@mui/material/Box';
import Tab from '@mui/material/Tab';
import Tabs from '@mui/material/Tabs';
import Profile from './Profile';
import SignInForm from './SignInForm';
import SignUpForm from './SignUpForm';

const SignInSignUp = () => {
  const {loggedInUserId} = useData();
  
  const [selectedTab, setSelectedTab] = useState(0);

  const handleTabChange = (event, newValue) => {
    setSelectedTab(newValue);
  };

  return (
    <Box sx={{ maxWidth: 400, mt: 2 }}>
      {loggedInUserId != null ? (
        <Profile />
      ) : (
        <>
          <Box sx={{ borderBottom: 2, borderColor: '#dfdfdf' }}>
            <Tabs value={selectedTab} onChange={handleTabChange} sx={{marginLeft:'15%'}}>
              <Tab label="Sign In" sx={{ color: 'gray' }}/>
              <Tab label="Sign Up" sx={{ color: 'gray' }} />
            </Tabs>
          </Box>
          <Box sx={{ p: 3 }}>
            {selectedTab === 0 && <SignInForm />}
            {selectedTab === 1 && <SignUpForm />}
          </Box>
        </>
      )}
    </Box>
  )
};

export default SignInSignUp;

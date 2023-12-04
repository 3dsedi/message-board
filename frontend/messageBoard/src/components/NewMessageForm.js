import React, { useRef } from 'react';
import { useData } from '../DataContext';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

const NewMessageForm = () => {
  const {handlePostMessage, loggedInUserId, updateMessages} = useData();

  const contentRef = useRef(null);
  

  const handlesubmitMessage = async (event) => {
    event.preventDefault();
    const messageData = {
      content: contentRef.current.value,
      user_id: loggedInUserId,
    };
    try {
      const message = await handlePostMessage(messageData); 
      if (message && loggedInUserId) {
        console.log('message post successful!', message);
        updateMessages()
      } else {
        console.log('Message Posting failed.');
      }
    } catch (error) {
      console.error('message posting error:', error);
    }
    contentRef.current.value = "";
  };
  return (
    <div style={{ display: 'flex', flexDirection: 'column' }}>
      <div>
        <Typography variant="h6" sx={{color:"#434343", fontSize: "16px" , fontWeight: "bold"}} className="title">Add New Message</Typography>
        <textarea
          ref={contentRef}
          placeholder="Type here ..."
          rows={2}
          className="form-control"
          style={{border: '1px solid #ccc', borderRadius: '4px', padding: '4px', marginTop: '8px'}}
        />
      </div>
      <div style={{ display: 'flex', justifyContent: 'flex-end', marginTop: '8px' }}>
        <Button
          variant="contained"
          onClick={handlesubmitMessage}
          style={{ color: "#f8ebeb", backgroundColor: "#785fbd", textTransform: "none", fontWeight: "bold", fontSize: "13px" }}
          type='submit'
        >
          Send Message
        </Button>
      </div>
    </div>
  );
  
};

export default NewMessageForm;

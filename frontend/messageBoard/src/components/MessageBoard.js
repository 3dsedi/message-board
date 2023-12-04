import React, { useEffect } from "react";
import MessageCard from "./MessageCard";
import { useData } from "../DataContext";
import Typography from '@mui/material/Typography';

const MessageBoard = () => {
  const { messages, fetchMessages, messageUpdated,loggedInUserId  } = useData()

  useEffect(() => {
    if (loggedInUserId){
      fetchMessages();
    }
  }, [messageUpdated]);

  const sortedMessages = messages.slice().sort((a, b) => {
    const timeA = new Date(a.CreatedAt).getTime();
    const timeB = new Date(b.CreatedAt).getTime();
    return timeB - timeA;
  });
  

  return (
    <div className="p-2" >
      <Typography variant="h5" className="title" sx={{color:"#434343", fontSize: "18px" , fontWeight: "bold"}}>All Messages</Typography>
      <div className="col p-1 mt-4" style={{ maxHeight: 'calc(50vh - 120px)', overflowY: 'auto', }}>
        {sortedMessages.map((message, index) => (
          <div key={message.id} >
            <MessageCard data={message}
            index={index} />
          </div>
        ))}
      </div>
    </div>
  );
};

export default MessageBoard;

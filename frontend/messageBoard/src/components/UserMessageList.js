import React from "react";
import "../appStyles.css";
import { useData } from "../DataContext";
import Typography from "@mui/material/Typography";
import UserMessageCard from "./UserMessageCard";

const UserMessageList = () => {
  const { userMessages } = useData();
  const sortedMessages = userMessages.slice().sort((a, b) => {
    const timeA = new Date(a.CreatedAt).getTime();
    const timeB = new Date(b.CreatedAt).getTime();
    return timeB - timeA;
  });
console.log("sortedMessages:", sortedMessages);
  return (
    <div className="p-2">
      <Typography variant="h5" className="title" sx={{color:"#434343", fontSize: "18px" , fontWeight: "bold"}}>
        Your Messages
      </Typography>
      {sortedMessages.length === 0 ? (
        <Typography variant="body1" className="message">
          Start by sending out your first message
        </Typography>
      ) : (
        <div
          className="col p-1 mt-4"
          style={{ maxHeight: "calc(50vh - 120px)", overflowY: "auto" }}
        >
          {sortedMessages.map((message) => (
            <div key={message.id} className="col-md-4 mb-3">
              <UserMessageCard data={message} />
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default UserMessageList;

import React, { useState, useRef } from "react";
import ReplyCard from "./ReplyCard";
import { Avatar, Button, TextField } from "@mui/material";
import { useData } from "../DataContext";

const MessageCard = ({ data, index }) => {
  const { id, user_name, content, CreatedAt, user_id } = data;

  const {
    handlePostMessage,
    loggedInUserId,
    handleDeleteMessage,
    updateMessages,
    getTimeDifference,
    setShowErrorPage,
    setErrorMessage,
    fetchReplies: fetchRepliesFromData,
  } = useData();

  const [showActions, setShowActions] = useState(false);
  const [replies, setReplies] = useState([]);

  const contentRef = useRef(null);

  const getInitials = (name) => {
    const nameArray = name.split(" ");
    return nameArray
      .map((word) => word.charAt(0))
      .join("")
      .toUpperCase();
  };

  const fetchReplies = async () => {
    const data = await fetchRepliesFromData(id);
    setReplies(data);
  };

  const toggleActions = async () => {
    await fetchReplies();
    setShowActions(!showActions);
  };

  const stopPropagation = (event) => {
    event.stopPropagation();
  };

  const handleSubmitRrply = async (event) => {
    event.preventDefault();
    event.stopPropagation();
    if (contentRef && contentRef.current && contentRef.current.value) {
      const messageData = {
        content: contentRef.current.value,
        user_id: loggedInUserId,
        parent_id: id,
      };
      console.log("messageData", messageData);
      try {
        const message = await handlePostMessage(messageData);
        if (message && loggedInUserId && id) {
          console.log("message reply successful!", message);
          await fetchReplies();
        } else {
          console.log("Message reply failed.");
        }
      } catch (error) {
        console.error("message reply error:", error);
      }
      contentRef.current.value = "";
    } else {
      console.error("No content available in the TextField.");
    }
  };

  const handleDelete = async (event) => {
    event.stopPropagation();
    if (loggedInUserId !== user_id) {
      setShowErrorPage(true);
      setErrorMessage("You are not authorized to delete this message.");
      console.error("You are not authorized to delete this message.");
      return;
    }
    if (id) {
      console.log("handleDeleteMessage id:", id);
      try {
        await handleDeleteMessage(id);
        updateMessages();
      } catch (error) {
        console.error("handleDeleteMessage error:", error);
      }
    }
  };

  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        width: "80%",
        flexDirection: index % 2 === 0 ? "row" : "row-reverse",
        marginLeft: "5%",
      }}
    >
      <div style={{ display: "flex", alignItems: "center" }}>
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            marginBottom: "3%",
          }}
        >
          <Avatar
            sx={{
              width: 30,
              height: 30,
              marginRight: "10px",
            }}
          >
            {getInitials(user_name)}
          </Avatar>
          <p className="msg-user-name">{user_name}</p>
        </div>
        {showActions ? (
          <div
            className="msd-card-reply text-white p-4"
            onClick={toggleActions}
          >
            <p className="text-black">{content}</p>
            <div>
              <TextField
                id="outlined-multiline-flexible"
                inputRef={contentRef}
                label="Your Reply"
                multiline
                maxRows={4}
                variant="outlined"
                onClick={stopPropagation}
                style={{ width: "100%", marginBottom: "10px" }}
              />
            </div>
            <div
              style={{
                display: "flex",
                justifyContent: "flex-end",
                marginTop: "10px",
              }}
            >
              <Button
                variant="contained"
                onClick={handleSubmitRrply}
                style={{
                  marginRight: "10px",
                  color: "white",
                  backgroundColor: "#785fbd",
                }}
              >
                Reply
              </Button>
              <Button
                variant="contained"
                onClick={handleDelete}
                style={{ color: "white", backgroundColor: "#785fbd" }}
              >
                Delete
              </Button>
            </div>
            <div>
              {replies.map((reply) => (
                <div key={reply.id} className="col-md-4 mb-3">
                  <ReplyCard
                    time={getTimeDifference(reply.CreatedAt)}
                    name={reply.user_name}
                    content={reply.content}
                    avatar={getInitials(reply.user_name)}
                    handleDelete={handleDelete}
                    id={reply.id}
                    user_id={reply.user_id}
                    setShowActions={setShowActions}
                  />
                </div>
              ))}
            </div>
          </div>
        ) : (
          <div
            style={{
              display: "flex",
              flexDirection: "column",
            }}
          >
            <p
              onClick={toggleActions}
              className={
                index % 2 === 0 ? "evenClass text-white p-3" : "oddClass"} >
              {content}
            </p>
            <p className="msg-time">{getTimeDifference(CreatedAt)}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default MessageCard;

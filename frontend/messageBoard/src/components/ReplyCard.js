import React from "react";
import { useData } from "../DataContext";
import { Avatar } from "@mui/material";
import IconButton from "@mui/material/IconButton";
import DeleteIcon from "@mui/icons-material/Delete";

const ReplyCard = ({name, time,content,avatar,id,setShowActions,user_id,}) => {
  const {
    handleDeleteMessage,
    updateMessages,
    loggedInUserId,
    setShowErrorPage,
    setErrorMessage,
  } = useData();

  const handleReplyDelete = async (event) => {
    event.stopPropagation();
    if (loggedInUserId !== user_id) {
      setShowErrorPage(true);
      setErrorMessage("You are not authorized to delete this message.");
      console.error("You are not authorized to delete this message.");
      return;
    }
    if (id) {
      try {
        await handleDeleteMessage(id);
        updateMessages();
        setShowActions(false);
        console.log("handleDeleteMessage id:", id);
      } catch (error) {
        console.error("handleDeleteMessage error:", error);
      }
    }
  };
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-start",
        width: "200%",
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          marginBottom: "5px",
        }}
      >
        <Avatar sx={{ width: 25, height: 25, marginRight: "5px" }}>
          {avatar}
        </Avatar>
        <p className="msg-user-name">{name}</p>
      </div>

      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "flex-start",
        }}
      >
        <div
          style={{
            display: "flex",
            flexDirection: "row",
            alignItems: "flex-start",
          }}
        >
          <p className="rply-card p-3">{content}</p>
          <IconButton
            variant="contained"
            color="error"
            size="small"
            onClick={handleReplyDelete}
            style={{ marginTop: "10px" }}
          >
            <DeleteIcon />
          </IconButton>
        </div>
        <p className="msg-time">{time}</p>
      </div>
    </div>
  );
};

export default ReplyCard;

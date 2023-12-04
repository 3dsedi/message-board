import React, {useState} from 'react'
import { useData } from '../DataContext'
import IconButton from "@mui/material/IconButton";
import DeleteIcon from "@mui/icons-material/Delete";

const UserMessageCard = (props) => {
    const {getTimeDifference, handleDeleteMessage, loggedInUserId, updateMessages} = useData()
    const {content, CreatedAt, id, user_id } = props.data
    const [showActions, setShowActions] = useState(false);

    const toggleActions = async () => {
      setShowActions(!showActions);
    };

    const handleDelete = async () => {
        if (loggedInUserId !== user_id) {
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
    <div style={{ display: "flex", alignItems: "center" }} onClick={toggleActions}>
    {showActions ? (
      <div style={{ display: "flex", flexDirection: "column", alignItems: "flex-start" }}>
        <div style={{ display: "flex", flexDirection: "row", alignItems: "flex-start" }}>
          <p className="msg-user-card text-black p-3">{content}</p>
          <div style={{ width: "200px", marginLeft: "0px" }}> 
            <IconButton
              variant="contained"
              color="error"
              size="small"
              onClick={handleDelete}
              style={{ marginTop: "10px" }}
            >
              <DeleteIcon />
            </IconButton>
          </div>
        </div>
        <p className="msg-user-time">{getTimeDifference(CreatedAt)}</p>
      </div>
    ) : (
      <div style={{ display: "flex", flexDirection: "column", alignItems: "flex-start" }}>
        <p className="msg-user-card text-black p-3">{content}</p>
        <p className="msg-user-time">{getTimeDifference(CreatedAt)}</p>
      </div>
    )}
  </div>
  
  );
}

export default UserMessageCard
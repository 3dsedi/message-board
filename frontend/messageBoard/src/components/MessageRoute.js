import React from "react";
import { useData } from "../DataContext";
import UserMessageList from "./UserMessageList";
import MessageBoard from "./MessageBoard";
import NewMessageForm from "./NewMessageForm";

const MessageRoute = () => {
    const {selectMessage} = useData();
    
    return (
        <div style={{ background: "#f5f5f5", width: "100%", height: "100%", display: "flex", flexDirection: "column", padding:'2%' }}>
            <div style={{ flex: 1 }}>
                <div>
                    {selectMessage === "all" && <MessageBoard />}
                    {selectMessage === "users" && <UserMessageList />}
                </div>
            </div>

            <div style={{}}>
                <NewMessageForm />
            </div>
        </div>
    );
};

export default MessageRoute;

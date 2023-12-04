import React, { createContext, useState, useContext, useEffect } from "react";

const DataContext = createContext();

export const DataProvider = ({ children }) => {
  const apiBaseUrl = `http://localhost:8080`;

  const [loggedInUserId, setLoggedInUserId] = useState(null);
  const [loggedInUserData, setLoggedInUserData] = useState(null);

  const [messages, setMessages] = useState([]);
  const [replies, setReplies] = useState([]);
  const [userMessages, setUserMessages] = useState([]);
  const [messageUpdated, setMessageUpdated] = useState(false);

  const [showErrorPage, setShowErrorPage] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  
  const [selectMessage, setSelectMessage] = useState("all");

  const closeErrorPage = () => {
    setShowErrorPage(false);
  };

  function getTimeDifference(timestamp) {
    const createdTime = new Date(timestamp).getTime();
    const currentTime = new Date().getTime();
    const difference = currentTime - createdTime;

    const seconds = Math.floor(difference / 1000);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);
    const days = Math.floor(hours / 24);

    if (days > 0) {
      return `${days} day${days !== 1 ? "s" : ""} ago`;
    } else if (hours > 0) {
      return `${hours} hour${hours !== 1 ? "s" : ""} ago`;
    } else if (minutes > 0) {
      return `${minutes} minute${minutes !== 1 ? "s" : ""} ago`;
    } else {
      return `${seconds} second${seconds !== 1 ? "s" : ""} ago`;
    }
  }

  const handleLogin = async (formData) => {
    try {
      const response = await fetch(`${apiBaseUrl}/user/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });
      const data = await response.json();
      if (data && data.id) {
        setLoggedInUserId(data.id);
        setLoggedInUserData(data);
        try {
          const messagesResponse = await fetch(
            `${apiBaseUrl}/message/${data.id}`
          );
          const messagesData = await messagesResponse.json();
          setUserMessages(messagesData);
        } catch (error) {
          console.error("Error fetching messages:", error);
        }
      }
      return data;
    } catch (error) {
      console.error("Login error:", error);
    }
  };

  const logout = () => {
    setLoggedInUserId(null);
    setLoggedInUserData(null);
  };

  const handleSignUp = async (formData) => {
    try {
      const response = await fetch(`${apiBaseUrl}/user`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });
      if (response.ok) {
        handleLogin(formData);
      } else {
        console.error("Sign up failed");
      }
    } catch (error) {
      console.error("Error during sign-up:", error);
      throw error;
    }
  };

  const fetchMessages = async () => {
    try {
      const response = await fetch(`${apiBaseUrl}/message`);
      const data = await response.json();
      setMessages(data);
    } catch (error) {}
  };

  const updateMessages = () => {
    setMessageUpdated((prevState) => !prevState);
  };

  useEffect(() => {
    console.log("replies", replies);
  }, [replies]);

  const fetchReplies = async (messageId) => {
    try {
      const response = await fetch(
        `http://localhost:8080/message/reply/${messageId}`
      );
      const data = await response.json();
      setReplies(data);
      return data;
    } catch (error) {}
  };

  const handlePostMessage = async (messageData) => {
    try {
      const response = await fetch(`http://localhost:8080/message`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(messageData),
      });

      if (!response.ok) {
        throw new Error("Failed to post message");
      }

      const data = await response.json();
      const messagesResponse = await fetch(
        `http://localhost:8080/message/${loggedInUserId}`
      );
      const messagesData = await messagesResponse.json();
      setUserMessages(messagesData);
      return data;
    } catch (error) {
      console.error("Error during posting message:", error);
      throw error;
    }
  };

  const handleDeleteMessage = async (messageId) => {
    try {
      const response = await fetch(
        `http://localhost:8080/message/${messageId}`,
        {
          method: "DELETE",
        }
      );
      if (!response.ok) {
        throw new Error("Failed to delete message");
      }
      setUserMessages((prevMessages) =>
        prevMessages.filter((message) => message.id !== messageId)
      );
    } catch (error) {
      console.error("Error during deleting message:", error);
      throw error;
    }
  };

  return (
    <DataContext.Provider
      value={{
        messages,
        replies,
        loggedInUserId,
        loggedInUserData,
        messageUpdated,
        userMessages,
        errorMessage,
        showErrorPage,
        selectMessage,
        setSelectMessage,
        setShowErrorPage,
        setErrorMessage,
        setLoggedInUserId,
        fetchMessages,
        fetchReplies,
        handleLogin,
        handleSignUp,
        logout,
        handlePostMessage,
        handleDeleteMessage,
        updateMessages,
        getTimeDifference,
        closeErrorPage,
      }}
    >
      {children}
    </DataContext.Provider>
  );
};

export const useData = () => {
  return useContext(DataContext);
};

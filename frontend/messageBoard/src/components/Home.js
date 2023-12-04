import React from "react";
import { useData } from "../DataContext";
import SignInSignUp from "./SignInSignUp";
import "../appStyles.css";
import MessageRoute from "./MessageRoute";

const Home = () => {
  const { loggedInUserId } = useData();
  return (
    <div className="main-container">
      <div className="centered-card">
        <div className="components-wrapper">
          {!loggedInUserId && (
            <>
              <div className="login-main-container">
                <SignInSignUp />
              </div>
              <div className="start-page-msg">
                <p className="start-page-msg-text">
                  Welcome to the message board
                </p>
                <p className="start-page-msg-text">
                  Please login to see the messages
                </p>
              </div>
            </>
          )}
          {loggedInUserId && (
            <>
              <div className="login-main-container">
                <SignInSignUp />
              </div>
              <div className="start-page-msg">
                <MessageRoute />
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
};

export default Home;

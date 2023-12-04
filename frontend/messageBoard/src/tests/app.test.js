import React from "react";
import { render, fireEvent } from "@testing-library/react";
import '@testing-library/jest-dom';
import SignUpForm from "../components/SignUpForm";
import NewMessageForm from "../components/NewMessageForm";
import MessageBoard from "../components/MessageBoard";
import SignInForm from "../components/SignInForm";
import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

const mockHandleSignUp = jest.fn();
const mockHandlePostMessage = jest.fn();
const mockFetchMessages = jest.fn();
const mockHandleDeleteMessage = jest.fn();
const messages = [];

let loggedInUserId = "123e4567-e89b-12d3-a456-426614174000";


jest.mock("../DataContext", () => ({
  useData: () => ({
    handleSignUp: mockHandleSignUp,
    handlePostMessage: mockHandlePostMessage,
    fetchMessages: mockFetchMessages,
    loggedInUserId: loggedInUserId,
    handleDeleteMessage: mockHandleDeleteMessage,
    messages: messages,
  }),
}));

test("Sign Up a New User", () => {
  render(<SignUpForm />);

  const nameInput = screen.getByLabelText("Name");
  const emailInput = screen.getByLabelText("Email");
  const passwordInput = screen.getByLabelText("Password");
  const submitButton = screen.getByRole("button", { name: /Sign Up/i });

  userEvent.type(nameInput, "Test User");
  userEvent.type(emailInput, "test@example.com");
  userEvent.type(passwordInput, "password123");

  fireEvent.click(submitButton);

  expect(mockHandleSignUp).toHaveBeenCalledWith({
    user_name: "Test User",
    email: "test@example.com",
    password: "password123",
  });
});

test("Post a New Message", () => {
  render(<NewMessageForm />);

  const messageInput = screen.getByPlaceholderText("Type here ...");
  const submitButton = screen.getByRole("button", { name: /Send Message/i });

  userEvent.type(messageInput, "Test message");
  fireEvent.click(submitButton);

  expect(mockHandlePostMessage).toHaveBeenCalledWith({
    content: "Test message",
    user_id: loggedInUserId,
  });
});

test("Fetch Messages", () => {
  render(<MessageBoard />);

  expect(mockFetchMessages).toHaveBeenCalled();
});

test("Accses Login Form if Not Loged in", () => {
  render(<SignInForm />);

  const emailInput = screen.getByLabelText("Email");
  const passwordInput = screen.getByLabelText("Password");
  const loginButton = screen.getByRole("button", { name: "Login" });

  expect(emailInput).toBeInTheDocument();
  expect(passwordInput).toBeInTheDocument();
  expect(loginButton).toBeInTheDocument();
});


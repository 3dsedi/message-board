import React, { useRef } from "react";
import { useData } from "../DataContext";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";

const SignUpForm = () => {
  const { handleSignUp, setShowErrorPage, setErrorMessage } = useData();

  const emailRef = useRef(null);
  const passwordRef = useRef(null);
  const nameRef = useRef(null);

  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = {
      email: emailRef.current.value,
      password: passwordRef.current.value,
      user_name: nameRef.current.value,
    };
    if (!emailRef || !passwordRef || !nameRef) {
      setErrorMessage("Please fill in all fields");
      setShowErrorPage(true);
      return;
    }
    try {
      await handleSignUp(formData);
      console.log("User created!");
    } catch (error) {
      console.error("SignUp error:", error);
      setErrorMessage("SignUp error occurred. Please try again.");
      setShowErrorPage(true);
    }
  };

  return (
    <div className="p-4">
      <form  style={{ marginTop: "30px" }}>
        <TextField
          type="text"
          label="Name"
          variant="outlined"
          inputRef={nameRef}
          fullWidth
          margin="normal"
          size="small"
          className="custom-textfield"
        />
        <TextField
          type="email"
          label="Email"
          variant="outlined"
          inputRef={emailRef}
          fullWidth
          margin="normal"
          size="small"
          className="custom-textfield"
        />
        <TextField
          type="password"
          label="Password"
          variant="outlined"
          inputRef={passwordRef}
          fullWidth
          margin="normal"
          size="small"
          className="custom-textfield"
        />
        <Button variant="contained" color="primary" type="submit" onClick={handleSubmit}>
          Sign Up
        </Button>
      </form>
    </div>
  );
};

export default SignUpForm;

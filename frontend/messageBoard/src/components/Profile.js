import React from "react";
import { useData } from "../DataContext";
import Avatar from "@mui/material/Avatar";
import Typography from "@mui/material/Typography";
import SettingsIcon from "@mui/icons-material/Settings";
import RateReviewIcon from "@mui/icons-material/RateReview";
import QuestionAnswerIcon from "@mui/icons-material/QuestionAnswer";
import NotificationsIcon from "@mui/icons-material/Notifications";
import IconButton from "@mui/material/IconButton";
import profilePic from "./S3D_1072.jpg";

const Profile = () => {
  const { loggedInUserData, logout, selectMessage, setSelectMessage } =
    useData();
  console.log(selectMessage);

  const handleIconClick = (icon) => {
    if (icon === "home") {
      setSelectMessage("all");
    } else if (icon === "mail") {
      setSelectMessage("users");
    }
  };

  const userName = loggedInUserData?.user_name;
  const userEmail = loggedInUserData?.email;

  const handleLogout = () => {
    logout();
  };

  return (
    <div
      style={{
        margin: "auto",
        marginTop: "30%",
        textAlign: "center",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <div style={{ position: "relative", display: "inline-block" }}>
        <Avatar
          src={profilePic}
          sx={{ bgcolor: "gray", width: 70, height: 70, margin: "auto" }}
        ></Avatar>
        <div
          style={{
            position: "absolute",
            top: "5px",
            right: "5px",
            borderRadius: "50%",
            width: "15px",
            height: "15px",
            backgroundColor: "green",
            border: "2px solid white",
          }}
        ></div>
      </div>

      <IconButton
        variant="contained"
        sx={{
          color: "#dfdfdf",
          fontSize: "10px",
          marginTop: "5%",
          fontWeight: "bold",
        }}
        onClick={handleLogout}
      >
        Logout
      </IconButton>
      <div style={{ marginTop: "20%" }}>
        <Typography
          variant="h5"
          gutterBottom
          sx={{ color: "#888888", fontSize: "15px" }}
        >
          {userName}
        </Typography>
        <Typography
          variant="body1"
          gutterBottom
          sx={{ color: "#888888", marginTop: "0%", fontSize: "15px" }}
        >
          {userEmail}
        </Typography>
      </div>
      <div style={{ marginTop: "10%" }}>
        <div style={{ textAlign: "center" }}>
          <IconButton
            sx={{ color: "gray" }}
            onClick={() => handleIconClick("home")}
          >
            <QuestionAnswerIcon />
          </IconButton>
          <Typography variant="body2" sx={{color: "#888888", fontSize:'10px'}}>All Messages</Typography>
        </div>

        <div style={{ textAlign: "center" }}>
          <IconButton
            sx={{ color: "gray" }}
            onClick={() => handleIconClick("mail")}
          >
            <RateReviewIcon />
          </IconButton>
          <Typography variant="body2" sx={{color: "#888888", fontSize:'10px'}}>Your Messages</Typography>
        </div>

        <div style={{ textAlign: "center" }}>
          <IconButton sx={{ color: "gray" }}>
            <NotificationsIcon />
          </IconButton>
          <Typography variant="body2" sx={{color: "#888888", fontSize:'10px'}} >Notifications</Typography>
        </div>
      </div>
      <IconButton sx={{ color: "gray", marginTop:'5%' }}>
        <SettingsIcon />
      </IconButton>
    </div>
  );
};

export default Profile;

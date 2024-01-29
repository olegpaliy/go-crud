import React from "react";

const UserDetail = ({ user }) => {
  return (
    <div>
      <h2>User Detail</h2>
      <p>Name: {user.name}</p>
    </div>
  );
};

export default UserDetail;

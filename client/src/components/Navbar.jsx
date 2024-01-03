import { Link } from "react-router-dom";

const Navbar = ({ user }) => {
  const logout = () => {
    window.open("http://localhost:5000/auth/logout", "_self");
  };
  return (
    <div className="navbar">
      <span className="logo">
        <Link className="link" to="/">
          SnapImg
        </Link>
      </span>
      {user ? (
        <ul className="list">
          <li className="listItem">
            <Link className="link" to="profile">
              <img src="./img/login.png" alt="" className="avatar" />
            </Link>
          </li>
          <li className="listItem">ChanChin</li>
          <li className="listItem" onClick={logout}>
            <Link className="link" to="logout">
              Logout
            </Link>
          </li>
        </ul>
      ) : (
        <Link className="link" to="login">
          Login
        </Link>
      )}
    </div>
  );
};

export default Navbar;

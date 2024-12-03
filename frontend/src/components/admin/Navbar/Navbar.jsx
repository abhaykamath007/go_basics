import React from 'react'
import { Link } from 'react-router-dom'
import "./Navbar.css"
import { useAuth } from '../../../context/AuthContext'


const Navbar = () => {

  const {logout,username} = useAuth();

  return (
    <div className='navbar'>
      <div className='navbar-logo'>
        <Link to="/admin">My Library</Link>
      </div>
      <ul className='navbar-links'>
        <li><Link to='/admin'>Home</Link></li>
        <li><Link to='/admin/manage-books'>Manage Books</Link></li>
        <li><Link to='/admin/manage-users'>Manage Users</Link></li>
      </ul>
      <div className='navbar-auth'>
        <span>{username}</span>
        <button className='navbar-button' onClick={logout}>Logout</button>
      </div>
    </div>
  )
}

export default Navbar

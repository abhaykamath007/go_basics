import React from 'react'
import { Link } from 'react-router-dom'
import "./Navbar.css"

const Navbar = () => {
  return (
    <div className='navbar'>
      <div className='navbar-logo'>
        <Link to="/">My Library</Link>
      </div>
      <ul className='navbar-links'>
        <li><Link to='/'>Home</Link></li>
        <li><Link to='/books'>Books</Link></li>
        <li><Link to='/account'>My Account</Link></li>
      </ul>
      <div className='navbar-auth'>
        <button className='navbar-button'>Logout</button>
      </div>
    </div>
  )
}

export default Navbar
